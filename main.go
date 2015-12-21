package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) < 4 {
		fmt.Printf("Syntax: %s <path> <offset> <length>\n", os.Args[0])
		fmt.Printf("This will write <length> random bytes at the <offset> given in the file at the <path> given.\n")
		os.Exit(1)
	}
	path := os.Args[1]
	offset, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	length, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fp, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = fp.Seek(offset, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := int64(0); i < length; i++ {
		_, err := fp.Write([]byte{byte(rand.Int())})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fp.Close()
}
