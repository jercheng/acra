// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Encryptor is an autogenerated mock type for the Encryptor type
type Encryptor struct {
	mock.Mock
}

// Decrypt provides a mock function with given fields: ctx, keyID, data, _a3
func (_m *Encryptor) Decrypt(ctx context.Context, keyID []byte, data []byte, _a3 []byte) ([]byte, error) {
	ret := _m.Called(ctx, keyID, data, _a3)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, []byte, []byte, []byte) []byte); ok {
		r0 = rf(ctx, keyID, data, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte, []byte, []byte) error); ok {
		r1 = rf(ctx, keyID, data, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encrypt provides a mock function with given fields: ctx, keyID, data, _a3
func (_m *Encryptor) Encrypt(ctx context.Context, keyID []byte, data []byte, _a3 []byte) ([]byte, error) {
	ret := _m.Called(ctx, keyID, data, _a3)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, []byte, []byte, []byte) []byte); ok {
		r0 = rf(ctx, keyID, data, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte, []byte, []byte) error); ok {
		r1 = rf(ctx, keyID, data, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewEncryptorT interface {
	mock.TestingT
	Cleanup(func())
}

// NewEncryptor creates a new instance of Encryptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEncryptor(t NewEncryptorT) *Encryptor {
	mock := &Encryptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
