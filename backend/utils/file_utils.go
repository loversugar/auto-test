package utils

import (
	"auto-test/datamodals"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadFileToJson(filename string) ([]*datamodals.Data, error) {
	var datas []*datamodals.Data

	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	for {
		var data *datamodals.Data
		line, err := reader.ReadString('\n')
		fmt.Println(line)

		json.Unmarshal([]byte(line), &data)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		datas = append(datas, data)
	}

	return datas, err
}
