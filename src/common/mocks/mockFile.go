package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockFileService struct {
	mock.Mock
}

func (service *MockFileService) ReadLines(path string) []string {
	args := service.Called(path)
	return args.Get(0).([]string)
}
