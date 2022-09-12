package encoder

import (
	"testing"
	"unsafe"

	"github.com/CarosDrean/go-json/internal/defaults"
)

func TestDumpOpcode(t *testing.T) {
	ctx := TakeRuntimeContext()
	defer ReleaseRuntimeContext(ctx)
	var v interface{} = 1
	header := (*emptyInterface)(unsafe.Pointer(&v))
	typ := header.typ
	typeptr := uintptr(unsafe.Pointer(typ))
	codeSet, err := CompileToGetCodeSet(ctx, typeptr, defaults.DefaultTag)
	if err != nil {
		t.Fatal(err)
	}
	codeSet.EscapeKeyCode.Dump()
}
