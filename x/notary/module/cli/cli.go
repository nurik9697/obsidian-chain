package cli

import (
	"github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"

)

// GetTxCmd объединяет все транзакции модуля под командой "notary"
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "notary",
		Short:                      "Notary transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// Регистрируем подкоманду создания документа
	cmd.AddCommand(CmdCreateDocument())

	return cmd
}



// GetQueryCmd возвращает команды для чтения данных (пока базово)
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "notary",
		Short:                      "Querying commands for the notary module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// ДОБАВЬ ЭТУ СТРОКУ (замени CmdListDocument на точное название твоей функции)
	cmd.AddCommand(CmdListDocument())

	return cmd
}

