package Storage

type StorageProcess interface {
	Save(filename string)
	//Load(filename string) (utils.Data,error)
	//Flush(key string,keyType string,value interface{})
}



