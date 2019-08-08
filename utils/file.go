package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"strings"
)

func ReadUnicode(path string)  {
	f,err := os.Open(path)
	if err !=nil{
		fmt.Println("file", err)
	}
	data,err := ioutil.ReadAll(f)
	if err !=nil{
		fmt.Println("read",err)
	}
	fmt.Print(data)
	data = data[3:]
	//b := make([]byte, utf8.UTFMax)
	data, err = Decode(data)
	rstInfo := make(map[string]interface{})
	fmt.Println(string(data))
	var splitFlag = "\r\n"
	for _, line := range strings.Split(string(data), splitFlag) {
		index := strings.IndexAny(line, "=")
		if index > 0 {
			fmt.Print(line[:index])
			fmt.Println(line[index+1:])
			rstInfo[line[:index]] = line[index+1:len(line)-2]
		}
	}
	fmt.Println(rstInfo)
	//result,err := json.MarshalIndent(rstInfo,"","")
	result, err := json.Marshal(rstInfo)
	fmt.Println("result",string(result))

	m := make(map[string]interface{}, 4) //因为类型多，可以用interface空接口
	m["address"] = "北京"
	m["languages"] = []string{"Golang", "PHP", "Java", "Python"}
	m["status"] = true
	m["price"] = 666.666

	result, err = json.Marshal(m)
	fmt.Println("result",string(result))
}


func Decode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.HZGB2312.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Encode(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GB18030.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func RunReadUnicode(){
	ReadUnicode("/home/too/Downloads/PutStr1_p31.info")
}
