package message

type SuccesfulLogin struct {
	Rank    int
	Flagged bool
}

type FailedLogin struct {
	ResponseCode int
}
