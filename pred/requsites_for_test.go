package pred

import "gopkg.in/workanator/go-floc.v2"

func alwaysTrue(floc.Context) bool {
	return true
}

func alwaysFalse(floc.Context) bool {
	return false
}
