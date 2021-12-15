package util

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"syscall"
)

const PORT = 9999

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""

}

func GetLocalIP1() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP.To4().String()
}

func GetLogicalDrives() []string {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
	n, _, _ := GetLogicalDrives.Call()
	s := strconv.FormatInt(int64(n), 2)
	var drives_all = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:", "J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：", "U：", "V：", "W：", "X：", "Y：", "Z："}
	temp := drives_all[0:len(s)]
	var d []string
	for i, v := range s {
		if v == 49 {
			l := len(s) - i - 1
			d = append(d, temp[l])
		}
	}
	var drives []string
	for i, v := range d {
		drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
	}
	return drives
}

func GetVisitHTML() string {
	template := ""
	ip := GetLocalIP1()
	port := PORT
	http_host := fmt.Sprintf("http://%s:%d", ip, port)
	drives := GetLogicalDrives()
	for _, drive := range drives {
		d := strings.ToLower(drive)
		d = d[0:1]
		link := http_host + "/" + d + "/"
		node := fmt.Sprintf("<h3><a href='%s'>%s</a></h3>", link, d)
		template += node
	}
	return template
}
