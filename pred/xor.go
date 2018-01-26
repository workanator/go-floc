package pred

import (
	"gopkg.in/devishot/go-floc.v2"
)

// Xor returns a predicate which chains multiple predicates into a condition
// with XOR logic. The result predicate finishes calculation of
// the condition as fast as the result is known.
//
// The result predicate tests the condition as follows.
//   (([PRED_1] XOR [PRED_2]) ... XOR [PRED_N])
func Xor(predicates ...floc.Predicate) floc.Predicate {
	count := len(predicates)
	if count > 2 {
		// More than 2 predicates
		return func(ctx floc.Context) bool {
			result := predicates[0](ctx) != predicates[1](ctx)

			for i := 2; i < count; i++ {
				result = result != predicates[i](ctx)
			}

			return result
		}
	} else if count == 2 {
		// 2 predicates
		return func(ctx floc.Context) bool {
			return predicates[0](ctx) != predicates[1](ctx)
		}
	}

	// Require at least 2 predicates
	panic("Xor requires at least 2 predicates")
}
