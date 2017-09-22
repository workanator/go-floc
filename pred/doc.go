/*
Package pred provides predicates for basic logic operations.

Predicates with conditional jobs like run.If allows to make non-linear
flow. In terms of floc predicate should return true or false depending
on context.

  const idReadyFlag = 100

  testReady := func(ctx floc.Context) bool {
	if flag := ctx.Value(idReadyFlag); flag != nil {
      return flag.(bool)
	}

    return false
  }

  flow := run.Sequence(
    ..., // Some job done here
    job.If(testReady, job.Background(writeToDisk)), // Write some data ready to disk in background
    ..., // Some job more
  )
*/
package pred
