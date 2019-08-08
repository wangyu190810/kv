package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	ValueData map[string]interface{}
}

func (data Data) Save(filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		ProcessError(err)
	}
	var GolangGob bytes.Buffer
	gob.Register(map[string]interface{}{})
	gob.Register(SetData{})
	gob.Register(HsetData{})
	enc := gob.NewEncoder(&GolangGob)
	err = enc.Encode(data)
	ProcessError(err)
	_, err = f.Write(GolangGob.Bytes())
	f.Sync()
	ProcessError(err)
	defer f.Close()

}

func Load(filename string) (Data, error) {
	f, err := os.Open(filename)
	var data Data
	if err != nil {
		ProcessError(err)
		//data = Data{make(map[string]interface{})}
		data = Data{map[string]interface{}{}}
		return data, err
	}
	gob.Register(map[string]interface{}{})
	gob.Register(SetData{})
	gob.Register(HsetData{})
	dec := gob.NewDecoder(f)
	var GolangGob Data
	err = dec.Decode(&GolangGob)
	ProcessError(err)
	data = GolangGob
	return data, err
}

func (self Data) String() string {
	return fmt.Sprintf("%v", self.ValueData)
}

func ProcessError(err error) bool {

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}
}

func (data Data) Command(contentLine string) {
	cmds := strings.Split(contentLine, " ")
	if len(cmds) <= 1 {
		fmt.Print("command err")
	} else {
		if len(cmds) == 2 {
			if "get" == cmds[0] {
				data.Get(cmds[1])
			} else if "del" == cmds[0] {
				data.Del(cmds[1])
			}
		} else if len(cmds) == 3 {
			if "set" == cmds[0] {
				data.Add(cmds[1], SetData{cmds[2]})
			} else if "hget" == cmds[0] {
				data.HGet(cmds[1], cmds[2])
			}
		} else if len(cmds) == 4 {
			if "hset" == cmds[0] {
				hsetValue := make(map[string]string)
				hsetValue[cmds[2]] = cmds[3]
				data.HAdd(cmds[1], HsetData{hsetValue})
			}
		}
		fmt.Print(data)
	}
}
