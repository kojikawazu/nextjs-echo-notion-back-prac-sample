package service_notion

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	model_notion "github.com/kojikawazu/backend/model/notion"
	utils_notion "github.com/kojikawazu/backend/utils/notion"
)

// GetDatasFromNotion はNotionからデータを取得する
func (ns *NotionServiceImpl) GetDatasFromNotion() ([]model_notion.ProcessedNotionResponse, error) {
	fmt.Println("GetDataFromNotion")

	// Notion APIのURLを取得
	notionDatabaseUrl := os.Getenv("NOTION_DATABASE_URL")
	notionDatabaseId := os.Getenv("NOTION_DATABASE_ID")
	notionApiUrl := notionDatabaseUrl + "/" + notionDatabaseId + "/query"

	// リクエストボディをJSONに変換（フィルターやソートの条件を含められます）
	jsonData, err := json.Marshal(map[string]interface{}{
		"page_size": 100, // 一度に取得するページ数（最大100）
		"sorts": []map[string]string{
			{
				"property":  "created_at",
				"direction": "descending",
			},
		},
	})
	if err != nil {
		fmt.Printf("failed to marshal request body: %v", err)
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	// リクエストを作成
	reqHTTP, err := utils_notion.CreateNotionRequest(notionApiUrl, jsonData, "POST")
	if err != nil {
		fmt.Printf("failed to create request: %v", err)
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// リクエストを送信
	resp, err := ns.HTTPClient.Do(reqHTTP)
	if err != nil {
		fmt.Printf("failed to send request: %v", err)
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// エラーレスポンスのチェック
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Printf("notion API error: status=%d, body=%s", resp.StatusCode, string(bodyBytes))
		return nil, fmt.Errorf("notion API error: status=%d, body=%s", resp.StatusCode, string(bodyBytes))
	}

	// レスポンスを構造体にデコード
	var notionResp model_notion.NotionGetDatasResponse
	if err := json.NewDecoder(resp.Body).Decode(&notionResp); err != nil {
		fmt.Printf("failed to decode response: %v", err)
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	// 結果を ProcessedNotionResponse スライスに変換
	results := utils_notion.ChangeNotionResponse(notionResp)

	fmt.Printf("Retrieved %d records from Notion\n", len(results))
	return results, nil
}

// AddDataToNotion はNotionにページを追加する
func (ns *NotionServiceImpl) AddDataToNotion(req model_notion.NotionCreateRequest) error {
	fmt.Println("AddPageToNotion")

	// リクエストボディをJSONに変換
	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Printf("failed to marshal request body: %v", err)
		return fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Notion APIのURLを取得
	notionApiUrl := os.Getenv("NOTION_API_URL")

	// リクエストを作成
	reqHTTP, err := utils_notion.CreateNotionRequest(notionApiUrl, jsonData, "POST")
	if err != nil {
		fmt.Printf("failed to create request: %v", err)
		return err
	}

	// リクエストを送信
	resp, err := ns.HTTPClient.Do(reqHTTP)
	if err != nil {
		fmt.Printf("failed to send request: %v", err)
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// レスポンスをハンドル
	if err := utils_notion.HandleNotionResponse(resp); err != nil {
		fmt.Printf("failed to handle response: %v", err)
		return err
	}

	fmt.Println("AddPage successed.")
	return nil
}
