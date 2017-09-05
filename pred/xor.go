package pred

import floc "gopkg.in/workanator/go-floc.v1"

// Xor returns a predicate which chains multiple predicates into a contidion
// with XOR logics. The result predicate finishes calculation of
// the condition as fast as the result is known.
//
// The result predicate tests the condition as follows.
//   (([PRED_1] XOR [PRED_2]) ... XOR [PRED_N])
func Xor(predicates ...floc.Predicate) floc.Predicate {
	count := len(predicates)
	if count > 2 {
		// More than 2 predicates
		return func(state floc.State) bool {
			result := predicates[0](state) != predicates[1](state)

			for i := 2; i < count; i++ {
				result = (result != predicates[i](state))
			}

			return result
		}
	} else if count == 2 {
		// 2 predicates
		return func(state floc.State) bool {
			return predicates[0](state) != predicates[1](state)
		}
	}

	// Require at least 2 predicates
	panic("Xor requires at least 2 predicates")
}
