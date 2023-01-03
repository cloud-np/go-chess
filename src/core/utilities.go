package core

// We make our own assert to simplify the code
func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
