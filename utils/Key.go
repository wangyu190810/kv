package utils

var Keydata KeyData

type State int

const (
	NoExist State = iota // value --> 0
	SetKey               // value --> 1
	HsetKey              // value --> 2
	ListKey              // value --> 3
)

// var KeyList = []int8{SetKey,HsetKey,ListKey}
var KeyList = make([]State, 3)

func init() {
	KeyList = append(KeyList, SetKey)
	KeyList = append(KeyList, HsetKey)
	KeyList = append(KeyList, ListKey)
}

type KeyData struct {
	Data map[string]int8
}

func (self KeyData) Exist(key string) int8 {
	value, flag := self.Data[key]
	if flag == false {
		return 0
	} else {
		return value
	}
}

func (self KeyData) Add(key string, keyType int8) int8 {
	value := self.Exist(key)
	var add_value int8
	if value == keyType {
		for _, v := range KeyList {
			if value == int8(v) {
				add_value = value
				break
			}
		}

	} else {
		add_value = -1
	}
	return add_value
}

func (self KeyData) KeyType(key string) int8  {
	return self.Exist(key)
}
