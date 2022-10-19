package service

import (
	"context"
	"github.com/MenciusCheng/code-gen-service/conf"
	"github.com/MenciusCheng/code-gen-service/dao"
)

type Service struct {
	c *conf.Config

	// dao: db handler
	dao *dao.Dao
}

func New(c *conf.Config) *Service {
	return &Service{
		c:   c,
		dao: dao.New(c),
	}
}

// Ping check service's resource status
func (s *Service) Ping(ctx context.Context) error {
	return s.dao.Ping(ctx)
}

// Close close the resource
func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}
}
