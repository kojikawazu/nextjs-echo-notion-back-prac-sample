package model_notion

// NotionGetDatasResponse はNotionからのレスポンス全体を表す
type NotionGetDatasResponse struct {
	Object     string       `json:"object"`
	Results    []NotionPage `json:"results"`
	NextCursor string       `json:"next_cursor"`
	HasMore    bool         `json:"has_more"`
}

// NotionPage は個々のページデータを表す
type NotionPage struct {
	Object         string     `json:"object"`
	ID             string     `json:"id"`
	CreatedTime    string     `json:"created_time"`
	LastEditedTime string     `json:"last_edited_time"`
	Properties     Properties `json:"properties"`
}

// Properties はページのプロパティを表す
type Properties struct {
	ID          PropertyValue `json:"id"`
	NotionTitle PropertyValue `json:"notion-title"`
	Contents    PropertyValue `json:"contents"`
	Kind        PropertyValue `json:"kind"`
	CreatedAt   PropertyValue `json:"created_at"`
}

// PropertyValue はプロパティの値を表す
type PropertyValue struct {
	ID       string  `json:"id"`
	Type     string  `json:"type"`
	Title    []Text  `json:"title,omitempty"`
	RichText []Text  `json:"rich_text,omitempty"`
	Select   *Select `json:"select,omitempty"`
	Date     *Date   `json:"date,omitempty"`
}

// NotionResponse は必要なデータのみを含むレスポンス構造体
type ProcessedNotionResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Contents  string `json:"contents"`
	Kind      string `json:"kind"`
	CreatedAt string `json:"created_at"`
}

// NotionCreateRequest はNotionにページを追加するためのリクエストを表す
type NotionCreateRequest struct {
	Parent     Parent              `json:"parent"`
	Properties map[string]Property `json:"properties"`
}

// Parent はページの親を表す
type Parent struct {
	DatabaseID string `json:"database_id"`
}

// Property はページのプロパティを表す
type Property struct {
	Title    []Text  `json:"title,omitempty"`
	RichText []Text  `json:"rich_text,omitempty"`
	Select   *Select `json:"select,omitempty"`
	Date     *Date   `json:"date,omitempty"`
}

// Text はテキストを表す
type Text struct {
	Text Content `json:"text"`
}

// Content はコンテンツを表す
type Content struct {
	Content string `json:"content"`
}

// Select は選択肢を表す
type Select struct {
	Name string `json:"name"`
}

// Date は日付を表す
type Date struct {
	Start string `json:"start"`
}

// CreatePageRequest はNotionにページを追加するためのリクエストを表す
type CreatePageRequest struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Contents  string `json:"contents"`
	Kind      string `json:"kind"`
	CreatedAt string `json:"created_at"`
}
