package erx

import (
	"strings"
	"strconv"
	"fmt"
	"os"
)

type StringFormatter struct {
	indent string
}

func NewStringFormatter(indent string) *StringFormatter {
	formatter := new(StringFormatter)
	formatter.indent = indent
	return formatter
}

func (f *StringFormatter) Format(err Error) string {
	return f.formatLevel(err, 0)
}

func (f *StringFormatter) formatLevel(err Error, level int) string {
	result := ""
	result += strings.Repeat(f.indent, level)
	result += transformPath(err.File()) + ":" + strconv.Itoa(err.Line()) + " " + err.Message()
	result += "\n"
	result += strings.Repeat(f.indent, level)
	funcFile, funcLine := err.Func().FileLine(err.Func().Entry())
	result += transformPath(funcFile) + ":" + strconv.Itoa(funcLine) + " " + err.Func().Name()
	result += "\n"
	level++
	if len(err.Variables())>0 {
		result += strings.Repeat(f.indent, level)
		result += "Scope variables:\n"
		for name, val := range err.Variables() {
			result += strings.Repeat(f.indent, level+1)
			result += name + "\t: "
			switch i := val.(type) {
				case string :
					result += i
				case fmt.Stringer :
					result += i.String()
				default :
					result += fmt.Sprint(i)
			}
			result += "\n"
		}
	}
	
	curErr := err.Errors().Front()
	if curErr!=nil {
		result += strings.Repeat(f.indent, level)
		result += "Scope errors:\n"
		for curErr!=nil {
			switch i := curErr.Value.(type) {
				case Error :
					result += f.formatLevel(i, level+1)
				case os.Error :
					result += strings.Repeat(f.indent, level+1)
					result += i.String()
				default :
					result += "???\n"
			}
			curErr = curErr.Next()
		}
	}
	return result
}

// Cut from path first dirs
func transformPath(path string) string {
	// finding path in pathCuts to cut
	for curPath := pathCuts.Front(); curPath!=nil; curPath = curPath.Next() {
		if pathStr, isString := curPath.Value.(string); isString {
			if len(pathStr)<=len(path) && path[0:len(pathStr)]==pathStr {
				return path[len(pathStr):]
			}
		}
	}
	return path
}

func PanicPrinter() {
	if e := recover(); e!=nil {
		if erxErr, ok := e.(Error); ok {
			formatter := NewStringFormatter("  ")
			fmt.Println(formatter.Format(erxErr) )
		}
	}	
}
