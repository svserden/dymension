package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dymensionxyz/dymension/x/rollapp/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group rollapp queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListRollapp())
	cmd.AddCommand(CmdShowRollapp())
	cmd.AddCommand(CmdListStateInfo())
	cmd.AddCommand(CmdShowStateInfo())
	cmd.AddCommand(CmdListLatestStateInfoIndex())
	cmd.AddCommand(CmdShowLatestStateInfoIndex())
	cmd.AddCommand(CmdListBlockHeightToFinalizationQueue())
	cmd.AddCommand(CmdShowBlockHeightToFinalizationQueue())
	cmd.AddCommand(CmdShowLatestFinalizedStateInfo())

	// this line is used by starport scaffolding # 1

	return cmd
}
