package erx

import (
	"strings"
	"strconv"
	"fmt"
	"os"
	"io"
)

func FormatConsole(w io.Writer, err Error, tab string) {
	formatConsole_gen(w, err, tab, 0)
}

func formatConsole_gen(w io.Writer, err Error, tab string, level int) {
	w.Write([]uint8(strings.Repeat(tab, level) + transformPath(err.File()) + ":"))
	w.Write([]uint8(strconv.Itoa(err.Line()) + " " + err.Message() + "\n"))
	funcFile, funcLine := err.Func().FileLine(err.Func().Entry())
	w.Write([]uint8(strings.Repeat(tab, level) + transformPath(funcFile) + ":"))
	w.Write([]uint8(strconv.Itoa(funcLine) + " " + err.Func().Name() + "\n"))
	level++
	if len(err.Variables()) > 0 {
		w.Write([]uint8(strings.Repeat(tab, level) + "Scope variables:\n"))
		for name, val := range err.Variables() {
			w.Write([]uint8(strings.Repeat(tab, level+1) + name + "\t: "))
			switch i := val.(type) {
			case string:
				w.Write([]uint8(i))
			case fmt.Stringer:
				w.Write([]uint8(i.String()))
			default:
				w.Write([]uint8(fmt.Sprint(i)))
			}
			w.Write([]uint8("\n"))
		}
	}

	curErr := err.Errors().Front()
	if curErr != nil {
		w.Write([]uint8(strings.Repeat(tab, level) + "Scope errors:\n"))
		for curErr != nil {
			switch i := curErr.Value.(type) {
			case Error:
				w.Write([]uint8("------"))
				formatConsole_gen(w, i, tab, level+1)
			case os.Error:
				w.Write([]uint8(strings.Repeat(tab, level+1) + i.String()))
			default:
				w.Write([]uint8("???\n"))
			}
			curErr = curErr.Next()
		}
	}
}


// Cut from path first dirs
func transformPath(path string) string {
	// finding path in pathCuts to cut
	for curPath := pathCuts.Front(); curPath != nil; curPath = curPath.Next() {
		if pathStr, isString := curPath.Value.(string); isString {
			if len(pathStr) <= len(path) && path[0:len(pathStr)] == pathStr {
				return path[len(pathStr):]
			}
		}
	}
	return path
}
