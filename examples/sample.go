package main

import (
	"fmt"
	"erx"
)

const (
	ERROR_ONE = iota
	ERROR_TWO
	ERROR_THREE
	ERROR_FOUR
)

const scope = "main"

var errors = map[uint]string {
	ERROR_ONE   : "error one",
	ERROR_TWO   : "error two",
	ERROR_THREE : "error three",
	ERROR_FOUR  : "error four",
}


func init() {
	fmt.Println("init called")
	erx.RegisterScopeMessages(scope, errors)
	erx.AddPathCut("/home/steplg/quickdoc/workspaces/jtt/")
}

func main() {
	fmt.Println("main called")
	err := erx.NewError(scope, ERROR_ONE)
	err.AddV("var1", "444")
	err1 := erx.NewSequent(scope, ERROR_TWO, err)
	formatter := erx.NewStringFormatter("  ")
	fmt.Println(formatter.Format(err))
	fmt.Println(formatter.Format(err1))
}
