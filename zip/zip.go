package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

func ExampleNewWriter() {
	var b bytes.Buffer

	w := zlib.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()
	fmt.Println(b.Bytes())
	// Output: [120 156 202 72 205 201 201 215 81 40 207 47 202 73 225 2 4 0 0 255 255 33 231 4 147]
}

func ExampleNewReader() {
	buff := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207, 47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	b := bytes.NewReader(buff)

	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}

	var out bytes.Buffer
	io.Copy(&out, r)
	fmt.Print(out.String())
	fmt.Print(out.Bytes())
	// Output: hello, world
	r.Close()
}

func main() {
	ExampleNewWriter()
	fmt.Println("###########################")
	ExampleNewReader()
}
