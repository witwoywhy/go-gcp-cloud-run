package greeting

import "github.com/witwoywhy/go-cores/logger"

type Service interface {
	Execute(request Request, l logger.Logger) (*Response, error)
}

type Request struct {
	Name string
}

type Response struct {
	Message string `json:"message"`
}
