package main

import (
	"fmt"
	"unsafe"

	foobarbaz "unsafe"
)

type Fake struct{}

func (Fake) Good() {}
func main() {
	unsafeM := Fake{}
	unsafeM.Good()
	intArray := [...]int{1, 2}
	fmt.Printf("\nintArray: %v\n", intArray)
	intPtr := &intArray[0]
	fmt.Printf("\nintPtr=%p, *intPtr=%d.\n", intPtr, *intPtr)
	// bad
	addressHolder := uintptr(foobarbaz.Pointer(intPtr)) + unsafe.Sizeof(intArray[0])
	// bad
	intPtr = (*int)(foobarbaz.Pointer(addressHolder))
	fmt.Printf("\nintPtr=%p, *intPtr=%d.\n\n", intPtr, *intPtr)
}
