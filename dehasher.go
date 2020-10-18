package main

/* Petit script en go permettant de faire des attaques par dictionnaire en md5
pas tres puissant mais fait le taff
ajout de regles a venir*/

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"time"
)

func logo() {
	fmt.Printf("  ____       _               _     _____      \n")
	fmt.Printf(" |  _ \\  ___| |__   __ _ ___| |__ |___ / _ __ \n")
	fmt.Printf(" | | | |/ _ \\ '_ \\ / _` / __| '_ \\  |_ \\| '__|\n")
	fmt.Printf(" | |_| |  __/ | | | (_| \\__ \\ | | |___) | |   \n")
	fmt.Printf(" |____/ \\___|_| |_|\\__,_|___/_| |_|____/|_|   \n")
	fmt.Printf("  V 0.1\n")
}

func hashfromstring(userstring string) string {
	hasher := md5.New()
	hasher.Write([]byte(userstring))
	return hex.EncodeToString(hasher.Sum(nil))
}

func attack(s string, l []string) {
	find := false
	for i := 0; i < len(l); i++ {
		data := hashfromstring(l[i])
		if data == s {
			find = true
			fmt.Printf("Password found : ")
			fmt.Printf("%s", l[i])
			//fmt.Printf("\nIn position : %v", i)
			fmt.Printf("\nFrom hash : %s\n", s)
			fmt.Printf("Found in : ")
			break
		}
	}
	if !find {
		fmt.Printf("Hash not found !\n")
		fmt.Printf("May use a larger wordlist\n")
		fmt.Printf("Total time : ")
	}
}

/*func otozero(s []rune, r int) []string {
	list := make([]string, r)
	for i := r; i > 0; i-- {

	}
}

func attackwithrules(s string, l []string, r string) {
	lenght := len(l)
	find := false
	var rule []string
	count := 0
	for i := 0; i < lenght; i++ {
		tab := []rune(l[i])
		data := hashfromstring(l[i])
		if data == s {
			find = true
			break
		} else {
			for j := 0; j < len(tab)-1; j++ {
				if tab[j] == 'o' || tab[j] == 'O' {
					count++
				}
			}
			if count > 0 {
				count *= count
			}
		}
	}
}*/

func main() {
	logo()
	hash := flag.String("h", "", "--hash ; the hash you want to crack( only md5)")
	list := flag.String("l", "", "--list ; the wordlist you want to use (.txt)")
	//rules := flag.String("r", "", "--rules ; rules you want to add for bruteforce ")
	flag.Parse()
	hashpnt := *hash
	liststr := *list
	//rulespnt := *rules
	file, err := os.Open(liststr)
	if len(hashpnt) == 32 && err == nil {
		timeinit := time.Now()
		scan := bufio.NewScanner(file)
		scan.Split(bufio.ScanLines)
		var txtlines []string
		for scan.Scan() {
			txtlines = append(txtlines, scan.Text())
		}
		fmt.Printf("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n")
		/*if rulespnt != "" {
			attackwithrules(hashpnt, txtlines, rulespnt)
		} else {
			attack(hashpnt, txtlines)
		}*/
		attack(hashpnt, txtlines)
		timefin := time.Now()
		fmt.Printf("%v", timefin.Sub(timeinit))
		fmt.Printf("\n+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n")
		return
	} else {
		fmt.Printf("Invalid hash or list \n")
		flag.PrintDefaults()
	}
}
