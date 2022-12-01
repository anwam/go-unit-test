package student

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/anwam/go-unit-test/lib"
	"github.com/anwam/go-unit-test/mocks"
	"go.uber.org/zap"
)

func TestNewStudentService(t *testing.T) {
	type args struct {
		req    lib.Requester
		logger lib.Logger
		timer  lib.Timer
	}
	tests := []struct {
		name string
		args args
		want *StudentService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStudentService(tt.args.req, tt.args.logger, tt.args.timer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStudentService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentService_GetStudent(t *testing.T) {
	type args struct {
		id int
	}
	req := mocks.NewRequester(t)
	logger := mocks.NewLogger(t)
	timer := mocks.NewTimer(t)
	stdSvc := NewStudentService(req, logger, timer)
	tests := []struct {
		name    string
		s       *StudentService
		args    args
		want    *int
		wantErr bool
	}{
		{
			name: "Get Student",
			s:    stdSvc,
			args: args{
				id: 2,
			},
			want: func() *int { i := 2; return &i }(),
		},
		{
			name: "Get Student",
			s:    stdSvc,
			args: args{
				id: 5,
			},
			want: func() *int { i := 5; return &i }(),
		},
		{
			name: "Get Student",
			s:    stdSvc,
			args: args{
				id: 0,
			},
			want:    (*int)(nil),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.id == 0 {
				logger.EXPECT().Error("error", zap.String("url", fmt.Sprintf("http://localhost:8080/student/%d", tt.args.id)), zap.Int("id", tt.args.id), zap.String("method", "GetStudent"), zap.String("service", "StudentService")).Return()
				req.EXPECT().Get(fmt.Sprintf("http://localhost:8080/student/%d", tt.args.id)).Call.Return(nil, fmt.Errorf("error"))
			} else {
				req.EXPECT().Get(fmt.Sprintf("http://localhost:8080/student/%d", tt.args.id)).Call.Return(func(url string) *http.Response {
					return &http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewBufferString(fmt.Sprintf(`{"id": %d, "name": "John Doe"}`, tt.args.id))),
					}
				}, nil)
			}
			got, err := tt.s.GetStudent(tt.args.id)
			if err != nil {
				if tt.wantErr {
					return
				} else {
					t.Errorf("StudentService.GetStudent() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}

			if !reflect.DeepEqual(got.ID, *tt.want) {
				t.Errorf("StudentService.GetStudent() = %v, want %v", got.ID, *tt.want)
			}
		})
	}
}
