package handler_notion

import (
	"fmt"
	"net/http"
	"os"

	model_notion "github.com/kojikawazu/backend/model/notion"
	"github.com/labstack/echo/v4"
)

// GetDatasFromNotion はNotionからデータを取得する
func (nh *NotionHandler) GetDatasFromNotion(c echo.Context) error {
	fmt.Println("GetDatasFromNotion started.")

	// Notion APIからデータを取得
	notionDatas, err := nh.NotionService.GetDatasFromNotion()
	if err != nil {
		fmt.Printf("Error from Notion API: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// レスポンスを返す
	fmt.Println("GetDatasFromNotion successed.")
	return c.JSON(http.StatusOK, notionDatas)
}

// AddDataToNotion はNotionにページを追加する
func (nh *NotionHandler) AddDataToNotion(c echo.Context) error {
	fmt.Println("AddDataToNotion started.")

	// リクエストボディからJSONデータを取得
	var req model_notion.CreatePageRequest
	if err := c.Bind(&req); err != nil {
		fmt.Printf("Invalid request format: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

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
							Content: req.ID,
						},
					},
				},
			},
			"notion-title": {
				RichText: []model_notion.Text{
					{
						Text: model_notion.Content{
							Content: req.Title,
						},
					},
				},
			},
			"contents": {
				RichText: []model_notion.Text{
					{
						Text: model_notion.Content{
							Content: req.Contents,
						},
					},
				},
			},
			"kind": {
				RichText: []model_notion.Text{
					{
						Text: model_notion.Content{
							Content: req.Kind,
						},
					},
				},
			},
			"created_at": {
				RichText: []model_notion.Text{
					{
						Text: model_notion.Content{
							Content: req.CreatedAt,
						},
					},
				},
			},
		},
	}

	// Notionにページを追加
	if err := nh.NotionService.AddDataToNotion(notionReq); err != nil {
		fmt.Printf("Error from Notion API: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// レスポンスを返す
	fmt.Println("AddDataToNotion successed.")
	return c.JSON(http.StatusOK, map[string]string{"message": "Page added successfully"})
}
