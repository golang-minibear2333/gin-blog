package service

import "errors"

type AuthRequest struct {
	AppKey    string `header:"app_key" binding:"required"`
	AppSecret string `header:"app_secret" binding:"required"`
}

// CheckAuth 查询数据库只为了校验是否存在这条记录
func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errors.New("auth info does not exist.")
}
