package cond

import floc "github.com/workanator/go-floc"

// Or returns a predicate which tests multiple predicates with OR logics.
func Or(predicates ...floc.Predicate) floc.Predicate {
	return func(flow floc.Flow, state floc.State) bool {
		for _, predicate := range predicates {
			if predicate(flow, state) {
				return true
			}
		}

		return false
	}
}
