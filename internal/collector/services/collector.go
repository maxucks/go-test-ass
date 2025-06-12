package services

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"test/internal/collector/config"
	"test/internal/collector/dto"

	"github.com/nats-io/nats.go"
)

type CollectorService struct {
	nc       *nats.Conn
	sub      *nats.Subscription
	pack     []*dto.GoodsEvent
	packSize int
	mux      sync.Mutex
}

func NewCollector(nc *nats.Conn, packSize int) *CollectorService {
	return &CollectorService{
		nc:       nc,
		sub:      nil,
		pack:     make([]*dto.GoodsEvent, 0, packSize),
		packSize: packSize,
		mux:      sync.Mutex{},
	}
}

func (s *CollectorService) Init(cfg *config.Config) error {
	sub, err := s.nc.Subscribe(cfg.NatsGoodsTopic, s.onGoodsEvent)
	s.sub = sub
	return err
}

func (s *CollectorService) Close() {
	s.sub.Drain()
}

func (s *CollectorService) onGoodsEvent(m *nats.Msg) {
	var event dto.GoodsEvent
	if err := json.Unmarshal(m.Data, &event); err != nil {
		log.Printf("unable to unmarshal event: %s\n", err)
	}

	s.mux.Lock()
	s.pack = append(s.pack, &event)
	s.mux.Unlock()

	if len(s.pack) == s.packSize {
		s.sendEvents()
	}
}

func (s *CollectorService) sendEvents() {
	for _, event := range s.pack {
		fmt.Printf("id=%d sent\n", event.ID)
	}

	s.mux.Lock()
	s.pack = make([]*dto.GoodsEvent, 0, s.packSize)
	s.mux.Unlock()
}
