package login

import "github.com/witwoywhy/go-cores/logger"

type Service interface {
	Execute(request Request, l logger.Logger) (*Response, error)
}

type Request struct {
	Email string `json:"email"`
}

type Response struct {
	AccountNo   string `json:"accountNo"`
	AccessToken string `json:"accessToken"`
}
