// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	flowcontrol "k8s.io/client-go/util/flowcontrol"

	rest "k8s.io/client-go/rest"

	schema "k8s.io/apimachinery/pkg/runtime/schema"

	types "k8s.io/apimachinery/pkg/types"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// APIVersion provides a mock function with given fields:
func (_m *Interface) APIVersion() schema.GroupVersion {
	ret := _m.Called()

	var r0 schema.GroupVersion
	if rf, ok := ret.Get(0).(func() schema.GroupVersion); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(schema.GroupVersion)
	}

	return r0
}

// Delete provides a mock function with given fields:
func (_m *Interface) Delete() *rest.Request {
	ret := _m.Called()

	var r0 *rest.Request
	if rf, ok := ret.Get(0).(func() *rest.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Request)
		}
	}

	return r0
}

// Get provides a mock function with given fields:
func (_m *Interface) Get() *rest.Request {
	ret := _m.Called()

	var r0 *rest.Request
	if rf, ok := ret.Get(0).(func() *rest.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Request)
		}
	}

	return r0
}

// GetRateLimiter provides a mock function with given fields:
func (_m *Interface) GetRateLimiter() flowcontrol.RateLimiter {
	ret := _m.Called()

	var r0 flowcontrol.RateLimiter
	if rf, ok := ret.Get(0).(func() flowcontrol.RateLimiter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flowcontrol.RateLimiter)
		}
	}

	return r0
}

// Patch provides a mock function with given fields: pt
func (_m *Interface) Patch(pt types.PatchType) *rest.Request {
	ret := _m.Called(pt)

	var r0 *rest.Request
	if rf, ok := ret.Get(0).(func(types.PatchType) *rest.Request); ok {
		r0 = rf(pt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Request)
		}
	}

	return r0
}

// Post provides a mock function with given fields:
func (_m *Interface) Post() *rest.Request {
	ret := _m.Called()

	var r0 *rest.Request
	if rf, ok := ret.Get(0).(func() *rest.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Request)
		}
	}

	return r0
}

// Put provides a mock function with given fields:
func (_m *Interface) Put() *rest.Request {
	ret := _m.Called()

	var r0 *rest.Request
	if rf, ok := ret.Get(0).(func() *rest.Request); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Request)
		}
	}

	return r0
}

// Verb provides a mock function with given fields: verb
func (_m *Interface) Verb(verb string) *rest.Request {
	ret := _m.Called(verb)

	var r0 *rest.Request
	if rf, ok := ret.Get(0).(func(string) *rest.Request); ok {
		r0 = rf(verb)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.Request)
		}
	}

	return r0
}
