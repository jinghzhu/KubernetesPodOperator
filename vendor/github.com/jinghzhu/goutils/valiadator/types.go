package valiadator

import "regexp"

const (
	Int string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
)

var (
	rxInt = regexp.MustCompile(Int)
)
