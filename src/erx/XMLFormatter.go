package erx

import (
        "strings"
        "strconv"
        "fmt"
        "os"
	"io"
	"xml"
)

/**
 * err -- erx error
 * indent -- flag to add indentation to result xml file
*/
func FormatSimpleXML(w io.Writer, err Error, indent bool) {
	formatSimpleXML_gen(w, err, indent, 0)
}

func escape (w io.Writer, str string) {
	xml.Escape(w, []uint8(str))
}

/**
 *This function create XML format of error
 *return string
*/
func formatSimpleXML_gen(w io.Writer, err Error, indent bool ,level int) {
	tab := ""
	hyphen := ""
	if indent {
		tab = "\t"
		hyphen = "\n"
	}
	if level == 0 {
		w.Write( []uint8( "<?xml version='1.0' encoding='UTF-8'>" + hyphen ) )
	}
	w.Write( []uint8( strings.Repeat(tab, level) + "<error>" + hyphen ) )
	w.Write( []uint8( strings.Repeat(tab, level+1) + "<message>" ) )
		escape(w, err.Message())
			w.Write( []uint8( "</message>" + hyphen ) )
	w.Write( []uint8( strings.Repeat(tab, level+1) + "<file>" ) ) 
		escape(w, transformPath(err.File()))
			w.Write( []uint8( "</file>" + hyphen ) )
	w.Write( []uint8( strings.Repeat(tab, level+1) + "<line>" ) ) 
		escape(w, strconv.Itoa(err.Line()))
			w.Write( []uint8( "</line>" + hyphen ) )
	w.Write( []uint8( strings.Repeat(tab, level+1) + "<variables>" + hyphen ) )
	if len(err.Variables())>0 {
		for name, val := range err.Variables() {
			w.Write( []uint8( strings.Repeat(tab, level+2) + "<variable name='" ) ) 
				escape(w, name)
					w.Write( []uint8 ( "'>" ) )
			switch i := val.(type) {
				case string :
					escape(w, i)
				case fmt.Stringer :
					escape(w, i.String())
				default :
					escape(w, fmt.Sprint(i))
			}
			w.Write( []uint8( "</variable>" + hyphen ) )
		}
	}
	w.Write( []uint8( strings.Repeat(tab, level+1) + "</variables>" + hyphen ) )
 	curErr := err.Errors().Front()
	if curErr!=nil {
		w.Write( []uint8( strings.Repeat(tab, level) ) )
		for curErr!=nil {
			switch i := curErr.Value.(type) {
				case Error :
					formatSimpleXML_gen(w, i, indent, level+1)
						w.Write( []uint8( hyphen ))
				case os.Error :
					w.Write( []uint8( strings.Repeat(tab, level) + "<error type='2'>" + hyphen ) )
						w.Write( []uint8( strings.Repeat(tab, level+2) ) )
							 escape(w, i.String())
								w.Write( []uint8( hyphen ) )
					w.Write( []uint8( strings.Repeat(tab, level+1) + "</error>" + hyphen ) )
				default :
					w.Write( []uint8( "???" + hyphen ) )
			}
			curErr = curErr.Next()
		}
	}

	w.Write( []uint8( strings.Repeat(tab, level) + "</error>" ) )	
}

