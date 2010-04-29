package main

import (
	"strings"
	"strconv"
	"os"
	"bufio"
	"fmt"
	"runtime"
	"../../src/erx/_obj/erx"
)

func init() {
	// add path prefix to erx cut list
	_, file, _, _ := runtime.Caller(0)
	index := strings.LastIndex(file, "/")
	if index!= -1 {
		dirName := file[0:strings.LastIndex(file, "/")+1]
		fmt.Println(dirName)
		erx.AddPathCut(dirName)
	}
}

func fileSum(fileName string) (result int) {
	makeError := func(err interface{}) (res erx.Error) {
		res = erx.NewSequentLevel("Sum integers from file.", err, 1)
		res.AddV("file name", fileName)
		return
	}

	result = 0
	f, err := os.Open(fileName, os.O_RDONLY, 0000)
	if f==nil {
		panic(makeError(erx.NewSequent("Open file.", err)))
	}
	
	reader := bufio.NewReader(f)
	
	line, err := reader.ReadString('\n')
	lineNum := 1
	for err==nil || err==os.EOF {
		if strings.TrimSpace(line)!="" {
			chunkNum := 1
			for _, chunk := range strings.Split(line, " ", 0) {
				var curInt int
				curInt, err = strconv.Atoi(strings.TrimSpace(chunk))
				if err!=nil {
					errErx := erx.NewSequent("Convert string to integer.", err)
					errErx.AddV("chunk", chunk)
					errErx.AddV("line num", lineNum)
					errErx.AddV("chunk num", chunkNum)
					panic(makeError(errErx))
				}
				result += curInt
			}
		}		

		if err==os.EOF {
			break
		}
		line, err = reader.ReadString('\n')
		lineNum++
	}
	
	if err!=os.EOF {
		panic(makeError(erx.NewSequent("Reading from file.", err)))
	}
	return
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			if errErx, ok := err.(erx.Error); ok {
				formatter := erx.NewStringFormatter("  ")
				fmt.Println(formatter.Format(errErx))
			}
		}
	}()
	fileSum("ints.txt")
}
