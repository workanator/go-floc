package floc

// Run runs the job with the given environment. The only purpose of
// the function is to make code more expressive.
func Run(flow Flow, state State, update Update, job Job) {
	job(flow, state, update)
}
