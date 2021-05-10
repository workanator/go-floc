package pred

import (
	"github.com/workanator/go-floc/v3"
)

// Not returns the negated value of the predicate.
//
// The result predicate tests the condition as follows.
//   NOT [PRED]
func Not(predicate floc.Predicate) floc.Predicate {
	return func(ctx floc.Context) bool {
		return !predicate(ctx)
	}
}
