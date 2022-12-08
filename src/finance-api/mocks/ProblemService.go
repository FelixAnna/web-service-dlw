// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	entity "github.com/FelixAnna/web-service-dlw/finance-api/mathematicals/problem/entity"
	mock "github.com/stretchr/testify/mock"
)

// ProblemService is an autogenerated mock type for the ProblemService type
type ProblemService struct {
	mock.Mock
}

type ProblemService_Expecter struct {
	mock *mock.Mock
}

func (_m *ProblemService) EXPECT() *ProblemService_Expecter {
	return &ProblemService_Expecter{mock: &_m.Mock}
}

// GenerateProblem provides a mock function with given fields: criteria
func (_m *ProblemService) GenerateProblem(criteria ...interface{}) *entity.Problem {
	var _ca []interface{}
	_ca = append(_ca, criteria...)
	ret := _m.Called(_ca...)

	var r0 *entity.Problem
	if rf, ok := ret.Get(0).(func(...interface{}) *entity.Problem); ok {
		r0 = rf(criteria...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Problem)
		}
	}

	return r0
}

// ProblemService_GenerateProblem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateProblem'
type ProblemService_GenerateProblem_Call struct {
	*mock.Call
}

// GenerateProblem is a helper method to define mock.On call
//  - criteria ...interface{}
func (_e *ProblemService_Expecter) GenerateProblem(criteria ...interface{}) *ProblemService_GenerateProblem_Call {
	return &ProblemService_GenerateProblem_Call{Call: _e.mock.On("GenerateProblem",
		append([]interface{}{}, criteria...)...)}
}

func (_c *ProblemService_GenerateProblem_Call) Run(run func(criteria ...interface{})) *ProblemService_GenerateProblem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *ProblemService_GenerateProblem_Call) Return(_a0 *entity.Problem) *ProblemService_GenerateProblem_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewProblemService interface {
	mock.TestingT
	Cleanup(func())
}

// NewProblemService creates a new instance of ProblemService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProblemService(t mockConstructorTestingTNewProblemService) *ProblemService {
	mock := &ProblemService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}