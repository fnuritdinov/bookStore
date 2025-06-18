package configs

import (
	"bookStore/pkg/config"
	"fmt"
)

type Service struct {
	Port    int
	BaseURL string
	Stage   string
}

func (s *Service) NewService(cfg config.Config) {
	s.Port = cfg.GetInt("service.port")
	s.BaseURL = cfg.GetString("service.base_url")
	s.Stage = cfg.GetString("service.stage")
}

type DB struct {
	Name     string
	Schema   string
	Port     int
	User     string
	Password string
}

func (db *DB) URL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@localhost:%d/%s?sslmode=disable&search_path=%s",
		db.User, db.Password, db.Port, db.Name, db.Schema)
}
