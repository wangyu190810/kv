package main

import (
	"kv/kvnet"
	//"mkydb/Storage"

	//"mkydb/Storage"
	// "utils"
	"log"
)

func ExitsKey(keyType int8) {
	if keyType == -1 {
		log.Printf("key exits other type")
	}
}

// func Run() {
// 	//var storage Storage.StorageProcess
// 	//storage = new(utils.Data)

// 	log.Printf(data.String())
// 	//data := utils.Data{map[string]interface{}{}}
// 	set_data := utils.SetData{"test"}
// 	exits := data.Add("test", set_data)
// 	ExitsKey(exits)
// 	fliter := make(map[string]string)
// 	fliter["test1"] = "test"
// 	hset_data := utils.HsetData{fliter}
// 	exits = data.HAdd("test", hset_data)
// 	ExitsKey(exits)
// 	log.Printf(data.String())

// }

func main() {
	// fmt.Println("hello")
	//Run()
	//file, err := os.Create("test.log")
	//if err != nil {
	//	log.Fatalln("fail to create test.log file!")
	//}
	//defer file.Close()
	//logger = log.New(file, "[Debug]", log.LstdFlags|log.Lshortfile) // 日志文件格式:log包含时间及文件行数
	//log.Println("输出日志到命令行终端")
	//logger.Println("将日志写入文件")
	//
	//logger.SetFlags(log.LstdFlags | log.Lshortfile) // 设置日志格式
	
	kvnet.Runserver("0.0.0.0:8080")
	data.Save(filename)
}
