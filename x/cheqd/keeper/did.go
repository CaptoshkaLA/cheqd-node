package keeper

import (
	"github.com/cheqd/cheqd-node/x/cheqd/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"
)

// GetDidCount get the total number of did
func (k Keeper) GetDidCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidCountKey))
	byteKey := types.KeyPrefix(types.DidCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetDidCount set the total number of did
func (k Keeper) SetDidCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidCountKey))
	byteKey := types.KeyPrefix(types.DidCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendDid appends a did in the store with a new id and update the count
func (k Keeper) AppendDid(ctx sdk.Context, did types.Did, metadata *types.Metadata, timestamp string, txHash string) string {
	// Create the did
	count := k.GetDidCount(ctx)
	k.SetDid(ctx, did, metadata, timestamp, txHash)

	// Update did count
	k.SetDidCount(ctx, count+1)
	return did.Id
}

// SetDid set a specific did in the store
func (k Keeper) SetDid(ctx sdk.Context, did types.Did, metadata *types.Metadata, timestamp string, txHash string) {
	stateValue := types.StateValue{
		Metadata: metadata,
		Data: &types.StateValue_Did{
			Did: &did,
		},
		Timestamp: timestamp,
		TxHash:    txHash,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	b := k.cdc.MustMarshal(&stateValue)
	store.Set(GetDidIDBytes(did.Id), b)
}

// GetDid returns a did from its id
func (k Keeper) GetDid(ctx *sdk.Context, id string) (*types.StateValue, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))

	var value types.StateValue
	if err := k.cdc.Unmarshal(store.Get(GetDidIDBytes(id)), &value); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, err.Error())
	}

	return &value, nil
}

// HasDid checks if the did exists in the store
func (k Keeper) HasDid(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	return store.Has(GetDidIDBytes(id))
}

// GetDidIDBytes returns the byte representation of the ID
func GetDidIDBytes(id string) []byte {
	return []byte(id)
}
