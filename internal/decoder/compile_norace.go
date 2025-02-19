//go:build !race
// +build !race

package decoder

import (
	"unsafe"

	"github.com/CarosDrean/go-json/internal/runtime"
)

func CompileToGetDecoder(typ *runtime.Type, customTag string) (Decoder, error) {
	typeptr := uintptr(unsafe.Pointer(typ))
	if typeptr > typeAddr.MaxTypeAddr {
		return compileToGetDecoderSlowPath(typeptr, customTag, typ)
	}

	index := (typeptr - typeAddr.BaseTypeAddr) >> typeAddr.AddrShift
	if dec := cachedDecoder[index]; dec != nil {
		return dec, nil
	}

	dec, err := compileHead(typ, customTag, map[uintptr]Decoder{})
	if err != nil {
		return nil, err
	}
	cachedDecoder[index] = dec
	return dec, nil
}
