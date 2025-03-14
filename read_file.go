// +build wasm

package pluginpdk

import (
	pdk "github.com/extism/go-pdk"
)

//go:wasmimport extism:host/pluginengine ReadFile
func readFile(path uint64) uint64

// ReadFile
//
//
// This is a wrapper function which uses the imported callExtensionPointExtensions function implemented in the
// pluginengine plugin. This wrapper makes it easier for Go plugin developers to avoid the WASM memory management
func ReadFile(path string) ([]byte, error) {
	pdk.Log(pdk.LogDebug, "ReadFile")

	// allocate the memory for the string
	dta1 := pdk.AllocateString(path)
	// get the offset
	off1 := dta1.Offset()

	// call the imported host function with the off1 and get its response offset
	resp := readFile(off1)

	// returned from the call to the imported readFile, so lets grab its memory that stores the file data
	mem1 := pdk.FindMemory(resp)

	// get the actual []bytes
	filedata := mem1.ReadBytes()

	// all done return it
	// TODO: Do we need to free it here.. above with a defer() or let the calling plugin free it somehow?
	return filedata, nil
}
