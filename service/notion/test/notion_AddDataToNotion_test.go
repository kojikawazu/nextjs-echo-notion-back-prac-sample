package test_service_notion

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	model_notion "github.com/kojikawazu/backend/model/notion"
	service_notion "github.com/kojikawazu/backend/service/notion"
)

func TestAddDataToNotion_Success(t *testing.T) {
	// テスト環境変数の設定
	testToken := "test-token"
	os.Setenv("NOTION_API_KEY", testToken)
	os.Setenv("NOTION_API_URL", "https://api.notion.com/v1/pages")

	// モックのレスポンスを設定
	mockClient := &service_notion.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// メソッドの検証
			if req.Method != "POST" {
				t.Errorf("Expected POST request, got %s", req.Method)
			}

			// ヘッダーの検証
			expectedHeaders := map[string]string{
				"Authorization":  "Bearer " + testToken,
				"Content-Type":   "application/json",
				"Notion-Version": "2022-06-28",
			}

			for key, expected := range expectedHeaders {
				if actual := req.Header.Get(key); actual != expected {
					t.Errorf("Expected %s header '%s', got '%s'", key, expected, actual)
				}
			}

			// 成功レスポンスを返す
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`{"object": "page", "id": "test-page-id"}`)),
			}, nil
		},
	}

	// テスト用のサービスを作成
	service := &service_notion.NotionServiceImpl{
		Token:      testToken,
		HTTPClient: mockClient,
	}

	// テスト用のリクエストを作成
	req := model_notion.NotionCreateRequest{
		Parent: model_notion.Parent{
			DatabaseID: "test-database-id",
		},
		Properties: map[string]model_notion.Property{
			"notion-title": {
				Title: []model_notion.Text{
					{
						Text: model_notion.Content{
							Content: "テストタイトル",
						},
					},
				},
			},
			"contents": {
				RichText: []model_notion.Text{
					{
						Text: model_notion.Content{
							Content: "テスト内容",
						},
					},
				},
			},
			"kind": {
				Select: &model_notion.Select{
					Name: "テスト",
				},
			},
			"created_at": {
				Date: &model_notion.Date{
					Start: "2024-01-01",
				},
			},
		},
	}

	// テスト実行
	err := service.AddDataToNotion(req)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestAddDataToNotion_Error(t *testing.T) {
	// テスト環境変数の設定
	testToken := "test-token"
	os.Setenv("NOTION_API_KEY", testToken)
	os.Setenv("NOTION_API_URL", "https://api.notion.com/v1/pages")

	// エラーケース用のモックを設定
	mockClient := &service_notion.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			// メソッドの検証
			if req.Method != "POST" {
				t.Errorf("Expected POST request, got %s", req.Method)
			}

			// ヘッダーの検証
			expectedHeaders := map[string]string{
				"Authorization":  "Bearer " + testToken,
				"Content-Type":   "application/json",
				"Notion-Version": "2022-06-28",
			}

			for key, expected := range expectedHeaders {
				if actual := req.Header.Get(key); actual != expected {
					t.Errorf("Expected %s header '%s', got '%s'", key, expected, actual)
				}
			}

			return &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(bytes.NewBufferString(`{"object": "error", "status": 400, "message": "Invalid request"}`)),
			}, nil
		},
	}

	// テスト用のサービスを作成
	service := &service_notion.NotionServiceImpl{
		Token:      testToken,
		HTTPClient: mockClient,
	}

	// テスト用のリクエストを作成
	req := model_notion.NotionCreateRequest{
		Parent: model_notion.Parent{
			DatabaseID: "test-database-id",
		},
		Properties: map[string]model_notion.Property{
			"notion-title": {
				Title: []model_notion.Text{
					{
						Text: model_notion.Content{
							Content: "テストタイトル",
						},
					},
				},
			},
		},
	}

	// テスト実行
	err := service.AddDataToNotion(req)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
