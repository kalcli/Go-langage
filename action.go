package actions

import (
	"fmt"
	"runtime"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Go version : %s\n", runtime.version())
	fmt.Println("GOOS: %s\n", runtime.GOOS)
	fmt.Println("GOARCH: %s\n",runtime.GOARCH)

	fmt.Println(quote.Go())
}