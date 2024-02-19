package powc

/*
#cgo LDFLAGS: -lic -L. -lm
#include <ic.h>
*/
import "C"
import "unsafe"

func Fpgenc32(bufferIn []float32, bufferOut []byte) int64 {
	written := C.fpgenc32((*C.uint)(unsafe.Pointer(&bufferIn[0])), C.size_t(len(bufferIn)), (*C.uchar)(&bufferOut[0]), 0)
	return int64(written)
}

// size_t fpgdec32(    unsigned char *in, size_t n, uint32_t      *out, uint32_t start);
func Fpgdec32(bufferIn []byte, size int, bufferOut []float32) int64 {
	written := C.fpgdec32((*C.uchar)(unsafe.Pointer(&bufferIn[0])), C.size_t(size), (*C.uint)(unsafe.Pointer(&bufferOut[0])), 0)
	return int64(written)
}
func Fpc0enc32(bufferIn []float32, bufferOut []byte) int64 {
	written := C.fpc0enc32((*C.uint)(unsafe.Pointer(&bufferIn[0])), C.size_t(len(bufferIn)), (*C.uchar)(&bufferOut[0]), 0)
	return int64(written)
}
func Fpfcmenc32(bufferIn []float32, bufferOut []byte) int64 {
	written := C.fpfcmenc32((*C.uint)(unsafe.Pointer(&bufferIn[0])), C.size_t(len(bufferIn)), (*C.uchar)(&bufferOut[0]), 0)
	return int64(written)
}
func Fpdfcmenc32(bufferIn []float32, bufferOut []byte) int64 {
	written := C.fpdfcmenc32((*C.uint)(unsafe.Pointer(&bufferIn[0])), C.size_t(len(bufferIn)), (*C.uchar)(&bufferOut[0]), 0)
	return int64(written)
}

func P4nzzenc128v32(bufferIn []float32, bufferOut []byte) int64 {
	written := C.p4nzzenc128v32((*C.uint)(unsafe.Pointer(&bufferIn[0])), C.size_t(len(bufferIn)), (*C.uchar)(&bufferOut[0]), 0)
	return int64(written)
}
