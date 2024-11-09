package service_notion

import (
	"net/http"

	model_notion "github.com/kojikawazu/backend/model/notion"
)

// NotionService はNotionサービスのインターフェース
type NotionService interface {
	GetDatasFromNotion() ([]model_notion.ProcessedNotionResponse, error)
	AddDataToNotion(req model_notion.NotionCreateRequest) error
}

// インターフェースを追加
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NotionServiceImpl はNotionサービスの実装
type NotionServiceImpl struct {
	Token      string
	HTTPClient HTTPClient
}

// NewNotionService は新しいNotionサービスを作成する
func NewNotionService(token string, httpClient HTTPClient) *NotionServiceImpl {
	return &NotionServiceImpl{Token: token, HTTPClient: httpClient}
}
