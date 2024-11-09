package service_dify

import (
	"github.com/go-resty/resty/v2"
	model_dify "github.com/kojikawazu/backend/model/dify"
)

// DifyService is the interface for the Dify service
type DifyService interface {
	GetDatas(inputs map[string]interface{}) (*model_dify.DifyResponse, error)
}

// DifyServiceImpl is the implementation of the Dify service
type DifyServiceImpl struct {
	Client *resty.Client
	Token  string
}

// NewDifyService creates a new Dify service
func NewDifyService(token string) *DifyServiceImpl {
	client := resty.New()
	return &DifyServiceImpl{
		Client: client,
		Token:  token,
	}
}
