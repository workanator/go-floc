package pred

import floc "github.com/workanator/go-floc.v1"

func alwaysTrue(state floc.State) bool {
	return true
}

func alwaysFalse(state floc.State) bool {
	return false
}
