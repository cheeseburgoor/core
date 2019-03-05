package pay

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// Keeper manages transfers between accounts. It implements the Keeper interface.
type Keeper struct {
	bank.Keeper

	key sdk.StoreKey
	cdc *codec.Codec

	ak auth.AccountKeeper
	fk auth.FeeCollectionKeeper
	//paramSpace params.Subspace
}

// NewKeeper returns a new Keeper
func NewKeeper(
	key sdk.StoreKey,
	cdc *codec.Codec,
	ak auth.AccountKeeper,
	fk auth.FeeCollectionKeeper) Keeper {
	return Keeper{
		key: key,
		cdc: cdc,
		ak:  ak,
		fk:  fk,
	}
}

// SendCoins moves coins from one account to another
func (keeper Keeper) SendCoins(
	ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins,
) (sdk.Tags, sdk.Error) {

	taxes := calculateTaxes(ctx, keeper, amt)
	_, taxTags, err := subtractCoins(ctx, keeper, fromAddr, taxes)
	if err != nil {
		return nil, err
	}
	keeper.fk.AddCollectedFees(ctx, taxes)
	keeper.recordTaxProceeds(ctx, taxes)

	_, subTags, err := subtractCoins(ctx, keeper, fromAddr, amt)
	if err != nil {
		return nil, err
	}

	_, addTags, err := addCoins(ctx, keeper, toAddr, amt)
	if err != nil {
		return nil, err
	}

	taxTags = taxTags.AppendTags(subTags)

	return taxTags.AppendTags(addTags), nil
}

// InputOutputCoins handles a list of inputs and outputs
func (keeper Keeper) InputOutputCoins(ctx sdk.Context, inputs []bank.Input, outputs []bank.Output) (sdk.Tags, sdk.Error) {
	allTags := sdk.EmptyTags()

	for _, in := range inputs {
		_, tags, err := subtractCoins(ctx, keeper, in.Address, in.Coins)
		if err != nil {
			return nil, err
		}
		allTags = allTags.AppendTags(tags)
	}

	for _, out := range outputs {

		taxes := calculateTaxes(ctx, keeper, out.Coins)
		_, taxTags, err := subtractCoins(ctx, keeper, out.Address, taxes)
		if err != nil {
			return nil, err
		}
		keeper.fk.AddCollectedFees(ctx, taxes)
		keeper.recordTaxProceeds(ctx, taxes)
		allTags = allTags.AppendTags(taxTags)

		_, tags, err := addCoins(ctx, keeper, out.Address, out.Coins)
		if err != nil {
			return nil, err
		}
		allTags = allTags.AppendTags(tags)
	}

	return allTags, nil
}

func getCoins(ctx sdk.Context, k Keeper, addr sdk.AccAddress) sdk.Coins {
	acc := k.ak.GetAccount(ctx, addr)
	if acc == nil {
		return sdk.Coins{}
	}
	return acc.GetCoins()
}

func setCoins(ctx sdk.Context, k Keeper, addr sdk.AccAddress, amt sdk.Coins) sdk.Error {
	acc := k.ak.GetAccount(ctx, addr)
	if acc == nil {
		acc = k.ak.NewAccountWithAddress(ctx, addr)
	}
	err := acc.SetCoins(amt)
	if err != nil {
		// Handle w/ #870
		panic(err)
	}
	k.ak.SetAccount(ctx, acc)
	return nil
}

// HasCoins returns whether or not an account has at least amt coins.
func hasCoins(ctx sdk.Context, k Keeper, addr sdk.AccAddress, amt sdk.Coins) bool {
	return getCoins(ctx, k, addr).IsAllGTE(amt)
}

func getAccount(ctx sdk.Context, k Keeper, addr sdk.AccAddress) auth.Account {
	return k.ak.GetAccount(ctx, addr)
}

func setAccount(ctx sdk.Context, k Keeper, acc auth.Account) {
	k.ak.SetAccount(ctx, acc)
}

// subtractCoins subtracts amt coins from an account with the given address addr.
//
// CONTRACT: If the account is a vesting account, the amount has to be spendable.
func subtractCoins(ctx sdk.Context, k Keeper, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, sdk.Tags, sdk.Error) {
	oldCoins, spendableCoins := sdk.Coins{}, sdk.Coins{}

	acc := getAccount(ctx, k, addr)
	if acc != nil {
		oldCoins = acc.GetCoins()
		spendableCoins = acc.SpendableCoins(ctx.BlockHeader().Time)
	}

	// For non-vesting accounts, spendable coins will simply be the original coins.
	// So the check here is sufficient instead of subtracting from oldCoins.
	_, hasNeg := spendableCoins.SafeMinus(amt)
	if hasNeg {
		return amt, nil, sdk.ErrInsufficientCoins(
			fmt.Sprintf("insufficient account funds; %s < %s", spendableCoins, amt),
		)
	}

	newCoins := oldCoins.Minus(amt) // should not panic as spendable coins was already checked
	err := setCoins(ctx, k, addr, newCoins)
	k.subtractIssuance(ctx, amt)
	tags := sdk.NewTags(bank.TagKeySender, addr.String())

	return newCoins, tags, err
}

// AddCoins adds amt to the coins at the addr.
func addCoins(ctx sdk.Context, k Keeper, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, sdk.Tags, sdk.Error) {
	oldCoins := getCoins(ctx, k, addr)
	newCoins := oldCoins.Plus(amt)

	if newCoins.IsAnyNegative() {
		return amt, nil, sdk.ErrInsufficientCoins(
			fmt.Sprintf("insufficient account funds; %s < %s", oldCoins, amt),
		)
	}

	err := setCoins(ctx, k, addr, newCoins)
	k.addIssuance(ctx, amt)
	tags := sdk.NewTags(bank.TagKeyRecipient, addr.String())

	return newCoins, tags, err
}

// SendCoins moves coins from one account to another
func sendCoins(ctx sdk.Context, k Keeper, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) (sdk.Tags, sdk.Error) {
	// Safety check ensuring that when sending coins the keeper must maintain the
	// supply invariant.
	if !amt.IsValid() {
		return nil, sdk.ErrInvalidCoins(amt.String())
	}

	_, subTags, err := subtractCoins(ctx, k, fromAddr, amt)
	if err != nil {
		return nil, err
	}

	_, addTags, err := addCoins(ctx, k, toAddr, amt)
	if err != nil {
		return nil, err
	}

	return subTags.AppendTags(addTags), nil
}

// InputOutputCoins handles a list of inputs and outputs
// NOTE: Make sure to revert state changes from tx on error
func inputOutputCoins(ctx sdk.Context, k Keeper, inputs []bank.Input, outputs []bank.Output) (sdk.Tags, sdk.Error) {
	// Safety check ensuring that when sending coins the keeper must maintain the
	// supply invariant.
	if err := bank.ValidateInputsOutputs(inputs, outputs); err != nil {
		return nil, err
	}

	allTags := sdk.EmptyTags()

	for _, in := range inputs {
		_, tags, err := subtractCoins(ctx, k, in.Address, in.Coins)
		if err != nil {
			return nil, err
		}
		allTags = allTags.AppendTags(tags)
	}

	for _, out := range outputs {
		_, tags, err := addCoins(ctx, k, out.Address, out.Coins)
		if err != nil {
			return nil, err
		}
		allTags = allTags.AppendTags(tags)
	}

	return allTags, nil
}