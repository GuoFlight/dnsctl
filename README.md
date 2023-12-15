# 简介

dnsctl是dns相关的客户端 ，类似于nslookup、dig等命令
为什么要做一个dns客户端？初衷nslookup和dig命令的可移植性不高，需要依赖各种库。

# 作者

京城郭少

# 编译

```shell
go build -o dnsctl
chmod +x dnsctl
./dnsctl baidu.com
```

# Demo

```shell
./dnsctl baidu.com          # 解析baidu.com
```

# 参数

```shell
-v              # 输出dnsctl的版本号
-h 8.8.8.8      # 指定DNS服务器
-p 53           # 指定DNS端口
-t 10           # 指定超时时间(单位：s)
```