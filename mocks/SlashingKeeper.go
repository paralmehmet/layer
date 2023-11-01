// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	time "time"

	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// SlashingKeeper is an autogenerated mock type for the SlashingKeeper type
type SlashingKeeper struct {
	mock.Mock
}

// GetValidatorSigningInfo provides a mock function with given fields: ctx, address
func (_m *SlashingKeeper) GetValidatorSigningInfo(ctx types.Context, address types.ConsAddress) (slashingtypes.ValidatorSigningInfo, bool) {
	ret := _m.Called(ctx, address)

	var r0 slashingtypes.ValidatorSigningInfo
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, types.ConsAddress) (slashingtypes.ValidatorSigningInfo, bool)); ok {
		return rf(ctx, address)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.ConsAddress) slashingtypes.ValidatorSigningInfo); ok {
		r0 = rf(ctx, address)
	} else {
		r0 = ret.Get(0).(slashingtypes.ValidatorSigningInfo)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.ConsAddress) bool); ok {
		r1 = rf(ctx, address)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// JailUntil provides a mock function with given fields: ctx, consAddr, jailTime
func (_m *SlashingKeeper) JailUntil(ctx types.Context, consAddr types.ConsAddress, jailTime time.Time) {
	_m.Called(ctx, consAddr, jailTime)
}

// SetValidatorSigningInfo provides a mock function with given fields: ctx, address, info
func (_m *SlashingKeeper) SetValidatorSigningInfo(ctx types.Context, address types.ConsAddress, info slashingtypes.ValidatorSigningInfo) {
	_m.Called(ctx, address, info)
}

// NewSlashingKeeper creates a new instance of SlashingKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSlashingKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *SlashingKeeper {
	mock := &SlashingKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}