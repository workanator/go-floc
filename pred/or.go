package pred

import floc "github.com/workanator/go-floc"

// Or returns a predicate which chains multiple predicates into a contidion
// with OR logics. The result predicate finishes calculation of
// the condition as fast as the result is known.
//
// The result predicate tests the condition as follows.
//   [PRED_1] OR ... OR [PRED_N]
func Or(predicates ...floc.Predicate) floc.Predicate {
	// Require at least 2 predicates
	if len(predicates) < 2 {
		panic("Or requires at least 2 predicates")
	}

	return func(state floc.State) bool {
		for _, predicate := range predicates {
			if predicate(state) {
				return true
			}
		}

		return false
	}
}
