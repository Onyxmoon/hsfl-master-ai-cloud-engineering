// Code generated by mockery v2.39.1. DO NOT EDIT.

package products

import (
	mock "github.com/stretchr/testify/mock"
	model "hsfl.de/group6/hsfl-master-ai-cloud-engineering/product-service/products/model"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

type MockRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepository) EXPECT() *MockRepository_Expecter {
	return &MockRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0
func (_m *MockRepository) Create(_a0 *model.Product) (*model.Product, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Product) (*model.Product, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*model.Product) *model.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Product) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 *model.Product
func (_e *MockRepository_Expecter) Create(_a0 interface{}) *MockRepository_Create_Call {
	return &MockRepository_Create_Call{Call: _e.mock.On("Create", _a0)}
}

func (_c *MockRepository_Create_Call) Run(run func(_a0 *model.Product)) *MockRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Product))
	})
	return _c
}

func (_c *MockRepository_Create_Call) Return(_a0 *model.Product, _a1 error) *MockRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_Create_Call) RunAndReturn(run func(*model.Product) (*model.Product, error)) *MockRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0
func (_m *MockRepository) Delete(_a0 *model.Product) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Product) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 *model.Product
func (_e *MockRepository_Expecter) Delete(_a0 interface{}) *MockRepository_Delete_Call {
	return &MockRepository_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *MockRepository_Delete_Call) Run(run func(_a0 *model.Product)) *MockRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Product))
	})
	return _c
}

func (_c *MockRepository_Delete_Call) Return(_a0 error) *MockRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Delete_Call) RunAndReturn(run func(*model.Product) error) *MockRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields:
func (_m *MockRepository) FindAll() ([]*model.Product, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []*model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.Product, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockRepository_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *MockRepository_Expecter) FindAll() *MockRepository_FindAll_Call {
	return &MockRepository_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *MockRepository_FindAll_Call) Run(run func()) *MockRepository_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRepository_FindAll_Call) Return(_a0 []*model.Product, _a1 error) *MockRepository_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindAll_Call) RunAndReturn(run func() ([]*model.Product, error)) *MockRepository_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindByEan provides a mock function with given fields: ean
func (_m *MockRepository) FindByEan(ean string) (*model.Product, error) {
	ret := _m.Called(ean)

	if len(ret) == 0 {
		panic("no return value specified for FindByEan")
	}

	var r0 *model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Product, error)); ok {
		return rf(ean)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Product); ok {
		r0 = rf(ean)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ean)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindByEan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByEan'
type MockRepository_FindByEan_Call struct {
	*mock.Call
}

// FindByEan is a helper method to define mock.On call
//   - ean string
func (_e *MockRepository_Expecter) FindByEan(ean interface{}) *MockRepository_FindByEan_Call {
	return &MockRepository_FindByEan_Call{Call: _e.mock.On("FindByEan", ean)}
}

func (_c *MockRepository_FindByEan_Call) Run(run func(ean string)) *MockRepository_FindByEan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRepository_FindByEan_Call) Return(_a0 *model.Product, _a1 error) *MockRepository_FindByEan_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindByEan_Call) RunAndReturn(run func(string) (*model.Product, error)) *MockRepository_FindByEan_Call {
	_c.Call.Return(run)
	return _c
}

// FindById provides a mock function with given fields: id
func (_m *MockRepository) FindById(id uint64) (*model.Product, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*model.Product, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) *model.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type MockRepository_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - id uint64
func (_e *MockRepository_Expecter) FindById(id interface{}) *MockRepository_FindById_Call {
	return &MockRepository_FindById_Call{Call: _e.mock.On("FindById", id)}
}

func (_c *MockRepository_FindById_Call) Run(run func(id uint64)) *MockRepository_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *MockRepository_FindById_Call) Return(_a0 *model.Product, _a1 error) *MockRepository_FindById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindById_Call) RunAndReturn(run func(uint64) (*model.Product, error)) *MockRepository_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0
func (_m *MockRepository) Update(_a0 *model.Product) (*model.Product, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *model.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Product) (*model.Product, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*model.Product) *model.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Product) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 *model.Product
func (_e *MockRepository_Expecter) Update(_a0 interface{}) *MockRepository_Update_Call {
	return &MockRepository_Update_Call{Call: _e.mock.On("Update", _a0)}
}

func (_c *MockRepository_Update_Call) Run(run func(_a0 *model.Product)) *MockRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Product))
	})
	return _c
}

func (_c *MockRepository_Update_Call) Return(_a0 *model.Product, _a1 error) *MockRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_Update_Call) RunAndReturn(run func(*model.Product) (*model.Product, error)) *MockRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
