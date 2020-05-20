package service

import (
	"github.com/lhlyu/nrand/controller/dto"
	"github.com/lhlyu/nrand/service/gen"
	"github.com/lhlyu/nrand/trace"
)

type IndexService struct {
	trace.BaseTracker
}

func NewIndexService(tracker trace.ITracker) *IndexService {
	return &IndexService{
		BaseTracker: trace.NewBaseTracker(tracker),
	}
}

func (s *IndexService) GenName(req *dto.NameDto) []string {
	g := gen.GenFactory(req.Target)
	if g == nil {
		return nil
	}
	return g.Handler(req)
}
