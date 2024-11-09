package utils_notion

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	model_notion "github.com/kojikawazu/backend/model/notion"
)

// ChangeNotionResponse はNotionのレスポンスをProcessedNotionResponseに変換する
func ChangeNotionResponse(notionResp model_notion.NotionGetDatasResponse) []model_notion.ProcessedNotionResponse {
	// 結果を interface{} スライスに変換
	var results []model_notion.ProcessedNotionResponse
	for _, result := range notionResp.Results {
		processed := model_notion.ProcessedNotionResponse{}

		// IDのチェック
		if len(result.Properties.ID.Title) > 0 {
			processed.ID = result.Properties.ID.Title[0].Text.Content
		}

		// Titleのチェック
		if len(result.Properties.NotionTitle.RichText) > 0 {
			processed.Title = result.Properties.NotionTitle.RichText[0].Text.Content
		}

		// Contentsのチェック
		if len(result.Properties.Contents.RichText) > 0 {
			processed.Contents = result.Properties.Contents.RichText[0].Text.Content
		}

		// Kindのチェック
		if len(result.Properties.Kind.RichText) > 0 {
			processed.Kind = result.Properties.Kind.RichText[0].Text.Content
		}

		// CreatedAtのチェック
		if len(result.Properties.CreatedAt.RichText) > 0 {
			processed.CreatedAt = result.Properties.CreatedAt.RichText[0].Text.Content
		}

		results = append(results, processed)
	}

	return results
}

// CreateNotionRequest creates a request to Notion API
func CreateNotionRequest(url string, jsonData []byte, method string) (*http.Request, error) {
	notionApiKey := os.Getenv("NOTION_API_KEY")

	// リクエストボディを作成
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// ヘッダーを設定
	req.Header.Set("Authorization", "Bearer "+notionApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2022-06-28")

	return req, nil
}

// HandleNotionResponse handles the response from Notion API
func HandleNotionResponse(resp *http.Response) error {
	// レスポンスのステータスコードが200または201でない場合はエラー
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to add page, status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}
	return nil
}
