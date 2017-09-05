package run

import floc "gopkg.in/workanator/go-floc.v1"

func getCounter(state floc.State) int {
	data, locker := state.DataWithReadLocker()
	counter := data.(*int)

	locker.Lock()
	defer locker.Unlock()

	return *counter
}

func updateCounter(flow floc.Flow, state floc.State, key string, value interface{}) {
	data, locker := state.DataWithWriteLocker()
	counter := data.(*int)

	locker.Lock()
	defer locker.Unlock()

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
