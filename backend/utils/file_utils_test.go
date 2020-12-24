package utils

import "testing"

func TestReadFileToJson(t *testing.T) {
	datas, err := ReadFileToJson("/home/C5311429/go/src/auto-test/record.json")

	if err != nil {
		t.Error(err)
	}

	for _, data := range datas {
		t.Log(data)
	}
}