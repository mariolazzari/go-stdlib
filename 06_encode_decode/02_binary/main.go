package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/gob"
	"fmt"
)

func main() {
	// base 64
	data := []byte("Mario Lazzari")
	fmt.Println("data", data)
	decoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("decoded", decoded)

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, int32(42))
	fmt.Println("buf", buf)

	// gob
	buf = new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	data2 := Data{X: 100, Y: 200}
	enc.Encode(data2)
	fmt.Println("enc", enc)

}

type Data struct {
	X int
	Y int
}
