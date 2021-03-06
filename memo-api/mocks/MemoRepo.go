// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	entity "github.com/FelixAnna/web-service-dlw/memo-api/memo/entity"
	mock "github.com/stretchr/testify/mock"
)

// MemoRepo is an autogenerated mock type for the MemoRepo type
type MemoRepo struct {
	mock.Mock
}

type MemoRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *MemoRepo) EXPECT() *MemoRepo_Expecter {
	return &MemoRepo_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: memo
func (_m *MemoRepo) Add(memo *entity.Memo) (string, error) {
	ret := _m.Called(memo)

	var r0 string
	if rf, ok := ret.Get(0).(func(*entity.Memo) string); ok {
		r0 = rf(memo)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Memo) error); ok {
		r1 = rf(memo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MemoRepo_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type MemoRepo_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//  - memo *entity.Memo
func (_e *MemoRepo_Expecter) Add(memo interface{}) *MemoRepo_Add_Call {
	return &MemoRepo_Add_Call{Call: _e.mock.On("Add", memo)}
}

func (_c *MemoRepo_Add_Call) Run(run func(memo *entity.Memo)) *MemoRepo_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Memo))
	})
	return _c
}

func (_c *MemoRepo_Add_Call) Return(_a0 string, _a1 error) *MemoRepo_Add_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *MemoRepo) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MemoRepo_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MemoRepo_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - id string
func (_e *MemoRepo_Expecter) Delete(id interface{}) *MemoRepo_Delete_Call {
	return &MemoRepo_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *MemoRepo_Delete_Call) Run(run func(id string)) *MemoRepo_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MemoRepo_Delete_Call) Return(_a0 error) *MemoRepo_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetByDateRange provides a mock function with given fields: start, end, userId
func (_m *MemoRepo) GetByDateRange(start string, end string, userId string) ([]entity.Memo, error) {
	ret := _m.Called(start, end, userId)

	var r0 []entity.Memo
	if rf, ok := ret.Get(0).(func(string, string, string) []entity.Memo); ok {
		r0 = rf(start, end, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Memo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(start, end, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MemoRepo_GetByDateRange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByDateRange'
type MemoRepo_GetByDateRange_Call struct {
	*mock.Call
}

// GetByDateRange is a helper method to define mock.On call
//  - start string
//  - end string
//  - userId string
func (_e *MemoRepo_Expecter) GetByDateRange(start interface{}, end interface{}, userId interface{}) *MemoRepo_GetByDateRange_Call {
	return &MemoRepo_GetByDateRange_Call{Call: _e.mock.On("GetByDateRange", start, end, userId)}
}

func (_c *MemoRepo_GetByDateRange_Call) Run(run func(start string, end string, userId string)) *MemoRepo_GetByDateRange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MemoRepo_GetByDateRange_Call) Return(_a0 []entity.Memo, _a1 error) *MemoRepo_GetByDateRange_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetById provides a mock function with given fields: id
func (_m *MemoRepo) GetById(id string) (*entity.Memo, error) {
	ret := _m.Called(id)

	var r0 *entity.Memo
	if rf, ok := ret.Get(0).(func(string) *entity.Memo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Memo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MemoRepo_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MemoRepo_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//  - id string
func (_e *MemoRepo_Expecter) GetById(id interface{}) *MemoRepo_GetById_Call {
	return &MemoRepo_GetById_Call{Call: _e.mock.On("GetById", id)}
}

func (_c *MemoRepo_GetById_Call) Run(run func(id string)) *MemoRepo_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MemoRepo_GetById_Call) Return(_a0 *entity.Memo, _a1 error) *MemoRepo_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByUserId provides a mock function with given fields: userId
func (_m *MemoRepo) GetByUserId(userId string) ([]entity.Memo, error) {
	ret := _m.Called(userId)

	var r0 []entity.Memo
	if rf, ok := ret.Get(0).(func(string) []entity.Memo); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Memo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MemoRepo_GetByUserId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUserId'
type MemoRepo_GetByUserId_Call struct {
	*mock.Call
}

// GetByUserId is a helper method to define mock.On call
//  - userId string
func (_e *MemoRepo_Expecter) GetByUserId(userId interface{}) *MemoRepo_GetByUserId_Call {
	return &MemoRepo_GetByUserId_Call{Call: _e.mock.On("GetByUserId", userId)}
}

func (_c *MemoRepo_GetByUserId_Call) Run(run func(userId string)) *MemoRepo_GetByUserId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MemoRepo_GetByUserId_Call) Return(_a0 []entity.Memo, _a1 error) *MemoRepo_GetByUserId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Update provides a mock function with given fields: memo
func (_m *MemoRepo) Update(memo entity.Memo) error {
	ret := _m.Called(memo)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Memo) error); ok {
		r0 = rf(memo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MemoRepo_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MemoRepo_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//  - memo entity.Memo
func (_e *MemoRepo_Expecter) Update(memo interface{}) *MemoRepo_Update_Call {
	return &MemoRepo_Update_Call{Call: _e.mock.On("Update", memo)}
}

func (_c *MemoRepo_Update_Call) Run(run func(memo entity.Memo)) *MemoRepo_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entity.Memo))
	})
	return _c
}

func (_c *MemoRepo_Update_Call) Return(_a0 error) *MemoRepo_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMemoRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewMemoRepo creates a new instance of MemoRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMemoRepo(t mockConstructorTestingTNewMemoRepo) *MemoRepo {
	mock := &MemoRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
