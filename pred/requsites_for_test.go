package pred

import "github.com/workanator/go-floc/v3"

func alwaysTrue(floc.Context) bool {
	return true
}

func alwaysFalse(floc.Context) bool {
	return false
}
