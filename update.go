package floc

// Update is the function which may be invoked by job to update
// the state and/or the flow. It's up to the direct implementation how to
// interpret key and value.
type Update func(flow Flow, state State, key string, value interface{})
