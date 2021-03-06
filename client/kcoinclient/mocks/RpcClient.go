// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "context"

import mock "github.com/stretchr/testify/mock"
import rpc "github.com/kowala-tech/kcoin/client/rpc"

// RpcClient is an autogenerated mock type for the RpcClient type
type RpcClient struct {
	mock.Mock
}

// BatchCall provides a mock function with given fields: b
func (_m *RpcClient) BatchCall(b []rpc.BatchElem) error {
	ret := _m.Called(b)

	var r0 error
	if rf, ok := ret.Get(0).(func([]rpc.BatchElem) error); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BatchCallContext provides a mock function with given fields: ctx, b
func (_m *RpcClient) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	ret := _m.Called(ctx, b)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []rpc.BatchElem) error); ok {
		r0 = rf(ctx, b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Call provides a mock function with given fields: result, method, args
func (_m *RpcClient) Call(result interface{}, method string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, result, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(result, method, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CallContext provides a mock function with given fields: ctx, result, method, args
func (_m *RpcClient) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, result, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, result, method, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *RpcClient) Close() {
	_m.Called()
}

// KowalaSubscribe provides a mock function with given fields: ctx, channel, args
func (_m *RpcClient) KowalaSubscribe(ctx context.Context, channel interface{}, args ...interface{}) (*rpc.ClientSubscription, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, channel)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *rpc.ClientSubscription
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...interface{}) *rpc.ClientSubscription); ok {
		r0 = rf(ctx, channel, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.ClientSubscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...interface{}) error); ok {
		r1 = rf(ctx, channel, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: ctx, namespace, channel, args
func (_m *RpcClient) Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (*rpc.ClientSubscription, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, namespace, channel)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *rpc.ClientSubscription
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, ...interface{}) *rpc.ClientSubscription); ok {
		r0 = rf(ctx, namespace, channel, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.ClientSubscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}, ...interface{}) error); ok {
		r1 = rf(ctx, namespace, channel, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SupportedModules provides a mock function with given fields:
func (_m *RpcClient) SupportedModules() (map[string]string, error) {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
