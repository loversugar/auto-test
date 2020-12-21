package main

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func AddData(w http.ResponseWriter, req *http.Request) {
	//io.WriteString(w, "hello, world!\n")
	fmt.Println("HelloWorld")

	var f *os.File
	var err1 error

	filename := "/home/C5311429/go/src/auto-test/record.json"
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}

	fmt.Println(f.Stat())

	if err1 != nil {
		fmt.Println(err1)
	}
	io.WriteString(f, "test") //写入文件(字符串)

	w.Write([]byte("success"))
}
