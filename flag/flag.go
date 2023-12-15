package flag

import (
	"flag"
	"fmt"
	"github.com/miekg/dns"
	"os"
)

var (
	Version = flag.Bool("v", false, "版本号")
	Timeout = flag.Int("t", 10, "超时时间")
	DNS     = flag.String("h", "", "指定DNS服务器ip")
	Port    = flag.String("p", "53", "指定DNS端口")
	Args    []string
)

func InitFlag() {

	//解析参数
	flag.Parse() //解析命令行参数
	Args = flag.Args()

	// 校验参数
	if len(Args) < 1 {
		fmt.Println("参数缺失，请输出ip")
		os.Exit(1)
	}

	// 初始化dns参数
	if *DNS == "" {
		dnsConfig, err := dns.ClientConfigFromFile("/etc/resolv.conf")
		if err != nil {
			fmt.Println("无法读取resolv.conf文件:", err)
			os.Exit(1)
		}
		if len(dnsConfig.Servers) == 0 {
			fmt.Println("resolv.conf文件中未配置dns")
			os.Exit(1)
		}
		*DNS = dnsConfig.Servers[0]
	}

	// 检查域名
	Args[0] = dns.Fqdn(Args[0])
}
