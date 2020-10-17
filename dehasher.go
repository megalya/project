package main

/* Petit script en go permettant de faire des attaques par dictionnaire en md5
pas tre puissant mais fait le taff*/

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
			fmt.Printf("\nFrom hash : %s\n", s)
			fmt.Printf("Found in : ")
			break
		}
	}
	if !find {
		fmt.Printf("Hash not found !\n")
		fmt.Printf("Total time : ")
	}
}

func main() {
	logo()
	hash := flag.String("h", "", "--hash ; the hash you want to crack( only md5)")
	list := flag.String("l", "", "--list ; the wordlist you want to use (.txt)")
	flag.Parse()
	hashstr := *hash
	liststr := *list
	file, err := os.Open(liststr)
	if len(hashstr) == 32 && err == nil {
		timeinit := time.Now()
		scan := bufio.NewScanner(file)
		scan.Split(bufio.ScanLines)
		var txtlines []string
		for scan.Scan() {
			txtlines = append(txtlines, scan.Text())
		}
		fmt.Printf("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n")
		attack(hashstr, txtlines)
		timefin := time.Now()
		fmt.Printf("%v", timefin.Sub(timeinit))
		fmt.Printf("\n+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n")
		return
	} else {
		fmt.Printf("Invalid hash or list \n")
		flag.PrintDefaults()
	}
}
