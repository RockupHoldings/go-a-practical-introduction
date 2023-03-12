package main

import "os"

func Play() error {
	f, _ := os.OpenFile("deleteme", os.O_RDWR|os.O_CREATE, 0666)

	f.Truncate(0)
	f.Write([]byte("420"))
	return nil

}
