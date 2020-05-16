package main

// throw calls panic when it means throwing exception (i.e. recovered later).
func throw(err error) {
	panic(err) // will be handled by mow.cli
}
