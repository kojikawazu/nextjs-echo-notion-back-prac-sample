package handler_dify

import (
	service_dify "github.com/kojikawazu/backend/service/dify"
	service_notion "github.com/kojikawazu/backend/service/notion"
)

// DifyHandler is the handler for the Dify service
type DifyHandler struct {
	DifyService   service_dify.DifyService
	NotionService service_notion.NotionService
}

// NewDifyHandler creates a new Dify handler
func NewDifyHandler(
	difyService service_dify.DifyService,
	notionService service_notion.NotionService,
) *DifyHandler {
	return &DifyHandler{
		DifyService:   difyService,
		NotionService: notionService,
	}
}
