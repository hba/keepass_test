package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/tobischo/gokeepasslib"
)

func main() {
	file, _ := os.Open("./example.kdbx")

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials("abcdefg12345678")
	_ = gokeepasslib.NewDecoder(file).Decode(db)
	db.UnlockProtectedEntries()
	
	for number, _ := range db.Content.Root.Groups[0].Groups {
		for _, entry := range db.Content.Root.Groups[0].Groups[number].Entries{
			if strings.HasPrefix(entry.GetTitle(),"foo"){
				fmt.Println("without deref -> "+entry.GetPassword())
			}
		}
	}
}