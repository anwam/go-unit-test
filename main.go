package main

import (
	"net/http"

	"github.com/anwam/go-unit-test/internal/student"
	"github.com/anwam/go-unit-test/lib"
)

func main() {
	logger := lib.NewWrapLogger()
	go func() {
		logger.Fatal(http.ListenAndServe("localhost:6060", nil).Error())
	}()

	studentServ := student.NewStudentService(lib.NewHttpClient(http.DefaultClient), logger, lib.NewTimer())
	if std, err := studentServ.GetStudent(1); err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info(std.Name)
	}
}
