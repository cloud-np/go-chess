package core

// We make our own assert to simplify the code
func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}
