package entities

import (
	"gorm-canban/constant"
	"time"
	//uuid "github.com/satori/go.uuid"
)

type Wallet struct {
	ID                  uint64     `json:"-"`
	CurrencyID          uint8      `json:"-"`
	UserID              uint64     `gorm:"unique" json:"-"`
	BalanceAccumulative uint       `json:"balanceAccumulative"`
	BalanceCurrent      uint       `json:"balanceCurrent"`
	BalanceKYCApproach  uint       `json:"balanceKYCApproach"` //CHange this to `json:"-"` in production
	NegativeAmount      int        `json:"negative_amount"`
	CreatedAt           time.Time  `json:"-"`
	UpdatedAt           time.Time  `json:"updatedAt"`
	DeletedAt           *time.Time `json:"-"`
}

type PointPendingTransaction struct {
	ID            uint64
	WalletID      uint64
	Point         int
	Rate          float64
	ProcessedType constant.PendingPointProcessedType
	Type          constant.PointUpdatedFromType
	TypeID        uint64
	CreatedAt     time.Time
	UpdatedAt     time.Time

}

type PointAvailableTransaction struct {
	ID            uint64
	WalletID      uint64
	Point         int
	BeforePoint   int
	AfterPoint    int
	Type          constant.PointUpdatedFromType
	TypeID        uint64
	OriginalID    *uint64
	//IdentifierAPI UUID                          `gorm:"type:bytea;unique"`
	CreatedAt     time.Time

//	IssuedCreditCardTransaction          *IssuedCreditCardTransaction          `gorm:"foreignkey:TypeID"`
//	IssuedCreditCardTransactionReceiveds []IssuedCreditCardTransactionReceived `gorm:"foreignkey:IssuedCreditCardTransactionID;association_foreignkey:TypeID"`
}

