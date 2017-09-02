package floc

/*
Releaser is responsible for releasing underlying resources.
*/
type Releaser interface {
	// Release should be called once when the object is not needed anymore.
	Release()
}
