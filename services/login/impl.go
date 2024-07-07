package login

import (
	"github.com/witwoywhy/go-cores/logger"
	"github.com/witwoywhy/go-cores/logs"
)

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) Execute(request Request, l logger.Logger) (*Response, error) {
	spl, end := logs.NewSpanLogAction(l, "login")
	defer end()

	spl.Info("SOME ONE TRY LOGIN")

	return &Response{
		AccountNo:   "0-0000-0000",
		AccessToken: "ABCDEFG",
	}, nil
}
