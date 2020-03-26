package constant

type PendingPointProcessedType uint

const (
	PendingPointProcessedTypeUnprocessed PendingPointProcessedType = iota
	PendingPointProcessedTypeByReversal
	PendingPointProcessedTypeBySalesDraft
	PendingPointProcessedTypeByBlacklistedMerchant
)

type PointUpdatedFromType uint

const (
	PointUpdatedFromTypeVisa PointUpdatedFromType = iota + 1
	PointUpdatedFromTypeAdminOperation
	PointUpdatedFromTypeToWallet
	PointUpdatedFromTypeGift
	PointUpdatedFromTypeGovCashless
	PointUpdatedFromTypeQuicPayPlus
)
