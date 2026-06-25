package transaction

type Status string

const (
	Pending  Status = "PENDING"
	Success  Status = "SUCCESS"
	Failed   Status = "FAILED"
	Reversed Status = "REVERSED"
)

var StatusFlow = map[Status][]Status{
	Pending: {Success, Failed},
	Success: {Reversed},
	Failed:  {Success},
}

func (s Status) CanTransitionTo(next Status) bool {
	allowedNext, ok := StatusFlow[s]
	if !ok {
		return false
	}

	for _, v := range allowedNext {
		if v == next {
			return true
		}
	}

	return false
}
