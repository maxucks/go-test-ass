package dto

type GoodsEvent struct {
	ID          int    `json:"id"`
	ProjectID   int    `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Removed     bool   `json:"removed"`
	EventTime   string `json:"event_time"`
}
