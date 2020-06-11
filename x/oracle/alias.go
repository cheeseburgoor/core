// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/terra-project/core/x/oracle/internal/types/
// ALIASGEN: github.com/terra-project/core/x/oracle/internal/keeper/
package oracle

import (
	"github.com/terra-project/core/x/oracle/internal/keeper"
	"github.com/terra-project/core/x/oracle/internal/types"
)

const (
	ModuleName                      = types.ModuleName
	StoreKey                        = types.StoreKey
	RouterKey                       = types.RouterKey
	QuerierRoute                    = types.QuerierRoute
	DefaultParamspace               = types.DefaultParamspace
	DefaultVotePeriod               = types.DefaultVotePeriod
	DefaultSlashWindow              = types.DefaultSlashWindow
	DefaultRewardDistributionWindow = types.DefaultRewardDistributionWindow
	QueryParameters                 = types.QueryParameters
	QueryExchangeRate               = types.QueryExchangeRate
	QueryExchangeRates              = types.QueryExchangeRates
	QueryActives                    = types.QueryActives
	QueryPrevotes                   = types.QueryPrevotes
	QueryVotes                      = types.QueryVotes
	QueryFeederDelegation           = types.QueryFeederDelegation
	QueryMissCounter                = types.QueryMissCounter
	QueryAggregatePrevote           = types.QueryAggregatePrevote
	QueryAggregateVote              = types.QueryAggregateVote
	QueryVoteTargets                = types.QueryVoteTargets
	QueryTobinTax                   = types.QueryTobinTax
	QueryTobinTaxes                 = types.QueryTobinTaxes
)

var (
	// functions aliases
	NewVoteForTally                    = types.NewVoteForTally
	NewClaim                           = types.NewClaim
	RegisterCodec                      = types.RegisterCodec
	NewGenesisState                    = types.NewGenesisState
	DefaultGenesisState                = types.DefaultGenesisState
	ValidateGenesis                    = types.ValidateGenesis
	GetVoteHash                        = types.GetVoteHash
	VoteHashFromHexString              = types.VoteHashFromHexString
	GetAggregateVoteHash               = types.GetAggregateVoteHash
	AggregateVoteHashFromHexString     = types.AggregateVoteHashFromHexString
	GetExchangeRatePrevoteKey          = types.GetExchangeRatePrevoteKey
	GetVoteKey                         = types.GetVoteKey
	GetExchangeRateKey                 = types.GetExchangeRateKey
	GetFeederDelegationKey             = types.GetFeederDelegationKey
	GetMissCounterKey                  = types.GetMissCounterKey
	GetAggregateExchangeRatePrevoteKey = types.GetAggregateExchangeRatePrevoteKey
	GetAggregateExchangeRateVoteKey    = types.GetAggregateExchangeRateVoteKey
	GetTobinTaxKey                     = types.GetTobinTaxKey
	ExtractDenomFromTobinTaxKey        = types.ExtractDenomFromTobinTaxKey
	NewMsgExchangeRatePrevote          = types.NewMsgExchangeRatePrevote
	NewMsgExchangeRateVote             = types.NewMsgExchangeRateVote
	NewMsgDelegateFeedConsent          = types.NewMsgDelegateFeedConsent
	NewMsgAggregateExchangeRatePrevote = types.NewMsgAggregateExchangeRatePrevote
	NewMsgAggregateExchangeRateVote    = types.NewMsgAggregateExchangeRateVote
	DefaultParams                      = types.DefaultParams
	ParamKeyTable                      = types.ParamKeyTable
	NewQueryExchangeRateParams         = types.NewQueryExchangeRateParams
	NewQueryPrevotesParams             = types.NewQueryPrevotesParams
	NewQueryVotesParams                = types.NewQueryVotesParams
	NewQueryFeederDelegationParams     = types.NewQueryFeederDelegationParams
	NewQueryMissCounterParams          = types.NewQueryMissCounterParams
	NewQueryAggregatePrevoteParams     = types.NewQueryAggregatePrevoteParams
	NewQueryAggregateVoteParams        = types.NewQueryAggregateVoteParams
	NewQueryTobinTaxParams             = types.NewQueryTobinTaxParams
	NewExchangeRatePrevote             = types.NewExchangeRatePrevote
	NewExchangeRateVote                = types.NewExchangeRateVote
	NewAggregateExchangeRatePrevote    = types.NewAggregateExchangeRatePrevote
	ParseExchangeRateTuples            = types.ParseExchangeRateTuples
	NewAggregateExchangeRateVote       = types.NewAggregateExchangeRateVote
	NewKeeper                          = keeper.NewKeeper
	NewQuerier                         = keeper.NewQuerier

	// variable aliases
	ModuleCdc                             = types.ModuleCdc
	ErrUnknowDenom                        = types.ErrUnknowDenom
	ErrInvalidExchangeRate                = types.ErrInvalidExchangeRate
	ErrNoPrevote                          = types.ErrNoPrevote
	ErrNoVote                             = types.ErrNoVote
	ErrNoVotingPermission                 = types.ErrNoVotingPermission
	ErrInvalidHash                        = types.ErrInvalidHash
	ErrInvalidHashLength                  = types.ErrInvalidHashLength
	ErrVerificationFailed                 = types.ErrVerificationFailed
	ErrRevealPeriodMissMatch              = types.ErrRevealPeriodMissMatch
	ErrInvalidSaltLength                  = types.ErrInvalidSaltLength
	ErrNoAggregatePrevote                 = types.ErrNoAggregatePrevote
	ErrNoAggregateVote                    = types.ErrNoAggregateVote
	ErrNoTobinTax                         = types.ErrNoTobinTax
	PrevoteKey                            = types.PrevoteKey
	VoteKey                               = types.VoteKey
	ExchangeRateKey                       = types.ExchangeRateKey
	FeederDelegationKey                   = types.FeederDelegationKey
	MissCounterKey                        = types.MissCounterKey
	AggregatePrevoteKey                   = types.AggregatePrevoteKey
	AggregateVoteKey                      = types.AggregateVoteKey
	TobinTaxKey                           = types.TobinTaxKey
	ParamStoreKeyVotePeriod               = types.ParamStoreKeyVotePeriod
	ParamStoreKeyVoteThreshold            = types.ParamStoreKeyVoteThreshold
	ParamStoreKeyRewardBand               = types.ParamStoreKeyRewardBand
	ParamStoreKeyRewardDistributionWindow = types.ParamStoreKeyRewardDistributionWindow
	ParamStoreKeyWhitelist                = types.ParamStoreKeyWhitelist
	ParamStoreKeySlashFraction            = types.ParamStoreKeySlashFraction
	ParamStoreKeySlashWindow              = types.ParamStoreKeySlashWindow
	ParamStoreKeyMinValidPerWindow        = types.ParamStoreKeyMinValidPerWindow
	DefaultVoteThreshold                  = types.DefaultVoteThreshold
	DefaultRewardBand                     = types.DefaultRewardBand
	DefaultTobinTax                       = types.DefaultTobinTax
	DefaultWhitelist                      = types.DefaultWhitelist
	DefaultSlashFraction                  = types.DefaultSlashFraction
	DefaultMinValidPerWindow              = types.DefaultMinValidPerWindow
)

type (
	VoteForTally                    = types.VoteForTally
	ExchangeRateBallot              = types.ExchangeRateBallot
	Claim                           = types.Claim
	Denom                           = types.Denom
	DenomList                       = types.DenomList
	StakingKeeper                   = types.StakingKeeper
	DistributionKeeper              = types.DistributionKeeper
	SupplyKeeper                    = types.SupplyKeeper
	GenesisState                    = types.GenesisState
	VoteHash                        = types.VoteHash
	AggregateVoteHash               = types.AggregateVoteHash
	MsgExchangeRatePrevote          = types.MsgExchangeRatePrevote
	MsgExchangeRateVote             = types.MsgExchangeRateVote
	MsgDelegateFeedConsent          = types.MsgDelegateFeedConsent
	MsgAggregateExchangeRatePrevote = types.MsgAggregateExchangeRatePrevote
	MsgAggregateExchangeRateVote    = types.MsgAggregateExchangeRateVote
	Params                          = types.Params
	QueryExchangeRateParams         = types.QueryExchangeRateParams
	QueryPrevotesParams             = types.QueryPrevotesParams
	QueryVotesParams                = types.QueryVotesParams
	QueryFeederDelegationParams     = types.QueryFeederDelegationParams
	QueryMissCounterParams          = types.QueryMissCounterParams
	QueryAggregatePrevoteParams     = types.QueryAggregatePrevoteParams
	QueryAggregateVoteParams        = types.QueryAggregateVoteParams
	QueryTobinTaxParams             = types.QueryTobinTaxParams
	ExchangeRatePrevote             = types.ExchangeRatePrevote
	ExchangeRatePrevotes            = types.ExchangeRatePrevotes
	ExchangeRateVote                = types.ExchangeRateVote
	ExchangeRateVotes               = types.ExchangeRateVotes
	AggregateExchangeRatePrevote    = types.AggregateExchangeRatePrevote
	ExchangeRateTuple               = types.ExchangeRateTuple
	ExchangeRateTuples              = types.ExchangeRateTuples
	AggregateExchangeRateVote       = types.AggregateExchangeRateVote
	Keeper                          = keeper.Keeper
)