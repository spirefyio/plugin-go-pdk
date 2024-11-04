// +build wasm

package pdk

import (
	pdk "github.com/extism/go-pdk"
)

//go:wasmimport extism:host/pluginengine GenerateFromTemplate
func generateFromTemplate(template, data uint64) uint64

// GenerateFromTemplate
//
// This function can be called by plugins to call one (or more) hooks matching the anchor id and version.
// If the slice of hooks is nil or empty, the first hook found is called if an anchor match is
// found.
//
// This is a wrapper function which uses the imported callHookForAnchor function implemented in the
// pluginengine plugin. This wrapper makes it easier for Go plugin developers to avoid the WASM memory management
func GenerateFromTemplate(template string, data []byte) ([]byte, error) {
	pdk.Log(pdk.LogDebug, "GenerateFromTemplate")

	dta1 := pdk.AllocateString(template)
	dta2 := pdk.AllocateBytes(data)

	off1 := dta1.Offset()
	off2 := dta2.Offset()

	resp := generateFromTemplate(off1, off2)

	mem1 := pdk.FindMemory(resp)
	bytes := mem1.ReadBytes()

	return bytes, nil
}
