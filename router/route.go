package router

import (
	"net/http"
	"os"

	handler_dify "github.com/kojikawazu/backend/handler/dify"
	handler_notion "github.com/kojikawazu/backend/handler/notion"
	service_dify "github.com/kojikawazu/backend/service/dify"
	service_notion "github.com/kojikawazu/backend/service/notion"

	"github.com/labstack/echo/v4"
)

// SetupRoutes sets up the routes for the API
func SetupRoutes(e *echo.Echo) {
	difyToken := os.Getenv("DIFY_API_KEY")
	notionToken := os.Getenv("NOTION_API_KEY")

	// サービス層
	difyService := service_dify.NewDifyService(difyToken)
	notionService := service_notion.NewNotionService(notionToken, &http.Client{})

	// ハンドラー層
	difyHandler := handler_dify.NewDifyHandler(difyService, notionService)
	notionHandler := handler_notion.NewNotionHandler(notionService)

	// Dify API
	e.POST("/dify", difyHandler.GetDifyDataHandler)
	e.POST("/dify/notion-create", difyHandler.GetDifyAndCreateNotionHandler)

	// Notion API
	e.GET("/notion", notionHandler.GetDatasFromNotion)
	e.POST("/notion/create", notionHandler.AddDataToNotion)
}
