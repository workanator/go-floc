package floc

// Job is the proptotype of function which do some piece of job in control flow.
type Job func(flow Flow, state State, update Update)
