package pred

import floc "github.com/workanator/go-floc"

// Not returns the negated value of the predicate.
//
// The result predicate tests the condition as follows.
//   NOT [PRED]
func Not(predicate floc.Predicate) floc.Predicate {
	return func(flow floc.Flow, state floc.State) bool {
		return !predicate(flow, state)
	}
}
