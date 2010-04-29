package erx

import (
	"container/list"
	"runtime"
)


type ErrorVariables map[string]interface{}

var pathCuts *list.List = list.New()

func AddPathCut(path string) {
	pathCuts.PushBack(path)
}

type Error interface {
	Message() string

	Errors() *list.List
	Variables() ErrorVariables
	
	File() string
	Line() int
	Func() *runtime.Func

	AddE(err interface{})
	AddV(name string, value interface{})
}

type error_realization struct {
	message string
	
	file string
	line int

	errors *list.List
	variables ErrorVariables
	
	funcInfo *runtime.Func
}

func NewError(msg string) Error {
	err := error_realization{msg, "", 0, list.New(), make(map[string] interface{}), nil}
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		err.file, err.line = file, line
	} else {
		err.file, err.line = "???", 666
	}
	
	err.funcInfo = runtime.FuncForPC(pc)

	return Error(&err)
}

func NewSequent(msg string, error interface{}) Error {
	err := error_realization{msg, "", 0, list.New(), make(map[string] interface{}), nil}
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		err.file, err.line = file, line
	} else {
		err.file, err.line = "???", 666
	}
	
	err.funcInfo = runtime.FuncForPC(pc)
	
	err.AddE(error)
	return Error(&err) 
}

func (e *error_realization) Message() string {
	return e.message;
}

func (e *error_realization) File() string {
	return e.file
}

func (e *error_realization) Line() int {
	return e.line
}

func (e *error_realization) Func() *runtime.Func {
	return e.funcInfo
}

func (e *error_realization) Errors() *list.List {
	return e.errors
}

func (e *error_realization) Variables() ErrorVariables {
	return e.variables
}

func (e *error_realization) AddV(name string, v interface{}) {
	e.variables[name] = v
}

func (e *error_realization) AddE(err interface{}) {
	e.errors.PushBack(err)
}
