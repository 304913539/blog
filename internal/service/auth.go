package service

import "errors"

type AuthRequest struct {
	AppKey    string `json:"app_key" validate:"required"`
	AppSecret string `json:"app_secret" validate:"required"`
}

func (svc *Service) CheckAuth(params *AuthRequest) error {
	auth, err := svc.dao.GetAuth(params.AppKey, params.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist")
}
