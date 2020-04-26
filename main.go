package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)
sudo
func main() {
	url := "https://ss.netnr.com/wallpaper"
	resp, err := http.Get(url)
	if err != nil {
		errors.New("get url failed!")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func m() {
	url := "http://p1.qhimg.com/bdm/1366_768_85/t0119918a57c366a5c5.jpg"
	resp, err := http.Get(url)
	if err != nil {
		errors.New("get url faild!")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.New("save jpg faild!")
	}
	ioutil.WriteFile("/home/edte/pictures/"+"a.jpg", data, 0644)
}
