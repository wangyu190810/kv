package utils;

import "testing"



func TestHadd( t *testing.T ){
	data := Data{make(map[string]interface{})}
	fliter := make(map[string]string)
	fliter["test"] = "test"
	newdata := HsetData{fliter}
	data.HAdd("test",newdata);
	t.Log(data);
}

func TestHget( t *testing.T ){
	data := Data{make(map[string]interface{})}
	fliter := make(map[string]string)
	fliter["test"] = "test"
	newdata := HsetData{fliter}
	data.HAdd("test",newdata);
	val,flag := data.HGet("test","test");
	if flag ==true{
		t.Log(val);
	}else{
		t.Error("get fail")
	}
	
}

func TestHdel( t *testing.T ){
	data := Data{make(map[string]interface{})}
	fliter := make(map[string]string)
	fliter["test"] = "test"
	newdata := HsetData{fliter}
	data.HAdd("test",newdata);
	flag := data.HDel("test","test");
	if flag ==0{
		t.Log("sucess");
	}else{
		t.Error("delete fail")
	}
	
}


