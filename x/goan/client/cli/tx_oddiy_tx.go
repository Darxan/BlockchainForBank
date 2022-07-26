package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"goan/x/goan/types"
)

var _ = strconv.Itoa(0)

func CmdOddiyTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oddiy-tx [sender] [receiver] [amount] [fee] [fee-receiver] [tx-type] [service-name]",
		Short: "Broadcast message oddiyTx",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSender := args[0]
			argReceiver := args[1]
			argAmount := args[2]
			argFee := args[3]
			argFeeReceiver := args[4]
			argTxType := args[5]
			argServiceName := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgOddiyTx(
				clientCtx.GetFromAddress().String(),
				argSender,
				argReceiver,
				argAmount,
				argFee,
				argFeeReceiver,
				argTxType,
				argServiceName,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
