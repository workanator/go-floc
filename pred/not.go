package pred

import floc "gopkg.in/workanator/go-floc.v1"

// Not returns the negated value of the predicate.
//
// The result predicate tests the condition as follows.
//   NOT [PRED]
func Not(predicate floc.Predicate) floc.Predicate {
	return func(state floc.State) bool {
		return !predicate(state)
	}
}
