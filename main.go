package main

import (
	"fmt"
	"math"
	"time"
	"unsafe"

	"example.com/float/powc"
	"github.com/ronanh/intcomp"
)

var width = 50
var height = 50
var days = 24
var data []float32 = make([]float32, width*height*days)
var dataOut []float32 = make([]float32, width*height*days)
var buffOut = make([]byte, width*height*days*4)

func init() {

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for z := 0; z < days; z++ {
				nbr := y*(width*days) + x*days + z
				floatVal := float32(math.Sin((1. / 24.) * float64(nbr)))
				data[nbr] = floatVal
			}
		}
	}
}

func main() {
	start := time.Now()
	written := powc.Fpc0enc32(data, buffOut)
	fmt.Printf("Fpc0enc32:            data length: %d bytes, written bytes: %d, ratio: %f ,time: %s\n", len(data)*4, written, float32(written)/float32(len(data)*4), time.Since(start))

	reset(buffOut)

	start = time.Now()
	written = powc.Fpdfcmenc32(data, buffOut)
	fmt.Printf("Fpdfcmenc32:          data length: %d bytes, written bytes: %d, ratio: %f ,time: %s\n", len(data)*4, written, float32(written)/float32(len(data)*4), time.Since(start))

	reset(buffOut)

	start = time.Now()
	written = powc.Fpfcmenc32(data, buffOut)
	fmt.Printf("Fpfcmenc32:           data length: %d bytes, written bytes: %d, ratio: %f ,time: %s\n", len(data)*4, written, float32(written)/float32(len(data)*4), time.Since(start))

	reset(buffOut)

	start = time.Now()
	written = powc.Fpgenc32(data, buffOut)
	fmt.Printf("Fpgenc32:             data length: %d bytes, written bytes: %d, ratio: %f ,time: %s\n", len(data)*4, written, float32(written)/float32(len(data)*4), time.Since(start))

	start = time.Now()
	read := powc.Fpgdec32(buffOut, len(dataOut), dataOut)
	fmt.Printf("Fpgdec32:             data length: %d bytes, read bytes: %d time: %s\n", len(data)*4, read, time.Since(start))

	reset(buffOut)
	reset(dataOut)

	// P4nzzenc128v32
	start = time.Now()
	written = powc.P4nzzenc128v32(data, buffOut)
	fmt.Printf("P4nzzenc128v32:       data length: %d bytes, written bytes: %d, ratio: %f ,time: %s\n", len(data)*4, written, float32(written)/float32(len(data)*4), time.Since(start))

	reset(buffOut)

	// int4 + P4nzzenc128v32
	start = time.Now()
	intData := ToInt40(data)
	written = powc.P4nzzenc128v32(ToFloat32(intData), buffOut)
	fmt.Printf("int + P4nzzenc128v32: data length: %d bytes, written bytes: %d, ratio: %f ,time: %s\n", len(data)*4, written, float32(written)/float32(len(data)*4), time.Since(start))

	start = time.Now()
	read = powc.P4nzzdec128v32(buffOut, len(dataOut), dataOut)
	floatData := FromInt40(ToInt32(dataOut))
	floatData = floatData
	fmt.Printf("int + P4nzzdec128v32: data length: %d bytes, read bytes: %d time: %s\n", len(data)*4, read, time.Since(start))

	reset(buffOut)
	reset(dataOut)

	reset(buffOut)

	start = time.Now()
	inint := ToInt32(data)
	int32Compressed := intcomp.CompressInt32(inint, []uint32{})
	fmt.Printf("CompressInt32:        data length: %d bytes, written bytes: %d, ratio: %f ,time: %s\n", len(data)*4, len(int32Compressed), float32(len(int32Compressed))/float32(len(data)), time.Since(start))

	start = time.Now()
	int32decrompressed := intcomp.UncompressInt32(int32Compressed, []int32{})
	fmt.Printf("UnCompressInt32:      data length: %d bytes, written bytes: %d, ratio: %f ,time: %s\n", len(data)*4, len(int32decrompressed), float32(len(int32decrompressed))/float32(len(data)), time.Since(start))
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

func reset[T float32 | float64 | int32 | int | byte](val []T) []T {
	for i := 0; i < len(val); i++ {
		val[i] = 0
	}
	return val
}

func ToInt40(in []float32) []int32 {
	outSlice := make([]int32, len(in))
	for idx, val := range in {
		outSlice[idx] = int32(val * 40)
	}
	return outSlice
}
func FromInt40(in []int32) []float32 {
	outSlice := make([]float32, len(in))
	for idx, val := range in {
		outSlice[idx] = float32(val) / 40
	}
	return outSlice
}

// func FromInt40(in []float32) []int32 {

// }

func ToByte(in []float32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&in[0])), len(in)*4)
}
func ToInt32(in []float32) []int32 {
	return unsafe.Slice((*int32)(unsafe.Pointer(&in[0])), len(in))
}

func ToFloat32(in []int32) []float32 {
	return unsafe.Slice((*float32)(unsafe.Pointer(&in[0])), len(in))
}

func ToUInt32(in []float32) []uint32 {
	return unsafe.Slice((*uint32)(unsafe.Pointer(&in[0])), len(in))
}
