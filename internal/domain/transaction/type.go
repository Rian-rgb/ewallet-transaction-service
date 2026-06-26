package transaction

type Type string

const (
	Topup    Type = "TOPUP"
	Purchase Type = "PURCHASE"
	Refund   Type = "REFUND"
)
