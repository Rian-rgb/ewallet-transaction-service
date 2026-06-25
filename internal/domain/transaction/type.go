package transaction

type Type string

const (
	Topup    Type = "TOPUP"
	Purchase Type = "PURCHASE"
	Refund   Type = "REFUND"
)

func (t Type) IsValid() bool {
	switch t {
	case Topup, Purchase, Refund:
		return true
	default:
		return false
	}
}
