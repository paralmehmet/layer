package reporter

import (
	modulev1 "github.com/tellor-io/layer/api/layer/reporter"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "Reporters",
					Use:            "reporters",
					Short:          "Query staked reporters",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},
				{
					RpcMethod:      "SelectorReporter",
					Use:            "selector-reporter [selector-address]",
					Short:          "Query reporter of a delegator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "selector_address"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateReporter",
					Use:            "create-reporter [commission-rate] [min-tokens-required]",
					Short:          "Execute the CreateReporter RPC method",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "commission_rate"}, {ProtoField: "min_tokens_required"}},
				},
				{
					RpcMethod:      "SelectReporter",
					Use:            "select-reporter [reporter-address]",
					Short:          "Execute the SelectReporter RPC method",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "reporter_address"}},
				},
				{
					RpcMethod:      "SwitchReporter",
					Use:            "switch-reporter [reporter-address]",
					Short:          "Execute the SwitchReporter RPC method",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "reporter_address"}},
				},
				{
					RpcMethod:      "RemoveSelector",
					Use:            "remove-selector [selector-address]",
					Short:          "Execute the RemoveSelector RPC method",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "selector_address"}},
				},
				{
					RpcMethod:      "UnjailReporter",
					Use:            "unjail-reporter [reporter-addr]",
					Short:          "Execute the UnjailReporter RPC method",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "reporter_address"}},
				},
				{
					RpcMethod:      "WithdrawTip",
					Use:            "withdraw-tip [selector-address] [validator-address]",
					Short:          "Send a WithdrawTip tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "selector_address"}, {ProtoField: "validator_address"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
