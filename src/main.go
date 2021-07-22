package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	type process_info struct {
		ApiServer     string  `json:"api_server"`
		AppId         string  `json:"app_id"`
		BasePath      string  `json:"base_path"`
		CpuUsedTime   string  `json:"cpu_used_time"`
		DeviceMemory  int16   `json:"device_memory"`
		Fid           string  `json:"fid"`
		FilePath      string  `json:"file_path"`
		HostCpu       float64 `json:"host_cpu"`
		HostMemory    float64 `json:"host_memory"`
		IsGroup       bool    `json:"is_group"`
		ItemVersionId string  `json:"item_version_id"`
		Name          string  `json:"name"`
		Pid           int64   `json:"pid"`
		StartTime     string  `json:"start_time"`
		Uid           string  `json:"uid"`
		User          string  `json:"user"`
	}

	type gpuinfo struct {
		Id          int64          `json:"id"`
		Name        string         `json:"name"`
		Free        int64          `json:"free"`
		Total       int64          `json:"total"`
		Used        int64          `json:"used"`
		UsedRate    string         `json:"used_rate"`
		Temperature int64          `json:"temperature"`
		Processes   []process_info `json:"processes"`
	}

	type ServerGpuInfo struct {
		Gpus []gpuinfo `json:"gpus"`
	}
	// resp, err := http.Get("http://172.16.201.19:8888/gpuinfo")
	// if err != nil {
	// 	fmt.Println("error occured ", err)
	// 	return
	// }

	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("error occured ", err)
	// 	return
	// }

	// fmt.Print(string(body))

	// 利用指定的method,url以及可选的body返回一个新的请求.如果body参数实现了io.Closer接口，
	// Request返回值的Body 字段会被设置为body，并会被Client类型的Do、Post和PostFOrm方法
	// 以及Transport.RoundTrip方法关闭。
	client := &http.Client{} //客户端,被Get,Head以及Post使用
	reqest, err := http.NewRequest("GET", "http://172.16.201.19:8888/gpuinfo", nil)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	//给一个key设定为响应的value.
	// reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value") //必须设定该参数,POST参数才能正常提交

	resp, err := client.Do(reqest) //发送请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer resp.Body.Close() //一定要关闭resp.Body
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	// fmt.Println(string(content))

	var gpus ServerGpuInfo
	error := json.Unmarshal(content, &gpus)
	if error != nil {
		fmt.Println("Fatal error ", error.Error())
	}

	fmt.Println(gpus)
}
