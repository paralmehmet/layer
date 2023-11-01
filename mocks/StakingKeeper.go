// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	math "cosmossdk.io/math"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper is an autogenerated mock type for the StakingKeeper type
type StakingKeeper struct {
	mock.Mock
}

// Delegate provides a mock function with given fields: ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount
func (_m *StakingKeeper) Delegate(ctx types.Context, delAddr types.AccAddress, bondAmt math.Int, tokenSrc stakingtypes.BondStatus, validator stakingtypes.Validator, subtractAccount bool) (math.LegacyDec, error) {
	ret := _m.Called(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)

	var r0 math.LegacyDec
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, math.Int, stakingtypes.BondStatus, stakingtypes.Validator, bool) (math.LegacyDec, error)); ok {
		return rf(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, math.Int, stakingtypes.BondStatus, stakingtypes.Validator, bool) math.LegacyDec); ok {
		r0 = rf(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.AccAddress, math.Int, stakingtypes.BondStatus, stakingtypes.Validator, bool) error); ok {
		r1 = rf(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteValidatorByPowerIndex provides a mock function with given fields: ctx, validator
func (_m *StakingKeeper) DeleteValidatorByPowerIndex(ctx types.Context, validator stakingtypes.Validator) {
	_m.Called(ctx, validator)
}

// GetAllDelegatorDelegations provides a mock function with given fields: ctx, delegator
func (_m *StakingKeeper) GetAllDelegatorDelegations(ctx types.Context, delegator types.AccAddress) []stakingtypes.Delegation {
	ret := _m.Called(ctx, delegator)

	var r0 []stakingtypes.Delegation
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress) []stakingtypes.Delegation); ok {
		r0 = rf(ctx, delegator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]stakingtypes.Delegation)
		}
	}

	return r0
}

// GetDelegation provides a mock function with given fields: ctx, delAddr, valAddr
func (_m *StakingKeeper) GetDelegation(ctx types.Context, delAddr types.AccAddress, valAddr types.ValAddress) (stakingtypes.Delegation, bool) {
	ret := _m.Called(ctx, delAddr, valAddr)

	var r0 stakingtypes.Delegation
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, types.ValAddress) (stakingtypes.Delegation, bool)); ok {
		return rf(ctx, delAddr, valAddr)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, types.ValAddress) stakingtypes.Delegation); ok {
		r0 = rf(ctx, delAddr, valAddr)
	} else {
		r0 = ret.Get(0).(stakingtypes.Delegation)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.AccAddress, types.ValAddress) bool); ok {
		r1 = rf(ctx, delAddr, valAddr)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetLastTotalPower provides a mock function with given fields: ctx
func (_m *StakingKeeper) GetLastTotalPower(ctx types.Context) math.Int {
	ret := _m.Called(ctx)

	var r0 math.Int
	if rf, ok := ret.Get(0).(func(types.Context) math.Int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	return r0
}

// GetValidator provides a mock function with given fields: ctx, addr
func (_m *StakingKeeper) GetValidator(ctx types.Context, addr types.ValAddress) (stakingtypes.Validator, bool) {
	ret := _m.Called(ctx, addr)

	var r0 stakingtypes.Validator
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress) (stakingtypes.Validator, bool)); ok {
		return rf(ctx, addr)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress) stakingtypes.Validator); ok {
		r0 = rf(ctx, addr)
	} else {
		r0 = ret.Get(0).(stakingtypes.Validator)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.ValAddress) bool); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetValidatorByConsAddr provides a mock function with given fields: ctx, consAddr
func (_m *StakingKeeper) GetValidatorByConsAddr(ctx types.Context, consAddr types.ConsAddress) (stakingtypes.Validator, bool) {
	ret := _m.Called(ctx, consAddr)

	var r0 stakingtypes.Validator
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, types.ConsAddress) (stakingtypes.Validator, bool)); ok {
		return rf(ctx, consAddr)
	}
	if rf, ok := ret.Get(0).(func(types.Context, types.ConsAddress) stakingtypes.Validator); ok {
		r0 = rf(ctx, consAddr)
	} else {
		r0 = ret.Get(0).(stakingtypes.Validator)
	}

	if rf, ok := ret.Get(1).(func(types.Context, types.ConsAddress) bool); ok {
		r1 = rf(ctx, consAddr)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// IterateDelegatorDelegations provides a mock function with given fields: ctx, delegator, cb
func (_m *StakingKeeper) IterateDelegatorDelegations(ctx types.Context, delegator types.AccAddress, cb func(stakingtypes.Delegation) bool) {
	_m.Called(ctx, delegator, cb)
}

// Jail provides a mock function with given fields: ctx, consAddr
func (_m *StakingKeeper) Jail(ctx types.Context, consAddr types.ConsAddress) {
	_m.Called(ctx, consAddr)
}

// RemoveDelegation provides a mock function with given fields: ctx, delegation
func (_m *StakingKeeper) RemoveDelegation(ctx types.Context, delegation stakingtypes.Delegation) error {
	ret := _m.Called(ctx, delegation)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, stakingtypes.Delegation) error); ok {
		r0 = rf(ctx, delegation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveValidatorTokens provides a mock function with given fields: ctx, validator, tokensToRemove
func (_m *StakingKeeper) RemoveValidatorTokens(ctx types.Context, validator stakingtypes.Validator, tokensToRemove math.Int) stakingtypes.Validator {
	ret := _m.Called(ctx, validator, tokensToRemove)

	var r0 stakingtypes.Validator
	if rf, ok := ret.Get(0).(func(types.Context, stakingtypes.Validator, math.Int) stakingtypes.Validator); ok {
		r0 = rf(ctx, validator, tokensToRemove)
	} else {
		r0 = ret.Get(0).(stakingtypes.Validator)
	}

	return r0
}

// RemoveValidatorTokensAndShares provides a mock function with given fields: ctx, validator, sharesToRemove
func (_m *StakingKeeper) RemoveValidatorTokensAndShares(ctx types.Context, validator stakingtypes.Validator, sharesToRemove math.LegacyDec) (stakingtypes.Validator, math.Int) {
	ret := _m.Called(ctx, validator, sharesToRemove)

	var r0 stakingtypes.Validator
	var r1 math.Int
	if rf, ok := ret.Get(0).(func(types.Context, stakingtypes.Validator, math.LegacyDec) (stakingtypes.Validator, math.Int)); ok {
		return rf(ctx, validator, sharesToRemove)
	}
	if rf, ok := ret.Get(0).(func(types.Context, stakingtypes.Validator, math.LegacyDec) stakingtypes.Validator); ok {
		r0 = rf(ctx, validator, sharesToRemove)
	} else {
		r0 = ret.Get(0).(stakingtypes.Validator)
	}

	if rf, ok := ret.Get(1).(func(types.Context, stakingtypes.Validator, math.LegacyDec) math.Int); ok {
		r1 = rf(ctx, validator, sharesToRemove)
	} else {
		r1 = ret.Get(1).(math.Int)
	}

	return r0, r1
}

// SetDelegation provides a mock function with given fields: ctx, delegation
func (_m *StakingKeeper) SetDelegation(ctx types.Context, delegation stakingtypes.Delegation) {
	_m.Called(ctx, delegation)
}

// SetValidator provides a mock function with given fields: ctx, validator
func (_m *StakingKeeper) SetValidator(ctx types.Context, validator stakingtypes.Validator) {
	_m.Called(ctx, validator)
}

// SetValidatorByPowerIndex provides a mock function with given fields: ctx, validator
func (_m *StakingKeeper) SetValidatorByPowerIndex(ctx types.Context, validator stakingtypes.Validator) {
	_m.Called(ctx, validator)
}

// TokensFromConsensusPower provides a mock function with given fields: ctx, power
func (_m *StakingKeeper) TokensFromConsensusPower(ctx types.Context, power int64) math.Int {
	ret := _m.Called(ctx, power)

	var r0 math.Int
	if rf, ok := ret.Get(0).(func(types.Context, int64) math.Int); ok {
		r0 = rf(ctx, power)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	return r0
}

// Validator provides a mock function with given fields: ctx, address
func (_m *StakingKeeper) Validator(ctx types.Context, address types.ValAddress) stakingtypes.ValidatorI {
	ret := _m.Called(ctx, address)

	var r0 stakingtypes.ValidatorI
	if rf, ok := ret.Get(0).(func(types.Context, types.ValAddress) stakingtypes.ValidatorI); ok {
		r0 = rf(ctx, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stakingtypes.ValidatorI)
		}
	}

	return r0
}

// NewStakingKeeper creates a new instance of StakingKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStakingKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *StakingKeeper {
	mock := &StakingKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}