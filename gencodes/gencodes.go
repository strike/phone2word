package phone2wordgen

import (
	"io/ioutil"
	"strings"
	"fmt"
	"flag"
	"os"
)


func main() {
	var infile *string = flag.String("f", "", "input file")
	var outfile *string = flag.String("o", "", "output file")
	var sql *bool = flag.Bool("sql", false, "create sql query")
	var sql_table = flag.String("table", "phone2word", "table name (if sql is true)")
	flag.Parse()

	if *infile == "" {
		fmt.Println("ERROR>\tPlease enter input file\n\tfor example, -f filename")
		os.Exit(1)
	}

	if *outfile == "" {
		if *sql {
			*outfile = *infile + ".sql"
		} else {
			*outfile = *infile + ".out"
		}
	}
	
	contents,_ := ioutil.ReadFile(*infile);

	words := strings.Split(string(contents),"\n")

	code, w := "", ""
	out := ""

	if *sql {
		out += "INSERT INTO `" + *sql_table + "` (`number`, `word`) VALUES \n"
	}

	first := true
	for _,word := range words {
		if *sql {
			if first {
				first = false			
			} else {
				out += ",\n"
			}
		}
		w = strings.Trim(word," ")
		code = ""
		for i:= 0; i < len(w); i++{
			code += getCode(string(w[i]))	
		}
		if *sql {
			out += "('" + code +"','" + w + "')"
		} else {
			out += code + "\t" + w + "\n"		
		}
	} 	
	if *sql {
		out += ";"
	}
	ioutil.WriteFile(*outfile, []byte(out), 0666)
	fmt.Println("saved to ", *outfile)
}

func getCode(s string) string{
	switch s {
		//rus
		case "а","б","в","г":
			return "2"
		case "д","е","ё","ж","з":
			return "3"

		case "и","й","к","л":
			return "4"

		case "м","н","о","п":
			return "5"

		case "р","с","т","у":
			return "6"

		case "ф","х","ц","ч":
			return "7"

		case "ш","щ","ъ","ы":
			return "8"

		case "ь","э","ю","я":
			return "9"

		//eng
		case "a","b","c":
			return "2"

		case "d","e","f":
			return "3"

		case "g","h","i":
			return "4"

		case "j","k","l":
			return "5"

		case "m","n","o":
			return "6"

		case "p","q","r","s":
			return "7"

		case "t","u","v":
			return "8"

		case "w","x","y","z":
			return "9"
	}
	return ""
}
