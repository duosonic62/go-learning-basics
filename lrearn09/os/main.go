package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	openFile()
	dir()
	osDirection()
}

func openFile() {
	bs := make([]byte, 128)
	f, err := os.Open("foc.txt")
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Read(bs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bs)
	for _, v := range bs {
		fmt.Print(v)
	}

	fi, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fi.Name())
	fmt.Println(fi.Size())
	fmt.Println(fi.Mode())
	fmt.Println(fi.ModTime())
	fmt.Println(fi.IsDir())
}

func dir() {
	f, err := os.Open("../../")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fis, err := f.Readdir(0) 

	for _, fi := range fis {
		if fi.IsDir() {
			fmt.Println(fi.Name())
		}
	}
}

func osDirection() {
	fmt.Println(os.Hostname())

	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}