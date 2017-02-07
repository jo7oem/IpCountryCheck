package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"./ipaddrtools"
	"log"
)
func main() {
	app:=cli.NewApp()
	app.Name="IP CLI TOOL"
	app.Author="jo7oem"
	app.Usage="IPアドレス関連ツール"
	app.Version="0.1.1"
	
	app.Before=func(c *cli.Context) error{
		return nil
	}
	app.After=func(c *cli.Context) error{
		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dryrun, d", // 省略指定 => d
			Usage: "グローバルオプション 未実装",
		},
	}
	app.Commands = []cli.Command{
		// コマンド設定
		{
			Name:    "checkaddr",
			Aliases: []string{"c"},
			Usage:   "ipアドレスかどうかを表示します。",
			Action:  Check_Ipaddr,
		},
		{
			Name:    "isaddr",
			Aliases: []string{"i"},
			Usage:   "ipアドレスかどうかを返す !!単一引数のみ!!",
			Action:  isIpaddr,
		},
		{
			Name:    "ExtendAddr",
			Aliases: []string{"e"},
			Usage:   "v6アドレスを展開する",
			Action:  ExtendAddr,
		},
		{
			Name:    "simplev6",
			Aliases: []string{"s"},
			Usage:   "v6アドレスを短縮する",
			Action:  SimpleV6,
		},
		{
			Name:    "convPTS",
			Aliases: []string{"c"},
			Usage:   "BindのIP逆引き形式に変換する",
			Action:  ConvPTS,
		},
	}
	app.Run(os.Args)
}
func Check_Ipaddr(c *cli.Context){
	for _,s:=range joinArg(c) {
		fmt.Printf("%s:v4=%s:v6=%s\n",s,ipaddrtools.Checkip4addr(s),ipaddrtools.Checkip6addr(s))
	}
}
func joinArg(c *cli.Context)[]string{
	var str []string
	str=append(str,c.Args().First())
	str=append(str,c.Args().Tail()...)
	if len(str)==0{
		fmt.Println("JJ")
	}
	return str
}
func isIpaddr(c *cli.Context){
	s:=c.Args().First()
	if ipaddrtools.Checkip4addr(s)||ipaddrtools.Checkip6addr(s) {
	}else{
		cli.NewExitError("no ip addr",1)
	}
}
func ExtendAddr(c *cli.Context){
	for _,s:=range joinArg(c){
		ea,err:=ipaddrtools.Extendip6addr(s,true,false)
		if err !=nil{
			log.Fatalln(err)
		}else{
			fmt.Println(ea)
		}
	}
}
func SimpleV6(c *cli.Context){
	for _,s:=range joinArg(c){
		ea,err:=ipaddrtools.SimplifyIp6(s)
		if err !=nil{
			log.Fatalln(err)
		}else{
			fmt.Println(ea)
		}
	}
}
func ConvPTS(c *cli.Context){
	for _,s:=range joinArg(c){
		ea,err:=ipaddrtools.ModePTS(s)
		if err !=nil{
			log.Fatalln(err)
		}else{
			fmt.Println(ea)
		}
	}
}