package handler_dify

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	model_dify "github.com/kojikawazu/backend/model/dify"
	model_notion "github.com/kojikawazu/backend/model/notion"
)

// GetDifyDataHandler gets data from the Dify API
func (dh *DifyHandler) GetDifyDataHandler(c echo.Context) error {
	fmt.Println("GetDifyDataHandler started.")

	// リクエストボディからJSONデータを取得
	var req model_dify.DifyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Invalid request format: %v", err),
		})
	}

	// データをDify APIに送信するためのパラメータ
	inputs := map[string]interface{}{
		"message": req.Message,
	}

	// Dify APIにリクエストを送信
	result, err := dh.DifyService.GetDatas(inputs)
	if err != nil {
		fmt.Printf("Error from Dify API: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Error from Dify API: %v", err),
		})
	}

	// レスポンスを返す
	fmt.Println("GetDifyDataHandler successed.")
	return c.JSON(http.StatusOK, result)
}

// GetDifyAndCreateNotionHandler gets data from the Dify API and creates a page in Notion
func (dh *DifyHandler) GetDifyAndCreateNotionHandler(c echo.Context) error {
	fmt.Println("GetDifyAndCreateNotionHandler started.")

	// ----------------------------------------------------------------------
	// 1. Dify APIからデータを取得
	// ----------------------------------------------------------------------

	// リクエストボディからJSONデータを取得
	var req model_dify.DifyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Invalid request format: %v", err),
		})
	}

	// データをDify APIに送信するためのパラメータ
	inputs := map[string]interface{}{
		"message": req.Message,
	}

	// Dify APIにリクエストを送信
	fmt.Println("Dify API request started.")
	result, err := dh.DifyService.GetDatas(inputs)
	if err != nil {
		fmt.Printf("Error from Dify API: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Error from Dify API: %v", err),
		})
	}

	// Dify APIのレスポンスをマップに格納
	difyTitle := result.Title
	difyContents := result.Contents
	difyKind := result.Kind
	fmt.Println("Dify API request successed.")
	// ----------------------------------------------------------------------
	// 2. Notion APIにデータを登録
	// ----------------------------------------------------------------------
	fmt.Println("Notion API request started.")
	notionDatabaseId := os.Getenv("NOTION_DATABASE_ID")

	// Notion APIに渡すためのリクエストを構築
	notionReq := model_notion.NotionCreateRequest{
		Parent: model_notion.Parent{
			DatabaseID: notionDatabaseId,
		},
		Properties: map[string]model_notion.Property{
			"id": {
				Title: []model_notion.Text{
					{
						Text: model_notion.Content{
							Content: uuid.New().String(),
						},
					},
				},
			},
			"notion-title": {
				RichText: []model_notion.Text{
					{
						Text: model_notion.Content{
							Content: difyTitle,
						},
					},
				},
			},
			"contents": {
				RichText: []model_notion.Text{
					{
						Text: model_notion.Content{
							Content: difyContents,
						},
					},
				},
			},
			"kind": {
				RichText: []model_notion.Text{
					{
						Text: model_notion.Content{
							Content: difyKind,
						},
					},
				},
			},
			"created_at": {
				RichText: []model_notion.Text{
					{
						Text: model_notion.Content{
							Content: time.Now().Format("2006/01/02"),
						},
					},
				},
			},
		},
	}

	// Notionにページを追加
	fmt.Println("Notion API request started.")
	if err := dh.NotionService.AddDataToNotion(notionReq); err != nil {
		fmt.Printf("Error from Notion API: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// レスポンスを返す
	fmt.Println("GetDifyAndCreateNotionHandler successed.")
	return c.JSON(http.StatusOK, result)
}
