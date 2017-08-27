package floc

/*
Data keeps shared data through execution of flow.
 */
type Data {
	// GetRead returns the contained data for read operation. The locker returned is already in locked state
	// so the caller is responsible for releasing the lock by calling `Unlock()`.
	//
	//     d, locker := data.GetRead()
	//     defer locker.Unlock()
	GetRead() (data interface{}, locker sync.Locker)

	// GetWrite returns the contained data for read and write operations. The locker returned is already in locked state
	// so the caller is responsible for releasing the lock by calling `Unlock()`.
	GetWrite() (data interface{}, locker sync.Locker)
}
