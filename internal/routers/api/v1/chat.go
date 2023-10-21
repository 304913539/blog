package v1

import (
	"blog-service/global"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type Chat struct {
}

func NewChat() *Chat {
	return &Chat{}
}

func (baidu *Chat) Session(c *gin.Context) {
	data := struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  "Success",
		Message: "",
		Data: map[string]interface{}{
			"auth":  false,
			"model": "文心一言",
		}}

	c.JSON(200, data)
}

func (baidu *Chat) ChatProcess(c *gin.Context) {
	all, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	requestData := map[string]interface{}{
		"messages": []map[string]string{
			{"role": "user", "content": string(all)},
		},
		"stream": true,
	}
	marshal, err := json.Marshal(requestData)
	if err != nil {
		return
	}
	url := "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/eb-instant?access_token=" + GetAccessToken()
	//payload := strings.NewReader(`{"messages":[{"role":"user","content":"你好"},{"role":"assistant","content":"你好，有什么我可以帮助你的吗？"}]}`)
	payload := strings.NewReader(string(marshal))
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Connection", "keep-alive")
	c.Header("Cache-Control", "no-cache")
	scanner := bufio.NewScanner(res.Body)
	str := ""
	c.Stream(func(w io.Writer) bool {

		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "data:") {
				responseBuffer := []byte(line[6:])
				var responseChunk map[string]interface{}
				err := json.Unmarshal(responseBuffer, &responseChunk)
				if err != nil {
					fmt.Println("解码响应块时出错:", err)
					return false
				}
				content, ok := responseChunk["result"].(string)
				if !ok {
					return false

				}
				isEnd, ok := responseChunk["is_end"].(bool)
				if !ok {
					return false
				}
				if isEnd {
					break
				}
				str += content
				responseChunk["text"] = str
				jsonData, err := json.Marshal(responseChunk)
				if err != nil {
					fmt.Println("转换为 JSON 失败:", err)
					return false
				}
				fmt.Fprintf(c.Writer, "%s\n", jsonData)
				c.Writer.Flush()
			} else {
				fmt.Println(line)
				fmt.Fprintf(c.Writer, "%s\n", line)

			}
		}
		return false
	})
}

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", global.BaiduChat.ApiKey, global.BaiduChat.SecretKey)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal(body, &accessTokenObj)
	return accessTokenObj["access_token"]
}
