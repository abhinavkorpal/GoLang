// Your code will fail to compile if you import a package without using any of its exported functions, interfaces, structures, or variables.
// If you really need the imported package you can use the blank identifier, _, as its package name to avoid this compilation failure.
// The blank identifier is used to import packages for their side effects.

package main

import (
	_ "fmt"
	"log"
	"time"
)

var _ = log.Println

func main() {
	_ = time.Now
}
