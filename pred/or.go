package pred

import (
	"gopkg.in/devishot/go-floc.v2"
)

// Or returns a predicate which chains multiple predicates into a condition
// with OR logic. The result predicate finishes calculation of
// the condition as fast as the result is known.
//
// The result predicate tests the condition as follows.
//   [PRED_1] OR ... OR [PRED_N]
func Or(predicates ...floc.Predicate) floc.Predicate {
	count := len(predicates)
	if count > 2 {
		// More than 2 predicates
		return func(ctx floc.Context) bool {
			for _, p := range predicates {
				if p(ctx) {
					return true
				}
			}

			return false
		}
	} else if count == 2 {
		// 2 predicates
		return func(ctx floc.Context) bool {
			return predicates[0](ctx) || predicates[1](ctx)
		}
	}

	// Require at least 2 predicates
	panic("Or requires at least 2 predicates")
}
