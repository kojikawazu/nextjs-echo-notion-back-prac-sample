package service_dify

import (
	"encoding/json"
	"fmt"
	"os"

	model_dify "github.com/kojikawazu/backend/model/dify"
)

// GetDatas gets data from the Dify API
func (ds *DifyServiceImpl) GetDatas(inputs map[string]interface{}) (*model_dify.DifyResponse, error) {
	fmt.Println("GetData started.")

	// 環境変数からAPIのURLとユーザーIDを取得
	apiUrl := os.Getenv("DIFY_API_URL")
	difyUser := os.Getenv("DIFY_USER")

	// Dify APIにリクエストを送信
	resp, err := ds.Client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", ds.Token)).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"inputs": map[string]interface{}{
				"input_data": inputs["message"],
			},
			"response_mode": "blocking",
			"user":          difyUser,
		}).
		Post(apiUrl)

	if err != nil {
		fmt.Printf("GetData failed to send request: %v", err)
		return nil, err
	}

	// レスポンスの内容をログに出力（デバッグ用）
	//fmt.Printf("Response: %s\n", resp.String())

	// レスポンスをパース
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(resp.String()), &result); err != nil {
		fmt.Printf("GetData failed to parse response: %v", err)
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// データの構造に合わせて修正
	if data, ok := result["data"].(map[string]interface{}); ok {
		if outputs, ok := data["outputs"].(map[string]interface{}); ok {
			response := &model_dify.DifyResponse{
				Title:    outputs["output_title"].(string),
				Kind:     outputs["output_kind"].(string),
				Contents: outputs["output_contents"].(string),
			}
			return response, nil
		}
	}

	fmt.Printf("GetData failed: %v", result)
	return nil, fmt.Errorf("unexpected response format: %v", result)
}
