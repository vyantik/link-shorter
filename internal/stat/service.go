package stat

import (
	"app/test/pkg/event"
	"log"
)

type StatService struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatServiceDeps struct {
	StatService *StatService
}

func NewStatService(deps *StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.StatService.EventBus,
		StatRepository: deps.StatService.StatRepository,
	}
}

func (s *StatService) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.LinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Printf("[Stat] - [Service] - [ERROR] invalid event data: %v", msg.Data)
				continue
			}
			s.StatRepository.AddClick(id)
		}
	}
}
