package main

import (
	"strings"
	"fmt"
	"strconv"
	"io"
	"errors"
)

func main() {
	var str string
	Checkip4addr("4.4.4.4")
	fmt.Scanf("%s\n",&str)
	Extendip6addr("::",true,true)
	fmt.Println(Extendip6addr(str,true,true))
}
func Checkip4addr(addr string) bool{
	if strings.Count(addr,".")==3{
		sbuf:=strings.Split(addr,".")
		for _,s := range sbuf {
			rs,err:=strconv.Atoi(s)
			if err==nil&&rs>=0&&rs<=255{
				
			}else{
				return false
			}
		}
		return true
	}
	return false
}
func Checkip6addr(addr string) bool{
	if 2<=strings.Count(addr,":")&&strings.Count(addr,":")<=7{
		sbuf:=strings.Split(addr,":")
		for _,s:=range sbuf {
			var add_frag int
			_,err:=fmt.Sscanf(s,"%x",&add_frag)
			if err!=nil&&err!=io.EOF{
				fmt.Printf("%s",err)
				return false
			}
		}
		return true
	}
	return false
}
func Extendip6addr(addr string,result_format,result_ALPHA bool) (string,error){
	FMT_STR:="%x:%x:%x:%x:%x:%x:%x:%x"
	if false==Checkip6addr(addr){
		return "",errors.New("noipv6addr")
	}
	for r:=strings.NewReplacer("::",":0::"); strings.Count(addr,":")<7 ;  {
		addr=r.Replace(addr)
	}
	addr=strings.Replace(addr,"::",":0:",-1)
	var ibuf [8]int64
	for n,s:=range (strings.Split(addr,":"))  {
		i,err:=strconv.ParseInt(s,16,64)
		if err==io.EOF{
			i=0
		}
		ibuf[n]=i
	}
	if result_format{
		FMT_STR=strings.Replace(FMT_STR,"%x","%.4x",-1)
	}
	if result_ALPHA{
		FMT_STR=strings.Replace(FMT_STR,"x","X",-1)
	}
	add_s:=fmt.Sprintf(FMT_STR,ibuf[0],ibuf[1],ibuf[2],ibuf[3],ibuf[4],ibuf[5],ibuf[6],ibuf[7])
	return add_s,nil
}
func SimplifyIp6(addr string)string{
	
}