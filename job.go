package floc

/*
Job is the small piece of flow.
*/
type Job func(ctx Context, ctrl Control) error
