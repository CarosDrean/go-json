//go:build !race
// +build !race

package encoder

func CompileToGetCodeSet(ctx *RuntimeContext, typeptr uintptr, customTag string) (*OpcodeSet, error) {
	if typeptr > typeAddr.MaxTypeAddr || typeptr < typeAddr.BaseTypeAddr {
		codeSet, err := compileToGetCodeSetSlowPath(typeptr, customTag)
		if err != nil {
			return nil, err
		}
		return getFilteredCodeSetIfNeeded(ctx, codeSet)
	}
	index := (typeptr - typeAddr.BaseTypeAddr) >> typeAddr.AddrShift
	if codeSet := cachedOpcodeSets[index]; codeSet != nil {
		filtered, err := getFilteredCodeSetIfNeeded(ctx, codeSet)
		if err != nil {
			return nil, err
		}
		return filtered, nil
	}
	codeSet, err := newCompiler().compile(typeptr, customTag)
	if err != nil {
		return nil, err
	}
	filtered, err := getFilteredCodeSetIfNeeded(ctx, codeSet)
	if err != nil {
		return nil, err
	}
	cachedOpcodeSets[index] = codeSet
	return filtered, nil
}
