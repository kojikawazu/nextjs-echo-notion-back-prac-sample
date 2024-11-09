package test_service_notion

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	service_notion "github.com/kojikawazu/backend/service/notion"
)

// 正常系のテスト
func TestGetDatasFromNotion_Success(t *testing.T) {
	// テスト用の環境変数設定
	os.Setenv("NOTION_DATABASE_URL", "https://api.notion.com/v1/databases")
	os.Setenv("NOTION_DATABASE_ID", "test-database-id")
	os.Setenv("NOTION_TOKEN", "test-token")

	// モックレスポンスの準備
	mockResponse := `{
		"results": [
			{
				"properties": {
					"id": {
						"title": [
							{
								"text": {
									"content": "test-id"
								}
							}
						]
					},
					"notion-title": {
						"rich_text": [
							{
								"text": {
									"content": "テストタイトル"
								}
							}
						]
					},
					"contents": {
						"rich_text": [
							{
								"text": {
									"content": "テスト内容"
								}
							}
						]
					},
					"kind": {
						"rich_text": [
							{
								"text": {
									"content": "テスト種類"
								}
							}
						]
					},
					"created_at": {
						"rich_text": [
							{
								"text": {
									"content": "2024-01-01"
								}
							}
						]
					}
				}
			}
		]
	}`

	// モックのHTTPクライアントを設定
	mockClient := &service_notion.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(mockResponse)),
			}, nil
		},
	}

	// NotionServiceの作成
	service := &service_notion.NotionServiceImpl{
		Token:      "test-token",
		HTTPClient: mockClient,
	}

	// テスト実行
	results, err := service.GetDatasFromNotion()

	// アサーション
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Len(t, results, 1)
	assert.Equal(t, "test-id", results[0].ID)
	assert.Equal(t, "テストタイトル", results[0].Title)
	assert.Equal(t, "テスト内容", results[0].Contents)
	assert.Equal(t, "テスト種類", results[0].Kind)
	assert.Equal(t, "2024-01-01", results[0].CreatedAt)
}

// エラー系のテスト
func TestGetDatasFromNotion_Error(t *testing.T) {
	// テスト用の環境変数設定
	os.Setenv("NOTION_DATABASE_URL", "https://api.notion.com/v1/databases")
	os.Setenv("NOTION_DATABASE_ID", "test-database-id")
	os.Setenv("NOTION_TOKEN", "test-token")

	tests := []struct {
		name           string
		mockStatusCode int
		mockResponse   string
		expectedError  string
	}{
		{
			name:           "APIエラー",
			mockStatusCode: http.StatusBadRequest,
			mockResponse:   `{"message": "Bad Request"}`,
			expectedError:  "notion API error",
		},
		{
			name:           "不正なJSONレスポンス",
			mockStatusCode: http.StatusOK,
			mockResponse:   `{invalid json}`,
			expectedError:  "failed to decode response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックのHTTPクライアントを設定
			mockClient := &service_notion.MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: tt.mockStatusCode,
						Body:       io.NopCloser(bytes.NewBufferString(tt.mockResponse)),
					}, nil
				},
			}

			// NotionServiceの作成
			service := &service_notion.NotionServiceImpl{
				Token:      "test-token",
				HTTPClient: mockClient,
			}

			// テスト実行
			results, err := service.GetDatasFromNotion()

			// アサーション
			assert.Error(t, err)
			assert.Nil(t, results)
			assert.Contains(t, err.Error(), tt.expectedError)
		})
	}
}
