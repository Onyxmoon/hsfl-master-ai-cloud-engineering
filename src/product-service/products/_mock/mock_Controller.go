// Code generated by mockery v2.39.1. DO NOT EDIT.

package products

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

// DeleteProduct provides a mock function with given fields: _a0, _a1
func (_m *MockController) DeleteProduct(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_DeleteProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProduct'
type MockController_DeleteProduct_Call struct {
	*mock.Call
}

// DeleteProduct is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) DeleteProduct(_a0 interface{}, _a1 interface{}) *MockController_DeleteProduct_Call {
	return &MockController_DeleteProduct_Call{Call: _e.mock.On("DeleteProduct", _a0, _a1)}
}

func (_c *MockController_DeleteProduct_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_DeleteProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_DeleteProduct_Call) Return() *MockController_DeleteProduct_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_DeleteProduct_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_DeleteProduct_Call {
	_c.Call.Return(run)
	return _c
}

// GetProductByEan provides a mock function with given fields: _a0, _a1
func (_m *MockController) GetProductByEan(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_GetProductByEan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProductByEan'
type MockController_GetProductByEan_Call struct {
	*mock.Call
}

// GetProductByEan is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) GetProductByEan(_a0 interface{}, _a1 interface{}) *MockController_GetProductByEan_Call {
	return &MockController_GetProductByEan_Call{Call: _e.mock.On("GetProductByEan", _a0, _a1)}
}

func (_c *MockController_GetProductByEan_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_GetProductByEan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_GetProductByEan_Call) Return() *MockController_GetProductByEan_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_GetProductByEan_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_GetProductByEan_Call {
	_c.Call.Return(run)
	return _c
}

// GetProductById provides a mock function with given fields: _a0, _a1
func (_m *MockController) GetProductById(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_GetProductById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProductById'
type MockController_GetProductById_Call struct {
	*mock.Call
}

// GetProductById is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) GetProductById(_a0 interface{}, _a1 interface{}) *MockController_GetProductById_Call {
	return &MockController_GetProductById_Call{Call: _e.mock.On("GetProductById", _a0, _a1)}
}

func (_c *MockController_GetProductById_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_GetProductById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_GetProductById_Call) Return() *MockController_GetProductById_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_GetProductById_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_GetProductById_Call {
	_c.Call.Return(run)
	return _c
}

// GetProducts provides a mock function with given fields: _a0, _a1
func (_m *MockController) GetProducts(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_GetProducts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProducts'
type MockController_GetProducts_Call struct {
	*mock.Call
}

// GetProducts is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) GetProducts(_a0 interface{}, _a1 interface{}) *MockController_GetProducts_Call {
	return &MockController_GetProducts_Call{Call: _e.mock.On("GetProducts", _a0, _a1)}
}

func (_c *MockController_GetProducts_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_GetProducts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_GetProducts_Call) Return() *MockController_GetProducts_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_GetProducts_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_GetProducts_Call {
	_c.Call.Return(run)
	return _c
}

// PostProduct provides a mock function with given fields: _a0, _a1
func (_m *MockController) PostProduct(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_PostProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostProduct'
type MockController_PostProduct_Call struct {
	*mock.Call
}

// PostProduct is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) PostProduct(_a0 interface{}, _a1 interface{}) *MockController_PostProduct_Call {
	return &MockController_PostProduct_Call{Call: _e.mock.On("PostProduct", _a0, _a1)}
}

func (_c *MockController_PostProduct_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_PostProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_PostProduct_Call) Return() *MockController_PostProduct_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_PostProduct_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_PostProduct_Call {
	_c.Call.Return(run)
	return _c
}

// PutProduct provides a mock function with given fields: _a0, _a1
func (_m *MockController) PutProduct(_a0 http.ResponseWriter, _a1 *http.Request) {
	_m.Called(_a0, _a1)
}

// MockController_PutProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutProduct'
type MockController_PutProduct_Call struct {
	*mock.Call
}

// PutProduct is a helper method to define mock.On call
//   - _a0 http.ResponseWriter
//   - _a1 *http.Request
func (_e *MockController_Expecter) PutProduct(_a0 interface{}, _a1 interface{}) *MockController_PutProduct_Call {
	return &MockController_PutProduct_Call{Call: _e.mock.On("PutProduct", _a0, _a1)}
}

func (_c *MockController_PutProduct_Call) Run(run func(_a0 http.ResponseWriter, _a1 *http.Request)) *MockController_PutProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockController_PutProduct_Call) Return() *MockController_PutProduct_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockController_PutProduct_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockController_PutProduct_Call {
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