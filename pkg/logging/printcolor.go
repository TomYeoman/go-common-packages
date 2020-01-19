package logging

import "fmt"

// Log any message
func PrintString(s string) {
	fmt.Print("YOU'VE CALLED LOGGER" + s)
}
