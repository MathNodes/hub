// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/deposit/types
// ALIASGEN: github.com/sentinel-official/hub/x/deposit/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/deposit/querier
package deposit

import (
	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/querier"
	"github.com/sentinel-official/hub/x/deposit/types"
)

const (
	ModuleName    = types.ModuleName
	QuerierRoute  = types.QuerierRoute
	QueryDeposit  = types.QueryDeposit
	QueryDeposits = types.QueryDeposits
)

var (
	// functions aliases
	RegisterCodec          = types.RegisterCodec
	NewGenesisState        = types.NewGenesisState
	DefaultGenesisState    = types.DefaultGenesisState
	DepositKey             = types.DepositKey
	NewQueryDepositParams  = types.NewQueryDepositParams
	NewQueryDepositsParams = types.NewQueryDepositsParams
	NewKeeper              = keeper.NewKeeper
	CreateTestInput        = keeper.CreateTestInput
	MakeTestCodec          = keeper.MakeTestCodec
	Querier                = querier.Querier

	// variable aliases
	ModuleCdc                     = types.ModuleCdc
	ErrorMarshal                  = types.ErrorMarshal
	ErrorUnmarshal                = types.ErrorUnmarshal
	ErrorUnknownMsgType           = types.ErrorUnknownMsgType
	ErrorUnknownQueryType         = types.ErrorUnknownQueryType
	ErrorInvalidField             = types.ErrorInvalidField
	ErrorInsufficientDepositFunds = types.ErrorInsufficientDepositFunds
	ErrorDepositDoesNotExist      = types.ErrorDepositDoesNotExist
	RouterKey                     = types.RouterKey
	StoreKey                      = types.StoreKey
	EventModuleName               = types.EventModuleName
	DepositKeyPrefix              = types.DepositKeyPrefix
)

type (
	Deposit             = types.Deposit
	Deposits            = types.Deposits
	GenesisState        = types.GenesisState
	QueryDepositParams  = types.QueryDepositParams
	QueryDepositsParams = types.QueryDepositsParams
	Keeper              = keeper.Keeper
)
