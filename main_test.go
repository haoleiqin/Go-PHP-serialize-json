package main

import (
	"fmt"
	"testing"
)

func TestPhpserialize2Json(t *testing.T) {
	testSerializeStr := `a:1:{i:0;a:3:{s:2:"id";i:1;s:5:"price";i:80;s:4:"name";s:12:"一个名字";}}`
	var serObj *serObj
	jsonRes := serObj.Phpserialize2Json(testSerializeStr)
	fmt.Print(jsonRes)
	t.Log(jsonRes)
}
