package models

type Goods struct {
	Id          int    `json:"id"`
	ProjectId   int    `json:"projectId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Removed     bool   `json:"removed"`
	CreatedAt   string `json:"createdAt"`
}

type ShortGoods struct {
	Id        int  `json:"id"`
	ProjectId int  `json:"projectId"`
	Removed   bool `json:"removed"`
}

type ReprioritizedGoods struct {
	Id       int `json:"id"`
	Priority int `json:"priority"`
}

type PaginationMeta struct {
	Total   int `json:"total"`
	Removed int `json:"removed"`
}
