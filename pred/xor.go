package pred

import floc "github.com/workanator/go-floc"

// Xor returns a predicate which chains multiple predicates into a contidion
// with XOR logics. The result predicate finishes calculation of
// the condition as fast as the result is known.
//
// The result predicate tests the condition as follows.
//   [PRED_1] XOR ... XOR [PRED_N]
func Xor(predicates ...floc.Predicate) floc.Predicate {
	// Require at least 2 predicates
	if len(predicates) < 2 {
		panic("Xor requires at least 2 predicates")
	}

	count := len(predicates)

	return func(state floc.State) bool {
		result := predicates[0](state)
		for i := 1; i < count; i++ {
			result = (result != predicates[i](state))
		}

		return result
	}
}
