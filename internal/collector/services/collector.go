package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"test/internal/collector/config"
	"test/internal/collector/dto"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/nats-io/nats.go"
)

const insertQuery = `
	INSERT INTO events (Id, ProjectId, Name, Description, Priority, Removed, EventTime)
`

type CollectorService struct {
	clickhouse driver.Conn
	nc         *nats.Conn
	sub        *nats.Subscription
	pack       []*dto.GoodsEvent
	packSize   int
	mux        sync.Mutex
}

func NewCollector(clickhouse driver.Conn, nc *nats.Conn, packSize int) *CollectorService {
	return &CollectorService{
		clickhouse: clickhouse,
		nc:         nc,
		sub:        nil,
		pack:       make([]*dto.GoodsEvent, 0, packSize),
		packSize:   packSize,
		mux:        sync.Mutex{},
	}
}

func (s *CollectorService) Run(ctx context.Context, cfg *config.Config) error {
	sub, err := s.nc.Subscribe(cfg.NatsGoodsTopic, func(msg *nats.Msg) {
		var event dto.GoodsEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("unable to unmarshal event: %s\n", err)
		}
		s.onGoodsEvent(ctx, &event)
	})
	s.sub = sub
	return err
}

func (s *CollectorService) Close() {
	s.sub.Drain()
}

func (s *CollectorService) onGoodsEvent(ctx context.Context, event *dto.GoodsEvent) {
	s.mux.Lock()
	s.pack = append(s.pack, event)
	s.mux.Unlock()

	if len(s.pack) == s.packSize {
		if err := s.sendEvents(ctx); err != nil {
			log.Printf("failed to send event: %s\n", err)
		}
	}
}

func (s *CollectorService) sendEvents(ctx context.Context) error {
	batch, err := s.clickhouse.PrepareBatch(ctx, insertQuery)
	if err != nil {
		log.Printf("failed to prepare batch (%s)", err)
	}

	for _, event := range s.pack {
		eventTime, err := time.Parse(time.RFC3339, event.EventTime)
		if err != nil {
			return fmt.Errorf("failed to convert event time string (%s)", err)
		}

		err = batch.Append(
			event.ID,
			event.ProjectID,
			event.Name,
			event.Description,
			event.Priority,
			event.Removed,
			eventTime,
		)
		if err != nil {
			return fmt.Errorf("failed to append batch (%s)", err)
		}
	}

	if err := batch.Send(); err != nil {
		return fmt.Errorf("failed to send batch (%s)", err)
	}

	s.mux.Lock()
	s.pack = make([]*dto.GoodsEvent, 0, s.packSize)
	s.mux.Unlock()

	return nil
}
