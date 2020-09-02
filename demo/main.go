package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
)

func httppost()  {
	data :=`{"userName":"admin","passWord":"1"}`
	request, _ := http.NewRequest("POST", "http://pls.platform.dev.ffpms.net/api/v1/aUnifyLogin/unifyLogin/", strings.NewReader(data))
	//post数据并接收http响应
	resp,err :=http.DefaultClient.Do(request)
	if err!=nil{
		fmt.Printf("post data error:%v\n",err)
	}else {
		fmt.Println("post a data successful.")
		respBody,_ :=ioutil.ReadAll(resp.Body)
		fmt.Printf("response data:%v\n",string(respBody))
	}
}

func main() {
	httppost()
}