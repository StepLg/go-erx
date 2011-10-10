package main

import (
	"strings"
	"strconv"
	"os"
	"bufio"
	"fmt"
	"runtime"
	"../../src/erx/_obj/erx"
	"flag"
)

func init() {
	// add path prefix to erx cut list
	_, file, _, _ := runtime.Caller(0)
	index := strings.LastIndex(file, "/")
	if index != -1 {
		dirName := file[0 : strings.LastIndex(file, "/")+1]
		fmt.Println(dirName)
		erx.AddPathCut(dirName)
	}

}

func fileSum(fileName string) (result int) {
	defer func() {
		if err := recover(); err != nil {
			res := erx.NewSequent("Sum integers from file.", err)
			res.AddV("file name", fileName)
			panic(res)
		}
	}()

	result = 0
	f, err := os.Open(fileName, os.O_RDONLY, 0000)
	if f == nil {
		panic(erx.NewSequent("Open file.", err))
	}

	reader := bufio.NewReader(f)

	line, err := reader.ReadString('\n')
	lineNum := 1
	for err == nil || err == os.EOF {
		if strings.TrimSpace(line) != "" {
			chunkNum := 1
			for _, chunk := range strings.Split(line, " ") {
				var curInt int
				curInt, err = strconv.Atoi(strings.TrimSpace(chunk))
				if err != nil {
					errErx := erx.NewSequent("Convert string to integer.", err)
					errErx.AddV("chunk", chunk)
					errErx.AddV("line num", lineNum)
					errErx.AddV("chunk num", chunkNum)
					panic(errErx)
				}
				result += curInt
			}
		}

		if err == os.EOF {
			break
		}
		line, err = reader.ReadString('\n')
		lineNum++
	}

	if err != os.EOF {
		panic(erx.NewSequent("Reading from file.", err))
	}
	return
}

func flagMode_get() string {
	var mode *string = flag.String("mode", "console", "mode out: XML, console")
	flag.Parse()
	return *mode
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			if errErx, ok := err.(erx.Error); ok {
				erx.AutoOutput(os.Stdout, flagMode_get(), errErx)
			}
		}
	}()

	fileSum("ints.txt")
}
