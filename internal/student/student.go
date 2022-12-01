package student

import (
	"encoding/json"
	"fmt"

	"github.com/anwam/go-unit-test/lib"
	"go.uber.org/zap"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type StudentService struct {
	req    lib.Requester
	logger lib.Logger
}

func NewStudentService(req lib.Requester, logger lib.Logger) *StudentService {
	return &StudentService{
		req:    req,
		logger: logger,
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
	return &std, nil
}
