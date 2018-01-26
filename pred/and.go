package pred

import (
	"gopkg.in/devishot/go-floc.v2"
)

// And returns a predicate which chains multiple predicates into a condition
// with AND logic. The result predicate finishes calculation of
// the condition as fast as the result is known. The function panics if
// the number of predicates is less than 2.
//
// The result predicate tests the condition as follows.
//   [PRED_1] AND ... AND [PRED_N]
func And(predicates ...floc.Predicate) floc.Predicate {
	count := len(predicates)
	if count > 2 {
		// More than 2 predicates
		return func(ctx floc.Context) bool {
			for _, p := range predicates {
				if !p(ctx) {
					return false
				}
			}

			return true
		}
	} else if count == 2 {
		// 2 predicates
		return func(ctx floc.Context) bool {
			return predicates[0](ctx) && predicates[1](ctx)
		}
	}

	// Require at least 2 predicates
	panic("And requires at least 2 predicates")
}
