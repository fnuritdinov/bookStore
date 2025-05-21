package configs

import "bookStore/pkg/config"

type Service struct {
	Port    int
	BaseURL string
	Stage   string
}

func NewService(cfg config.Config) Service {
	return Service{
		Port:    cfg.GetInt("service.port"),
		BaseURL: cfg.GetString("service.base_url"),
		Stage:   cfg.GetString("service.stage"),
	}
}
