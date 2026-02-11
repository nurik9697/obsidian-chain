package notary

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	//modulev1 "obsidian/api/obsidian/notary"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
        return nil
}
