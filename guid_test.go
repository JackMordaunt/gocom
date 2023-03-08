package main

import (
	"testing"
)

func TestNewCLSIDFromString(t *testing.T) {
	testCases := []struct {
		input    string
		expected GUID
	}{
		{"{00000000-0000-0000-C000-000000000046}", GUID{Data1: 0x00000000, Data2: 0x0000, Data3: 0x0000, Data4: [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}},
		{"{11111111-1111-1111-1111-111111111111}", GUID{Data1: 0x11111111, Data2: 0x1111, Data3: 0x1111, Data4: [8]byte{0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11}}},
		{"{22222222-2222-2222-2222-222222222222}", GUID{Data1: 0x22222222, Data2: 0x2222, Data3: 0x2222, Data4: [8]byte{0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22}}},
	}

	for _, tc := range testCases {
		got, err := NewGUIDFromString(tc.input)
		if err != nil {
			t.Errorf("NewGUIDFromString(%q) returned an error: %v", tc.input, err)
		}

		if got != tc.expected {
			t.Errorf("NewGUIDFromString(%q) = %v, expected %v", tc.input, got, tc.expected)
		}
	}
}
