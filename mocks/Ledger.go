// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	ids "github.com/luxdefi/node/ids"

	mock "github.com/stretchr/testify/mock"
)

// Ledger is an autogenerated mock type for the Ledger type
type Ledger struct {
	mock.Mock
}

// Address provides a mock function with given fields: displayHRP, addressIndex
func (_m *Ledger) Address(displayHRP string, addressIndex uint32) (ids.ShortID, error) {
	ret := _m.Called(displayHRP, addressIndex)

	var r0 ids.ShortID
	if rf, ok := ret.Get(0).(func(string, uint32) ids.ShortID); ok {
		r0 = rf(displayHRP, addressIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ShortID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint32) error); ok {
		r1 = rf(displayHRP, addressIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Addresses provides a mock function with given fields: _a0
func (_m *Ledger) Addresses(_a0 []uint32) ([]ids.ShortID, error) {
	ret := _m.Called(_a0)

	var r0 []ids.ShortID
	if rf, ok := ret.Get(0).(func([]uint32) []ids.ShortID); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ids.ShortID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint32) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Disconnect provides a mock function with given fields:
func (_m *Ledger) Disconnect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignHash provides a mock function with given fields: hash, addressIndices
func (_m *Ledger) SignHash(hash []byte, addressIndices []uint32) ([][]byte, error) {
	ret := _m.Called(hash, addressIndices)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func([]byte, []uint32) [][]byte); ok {
		r0 = rf(hash, addressIndices)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, []uint32) error); ok {
		r1 = rf(hash, addressIndices)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Version provides a mock function with given fields:
func (_m *Ledger) Version() (string, string, string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func() string); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(string)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func() error); ok {
		r3 = rf()
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

type mockConstructorTestingTNewLedger interface {
	mock.TestingT
	Cleanup(func())
}

// NewLedger creates a new instance of Ledger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLedger(t mockConstructorTestingTNewLedger) *Ledger {
	mock := &Ledger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
