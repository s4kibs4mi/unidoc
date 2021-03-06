/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package core

import (
	"testing"
)

func TestHexStringWriteBasic(t *testing.T) {
	testcases := map[string]string{
		" ": "<20>",
	}

	for src, expected := range testcases {
		strObj := MakeHexString(src)
		ws := strObj.DefaultWriteString()

		if ws != expected {
			t.Fatalf("%s: '%s' != '%s'\n", src, ws, expected)
		}
	}
}

// Test writing and parsing back of hexadecimal and regular strings.
func TestHexStringMulti(t *testing.T) {
	testcases := []string{
		"This is a string",
		"Strings may contain\n newlines and such",
		string([]byte{0x50, 0x01, 0x00, 0x90, 0xff, 0x49, 0xdf, 0x20, 0x32}),
		"",
	}

	for _, testcase := range testcases {
		// Make *PdfObject representations for regular and hexadecimal strings.
		s := MakeString(testcase)
		shex := MakeHexString(testcase)

		// Write out.
		writestr := s.DefaultWriteString()
		writestrhex := shex.DefaultWriteString()

		// Parse back.
		parser1 := makeParserForText(writestr)
		parser2 := makeParserForText(writestrhex)

		// Check that representation is correct.
		obj1, err := parser1.parseObject()
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		strObj1, ok := obj1.(*PdfObjectString)
		if !ok {
			t.Fatalf("Type incorrect")
		}
		if strObj1.isHex != false {
			t.Fatalf("Should not be hex")
		}
		if strObj1.Str() != testcase {
			t.Fatalf("String mismatch")
		}

		obj2, err := parser2.parseObject()
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		strObj2, ok := obj2.(*PdfObjectString)
		if !ok {
			t.Fatalf("Type incorrect")
		}
		if strObj2.isHex != true {
			t.Fatalf("Should be hex")
		}
		if strObj2.Str() != testcase {
			t.Fatalf("String mismatch")
		}
	}
}
