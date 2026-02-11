package notary

import (
        "context"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
        "github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	modulev1 "obsidian/api/obsidian/notary/v1"
	"obsidian/x/notary/keeper"
        "obsidian/x/notary/types"
        cli "obsidian/x/notary/module/cli"
        "github.com/spf13/cobra"
        _ "obsidian/api/obsidian/notary/v1"
)

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}


type ModuleInputs struct {
	depinject.In

	StoreService store.KVStoreService
	Cdc          codec.Codec
	Config       *modulev1.Module
	Logger       log.Logger

	AccountKeeper authkeeper.AccountKeeper
	BankKeeper    bankkeeper.Keeper
}

type ModuleOutputs struct {
	depinject.Out

	NotaryKeeper keeper.Keeper
	Module       appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	k := keeper.NewKeeper(
		in.Cdc,
		in.StoreService,
		in.Logger,
		authority.String(),
		in.BankKeeper,
		in.AccountKeeper,
	)

	m := NewAppModule(in.Cdc, k, in.AccountKeeper, in.BankKeeper)

	return ModuleOutputs{NotaryKeeper: k, Module: m}
}

type AppModule struct {
	cdc           codec.Codec
	keeper        keeper.Keeper
	accountKeeper authkeeper.AccountKeeper
	bankKeeper    bankkeeper.Keeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
) AppModule {
	return AppModule{
		cdc:           cdc,
		keeper:        keeper,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}
// GetQueryCmd возвращает команды для чтения данных (q)
func (AppModule) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// GetTxCmd возвращает команды для отправки транзакций (tx)
func (AppModule) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// Методы для реализации интерфейса appmodule.AppModule
func (AppModule) IsAppModule() {}
func (AppModule) IsOnePerModuleType() {} // ЭТОТ МЕТОД ИСПРАВИТ ОШИБКУ

func (AppModule) Name() string { return "notary" }

func (am AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

func (am AppModule) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (am AppModule) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

func (am AppModule) ConsensusVersion() uint64 {
	return 1
}

// RegisterServices регистрирует обработчики сообщений (MsgServer) и запросов (QueryServer)
func (am AppModule) RegisterServices(cfg module.Configurator) {
	// Эта строка связывает твои MsgCreateDocument с логикой в Keeper
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))

	// Эта строка регистрирует запросы (Query)
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}
