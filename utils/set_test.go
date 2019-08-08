package utils;

import "testing"



func TestAdd( t *testing.T ){
	data := Data{make(map[string]interface{})}
	set_data := SetData{"test"}
	data.Add("test",set_data);
	t.Log(data);
}

func TestGet( t *testing.T ){
	data := Data{make(map[string]interface{})}
	set_data := SetData{"test"}
	data.Add("test",set_data);
	val,flag := data.Get("test");;
	if flag ==true{
		t.Log(val);
	}else{
		t.Error("get fail")
	}
	
}

func TestDel( t *testing.T ){
	data := Data{make(map[string]interface{})};
	set_data := SetData{"test"};
	data.Add("test",set_data);
	flag := data.Del("test");
	if flag == 0{
		t.Log("sucess");
	}else{
		t.Error("delete fail")
	}
	
}


