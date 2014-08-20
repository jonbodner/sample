package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"sync"
)

func main() {
	fmt.Println("hello, world")
	for k, v := range os.Args {
		fmt.Println(k, v)
	}
	content, err := ioutil.ReadFile("files.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileList := string(content)
	fmt.Println(fileList)
	files := strings.Split(fileList, "\n")
	fmt.Println(files)
	fileChan := make(chan string)
	var wg sync.WaitGroup
	for _, v := range files {
		wg.Add(1)
		go func(fileName string) {
			defer wg.Done()
			content, err := ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Println(err)
				return
			}
			fileChan <- string(content)
		}(v)
	}
	go func() {
		wg.Wait()
		close(fileChan)
	}()

	contents := []string{}
	for c := range fileChan {
		contents = append(contents, c)
	}
	fmt.Println(contents)

}
