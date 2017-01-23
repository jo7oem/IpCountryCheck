package main
import (
	"strings"
	"fmt"
	"strconv"
	"errors"
	"regexp"
)
func main() {
	var str string
	Checkip4addr("4.4.4.4")
	fmt.Scanf("%s\n",&str)
	fmt.Println(SimplifyIp6(str))
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
	ipblockn:=strings.Count(addr,":")
	if 2<=ipblockn&&ipblockn<=7{
		if ipblockn<7 {
			if false == strings.Contains(addr, "::") {
				return false
			}
		}
		sbuf:=strings.Split(addr,":")
		for _,s:=range sbuf {
			if 0==strings.Compare(s,""){
				s="0"
			}
			i,err:=strconv.ParseInt(s,16,64)
			if err!=nil{
				return false
			}
			if 0x0>i||i>0xffff{
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
		if err!=nil{
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
func SimplifyIp6(addr string)(string,error){
	if false==Checkip6addr(addr){
		return "",errors.New("noipv6addr")
	}
	addr,_=Extendip6addr(addr,false,false)
	re := regexp.MustCompile("(:|^)(0:)+0")
	ibuf:=re.FindAllStringIndex(addr,-1)
	if ibuf!=nil{
		var ilen int
		var imax []int
		for n,i :=range ibuf{
			if n==0{
				imax=i
				ilen=i[1]-i[0]
				continue
			}
			if ilen<(i[1]-i[0]){
				imax=i
				ilen=i[1]-i[0]
			}
		}
		addr=strings.Replace(addr,addr[imax[0]:imax[1]],":",1)
	}
	if strings.Count(addr,":")==1{
		addr=strings.Replace(addr,":","::",1)
	}
	return addr,nil
}