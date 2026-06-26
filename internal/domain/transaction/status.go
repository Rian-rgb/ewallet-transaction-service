package transaction

type Status string

const (
	Pending  Status = "PENDING"
	Success  Status = "SUCCESS"
	Failed   Status = "FAILED"
	Reversed Status = "REVERSED"
)
