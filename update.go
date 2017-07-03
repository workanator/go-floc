package floc

// Update is the function which may be invoked by job any time to update
// the state or the flow.
type Update func(flow Flow, state State, key string, value interface{})
