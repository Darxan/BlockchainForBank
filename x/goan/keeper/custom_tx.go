package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"goan/x/goan/types"
)

// GetCustomTxCount get the total number of customTx
func (k Keeper) GetCustomTxCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CustomTxCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCustomTxCount set the total number of customTx
func (k Keeper) SetCustomTxCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.CustomTxCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendCustomTx appends a customTx in the store with a new id and update the count
func (k Keeper) AppendCustomTx(
	ctx sdk.Context,
	customTx types.CustomTx,
) uint64 {
	// Create the customTx
	count := k.GetCustomTxCount(ctx)

	// Set the ID of the appended value
	customTx.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CustomTxKey))
	appendedValue := k.cdc.MustMarshal(&customTx)
	store.Set(GetCustomTxIDBytes(customTx.Id), appendedValue)

	// Update customTx count
	k.SetCustomTxCount(ctx, count+1)

	return count
}

// SetCustomTx set a specific customTx in the store
func (k Keeper) SetCustomTx(ctx sdk.Context, customTx types.CustomTx) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CustomTxKey))
	b := k.cdc.MustMarshal(&customTx)
	store.Set(GetCustomTxIDBytes(customTx.Id), b)
}

// GetCustomTx returns a customTx from its id
func (k Keeper) GetCustomTx(ctx sdk.Context, id uint64) (val types.CustomTx, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CustomTxKey))
	b := store.Get(GetCustomTxIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCustomTx removes a customTx from the store
func (k Keeper) RemoveCustomTx(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CustomTxKey))
	store.Delete(GetCustomTxIDBytes(id))
}

// GetAllCustomTx returns all customTx
func (k Keeper) GetAllCustomTx(ctx sdk.Context) (list []types.CustomTx) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CustomTxKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CustomTx
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCustomTxIDBytes returns the byte representation of the ID
func GetCustomTxIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetCustomTxIDFromBytes returns ID in uint64 format from a byte array
func GetCustomTxIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
