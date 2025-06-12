package managers

import (
	"encoding/json"
	"log"
	"test/internal/app/models"
	"time"

	"github.com/nats-io/nats.go"
)

type goodsEvent struct {
	Id          int    `json:"id"`
	ProjectId   int    `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Removed     bool   `json:"removed"`
	EventTime   string `json:"event_time"`
}

type Publisher struct {
	nc    *nats.Conn
	topic string
}

func NewPublisher(nc *nats.Conn, topic string) *Publisher {
	return &Publisher{nc, topic}
}

func (p *Publisher) PublishGoods(goods *models.Goods) {
	if goods == nil {
		return
	}

	event := goodsEvent{
		Id:          goods.Id,
		ProjectId:   goods.ProjectId,
		Name:        goods.Name,
		Description: goods.Description,
		Priority:    goods.Priority,
		Removed:     goods.Removed,
		EventTime:   time.Now().Format(time.RFC3339),
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Printf("failed to marshal event: %s\n", err)
	}

	p.nc.Publish(p.topic, eventJSON)
}
