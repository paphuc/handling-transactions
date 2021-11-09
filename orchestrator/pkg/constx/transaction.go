package constx

import "time"

const (
	RollBack       = "R"
	Commit         = "C"
	True           = "T"
	False          = "F"
	ExpiredMinutes = time.Second * 30
)
