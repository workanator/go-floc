package floc

import "context"

/*
Context provides safe access to underlying context.
 */
type Context interface {
	Get() context.Context
	Update(ctx context.Context)
}
