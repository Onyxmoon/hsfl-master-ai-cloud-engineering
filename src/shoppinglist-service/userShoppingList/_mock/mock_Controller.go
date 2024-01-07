// Code generated by mockery v2.39.1. DO NOT EDIT.

package userShoppingList

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

// DeleteList provides a mock function with given fields: _a0, _a1
func (_m *MockController) DeleteList(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_DeleteList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteList'
type MockController_DeleteList_Call struct {
	*mock.Call
}

// DeleteList is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) DeleteList(_a0 interface{}, _a1 interface{}) *MockController_DeleteList_Call {
	return &MockController_DeleteList_Call{Call: _e.mock.On("DeleteList", _a0, _a1)}
}

func (_c *MockController_DeleteList_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_DeleteList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_DeleteList_Call) Return() *MockController_DeleteList_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_DeleteList_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_DeleteList_Call {
	_c.Call.Return(run)
	return _c
}

// GetList provides a mock function with given fields: _a0, _a1
func (_m *MockController) GetList(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_GetList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetList'
type MockController_GetList_Call struct {
	*mock.Call
}

// GetList is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) GetList(_a0 interface{}, _a1 interface{}) *MockController_GetList_Call {
	return &MockController_GetList_Call{Call: _e.mock.On("GetList", _a0, _a1)}
}

func (_c *MockController_GetList_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_GetList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_GetList_Call) Return() *MockController_GetList_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_GetList_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_GetList_Call {
	_c.Call.Return(run)
	return _c
}

// GetLists provides a mock function with given fields: _a0, _a1
func (_m *MockController) GetLists(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_GetLists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLists'
type MockController_GetLists_Call struct {
	*mock.Call
}

// GetLists is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) GetLists(_a0 interface{}, _a1 interface{}) *MockController_GetLists_Call {
	return &MockController_GetLists_Call{Call: _e.mock.On("GetLists", _a0, _a1)}
}

func (_c *MockController_GetLists_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_GetLists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_GetLists_Call) Return() *MockController_GetLists_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_GetLists_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_GetLists_Call {
	_c.Call.Return(run)
	return _c
}

// PostList provides a mock function with given fields: _a0, _a1
func (_m *MockController) PostList(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_PostList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostList'
type MockController_PostList_Call struct {
	*mock.Call
}

// PostList is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) PostList(_a0 interface{}, _a1 interface{}) *MockController_PostList_Call {
	return &MockController_PostList_Call{Call: _e.mock.On("PostList", _a0, _a1)}
}

func (_c *MockController_PostList_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_PostList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_PostList_Call) Return() *MockController_PostList_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_PostList_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_PostList_Call {
	_c.Call.Return(run)
	return _c
}

// PutList provides a mock function with given fields: _a0, _a1
func (_m *MockController) PutList(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_PutList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutList'
type MockController_PutList_Call struct {
	*mock.Call
}

// PutList is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) PutList(_a0 interface{}, _a1 interface{}) *MockController_PutList_Call {
	return &MockController_PutList_Call{Call: _e.mock.On("PutList", _a0, _a1)}
}

func (_c *MockController_PutList_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_PutList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_PutList_Call) Return() *MockController_PutList_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_PutList_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_PutList_Call {
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