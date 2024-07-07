package greeting

import (
	"fmt"

	"github.com/witwoywhy/go-cores/logger"
)

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) Execute(request Request, l logger.Logger) (*Response, error) {
	greeting := request.Name
	if request.Name == "" {
		greeting = "World!"
	}

	return &Response{
		Message: fmt.Sprintf("Hello %v", greeting),
	}, nil
}
