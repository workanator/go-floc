/*
Package state contains implementations of floc.State interface.

The default implementation allows the state to contain any arbitrary data.
So data can either be of primitive type or complex structure or even
interface or function. What the state should contain depends on task.

  type Events struct {
    HeaderReady bool
    BodyReady bool
    DataReady bool
  }

  theState := state.New(new(Events))
*/
package state
