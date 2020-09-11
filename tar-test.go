package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	tarball, err := ioutil.ReadFile("something.tar.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	tarBuffer := bytes.NewBuffer(tarball)
	tr := tar.NewReader(tarBuffer)
	for {
		fmt.Println("in loop")
		fileEntry, err := tr.Next()
		if err == io.EOF {
			fmt.Println("end of file", err.Error())
			break
		}
		if err != nil {
			fmt.Println("err was not nil")
			return
		}
		fmt.Println("file", fileEntry.Name)
	}
	return
}
