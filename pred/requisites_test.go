package pred

import floc "github.com/workanator/go-floc"

func alwaysTrue(flow floc.Flow, state floc.State) bool {
	return true
}

func alwaysFalse(flow floc.Flow, state floc.State) bool {
	return false
}
