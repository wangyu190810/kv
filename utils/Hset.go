package utils


type HsetData struct {
	FilterData map[string]string
}

func map_zip(data HsetData, newdata HsetData) HsetData {
	enddata := data
	for key, _ := range data.FilterData {
		value, flag := newdata.FilterData[key]
		if flag == true {
			enddata.FilterData[key] = value
		}
	}
	return enddata
}

func (self Data) HAdd(key string, data HsetData) int8 {
	if Keydata.Add(key, int8(HsetKey)) > -1 {
		value, flag := self.ValueData[key].(HsetData)
		if flag == false {
			self.ValueData[key] = data
			return 0
		} else {
			self.ValueData[key] = map_zip(value, data)
			return 1
		}
	} else {
		return -1
	}
}

func (self Data) HGet(key string, filter string) (string, bool) {

	value, flag := self.ValueData[key]
	if flag == false {
		return "", false
	} else {
		var hval string
		hval, flag = value.(HsetData).FilterData[filter]
		if flag == false {
			return "", false
		} else {
			return hval, true
		}
	}
}

func (self Data) HDel(key string, filter string) int {
	value, flag := self.ValueData[key]
	if flag == false {
		return 1
	} else {
		_, flag = value.(HsetData).FilterData[filter]
		if flag == false {
			return 1
		} else {
			delete(value.(HsetData).FilterData, filter)
			return 0
		}
	}

}

