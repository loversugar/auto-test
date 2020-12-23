package main

import (
	"auto-test/backend-1/util"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func AddData(w http.ResponseWriter, req *http.Request) {
	fmt.Println("HelloWorld")

	if !strings.EqualFold(req.Method, "options") {
		data := util.ParseHttpRequest(req);

		fmt.Println(data)

		var f *os.File
		var err1 error

		filename := "/Users/totti/go/src/auto-test/record.json"
		if util.CheckFileIsExist(filename) { //如果文件存在
			f, err1 = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666) //打开文件
			fmt.Println("文件存在")
		} else {
			f, err1 = os.Create(filename) //创建文件
			fmt.Println("文件不存在")
		}

		if err1 != nil {
			fmt.Println(err1)
		}

		byteToJson, _ := json.Marshal(data)
		io.WriteString(f, string(byteToJson) + "\n") //写入文件(字符串)
	}

	util.ResponseWithOrigin(w, req, []byte("success"))
}
