package cond

import floc "github.com/workanator/go-floc"

// Not returns the negated value of predicate.
func Not(predicate floc.Predicate) floc.Predicate {
	return func(flow floc.Flow, state floc.State) bool {
		return !predicate(flow, state)
	}
}
