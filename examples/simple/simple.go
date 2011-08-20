package main

import (
	"fmt"
	"../../src/erx/_obj/erx"
	"os"
	"runtime"
	"strings"
)

type MyType struct {
    v int
}

func (m *MyType) String() string {
    return "Hello error!"
}

func init() {
	_, file, _, _ := runtime.Caller(0)
	index := strings.LastIndex(file, "/")
	if index!= -1 {
		dirName := file[0:strings.LastIndex(file, "/")+1]
		fmt.Println(dirName)
		erx.AddPathCut(dirName)
	}
}

func main() {
	var m MyType
	_, osError := os.OpenFile("nonExistedFile.tmp", os.O_RDONLY, 0000)
	err := erx.NewSequent("Sequent error", osError)
	err.AddV("var1", "444")
	err.AddV("var2", &m)
	err1 := erx.NewSequent("Simple error", err)
	formatter := erx.NewStringFormatter("  ")
	fmt.Println(formatter.Format(err))
	fmt.Println(formatter.Format(err1))
}
