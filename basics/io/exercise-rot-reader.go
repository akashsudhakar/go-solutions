package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

const FROM = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz "
const TO = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm "

func getMap(from string, to string) map[byte]byte {
	aMap := make(map[byte]byte)

	for i := 0; i < len(from); i++ {
		aMap[from[i]] = to[i]
	}

	return aMap
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	if err != nil {
		return -1, err
	}

	aMap := getMap(FROM, TO)

	for i := 0; i < n; i++ {
		b[i] = aMap[b[i]]
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
