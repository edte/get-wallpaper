package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const (
	basePageUrl = "https://wall.alphacoders.com/by_resolution.php?w=1920&h=1080&lang=Chinese&page="
	numsReg     = "<title>(.*?)\\s1920x1080\\s高清壁纸"
	PagesReg    = "桌面背景\\sID:(.*?)\""
	perPageReg  = "width=\"1920\" height=\"1080\" src=\"(.*?)\""
)

var (
	PageNum int
	ids     []string
	urls    []string
	imgUlrs []string
)

type Url struct {
	url     string
	baseUrl string
	pageNum int
}

func (u *Url) Url() string {
	return u.url
}

func (u *Url) SetPageNum(pageNum int) {
	u.pageNum = pageNum
}

func NewUrl(page int) *Url {
	return &Url{url: basePageUrl + strconv.Itoa(page)}
}

func (u *Url) GetBody() (string, error) {
	resp, err := http.Get(u.Url())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetPagesNum() {
	body, _ := NewUrl(1).GetBody()
	num := regexp.MustCompile(numsReg)
	strs := num.FindAllStringSubmatch(body, -1)[0][1]
	PageNum, _ = strconv.Atoi(strs)
}

func GetPagesUrl() {
	body, _ := NewUrl(1).GetBody()
	pages := regexp.MustCompile(PagesReg)
	t := pages.FindAllStringSubmatch(body, -1)
	for _, v := range t {
		ids = append(ids, v[1])
	}
}

func GetUrl(id string) string {
	return "https://wall.alphacoders.com/big.php?i=" + id + "&lang=Chinese"
}

func GetUrls() {
	for _, v := range ids {
		urls = append(urls, GetUrl(v))
	}
}

func GetImgUrls() {
	for _, url := range urls {
		fmt.Println(url)
		reps, _ := http.Get(url)
		defer reps.Body.Close()
		body, _ := ioutil.ReadAll(reps.Body)
		page := regexp.MustCompile(perPageReg)
		s := page.FindAllStringSubmatch(string(body), -1)
		imgUlrs = append(imgUlrs, s[0][1])
	}
	fmt.Println(imgUlrs)
}

func SaveImgByUrl() {
	for _, v := range imgUlrs {
		fmt.Println(v)
		resp, _ := http.Get(v)
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		fmt.Println(err)
		ioutil.WriteFile("/home/edte/pictures/"+GetTimeUnix()+".jpg", data, 0644)
	}
}

func main() {
	GetPagesUrl()
	fmt.Println("#")
	GetUrls()
	fmt.Println("##")
	GetImgUrls()
	fmt.Println("###")
	SaveImgByUrl()
}

func GetTimeUnix() string {
	t := time.Now().Unix()
	return strconv.FormatInt(t, 10)
}
