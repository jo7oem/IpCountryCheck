package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)
func main() {
	app:=cli.NewApp()
	app.Name="IP CLI TOOL"
	app.Author="jo7oem"
	app.Usage="IPアドレス関連ツール"
	app.Version="0.0.1"
	
	app.Before=func(c *cli.Context) error{
		
		fmt.Println("Start")
		return nil
	}
	app.After=func(c *cli.Context) error{
		
		fmt.Println("Stop")
		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dryrun, d", // 省略指定 => d
			Usage: "グローバルオプション dryrunです。",
		},
	}
	
	app.Commands = []cli.Command{
		// コマンド設定
		{
			Name:    "hello",
			Aliases: []string{"h"},
			Usage:   "hello world を表示します",
			Action:  helloAction,
		},
	}
	
	app.Run(os.Args)
}
func helloAction(c *cli.Context) {
	
	// グローバルオプション
	var isDry = c.GlobalBool("dryrun")
	if isDry {
		fmt.Println("this is dry-run")
	}
	
	// パラメータ
	var paramFirst = ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First() // c.Args()[0] と同じ意味
	}
	
	fmt.Printf("Hello world! %s\n", paramFirst)
}