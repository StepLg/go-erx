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
	Code() uint
	Scope() string

	Errors() *list.List
	Variables() ErrorVariables
	
	File() string
	Line() int

	AddE(interface{})
	AddV(name string, value interface{})
}

type error_realization struct {
	code uint
	scope string
	
	file string
	line int

	errors *list.List
	variables ErrorVariables
}

func NewError(scope string, code uint) (Error) {
	err := error_realization{code, scope, "", 0, list.New(), make(map[string] interface{})}
	_, file, line, ok := runtime.Caller(1)
	if ok {
		err.file, err.line = file, line
	} else {
		err.file, err.line = "???", 666
	}
	
	// ищем начало пути файла в pathCuts, чтобы их вырезать
	for curPath := pathCuts.Front(); curPath!=nil; curPath = curPath.Next() {
		if pathStr, isString := curPath.Value.(string); isString {
			if len(pathStr)<=len(err.file) && err.file[0:len(pathStr)]==pathStr {
				err.file = err.file[len(pathStr):len(err.file)]
			}
		}
	}
	return &err 
}

func NewSequent(scope string, code uint, error interface{}) (Error) {
	err := error_realization{code, scope, "", 0, list.New(), make(map[string] interface{})}
	_, file, line, ok := runtime.Caller(1)
	if ok {
		err.file, err.line = file, line
	} else {
		err.file, err.line = "???", 666
	}
	
	// ищем начало пути файла в pathCuts, чтобы их вырезать
	for curPath := pathCuts.Front(); curPath!=nil; curPath = curPath.Next() {
		if pathStr, isString := curPath.Value.(string); isString {
			if len(pathStr)<=len(err.file) && err.file[0:len(pathStr)]==pathStr {
				err.file = err.file[len(pathStr):len(err.file)]
			}
		}
	}
	err.AddE(error)
	return &err 
}

func (e *error_realization) Code() uint {
	return e.code
}

func (e *error_realization) Scope() string {
	return e.scope
}

func (e *error_realization) File() string {
	return e.file
}

func (e *error_realization) Line() int {
	return e.line
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
