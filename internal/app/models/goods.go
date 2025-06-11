package models

type Goods struct {
	Id          int    `json:"id"`
	ProjectId   string `json:"projectId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Removed     bool   `json:"removed"`
	CreatedAt   string `json:"createdAt"`
}

type ReprioritizedGoods struct {
	Id       int `json:"id"`
	Priority int `json:"priority"`
}

type PaginationMeta struct {
	Total   int `json:"total"`
	Removed int `json:"removed"`
	Limit   int `json:"limit"`
	Offset  int `json:"offset"`
}
