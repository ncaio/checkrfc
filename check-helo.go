package main

//
//
//
import (
	"fmt"
	"github.com/bobesa/go-domain-util/domainutil"
	"net"
	"os"
	"net/smtp"
)

//
//
//
func main() {
	arg := os.Args[1]
	rdns, _ := net.LookupAddr(arg)
	ips, _ := net.LookupIP(rdns[0])
	size := len(rdns[0])
	if size > 0 && rdns[0][size-1] == '.' {
		rdns[0] = rdns[0][:size-1]
	}
	fmt.Println("------------------------------")
	fmt.Println("rDns: ", rdns[0])
//
//
//
	fmt.Println("------------------------------")
	fmt.Println("A or AAAA records: ")
	for l := range ips{
		fmt.Println(ips[l])
}
//
//
//
	domain := (domainutil.Domain(rdns[0]))
//
//
//
	fmt.Println("------------------------------")
	txt, _ := net.LookupTXT(domain)
	fmt.Println("TXT: ")
	for p := range txt{
		fmt.Println(txt[p])
	}
//
//
//
	fmt.Println("------------------------------")
	mx, _ := net.LookupMX(domain)
	fmt.Println("MX:")
	for i := range mx {
		fmt.Println(mx[i].Host)
		c, err := smtp.Dial(mx[i].Host+":25")
		c.Quit()
		if err != nil{
			fmt.Println("Erro ao conectar")
		}else{
			fmt.Println("Conexao OK")
        		conn, err := net.Dial("tcp", mx[i].Host+":25")
        		if err != nil {
                		fmt.Println(err)
                		os.Exit(1)
        		}
        		defer conn.Close()
        		var readbuf [512]byte
        		n, _ := conn.Read(readbuf[0:])
        		os.Stdout.Write(readbuf[0:n])
		}			
	}
	fmt.Println("------------------------------")
}
