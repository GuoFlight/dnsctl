package main

import (
	"dnsctl/conf"
	"dnsctl/flag"
	"fmt"
	"github.com/miekg/dns"
	"os"
	"time"
)

func main() {
	//解析命令行
	flag.InitFlag()

	//功能导航
	if *flag.Version {
		fmt.Println(conf.VERSION)
		return
	}

	err := queryDNS()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func queryDNS() error {
	// 创建DNS请求的Client
	c := dns.Client{
		Timeout: time.Duration(*flag.Timeout) * time.Second,
	}
	//创建请求消息
	m := dns.Msg{}
	m.SetQuestion(flag.Args[0], dns.TypeA)
	// 发起请求
	r, _, err := c.Exchange(&m, *flag.DNS+":"+*flag.Port)
	if err != nil {
		return err
	}
	fmt.Printf("Server:\t%s\t%s\n", *flag.DNS, *flag.Port)

	// 从返回结果中得到A记录
	for _, ans := range r.Answer {
		A, ok := ans.(*dns.A)
		if ok {
			fmt.Printf("A:\t%s\n", A.A)
		}
	}
	return nil
}
