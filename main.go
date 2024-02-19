package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
	"unsafe"

	"example.com/float/powc"
	"github.com/ronanh/intcomp"
)

var width = 50
var height = 50
var days = 24
var data []float32 = make([]float32, width*height*days)
var decode = make([]float32, width*height*days)
var buffOut = make([]byte, width*height*days*4)

func init() {

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for z := 0; z < days; z++ {
				floatVal := float32(math.Sin((1. / 24.) * (float64(x)*float64(width) + (float64(y) * float64(height)) + float64(z))))
				data[x*width+y] = floatVal
			}
		}
	}
}

func main() {
	// fmt.Printf("%f\n", in)
	// written := powc.P4nzenc128v32(in, buffOut)
	start := time.Now()
	written := powc.Fpc0enc32(data, buffOut)
	fmt.Printf("Fpc0enc32:      data length: %d written bytes: %d time: %s\n", len(data)*4, written, time.Since(start))

	start = time.Now()
	written = powc.Fpdfcmenc32(data, buffOut)
	fmt.Printf("Fpdfcmenc32:    data length: %d written bytes: %d time: %s\n", len(data)*4, written, time.Since(start))

	start = time.Now()
	written = powc.Fpfcmenc32(data, buffOut)
	fmt.Printf("Fpfcmenc32:     data length: %d written bytes: %d time: %s\n", len(data)*4, written, time.Since(start))

	start = time.Now()
	written = powc.Fpgenc32(data, buffOut)
	fmt.Printf("Fpgenc32:       data length: %d written bytes: %d time: %s\n", len(data)*4, written, time.Since(start))

	start = time.Now()
	written = powc.P4nzzenc128v32(data, buffOut)
	fmt.Printf("P4nzzenc128v32: data length: %d written bytes: %d time: %s\n", len(data)*4, written, time.Since(start))

	start = time.Now()
	inint := ToInt32(data)
	var int32out []uint32
	one := intcomp.CompressInt32(inint, int32out)
	fmt.Printf("CompressInt32:  data length: %d written bytes: %d time: %s\n", len(data)*4, len(one)*4, time.Since(start))

	// var int32out2 []int32
	// out := intcomp.UncompressInt32(one, int32out2)
	// floatout := int32tofloat32(out)
	// fmt.Printf("%x\n", floatout[0])
	// fmt.Printf("%x\n", decode)
	// written = powc.P4nzdec128v32(buffOut, 8, decode)
	// written = powc.Decodefp4(buffOut, 8, decode)
	// fmt.Printf("%d\n", written)
	// fmt.Printf("%f\n", decode)
	// file, _ := os.OpenFile("test.png", os.O_CREATE|os.O_RDWR, 0777)
	// png.Encode(file, img)

}

func int32tofloat32(in []int32) []float32 {
	p := unsafe.Pointer(&in[0])

	size := len(in)

	var data []float32
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = uintptr(p)
	sh.Len = size
	sh.Cap = size

	return data
}

func ToByte(in []float32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&in[0])), len(in)*4)
}
func ToInt32(in []float32) []int32 {
	return unsafe.Slice((*int32)(unsafe.Pointer(&in[0])), len(in))
}

func float32toint32(in []float32) []int32 {
	p := uintptr(unsafe.Pointer(&in[0]))

	size := len(in)

	var data []int32
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sh.Data = p
	sh.Len = size
	sh.Cap = size

	return data
}
