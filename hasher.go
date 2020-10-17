package main

/* V 0.1 ne permet que de hasher une string en 2 type (md5/sha256) .
D'autre type arrive bientot ainsi que du " déhashage"
*/

import (
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
)

func hash(s1 string) int {
	if s1 == "md5" {
		return 1
	} else if s1 == "sha256" {
		return 2
	} else {
		return 0
	}
}

func md5hash(s string) {
	data := []byte(s)
	fmt.Printf("%x", md5.Sum(data))

}

func sha256hash(s string) {
	data := []byte(s)
	fmt.Printf("%x", sha256.Sum256(data))
}

func logo() {
	fmt.Printf("  _    _           _                \n")
	fmt.Printf(" | |  | |         | |               \n")
	fmt.Printf(" | |__| | __ _ ___| |__   ___ _ __  \n")
	fmt.Printf(" |  __  |/ _` / __| '_ \\ / _ \\ '__| \n")
	fmt.Printf(" | |  | | (_| \\__ \\ | | |  __/ |    \n")
	fmt.Printf(" |_|  |_|\\__,_|___/_| |_|\\___|_|    \n")
	fmt.Printf(" v.0.1 ")
	fmt.Printf("\n")
}

func main() {
	logo()
	typehash := flag.String("t", "", "--type ; get the output hash type")
	hashstr := flag.String("s", "", "--string ; paste the string you want to hash")
	flag.Parse()
	typestr := *typehash
	str := *hashstr
	if hash(typestr) != 0 && len(str) > 0 {
		fmt.Printf("\n================================\n")
		fmt.Printf("type : %s \n", typestr)
		fmt.Printf("Hash : ")

		if hash(typestr) == 1 {
			md5hash(str)
		} else if hash(typestr) == 2 {
			sha256hash(str)
		}
		fmt.Printf("\n================================\n")
		return
	} else {
		fmt.Printf("\nInvalid type or string\n")
		flag.PrintDefaults()
	}
}
