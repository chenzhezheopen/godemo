package main

import (
	"fmt"
	"net"
	"os"
)

// func main() {
// 	fmt.Println(get_external())
// }
// func get_external() string {
// 	resp, err := http.Get("http://www.baidu.com")
// 	if err != nil {
// 		return ""
// 	}
// 	defer resp.Body.Close()
// 	content, _ := ioutil.ReadAll(resp.Body)
// 	//buf := new(bytes.Buffer)
// 	//buf.ReadFrom(resp.Body)
// 	//s := buf.String()
// 	return string(content)
// }
// func main() {

// 	domain := "www.baidu.com"
// 	ipAddr, err := net.ResolveIPAddr("ip", domain)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
// 		return
// 	}
// 	fmt.Fprintf(os.Stdout, "%s IP: %s Network: %s Zone: %s\n", ipAddr.String(), ipAddr.IP, ipAddr.Network(), ipAddr.Zone)

// 	// 百度，虽然只有一个域名，但实际上，他对应电信，网通，联通等又有多个IP地址
// 	ns, err := net.LookupHost(domain)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
// 		return
// 	}

// 	for _, n := range ns {
// 		fmt.Fprintf(os.Stdout, "%s", n) // 115.239.210.26    115.239.210.27 这2个地址打开都是百度
// 	}
// }
func main() {

	port, err := net.LookupPort("tcp", "https") // 查看telnet服务器使用的端口

	if err != nil {
		fmt.Fprintf(os.Stderr, "未找到指定服务")
		return
	}

	fmt.Fprintf(os.Stdout, "telnet port: %d", port)

}
