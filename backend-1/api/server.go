package main

import (
	"auto-test/backend-1/util"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main()  {
	if util.CheckFileIsExist(util.FileName) {
		if err := os.Remove(util.FileName); err != nil {
			fmt.Println("file Error", err)
		}
	}
	http.HandleFunc("/addData", AddData)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
