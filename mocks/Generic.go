// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Generic is an autogenerated mock type for the Generic type
type Generic struct {
	mock.Mock
}

type Generic_Expecter struct {
	mock *mock.Mock
}

func (_m *Generic) EXPECT() *Generic_Expecter {
	return &Generic_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, key, value
func (_m *Generic) Get(ctx context.Context, key string, value interface{}) (bool, error) {
	ret := _m.Called(ctx, key, value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) bool); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, key, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Generic_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Generic_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - key string
//  - value interface{}
func (_e *Generic_Expecter) Get(ctx interface{}, key interface{}, value interface{}) *Generic_Get_Call {
	return &Generic_Get_Call{Call: _e.mock.On("Get", ctx, key, value)}
}

func (_c *Generic_Get_Call) Run(run func(ctx context.Context, key string, value interface{})) *Generic_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *Generic_Get_Call) Return(_a0 bool, _a1 error) *Generic_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Set provides a mock function with given fields: ctx, key, value, ttl
func (_m *Generic) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	ret := _m.Called(ctx, key, value, ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, value, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Generic_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type Generic_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//  - ctx context.Context
//  - key string
//  - value interface{}
//  - ttl time.Duration
func (_e *Generic_Expecter) Set(ctx interface{}, key interface{}, value interface{}, ttl interface{}) *Generic_Set_Call {
	return &Generic_Set_Call{Call: _e.mock.On("Set", ctx, key, value, ttl)}
}

func (_c *Generic_Set_Call) Run(run func(ctx context.Context, key string, value interface{}, ttl time.Duration)) *Generic_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), args[3].(time.Duration))
	})
	return _c
}

func (_c *Generic_Set_Call) Return(_a0 error) *Generic_Set_Call {
	_c.Call.Return(_a0)
	return _c
}
