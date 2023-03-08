package main

import "C"
import "unsafe"

func main() {

}

type IThumbnailProviderVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type ICNSProvider struct {
	IThumbnailProviderVtbl
}

//export QueryInterface
func QueryInterface(this uintptr, riid uintptr, out uintptr) HRESULT {
	// Write our object to the out pointer.
	*(**ICNSProvider)(unsafe.Pointer(out)) = (*ICNSProvider)(unsafe.Pointer(this))
	return -1
}
