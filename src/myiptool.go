package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

type ipdate struct {
	IPaddr string
	Mask   int
}
type ipDataCountry struct {
	Country string
	ipdate
}

func GetHttp(url string) (html string, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	html = bytes.NewBuffer(body).String()
	return
}
func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func SplitIpData(instrings string) (IpDataList ipDataCountry, e error) {
	buf := strings.Split(instrings, "|")
	//buf="arin|US|ipv4|64.187.240.0|4096|20121210|assigned|45"
	//buf="arin|US|ipv6|2607:c880::|32|20161110|allocated|5d1f"
	if buf[1] == "" {
		e := errors.New("NoCountry")
		return IpDataList, e
	}
	IpDataList.Country, IpDataList.IPaddr, IpDataList.Mask = buf[1], buf[3], Atoi(buf[4])
	return IpDataList, e
}
func datamakeIp(instrings *[]string, kyeword string, ch chan *[]ipDataCountry) {
	var DataList []ipDataCountry
	for _, str := range *instrings {
		if strings.Contains(str, kyeword) {
			buf, err := SplitIpData(str)
			if err != nil {
				continue
			}
			DataList = append(DataList, buf)
		}
	}
	DataListP := &DataList
	ch <- DataListP
}
func GetIpCountryList(url string) (iv4list, iv6list *[]ipDataCountry, err error) {
	htdata, err := GetHttp(url)
	if err != nil {
		return nil, nil, err
	}
	gg := strings.Split(htdata, "\n")

	iv4ch := make(chan *[]ipDataCountry)
	iv6ch := make(chan *[]ipDataCountry)
	go datamakeIp(&gg, "ipv4", iv4ch)
	go datamakeIp(&gg, "ipv6", iv6ch)
	iv4list = <-iv4ch
	iv6list = <-iv6ch
	return iv4list, iv6list, err
}

/*func newww()  {

}
*/
func main() {
	runtime.GOMAXPROCS(4)
	url := "http://ftp.arin.net/pub/stats/arin/delegated-arin-extended-latest"

	iv4l, iv6l, err := GetIpCountryList(url)
	if err != nil {
		panic("!!!WARN!!!")
	}
	for _, ss := range *iv4l {
		fmt.Printf("%s %s %d\n", ss.Country, ss.IPaddr, ss.Mask)
	}
	for _, ss := range *iv6l {
		fmt.Printf("%s %s %d\n", ss.Country, ss.IPaddr, ss.Mask)
	}
}
