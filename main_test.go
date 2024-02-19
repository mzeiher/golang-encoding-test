package main_test

// import (
// 	"math"
// 	"testing"

// 	"example.com/float/powc"
// )

// // var width = 2880
// // var height = 5760
// var width = 100
// var height = 100
// var data []float32 = make([]float32, width*height)
// var decode = make([]float32, width*height)
// var buffOut = make([]byte, width*height*4)

// func init() {

// 	for x := 0; x < width; x++ {
// 		for y := 0; y < height; y++ {
// 			floatVal := float32(math.Sin((1. / 24.) * (float64(x)*float64(width) + float64(y))))
// 			data[x*width+y] = floatVal
// 		}
// 	}
// }

// func BenchmarkCompresssDecompress(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		powc.P4nzenc128v32(data, buffOut)
// 		powc.P4nzdec128v32(buffOut, int64(width)*int64(height), decode)
// 	}

// }
// func BenchmarkCompresss(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		powc.P4nzenc128v32(data, buffOut)
// 		// powc.P4nzdec128v32(buffOut, 8, decode)
// 	}

// }
// func BenchmarkDecompress(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		// powc.P4nzenc128v32(data, buffOut)
// 		powc.P4nzdec128v32(buffOut, int64(width)*int64(height), decode)
// 	}

// }
