package main

import "fmt"

// CLSID is the class id representing the object itself.
var CLSID_ICNS = must(NewGUIDFromString("{CA7786E8-C694-49AF-AFC1-DB80EDBCEDAE}"))

// IID is the interface id representing the vtable.
var IID_ICNS = must(NewGUIDFromString("{2D03DA50-13A7-4166-AD3D-E3EAE6244C1C}"))

var IID_IThumbnailProvider = must(NewGUIDFromString("{E357FCCD-A995-4576-B01F-234630154E96}"))

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

func NewGUIDFromString(s string) (GUID, error) {
	// Remove the surrounding curly braces and convert to bytes
	b := make([]byte, 16)
	_, err := fmt.Sscanf(s, "{%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x}",
		&b[3], &b[2], &b[1], &b[0],
		&b[5], &b[4],
		&b[7], &b[6],
		&b[8], &b[9],
		&b[10], &b[11], &b[12], &b[13], &b[14], &b[15])
	if err != nil {
		return GUID{}, err
	}

	// Convert to a GUID
	var guid GUID
	copy(guid.Data4[:], b[8:])
	guid.Data1 = uint32(b[3])<<24 | uint32(b[2])<<16 | uint32(b[1])<<8 | uint32(b[0])
	guid.Data2 = uint16(b[5])<<8 | uint16(b[4])
	guid.Data3 = uint16(b[7])<<8 | uint16(b[6])
	return guid, nil
}
