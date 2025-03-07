//go:build std
// +build std
// +build wasm


package pdk

import (
	"errors"
	"fmt"

	pdk "github.com/extism/go-pdk"
)

//go:wasmimport extism:host/pluginengine WriteFile
func writeFile(path uint64, contents uint64) uint64

// WriteFile
//
// This is a wrapper function which uses the imported writeFile function implemented in the
// pluginengine as a host function. This wrapper makes it easier for Go plugin developers to 
// avoid the WASM memory management
func WriteFile(path string, contents []byte) error {
	pdk.Log(pdk.LogDebug, "WriteFile")

	// allocate the memory for the string
	dta1 := pdk.AllocateString(path)
	// get the offset
	off1 := dta1.Offset()

	dta2 := pdk.AllocateBytes(contents)
	off2 := dta2.Offset()

	// call the imported host function with the off1 and get its response offset
	resp := writeFile(off1, off2)

	pdk.Log(pdk.LogDebug, "Return resp from writing file: " + fmt.Sprint(resp))
	// returned from the call to the imported readFile, so lets grab its memory that stores the file data

	if resp != 0 {
		return errors.New("Problem with writing to file")
	}
	
	return nil
}

