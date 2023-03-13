package main

import (
	"embed"
	"fmt"
	"os"
	"rock_ed/todo"
)

var (
	//go:embed store.json
	store embed.FS
)

func main() {
	db := openDB()
	todoStore := todo.NewFsStore(db)

	fmt.Println(todoStore.List(todo.Checked))
}

func openDB() *os.File {
	file, err := store.Open("store.json")
	exitIfErr(err)

	info, err := file.Stat()
	exitIfErr(err)

	db, err := os.OpenFile(info.Name(), os.O_RDWR, 0666)
	exitIfErr(err)

	return db
}

func exitIfErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
