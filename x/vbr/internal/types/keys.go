package types

const (
	ModuleName   = "vbr"
	StoreKey     = ModuleName
	QuerierRoute = ModuleName

	QueryBlockRewardsPoolFunds = "funds"

	MsgTypeIncrementBlockRewardsPool = "incrementBlockRewardsPool"

	PoolStoreKey       = StoreKey + ":pool:"
	YearlyPoolStoreKey = StoreKey + ":yearly_pool:"
	YearNumberStoreKey = StoreKey + ":year_number:"
)
