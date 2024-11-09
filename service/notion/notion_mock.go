package service_notion

import (
	"net/http"

	model_notion "github.com/kojikawazu/backend/model/notion"
	"github.com/stretchr/testify/mock"
)

// NotionService のモック
type MockNotionService struct {
	mock.Mock
}

// モックのHTTPクライアント
type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do はHTTPリクエストを実行する
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

// GetDatasFromNotion はNotionからデータを取得する
func (s *MockNotionService) GetDatasFromNotion() ([]model_notion.ProcessedNotionResponse, error) {
	args := s.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model_notion.ProcessedNotionResponse), nil
}

// AddDataToNotion はNotionにデータを追加する
func (s *MockNotionService) AddDataToNotion(req model_notion.NotionCreateRequest) error {
	args := s.Called(req)
	if args.Get(0) == nil {
		return args.Error(1)
	}
	return nil
}
