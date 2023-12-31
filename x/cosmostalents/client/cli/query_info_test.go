package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"github.com/cosmos-talents/cosmos-talents/testutil/network"
	"github.com/cosmos-talents/cosmos-talents/testutil/nullify"
	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents/client/cli"
	"github.com/cosmos-talents/cosmos-talents/x/cosmostalents/types"
)

func networkWithInfoObjects(t *testing.T) (*network.Network, types.Info) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	info := &types.Info{}
	nullify.Fill(&info)
	state.Info = info
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.Info
}

func TestShowInfo(t *testing.T) {
	net, obj := networkWithInfoObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc string
		args []string
		err  error
		obj  types.Info
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowInfo(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetInfoResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Info)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Info),
				)
			}
		})
	}
}
