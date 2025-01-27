// Code generated by mockery v2.39.1. DO NOT EDIT.

package user

import (
	mock "github.com/stretchr/testify/mock"
	model "hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/user/model"
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
func (_m *MockRepository) Create(_a0 *model.User) (*model.User, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.User) (*model.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
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
//   - _a0 *model.User
func (_e *MockRepository_Expecter) Create(_a0 interface{}) *MockRepository_Create_Call {
	return &MockRepository_Create_Call{Call: _e.mock.On("Create", _a0)}
}

func (_c *MockRepository_Create_Call) Run(run func(_a0 *model.User)) *MockRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.User))
	})
	return _c
}

func (_c *MockRepository_Create_Call) Return(_a0 *model.User, _a1 error) *MockRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_Create_Call) RunAndReturn(run func(*model.User) (*model.User, error)) *MockRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0
func (_m *MockRepository) Delete(_a0 *model.User) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
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
//   - _a0 *model.User
func (_e *MockRepository_Expecter) Delete(_a0 interface{}) *MockRepository_Delete_Call {
	return &MockRepository_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *MockRepository_Delete_Call) Run(run func(_a0 *model.User)) *MockRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.User))
	})
	return _c
}

func (_c *MockRepository_Delete_Call) Return(_a0 error) *MockRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepository_Delete_Call) RunAndReturn(run func(*model.User) error) *MockRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields:
func (_m *MockRepository) FindAll() ([]*model.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []*model.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*model.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
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

func (_c *MockRepository_FindAll_Call) Return(_a0 []*model.User, _a1 error) *MockRepository_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindAll_Call) RunAndReturn(run func() ([]*model.User, error)) *MockRepository_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindAllByRole provides a mock function with given fields: role
func (_m *MockRepository) FindAllByRole(role model.Role) ([]*model.User, error) {
	ret := _m.Called(role)

	if len(ret) == 0 {
		panic("no return value specified for FindAllByRole")
	}

	var r0 []*model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Role) ([]*model.User, error)); ok {
		return rf(role)
	}
	if rf, ok := ret.Get(0).(func(model.Role) []*model.User); ok {
		r0 = rf(role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(model.Role) error); ok {
		r1 = rf(role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindAllByRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllByRole'
type MockRepository_FindAllByRole_Call struct {
	*mock.Call
}

// FindAllByRole is a helper method to define mock.On call
//   - role model.Role
func (_e *MockRepository_Expecter) FindAllByRole(role interface{}) *MockRepository_FindAllByRole_Call {
	return &MockRepository_FindAllByRole_Call{Call: _e.mock.On("FindAllByRole", role)}
}

func (_c *MockRepository_FindAllByRole_Call) Run(run func(role model.Role)) *MockRepository_FindAllByRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.Role))
	})
	return _c
}

func (_c *MockRepository_FindAllByRole_Call) Return(_a0 []*model.User, _a1 error) *MockRepository_FindAllByRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindAllByRole_Call) RunAndReturn(run func(model.Role) ([]*model.User, error)) *MockRepository_FindAllByRole_Call {
	_c.Call.Return(run)
	return _c
}

// FindByEmail provides a mock function with given fields: email
func (_m *MockRepository) FindByEmail(email string) (*model.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for FindByEmail")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepository_FindByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByEmail'
type MockRepository_FindByEmail_Call struct {
	*mock.Call
}

// FindByEmail is a helper method to define mock.On call
//   - email string
func (_e *MockRepository_Expecter) FindByEmail(email interface{}) *MockRepository_FindByEmail_Call {
	return &MockRepository_FindByEmail_Call{Call: _e.mock.On("FindByEmail", email)}
}

func (_c *MockRepository_FindByEmail_Call) Run(run func(email string)) *MockRepository_FindByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRepository_FindByEmail_Call) Return(_a0 *model.User, _a1 error) *MockRepository_FindByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindByEmail_Call) RunAndReturn(run func(string) (*model.User, error)) *MockRepository_FindByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// FindById provides a mock function with given fields: id
func (_m *MockRepository) FindById(id uint64) (*model.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*model.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint64) *model.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
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

func (_c *MockRepository_FindById_Call) Return(_a0 *model.User, _a1 error) *MockRepository_FindById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_FindById_Call) RunAndReturn(run func(uint64) (*model.User, error)) *MockRepository_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0
func (_m *MockRepository) Update(_a0 *model.User) (*model.User, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.User) (*model.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
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
//   - _a0 *model.User
func (_e *MockRepository_Expecter) Update(_a0 interface{}) *MockRepository_Update_Call {
	return &MockRepository_Update_Call{Call: _e.mock.On("Update", _a0)}
}

func (_c *MockRepository_Update_Call) Run(run func(_a0 *model.User)) *MockRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.User))
	})
	return _c
}

func (_c *MockRepository_Update_Call) Return(_a0 *model.User, _a1 error) *MockRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepository_Update_Call) RunAndReturn(run func(*model.User) (*model.User, error)) *MockRepository_Update_Call {
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
