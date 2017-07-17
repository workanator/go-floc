package floc

// Run runs the job with the given environment. The only purpose of
// the function at the moment is to make code more expressive. In future
// releases the function may have additional functionality.
func Run(flow Flow, state State, update Update, job Job) {
	job(flow, state, update)
}
