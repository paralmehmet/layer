// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	math "cosmossdk.io/math"

	mock "github.com/stretchr/testify/mock"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper is an autogenerated mock type for the StakingKeeper type
type StakingKeeper struct {
	mock.Mock
}

// GetAllDelegatorDelegations provides a mock function with given fields: ctx, delegator
func (_m *StakingKeeper) GetAllDelegatorDelegations(ctx context.Context, delegator types.AccAddress) ([]stakingtypes.Delegation, error) {
	ret := _m.Called(ctx, delegator)

	var r0 []stakingtypes.Delegation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress) ([]stakingtypes.Delegation, error)); ok {
		return rf(ctx, delegator)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress) []stakingtypes.Delegation); ok {
		r0 = rf(ctx, delegator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]stakingtypes.Delegation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.AccAddress) error); ok {
		r1 = rf(ctx, delegator)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDelegation provides a mock function with given fields: ctx, delAddr, valAddr
func (_m *StakingKeeper) GetDelegation(ctx context.Context, delAddr types.AccAddress, valAddr types.ValAddress) (stakingtypes.Delegation, error) {
	ret := _m.Called(ctx, delAddr, valAddr)

	var r0 stakingtypes.Delegation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, types.ValAddress) (stakingtypes.Delegation, error)); ok {
		return rf(ctx, delAddr, valAddr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, types.ValAddress) stakingtypes.Delegation); ok {
		r0 = rf(ctx, delAddr, valAddr)
	} else {
		r0 = ret.Get(0).(stakingtypes.Delegation)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.AccAddress, types.ValAddress) error); ok {
		r1 = rf(ctx, delAddr, valAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastTotalPower provides a mock function with given fields: ctx
func (_m *StakingKeeper) GetLastTotalPower(ctx context.Context) (math.Int, error) {
	ret := _m.Called(ctx)

	var r0 math.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (math.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) math.Int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validator provides a mock function with given fields: ctx, address
func (_m *StakingKeeper) Validator(ctx context.Context, address types.ValAddress) (stakingtypes.ValidatorI, error) {
	ret := _m.Called(ctx, address)

	var r0 stakingtypes.ValidatorI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ValAddress) (stakingtypes.ValidatorI, error)); ok {
		return rf(ctx, address)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.ValAddress) stakingtypes.ValidatorI); ok {
		r0 = rf(ctx, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stakingtypes.ValidatorI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.ValAddress) error); ok {
		r1 = rf(ctx, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStakingKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewStakingKeeper creates a new instance of StakingKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStakingKeeper(t mockConstructorTestingTNewStakingKeeper) *StakingKeeper {
	mock := &StakingKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}