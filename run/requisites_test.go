package run

import floc "github.com/workanator/go-floc"

func getCounter(state floc.State) int {
	data, lock := state.Get()
	counter := data.(*int)

	lock.Lock()
	defer lock.Unlock()

	return *counter
}

func updateCounter(flow floc.Flow, state floc.State, key string, value interface{}) {
	data, lock := state.GetExclusive()
	counter := data.(*int)

	lock.Lock()
	defer lock.Unlock()

	*counter += value.(int)
}

func jobIncrement(flow floc.Flow, state floc.State, update floc.Update) {
	update(flow, state, "", 1)
}

func predCounterEquals(n int) floc.Predicate {
	return func(state floc.State) bool {
		return getCounter(state) == n
	}
}
