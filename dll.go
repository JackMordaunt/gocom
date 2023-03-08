package main 

import "C" 

import "unsafe"
import "fmt"
import "os"

type HRESULT int32

func (hr HRESULT) Error() string {
    return fmt.Sprintf("HRESULT 0x%X", int32(hr))
}

/*
	DLL entry points.
*/

//export DllGetClassObject
func DllGetClassObject(rclsid uintptr, riid uintptr, out uintptr) HRESULT {
	// This function should return the IClassFactory object that can build an ICNS object. 
	clsid := (*GUID)(unsafe.Pointer(rclsid))
	iid := (*GUID)(unsafe.Pointer(riid))

	outf, err := os.Open("dll-test-log.txt")
	if err != nil {
		return -1
	}

	defer outf.Close()

	if _, err := fmt.Fprintf(outf, "CLSID %v\n", *clsid); err != nil {
		return -1
	}
	if _, err := fmt.Fprintf(outf, "IID %v\n", *iid); err != nil {
		return -1
	}
	
	return -1
}

//export DllCanUnloadNow
func DllCanUnloadNow() HRESULT {
	return 0
}

//export DllRegisterServer
func DllRegisterServer() HRESULT {
	return 0
}

//export DllUnregisterServer
func DllUnregisterServer() HRESULT {
	return 0
}