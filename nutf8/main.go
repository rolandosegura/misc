package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatal("need a file name")
	}
	f, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal("error on open ", err)
	}
	defer f.Close()
	buf := make([]byte, 1024, 1024)
	var nc int
	for {
		n, err := f.Read(buf)
		if n == 0 && err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error reading ", err)
		}
		b := buf[:n]
		for len(b) > 0 {
			r, sz := utf8.DecodeRune(b)
			if r == utf8.RuneError && sz == 1 {
				fmt.Println()
				log.Printf("found invalid utf8 rune at byte %v %x", nc, b[0])
				b = b[1:]
				continue
			}
			fmt.Printf("%c", r)
			b = b[sz:]
			nc += sz
		}
	}
}
