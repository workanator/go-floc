package floc

// Predicate is the function which calculates the result of some condition.
type Predicate func(state State) bool
