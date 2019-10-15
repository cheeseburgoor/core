package types

import (
	"fmt"

	core "github.com/terra-project/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
)

// DefaultParamspace
const DefaultParamspace = ModuleName

// Parameter keys
var (
	ParamStoreKeyVotePeriod               = []byte("voteperiod")
	ParamStoreKeyVoteThreshold            = []byte("votethreshold")
	ParamStoreKeyRewardBand               = []byte("rewardband")
	ParamStoreKeyRewardDistributionPeriod = []byte("rewarddistributionperiod")
	ParamStoreKeyVotesWindow              = []byte("voteswindow")
	ParamStoreKeyMinValidVotesPerWindow   = []byte("minvalidvotesperwindow")
	ParamStoreKeySlashFraction            = []byte("slashfraction")
	ParamStoreKeyWhitelist                = []byte("whitelist")
)

// Default parameter values
const (
	DefaultVotePeriod  = core.BlocksPerMinute // 1 minute
	DefaultVotesWindow = int64(1000)          // 1000 oracle period
)

// Default parameter values
var (
	DefaultVoteThreshold            = sdk.NewDecWithPrec(50, 2)                                             // 50%
	DefaultRewardBand               = sdk.NewDecWithPrec(1, 2)                                              // 1%
	DefaultRewardDistributionPeriod = core.BlocksPerMonth                                                   // 432,000
	DefaultMinValidVotesPerWindow   = sdk.NewDecWithPrec(5, 2)                                              // 5%
	DefaultSlashFraction            = sdk.NewDecWithPrec(1, 4)                                              // 0.01%
	DefaultWhitelist                = DenomList{core.MicroKRWDenom, core.MicroSDRDenom, core.MicroUSDDenom} // ukrw, usdr, uusd
)

var _ subspace.ParamSet = &Params{}

// Params oracle parameters
type Params struct {
	VotePeriod               int64     `json:"vote_period" yaml:"vote_period"`                               // the number of blocks during which voting takes place.
	VoteThreshold            sdk.Dec   `json:"vote_threshold" yaml:"vote_threshold"`                         // the minimum percentage of votes that must be received for a ballot to pass.
	RewardBand               sdk.Dec   `json:"reward_band" yaml:"reward_band"`                               // the ratio of allowable price error that can be rewared.
	RewardDistributionPeriod int64     `json:"reward_distribution_period" yaml:"reward_distribution_period"` // the number of blocks of the the period during which seigiornage reward comes in and then is distributed.
	VotesWindow              int64     `json:"votes_window" yaml:"votes_window"`                             // the number of blocks units on which the penalty is based.
	MinValidVotesPerWindow   sdk.Dec   `json:"min_valid_votes_per_window" yaml:"min_valid_votes_per_window"` // the minimum number of blocks to avoid slashing in a window.
	SlashFraction            sdk.Dec   `json:"slash_fraction" yaml:"slash_fraction"`                         // the slashing ratio on the delegated token.
	Whitelist                DenomList `json:"whitelist" yaml:"whitelist"`                                   // the denom list that can be acitivated,
}

// DefaultParams creates default oracle module parameters
func DefaultParams() Params {
	return Params{
		VotePeriod:               DefaultVotePeriod,
		VoteThreshold:            DefaultVoteThreshold,
		RewardBand:               DefaultRewardBand,
		RewardDistributionPeriod: DefaultRewardDistributionPeriod,
		VotesWindow:              DefaultVotesWindow,
		MinValidVotesPerWindow:   DefaultMinValidVotesPerWindow,
		SlashFraction:            DefaultSlashFraction,
		Whitelist:                DefaultWhitelist,
	}
}

// Validate validates a set of params
func (params Params) Validate() error {
	if params.VotePeriod <= 0 {
		return fmt.Errorf("oracle parameter VotePeriod must be > 0, is %d", params.VotePeriod)
	}
	if params.VoteThreshold.LTE(sdk.NewDecWithPrec(33, 2)) {
		return fmt.Errorf("oracle parameter VoteTheshold must be greater than 33 percent")
	}
	if params.RewardBand.IsNegative() {
		return fmt.Errorf("oracle parameter RewardBand must be positive")
	}
	if params.RewardDistributionPeriod < params.VotePeriod {
		return fmt.Errorf("oracle parameter RewardBand must be bigger or equal than Voteperiod")
	}
	if params.VotesWindow <= 10 {
		return fmt.Errorf("oracle parameter VotesWindow must be > 0, is %d", params.VotesWindow)
	}
	if params.SlashFraction.GT(sdk.NewDecWithPrec(1, 2)) || params.SlashFraction.IsNegative() {
		return fmt.Errorf("oracle parameter SlashFraction must be smaller or equal than 1 percent and positive")
	}
	if params.MinValidVotesPerWindow.IsNegative() || params.MinValidVotesPerWindow.GT(sdk.OneDec()) {
		return fmt.Errorf("Min valid votes per window should be less than or equal to one and greater than zero, is %s", params.MinValidVotesPerWindow.String())
	}
	return nil
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of oracle module's parameters.
// nolint
func (params *Params) ParamSetPairs() subspace.ParamSetPairs {
	return subspace.ParamSetPairs{
		{Key: ParamStoreKeyVotePeriod, Value: &params.VotePeriod},
		{Key: ParamStoreKeyVoteThreshold, Value: &params.VoteThreshold},
		{Key: ParamStoreKeyRewardBand, Value: &params.RewardBand},
		{Key: ParamStoreKeyRewardDistributionPeriod, Value: &params.RewardDistributionPeriod},
		{Key: ParamStoreKeyVotesWindow, Value: &params.VotesWindow},
		{Key: ParamStoreKeyMinValidVotesPerWindow, Value: &params.MinValidVotesPerWindow},
		{Key: ParamStoreKeySlashFraction, Value: &params.SlashFraction},
		{Key: ParamStoreKeyWhitelist, Value: &params.Whitelist},
	}
}

// String implements fmt.Stringer interface
func (params Params) String() string {
	return fmt.Sprintf(`Treasury Params:
  VotePeriod:                  %d
  VoteThreshold:               %s
	RewardBand:                  %s
	RewardDistributionPeriod:    %d
	VotesWindow:                 %d
	MinValidVotesPerWindow:      %s
	SlashFraction:               %s
	Whitelist                    %s
	`, params.VotePeriod, params.VoteThreshold, params.RewardBand, params.RewardDistributionPeriod,
		params.VotesWindow, params.MinValidVotesPerWindow, params.SlashFraction, params.Whitelist)
}