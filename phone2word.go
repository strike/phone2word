// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phone2word

import(
	"flag"
	"fmt"
	"regexp"
	"os"
	"strings"
	"github.com/Philio/GoMySQL"
	"io/ioutil"
)

func main() {
	var p *string = flag.String("n", "495", "phone number")
	var du *string = flag.String("du", "root", "db user")
	var dp *string = flag.String("dp", "", "db password")
	var ddb *string = flag.String("db", "phone2word", "database")
	var file *string = flag.String("f", "", "take words from file")
	flag.Parse()


	fmt.Println("phone number:\t", *p)

	reg := regexp.MustCompile("[^0-9]")
	number := reg.ReplaceAllString(*p,"")

	fmt.Println("cleaned number:\t", number, "\t", len(number))

	if number == "" {
		fmt.Println("ERROR>\t Please enter correct phone number")
		os.Exit(1)
	}

	reg = regexp.MustCompile("1")
	number = reg.ReplaceAllString(number,"0")

	pieces := strings.Split(number,"0")

	query := "select * from words where "
	num := map[string]int {}
	
	fmt.Println("splited:")
	for _, numbers := range pieces {
			fmt.Print("\t",numbers, ":\n\t\t\t")
			//generare query
			for i:=1; i <= len(numbers); i++ {
				j := 0
				for {
					if *file == "" {
						query += "number = " + numbers[j:j+i]
					} else {
						num[numbers[j:j+i]] = 0
					}
					fmt.Print(numbers[j:j+i], " ")
					j++
					
					if *file == "" {	
						if j+i > len(numbers) {	
							if i+1 <= len(numbers) {
								query += " or "
							} else {
								query += ";"
							}
							break
						} else {
							query += " or "
						}
					} else {
						if j+i > len(numbers) {	
							break
						}
					}
				}
			}
				fmt.Println()
	}
	fmt.Println()
	
 	if *file == "" {
		fmt.Println(query)

		db, err := mysql.DialUnix(mysql.DEFAULT_SOCKET, *du, *dp, *ddb)  
		if err != nil {  
			fmt.Println("ERROR>\tCan't connect to database!")
			fmt.Println(err)
			fmt.Println("\tfor example: -du user -dp pass -db database")
			fmt.Println("\tdefault variables:")
			fmt.Println("\t\tuser = root \n\t\tdatabase = phone2word \n\t\tand empty password")
			os.Exit(1)  
		}  

		err = db.Query(query)  
		if err != nil {  
			os.Exit(1)  
		} 

		result, err := db.UseResult()  
		if err != nil {  
		    os.Exit(1)  
		}  

		fmt.Println("result")
		for {  
		    row := result.FetchRow()  
		    if row == nil {  
			break  
		    }  
			fmt.Println(row)
		}
	} else {
		contents,_ := ioutil.ReadFile(*file)	
		words := strings.Split(strings.ToLower(string(contents)),"\n")
		code := ""
		for _, w := range words {
			code = ""
			for _,ws := range w {
				code += getCode(string(ws))	
			}
			for c,_ := range num {
				if code == c {
					fmt.Println(code, "\t",w )
				}
			}
		}

	}  
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



