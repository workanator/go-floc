package cond

import floc "github.com/workanator/go-floc"

// And returns a predicate which tests multiple predicates with AND logics.
func And(predicates ...floc.Predicate) floc.Predicate {
	return func(flow floc.Flow, state floc.State) bool {
		if len(predicates) == 0 {
			return false
		}

		for _, predicate := range predicates {
			if !predicate(flow, state) {
				return false
			}
		}

		return true
	}
}
