package utils

type SetData struct {
	Data string
}

func (self Data) Add(key string, data SetData) int8 {
	if Keydata.Exist(key) > -1 {
		_, flag := self.ValueData[key].(SetData)
		self.ValueData[key] = data
		if flag == false {
			return 0
		} else {
			return 1
		}
	} else {
		return -1
	}

}

func (self Data) Get(key string) (SetData, bool) {
	value, flag := self.ValueData[key]
	if flag == false {
		return SetData{}, false
	} else {
		return value.(SetData), true
	}
}

func (self Data) Del(key string) int {
	_, flag := self.ValueData[key]
	if flag == false {
		return 1
	} else {
		delete(self.ValueData, key)
		return 0
	}

}


