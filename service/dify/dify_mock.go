package service_dify

import (
	model_dify "github.com/kojikawazu/backend/model/dify"
	"github.com/stretchr/testify/mock"
)

// DifyService のモック
type MockDifyService struct {
	mock.Mock
}

// GetDatas はDifyからデータを取得する
func (s *MockDifyService) GetDatas(inputs map[string]interface{}) (*model_dify.DifyResponse, error) {
	args := s.Called(inputs)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model_dify.DifyResponse), nil
}
