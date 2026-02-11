package cli

import (
	"strconv"
	"github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"obsidian/x/notary/types"
)

func CmdCreateDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-document [index] [fileHash] [owner] [timestamp]",
		Short: "Create a new document",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			timestamp, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDocument(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				args[2],
				int32(timestamp),
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
