package floc

// Job is the proptotype of function which do some piece of the overall work.
// With the parameters it has the implementation can control execution of flow
// and read/write state directly or with update function.
type Job func(flow Flow, state State, update Update)
