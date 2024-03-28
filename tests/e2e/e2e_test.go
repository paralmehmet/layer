package e2e_test

import (
	"encoding/hex"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tellor-io/layer/utils"
	minttypes "github.com/tellor-io/layer/x/mint/types"
	oraclekeeper "github.com/tellor-io/layer/x/oracle/keeper"
	oracletypes "github.com/tellor-io/layer/x/oracle/types"
	oracleutils "github.com/tellor-io/layer/x/oracle/utils"
	reporterkeeper "github.com/tellor-io/layer/x/reporter/keeper"
	reportertypes "github.com/tellor-io/layer/x/reporter/types"
)

func (s *E2ETestSuite) TestInitialMint() {
	require := s.Require()

	mintToTeamAcc := s.accountKeeper.GetModuleAddress(minttypes.MintToTeam)
	require.NotNil(mintToTeamAcc)
	balance := s.bankKeeper.GetBalance(s.ctx, mintToTeamAcc, s.denom)
	require.Equal(balance.Amount, math.NewInt(300*1e6))
}

func (s *E2ETestSuite) TestTransferAfterMint() {
	require := s.Require()

	mintToTeamAcc := s.accountKeeper.GetModuleAddress(minttypes.MintToTeam)
	require.NotNil(mintToTeamAcc)
	balance := s.bankKeeper.GetBalance(s.ctx, mintToTeamAcc, s.denom)
	require.Equal(balance.Amount, math.NewInt(300*1e6))

	// create 5 accounts
	type Accounts struct {
		PrivateKey secp256k1.PrivKey
		Account    sdk.AccAddress
	}
	accounts := make([]Accounts, 0, 5)
	for i := 0; i < 5; i++ {
		privKey := secp256k1.GenPrivKey()
		accountAddress := sdk.AccAddress(privKey.PubKey().Address())
		account := authtypes.BaseAccount{
			Address:       accountAddress.String(),
			PubKey:        codectypes.UnsafePackAny(privKey.PubKey()),
			AccountNumber: uint64(i + 1),
		}
		existingAccount := s.accountKeeper.GetAccount(s.ctx, accountAddress)
		if existingAccount == nil {
			s.accountKeeper.SetAccount(s.ctx, &account)
			accounts = append(accounts, Accounts{
				PrivateKey: *privKey,
				Account:    accountAddress,
			})
		}
	}

	// transfer 1000 tokens from team to all 5 accounts
	for _, acc := range accounts {
		startBalance := s.bankKeeper.GetBalance(s.ctx, acc.Account, s.denom).Amount
		err := s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, minttypes.MintToTeam, acc.Account, sdk.NewCoins(sdk.NewCoin(s.denom, math.NewInt(1000))))
		require.NoError(err)
		require.Equal(startBalance.Add(math.NewInt(1000)), s.bankKeeper.GetBalance(s.ctx, acc.Account, s.denom).Amount)
	}
	expectedTeamBalance := math.NewInt(300*1e6 - 1000*5)
	require.Equal(expectedTeamBalance, s.bankKeeper.GetBalance(s.ctx, mintToTeamAcc, s.denom).Amount)

	// transfer from account 0 to account 1
	s.bankKeeper.SendCoins(s.ctx, accounts[0].Account, accounts[1].Account, sdk.NewCoins(sdk.NewCoin(s.denom, math.NewInt(1000))))
	require.Equal(math.NewInt(0), s.bankKeeper.GetBalance(s.ctx, accounts[0].Account, s.denom).Amount)
	require.Equal(math.NewInt(2000), s.bankKeeper.GetBalance(s.ctx, accounts[1].Account, s.denom).Amount)

	// transfer from account 2 to team
	s.bankKeeper.SendCoinsFromAccountToModule(s.ctx, accounts[2].Account, minttypes.MintToTeam, sdk.NewCoins(sdk.NewCoin(s.denom, math.NewInt(1000))))
	require.Equal(math.NewInt(0), s.bankKeeper.GetBalance(s.ctx, accounts[2].Account, s.denom).Amount)
	require.Equal(expectedTeamBalance.Add(math.NewInt(1000)), s.bankKeeper.GetBalance(s.ctx, mintToTeamAcc, s.denom).Amount)

	// try to transfer more than balance from account 3 to 4
	err := s.bankKeeper.SendCoins(s.ctx, accounts[3].Account, accounts[4].Account, sdk.NewCoins(sdk.NewCoin(s.denom, math.NewInt(1001))))
	require.Error(err)
	require.Equal(s.bankKeeper.GetBalance(s.ctx, accounts[3].Account, s.denom).Amount, math.NewInt(1000))
	require.Equal(s.bankKeeper.GetBalance(s.ctx, accounts[4].Account, s.denom).Amount, math.NewInt(1000))
}

func (s *E2ETestSuite) TestValidateCycleList() {
	require := s.Require()

	// height 0
	_, err := s.app.BeginBlocker(s.ctx)
	require.NoError(err)
	firstInCycle, err := s.oraclekeeper.GetCurrentQueryInCycleList(s.ctx)
	require.NoError(err)
	require.Equal(strings.ToLower(ethQueryData[2:]), firstInCycle)
	require.Equal(s.ctx.BlockHeight(), int64(0))
	_, err = s.app.EndBlocker(s.ctx)
	require.NoError(err)

	// height 1
	s.ctx = s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1)
	_, err = s.app.BeginBlocker(s.ctx)
	require.NoError(err)
	require.Equal(s.ctx.BlockHeight(), int64(1))
	secondInCycle, err := s.oraclekeeper.GetCurrentQueryInCycleList(s.ctx)
	require.NoError(err)
	require.Equal(strings.ToLower(btcQueryData[2:]), secondInCycle)
	_, err = s.app.EndBlocker(s.ctx)
	require.NoError(err)

	// height 2
	s.ctx = s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1)
	_, err = s.app.BeginBlocker(s.ctx)
	require.NoError(err)
	require.Equal(s.ctx.BlockHeight(), int64(2))
	thirdInCycle, err := s.oraclekeeper.GetCurrentQueryInCycleList(s.ctx)
	require.NoError(err)
	require.Equal(strings.ToLower(trbQueryData[2:]), thirdInCycle)
	_, err = s.app.EndBlocker(s.ctx)
	require.NoError(err)

	// loop through more times
	list, err := s.oraclekeeper.GetCyclelist(s.ctx)
	require.NoError(err)
	for i := 0; i < 20; i++ {
		s.ctx = s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1)
		_, err = s.app.BeginBlocker(s.ctx)
		require.NoError(err)

		query, err := s.oraclekeeper.GetCurrentQueryInCycleList(s.ctx)
		require.NoError(err)
		require.Contains(list, query)

		_, err = s.app.EndBlocker(s.ctx)
		require.NoError(err)
	}
}

func (s *E2ETestSuite) TestSetUpValidatorAndReporter() {
	require := s.Require()

	// Create Validator Accounts
	numValidators := 10
	base := new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)
	_ = new(big.Int).Mul(big.NewInt(1000), base)

	// make addresses
	testAddresses := simtestutil.CreateIncrementalAccounts(numValidators)
	// mint 50k tokens to minter account and send to each address
	initCoins := sdk.NewCoin(s.denom, math.NewInt(5000*1e6))
	for _, addr := range testAddresses {
		s.NoError(s.bankKeeper.MintCoins(s.ctx, authtypes.Minter, sdk.NewCoins(initCoins)))
		s.NoError(s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, authtypes.Minter, addr, sdk.NewCoins(initCoins)))
	}
	// get val address for each test address
	valAddresses := simtestutil.ConvertAddrsToValAddrs(testAddresses)
	// create pub keys for each address
	pubKeys := simtestutil.CreateTestPubKeys(numValidators)

	// set each account with proper keepers
	for i, pubKey := range pubKeys {
		s.accountKeeper.NewAccountWithAddress(s.ctx, testAddresses[i])
		validator, err := stakingtypes.NewValidator(valAddresses[i].String(), pubKey, stakingtypes.Description{Moniker: strconv.Itoa(i)})
		require.NoError(err)
		s.stakingKeeper.SetValidator(s.ctx, validator)
		s.stakingKeeper.SetValidatorByConsAddr(s.ctx, validator)
		s.stakingKeeper.SetNewValidatorByPowerIndex(s.ctx, validator)

		randomStakeAmount := rand.Intn(5000-1000+1) + 1000
		require.True(randomStakeAmount >= 1000 && randomStakeAmount <= 5000, "randomStakeAmount is not within the expected range")
		_, err = s.stakingKeeper.Delegate(s.ctx, testAddresses[i], math.NewInt(int64(randomStakeAmount)*1e6), stakingtypes.Unbonded, validator, true)
		require.NoError(err)
		// call hooks for distribution init
		valBz, err := s.stakingKeeper.ValidatorAddressCodec().StringToBytes(validator.GetOperator())
		if err != nil {
			panic(err)
		}
		err = s.distrKeeper.Hooks().AfterValidatorCreated(s.ctx, valBz)
		require.NoError(err)
		err = s.distrKeeper.Hooks().BeforeDelegationCreated(s.ctx, testAddresses[i], valBz)
		require.NoError(err)
		err = s.distrKeeper.Hooks().AfterDelegationModified(s.ctx, testAddresses[i], valBz)
		require.NoError(err)
	}

	_, err := s.stakingKeeper.EndBlocker(s.ctx)
	s.NoError(err)

	// check that everyone is a bonded validator
	validatorSet, err := s.stakingKeeper.GetAllValidators(s.ctx)
	require.NoError(err)
	for _, val := range validatorSet {
		status := val.GetStatus()
		require.Equal(stakingtypes.Bonded.String(), status.String())
	}

	// create 3 delegators
	const (
		reporter     = "reporter"
		delegatorI   = "delegator1"
		delegatorII  = "delegator2"
		delegatorIII = "delegator3"
	)

	type Delegator struct {
		delegatorAddress sdk.AccAddress
		validator        stakingtypes.Validator
		tokenAmount      math.Int
	}

	numDelegators := 4
	// create random private keys for each delegator
	delegatorPrivateKeys := make([]secp256k1.PrivKey, numDelegators)
	for i := 0; i < numDelegators; i++ {
		pk := secp256k1.GenPrivKey()
		delegatorPrivateKeys[i] = *pk
	}
	// turn private keys into accounts
	delegatorAccounts := make([]sdk.AccAddress, numDelegators)
	for i, pk := range delegatorPrivateKeys {
		delegatorAccounts[i] = sdk.AccAddress(pk.PubKey().Address())
		// give each account tokens
		s.NoError(s.bankKeeper.MintCoins(s.ctx, authtypes.Minter, sdk.NewCoins(initCoins)))
		s.NoError(s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, authtypes.Minter, delegatorAccounts[i], sdk.NewCoins(initCoins)))
	}
	// define each delegator
	delegators := map[string]Delegator{
		reporter:     {delegatorAddress: delegatorAccounts[0], validator: validatorSet[1], tokenAmount: math.NewInt(100 * 1e6)},
		delegatorI:   {delegatorAddress: delegatorAccounts[1], validator: validatorSet[1], tokenAmount: math.NewInt(100 * 1e6)},
		delegatorII:  {delegatorAddress: delegatorAccounts[2], validator: validatorSet[1], tokenAmount: math.NewInt(100 * 1e6)},
		delegatorIII: {delegatorAddress: delegatorAccounts[3], validator: validatorSet[2], tokenAmount: math.NewInt(100 * 1e6)},
	}
	// delegate to validators
	for _, del := range delegators {
		_, err := s.stakingKeeper.Delegate(s.ctx, del.delegatorAddress, del.tokenAmount, stakingtypes.Unbonded, del.validator, true)
		require.NoError(err)
	}

	// set up reporter module msgServer
	msgServerReporter := reporterkeeper.NewMsgServerImpl(s.reporterkeeper)
	require.NotNil(msgServerReporter)
	// define reporter params
	var createReporterMsg reportertypes.MsgCreateReporter
	reporterAddress := delegators[reporter].delegatorAddress.String()
	amount := math.NewInt(100 * 1e6)
	source := reportertypes.TokenOrigin{ValidatorAddress: validatorSet[1].GetOperator(), Amount: math.NewInt(100 * 1e6)}
	commission := stakingtypes.NewCommissionWithTime(math.LegacyNewDecWithPrec(1, 1), math.LegacyNewDecWithPrec(3, 1),
		math.LegacyNewDecWithPrec(1, 1), s.ctx.BlockTime())
	// fill in createReporterMsg
	createReporterMsg.Reporter = reporterAddress
	createReporterMsg.Amount = amount
	createReporterMsg.TokenOrigins = []*reportertypes.TokenOrigin{&source}
	createReporterMsg.Commission = &commission
	// create reporter through msg server
	_, err = msgServerReporter.CreateReporter(s.ctx, &createReporterMsg)
	require.NoError(err)
	// check that reporter was created correctly
	oracleReporter, err := s.reporterkeeper.Reporters.Get(s.ctx, delegators[reporter].delegatorAddress)
	require.NoError(err)
	require.Equal(oracleReporter.Reporter, delegators[reporter].delegatorAddress.String())
	require.Equal(oracleReporter.TotalTokens, math.NewInt(100*1e6))
	require.Equal(oracleReporter.Jailed, false)

	// define delegation source
	source = reportertypes.TokenOrigin{ValidatorAddress: validatorSet[1].GetOperator(), Amount: math.NewInt(25 * 1e6)}
	delegation := reportertypes.NewMsgDelegateReporter(
		delegators[delegatorI].delegatorAddress.String(),
		delegators[reporter].delegatorAddress.String(),
		math.NewInt(25*1e6),
		[]*reportertypes.TokenOrigin{&source},
	)
	// self delegate as reporter
	_, err = msgServerReporter.DelegateReporter(s.ctx, delegation)
	require.NoError(err)
	delegationReporter, err := s.reporterkeeper.Delegators.Get(s.ctx, delegators[delegatorI].delegatorAddress)
	require.NoError(err)
	require.Equal(delegationReporter.Reporter, delegators[reporter].delegatorAddress.String())
}

// todo: claim tbr/tips for these reports
func (s *E2ETestSuite) TestBasicReporting() {
	require := s.Require()

	// create a validator
	// create account that will become a validator
	testAccount := simtestutil.CreateIncrementalAccounts(1)
	// mint 5000*1e6tokens for validator
	initCoins := sdk.NewCoin(s.denom, math.NewInt(5000*1e8))
	require.NoError(s.bankKeeper.MintCoins(s.ctx, authtypes.Minter, sdk.NewCoins(initCoins)))
	require.NoError(s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, authtypes.Minter, testAccount[0], sdk.NewCoins(initCoins)))
	// get val address
	testAccountValAddrs := simtestutil.ConvertAddrsToValAddrs(testAccount)
	// create pub key for validator
	pubKey := simtestutil.CreateTestPubKeys(1)
	// tell keepers about the new validator
	s.accountKeeper.NewAccountWithAddress(s.ctx, testAccount[0])
	validator, err := stakingtypes.NewValidator(testAccountValAddrs[0].String(), pubKey[0], stakingtypes.Description{Moniker: "created validator"})
	require.NoError(err)
	s.stakingKeeper.SetValidator(s.ctx, validator)
	s.stakingKeeper.SetValidatorByConsAddr(s.ctx, validator)
	s.stakingKeeper.SetNewValidatorByPowerIndex(s.ctx, validator)
	// self delegate from validator account to itself
	_, err = s.stakingKeeper.Delegate(s.ctx, testAccount[0], math.NewInt(int64(4000)*1e8), stakingtypes.Unbonded, validator, true)
	require.NoError(err)
	// call hooks for distribution init
	valBz, err := s.stakingKeeper.ValidatorAddressCodec().StringToBytes(validator.GetOperator())
	if err != nil {
		panic(err)
	}
	err = s.distrKeeper.Hooks().AfterValidatorCreated(s.ctx, valBz)
	require.NoError(err)
	err = s.distrKeeper.Hooks().BeforeDelegationCreated(s.ctx, testAccount[0], valBz)
	require.NoError(err)
	err = s.distrKeeper.Hooks().AfterDelegationModified(s.ctx, testAccount[0], valBz)
	require.NoError(err)
	_, err = s.stakingKeeper.EndBlocker(s.ctx)
	s.NoError(err)

	//create a self delegated reporter from a different account
	type Delegator struct {
		delegatorAddress sdk.AccAddress
		validator        stakingtypes.Validator
		tokenAmount      math.Int
	}
	pk := secp256k1.GenPrivKey()
	reporterAccount := sdk.AccAddress(pk.PubKey().Address())
	// mint 5000*1e6 tokens for reporter
	s.NoError(s.bankKeeper.MintCoins(s.ctx, authtypes.Minter, sdk.NewCoins(initCoins)))
	s.NoError(s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, authtypes.Minter, reporterAccount, sdk.NewCoins(initCoins)))
	// delegate to validator so reporterDelforVal can delegate to themselves
	reporterDelforVal := Delegator{delegatorAddress: reporterAccount, validator: validator, tokenAmount: math.NewInt(5000 * 1e6)}
	_, err = s.stakingKeeper.Delegate(s.ctx, reporterDelforVal.delegatorAddress, reporterDelforVal.tokenAmount, stakingtypes.Unbonded, reporterDelforVal.validator, true)
	require.NoError(err)
	// set up reporter module msgServer
	msgServerReporter := reporterkeeper.NewMsgServerImpl(s.reporterkeeper)
	require.NotNil(msgServerReporter)
	// define createReporterMsg params
	var createReporterMsg reportertypes.MsgCreateReporter
	reporterAddress := reporterDelforVal.delegatorAddress.String()
	amount := math.NewInt(4000 * 1e6)
	source := reportertypes.TokenOrigin{ValidatorAddress: validator.OperatorAddress, Amount: math.NewInt(4000 * 1e6)}
	commission := stakingtypes.NewCommissionWithTime(math.LegacyNewDecWithPrec(1, 1), math.LegacyNewDecWithPrec(3, 1),
		math.LegacyNewDecWithPrec(1, 1), s.ctx.BlockTime())
	// fill in createReporterMsg
	createReporterMsg.Reporter = reporterAddress
	createReporterMsg.Amount = amount
	createReporterMsg.TokenOrigins = []*reportertypes.TokenOrigin{&source}
	createReporterMsg.Commission = &commission
	// send createreporter msg
	_, err = msgServerReporter.CreateReporter(s.ctx, &createReporterMsg)
	require.NoError(err)
	// check that reporter was created in Reporters collections
	reporter, err := s.reporterkeeper.Reporters.Get(s.ctx, reporterDelforVal.delegatorAddress)
	require.NoError(err)
	require.Equal(reporter.Reporter, reporterDelforVal.delegatorAddress.String())
	require.Equal(reporter.TotalTokens, math.NewInt(4000*1e6))
	require.Equal(reporter.Jailed, false)
	// check on reporter in Delegators collections
	delegation, err := s.reporterkeeper.Delegators.Get(s.ctx, reporterDelforVal.delegatorAddress)
	require.NoError(err)
	require.Equal(delegation.Reporter, reporterDelforVal.delegatorAddress.String())
	require.Equal(delegation.Amount, math.NewInt(4000*1e6))

	// Report
	// setup oracle msgServer
	msgServerOracle := oraclekeeper.NewMsgServerImpl(s.oraclekeeper)
	require.NotNil(msgServerOracle)
	// case 1: commit/reveal for cycle list
	queryInCycleList1, err := s.oraclekeeper.GetCurrentQueryInCycleList(s.ctx)
	require.NoError(err)
	// create hash for commit
	saltEth, err := oracleutils.Salt(32)
	require.NoError(err)
	valueEth := encodeValue(4500)
	hashEth := oracleutils.CalculateCommitment(valueEth, saltEth)
	// create commitEth msg
	commitEth := oracletypes.MsgCommitReport{
		Creator:   reporter.Reporter,
		QueryData: queryInCycleList1,
		Hash:      hashEth,
	}
	// send commit tx
	commitResponseEth, err := msgServerOracle.CommitReport(s.ctx, &commitEth)
	require.NoError(err)
	require.NotNil(commitResponseEth)
	commitHeight := s.ctx.BlockHeight()
	require.Equal(int64(0), commitHeight)
	// advance 1 block
	s.ctx = s.ctx.WithBlockHeight(commitHeight + 1) // height 1
	// create reveal msg
	revealEth := oracletypes.MsgSubmitValue{
		Creator:   reporter.Reporter,
		QueryData: queryInCycleList1,
		Value:     valueEth,
		Salt:      saltEth,
	}
	// send reveal tx
	revealResponseEth, err := msgServerOracle.SubmitValue(s.ctx, &revealEth)
	require.NoError(err)
	require.NotNil(revealResponseEth)
	// advance time and block height to expire the query and aggregate report
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(7 * time.Second))
	_, err = s.app.EndBlocker(s.ctx)
	require.NoError(err)
	// get queryId for GetAggregatedReportRequest
	queryIdEth, err := utils.QueryIDFromDataString(queryInCycleList1)
	s.NoError(err)
	// check that aggregated report is stored
	getAggReportRequestEth := oracletypes.QueryGetCurrentAggregatedReportRequest{
		QueryId: hex.EncodeToString(queryIdEth),
	}
	resultEth, err := s.oraclekeeper.GetAggregatedReport(s.ctx, &getAggReportRequestEth)
	require.NoError(err)
	require.Equal(resultEth.Report.Height, int64(1))
	require.Equal(resultEth.Report.AggregateReportIndex, int64(0))
	require.Equal(resultEth.Report.AggregateValue, encodeValue(4500))
	require.Equal(resultEth.Report.AggregateReporter, reporter.Reporter)
	require.Equal(resultEth.Report.QueryId, hex.EncodeToString(queryIdEth))
	require.Equal(int64(4000), resultEth.Report.ReporterPower)

	// case 2: submit without committing for cycle list
	s.ctx = s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1) // height 2
	// get new cycle list query data
	queryInCycleList2, err := s.oraclekeeper.GetCurrentQueryInCycleList(s.ctx)
	require.NoError(err)
	require.NotEqual(queryInCycleList1, queryInCycleList2)
	// create reveal message
	valueBtc := encodeValue(100_000)
	require.NoError(err)
	revealBtc := oracletypes.MsgSubmitValue{
		Creator:   reporter.Reporter,
		QueryData: queryInCycleList2,
		Value:     valueBtc,
	}
	// send reveal message
	revealResponseBtc, err := msgServerOracle.SubmitValue(s.ctx, &revealBtc)
	require.NoError(err)
	require.NotNil(revealResponseBtc)
	// advance time and block height to expire the query and aggregate report
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(7 * time.Second))
	_, err = s.app.EndBlocker(s.ctx)
	require.NoError(err)
	// get queryId for GetAggregatedReportRequest
	queryIdBtc, err := utils.QueryIDFromDataString(queryInCycleList2)
	s.NoError(err)
	// create get aggregated report query
	getAggReportRequestBtc := oracletypes.QueryGetCurrentAggregatedReportRequest{
		QueryId: hex.EncodeToString(queryIdBtc),
	}
	// check that aggregated report is stored correctly
	resultBtc, err := s.oraclekeeper.GetAggregatedReport(s.ctx, &getAggReportRequestBtc)
	require.NoError(err)
	require.Equal(int64(0), resultBtc.Report.AggregateReportIndex)
	require.Equal(encodeValue(100_000), resultBtc.Report.AggregateValue)
	require.Equal(reporter.Reporter, resultBtc.Report.AggregateReporter)
	require.Equal(hex.EncodeToString(queryIdBtc), resultBtc.Report.QueryId)
	require.Equal(int64(4000), resultBtc.Report.ReporterPower)
	require.Equal(int64(2), resultBtc.Report.Height)

	// case 3: commit/reveal for tipped query
	s.ctx = s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1) // height 3
	// create tip msg
	tipAmount := sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(1))
	msgTip := oracletypes.MsgTip{
		Tipper:    reporter.Reporter,
		QueryData: queryInCycleList1,
		Amount:    tipAmount,
	}
	// send tip tx
	tipRes, err := msgServerOracle.Tip(s.ctx, &msgTip)
	require.NoError(err)
	require.NotNil(tipRes)
	// check that tip is in oracle module account
	tipModuleAcct := s.accountKeeper.GetModuleAddress(oracletypes.ModuleName)
	tipAcctBalance := s.bankKeeper.GetBalance(s.ctx, tipModuleAcct, sdk.DefaultBondDenom)
	require.Equal(tipAcctBalance, tipAmount)
	// create commit for tipped eth query
	saltEth, err = oracleutils.Salt(32)
	require.NoError(err)
	valueEth = encodeValue(5000)
	hashEth = oracleutils.CalculateCommitment(valueEth, saltEth)
	commitEth = oracletypes.MsgCommitReport{
		Creator:   reporter.Reporter,
		QueryData: queryInCycleList1,
		Hash:      hashEth,
	}
	// send commit tx
	commitResponseEth, err = msgServerOracle.CommitReport(s.ctx, &commitEth)
	require.NoError(err)
	require.NotNil(commitResponseEth)
	commitHeight = s.ctx.BlockHeight()
	require.Equal(int64(3), commitHeight)
	// advance 1 block
	s.ctx = s.ctx.WithBlockHeight(commitHeight + 1) // height 4
	// create reveal msg
	valueEth = encodeValue(5000)
	revealEth = oracletypes.MsgSubmitValue{
		Creator:   reporter.Reporter,
		QueryData: queryInCycleList1,
		Value:     valueEth,
		Salt:      saltEth,
	}
	// send reveal tx
	revealResponseEth, err = msgServerOracle.SubmitValue(s.ctx, &revealEth)
	require.NoError(err)
	require.NotNil(revealResponseEth)
	// advance time and block height to expire the query and aggregate report
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(7 * time.Second))
	_, err = s.app.EndBlocker(s.ctx)
	require.NoError(err)
	// create get aggreagted report query
	getAggReportRequestEth = oracletypes.QueryGetCurrentAggregatedReportRequest{
		QueryId: hex.EncodeToString(queryIdEth),
	}
	// check that the aggregated report is stored correctly
	resultEth, err = s.oraclekeeper.GetAggregatedReport(s.ctx, &getAggReportRequestEth)
	require.NoError(err)
	require.Equal(resultEth.Report.AggregateReportIndex, int64(0))
	require.Equal(resultEth.Report.AggregateValue, encodeValue(5000))
	require.Equal(resultEth.Report.AggregateReporter, reporter.Reporter)
	require.Equal(hex.EncodeToString(queryIdEth), resultEth.Report.QueryId)
	require.Equal(int64(4000), resultEth.Report.ReporterPower)
	require.Equal(int64(4), resultEth.Report.Height)

	// case 4: submit without committing for tipped query
	s.ctx = s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1) // height 5
	// create tip msg
	msgTip = oracletypes.MsgTip{
		Tipper:    reporter.Reporter,
		QueryData: trbQueryData,
		Amount:    tipAmount,
	}
	// send tip tx
	tipRes, err = msgServerOracle.Tip(s.ctx, &msgTip)
	require.NoError(err)
	require.NotNil(tipRes)
	// check that tip is in oracle module account
	tipModuleAcct = s.accountKeeper.GetModuleAddress(oracletypes.ModuleName)
	tipAcctBalance = s.bankKeeper.GetBalance(s.ctx, tipModuleAcct, sdk.DefaultBondDenom)
	require.Equal(tipAcctBalance, tipAmount)
	// create submit msg
	revealMsgTrb := oracletypes.MsgSubmitValue{
		Creator:   reporter.Reporter,
		QueryData: trbQueryData,
		Value:     encodeValue(1_000_000),
	}
	// send submit msg
	revealTrb, err := msgServerOracle.SubmitValue(s.ctx, &revealMsgTrb)
	require.NoError(err)
	require.NotNil(revealTrb)
	// advance time and block height to expire the query and aggregate report
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(7 * time.Second))
	_, err = s.app.EndBlocker(s.ctx)
	require.NoError(err)
	// get trb query id
	queryIdTrb, err := utils.QueryIDFromDataString(trbQueryData)
	s.NoError(err)
	// create get aggregated report query
	getAggReportRequestTrb := oracletypes.QueryGetCurrentAggregatedReportRequest{
		QueryId: hex.EncodeToString(queryIdTrb),
	}
	// query aggregated report
	resultTrb, err := s.oraclekeeper.GetAggregatedReport(s.ctx, &getAggReportRequestTrb)
	require.NoError(err)
	require.Equal(resultTrb.Report.AggregateReportIndex, int64(0))
	require.Equal(resultTrb.Report.AggregateValue, encodeValue(1_000_000))
	require.Equal(resultTrb.Report.AggregateReporter, reporter.Reporter)
	require.Equal(hex.EncodeToString(queryIdTrb), resultTrb.Report.QueryId)
	require.Equal(int64(4000), resultTrb.Report.ReporterPower)
	require.Equal(int64(5), resultTrb.Report.Height)
}

// todo: dispute test

func (s *E2ETestSuite) TestDisputes() {
	require := s.Require()
	_, msgServerDispute := s.disputeKeeper()
	require.NotNil(msgServerDispute)

}
