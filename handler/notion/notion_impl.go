package handler_notion

import service_notion "github.com/kojikawazu/backend/service/notion"

// NotionHandler is the handler for the Notion service
type NotionHandler struct {
	NotionService service_notion.NotionService
}

// NewNotionHandler creates a new Notion handler
func NewNotionHandler(notionService service_notion.NotionService) *NotionHandler {
	return &NotionHandler{NotionService: notionService}
}
