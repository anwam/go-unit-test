package student

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/anwam/go-unit-test/lib"
	"go.uber.org/zap"
)

type Student struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	RegisteredAt time.Time `json:"registered_at"`
}

type StudentService struct {
	req    lib.Requester
	logger lib.Logger
	timer  lib.Timer
}

func NewStudentService(req lib.Requester, logger lib.Logger, timer lib.Timer) *StudentService {
	return &StudentService{
		req:    req,
		logger: logger,
		timer:  timer,
	}
}

func (s *StudentService) GetStudent(id int) (*Student, error) {
	url := fmt.Sprintf("http://localhost:8080/student/%d", id)
	resp, err := s.req.Get(url)
	if err != nil {
		s.logger.Error(err.Error(), zap.String("url", url), zap.Int("id", id), zap.String("method", "GetStudent"), zap.String("service", "StudentService"))
		return nil, err
	}
	defer resp.Body.Close()
	std := Student{}
	if err := json.NewDecoder(resp.Body).Decode(&std); err != nil {
		s.logger.Error(err.Error(), zap.String("url", url), zap.Int("id", id), zap.String("method", "GetStudent"), zap.String("service", "StudentService"))
		return nil, err
	}
	std.RegisteredAt = s.timer.Now()
	return &std, nil
}
