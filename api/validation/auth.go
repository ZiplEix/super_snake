package validation

import "github.com/ZiplEix/super_snake/api/request_models"

func Login(req request_models.LoginReq) error {
	return validate.Struct(req)
}

func Register(req request_models.RegisterReq) error {
	return validate.Struct(req)
}
