// Code generated by mockery v2.39.1. DO NOT EDIT.

package user

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockController is an autogenerated mock type for the Controller type
type MockController struct {
	mock.Mock
}

type MockController_Expecter struct {
	mock *mock.Mock
}

func (_m *MockController) EXPECT() *MockController_Expecter {
	return &MockController_Expecter{mock: &_m.Mock}
}

// DeleteUser provides a mock function with given fields: _a0, _a1
func (_m *MockController) DeleteUser(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type MockController_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) DeleteUser(_a0 interface{}, _a1 interface{}) *MockController_DeleteUser_Call {
	return &MockController_DeleteUser_Call{Call: _e.mock.On("DeleteUser", _a0, _a1)}
}

func (_c *MockController_DeleteUser_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_DeleteUser_Call) Return() *MockController_DeleteUser_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_DeleteUser_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUser provides a mock function with given fields: _a0, _a1
func (_m *MockController) GetUser(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type MockController_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) GetUser(_a0 interface{}, _a1 interface{}) *MockController_GetUser_Call {
	return &MockController_GetUser_Call{Call: _e.mock.On("GetUser", _a0, _a1)}
}

func (_c *MockController_GetUser_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_GetUser_Call) Return() *MockController_GetUser_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_GetUser_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUsersByRole provides a mock function with given fields: _a0, _a1
func (_m *MockController) GetUsersByRole(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_GetUsersByRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUsersByRole'
type MockController_GetUsersByRole_Call struct {
	*mock.Call
}

// GetUsersByRole is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) GetUsersByRole(_a0 interface{}, _a1 interface{}) *MockController_GetUsersByRole_Call {
	return &MockController_GetUsersByRole_Call{Call: _e.mock.On("GetUsersByRole", _a0, _a1)}
}

func (_c *MockController_GetUsersByRole_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_GetUsersByRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_GetUsersByRole_Call) Return() *MockController_GetUsersByRole_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_GetUsersByRole_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_GetUsersByRole_Call {
	_c.Call.Return(run)
	return _c
}

// PutUser provides a mock function with given fields: _a0, _a1
func (_m *MockController) PutUser(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_PutUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutUser'
type MockController_PutUser_Call struct {
	*mock.Call
}

// PutUser is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) PutUser(_a0 interface{}, _a1 interface{}) *MockController_PutUser_Call {
	return &MockController_PutUser_Call{Call: _e.mock.On("PutUser", _a0, _a1)}
}

func (_c *MockController_PutUser_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_PutUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_PutUser_Call) Return() *MockController_PutUser_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_PutUser_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_PutUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockController creates a new instance of MockController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockController(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockController {
	mock := &MockController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
