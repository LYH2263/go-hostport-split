# go-hostport-split

IPv6 hostport 拆错

internal/netx/hostport.go 的 Split：按第一个冒号切割，未处理 [ipv6]:port
