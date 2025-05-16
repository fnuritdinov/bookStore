package configs

import "bookStore/pkg/config"

type Service struct {
	Port int
}

func NewService(cfg config.Config) Service {
	return Service{
		Port: cfg.GetInt("service.port"),
	}
}
