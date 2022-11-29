package server

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

var server *Server

type Server struct {
	client *redis.Client
}

func New() (*Server, error) {
	if server == nil {
		config, err := ReadConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to read config: %s", err)
		}

		server = &Server{}

		addr := fmt.Sprintf("%s:%s", config.Host, config.Port)

		server.client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: "",
			DB:       0,
		})

		log.Printf("Connected to %s", addr)
	}

	err := server.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect and ping server: %s", err)
	}

	return server, nil
}

func (s *Server) Set(key string, val interface{}) error {
	err := s.client.Set(key, val, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set value: %s", err)
	}

	return nil
}

func (s *Server) Get(key string) (interface{}, error) {
	val, err := s.client.Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get value for key=%s : %s", key, err)
	}

	return val, nil
}

func (s *Server) Ping() error {
	pong, err := s.client.Ping().Result()
	if err != nil {
		return fmt.Errorf("failed to ping server: %s", err)
	}

	log.Printf("Ping response %s", pong)

	return nil
}

func (s *Server) Close() error {
	err := s.client.Close()
	if err != nil {
		return fmt.Errorf("failed to ping server: %s", err)
	}

	log.Println("Server closed")

	return nil
}
