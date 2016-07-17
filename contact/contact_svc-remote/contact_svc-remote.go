// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"contact"
	"flag"
	"fmt"
	"git-wip-us.apache.org/repos/asf/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  Contact create(Contact contact)")
	fmt.Fprintln(os.Stderr, "  Contact read(string contactId)")
	fmt.Fprintln(os.Stderr, "  Contact update(Contact contact)")
	fmt.Fprintln(os.Stderr, "  void destroy(string contactId)")
	fmt.Fprintln(os.Stderr, "   fetch()")
	fmt.Fprintln(os.Stderr, "  void reset()")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := contact.NewContactSvcClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "create":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Create requires 1 args")
			flag.Usage()
		}
		arg15 := flag.Arg(1)
		mbTrans16 := thrift.NewTMemoryBufferLen(len(arg15))
		defer mbTrans16.Close()
		_, err17 := mbTrans16.WriteString(arg15)
		if err17 != nil {
			Usage()
			return
		}
		factory18 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt19 := factory18.GetProtocol(mbTrans16)
		argvalue0 := contact.NewContact()
		err20 := argvalue0.Read(jsProt19)
		if err20 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Create(value0))
		fmt.Print("\n")
		break
	case "read":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Read requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.Read(value0))
		fmt.Print("\n")
		break
	case "update":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Update requires 1 args")
			flag.Usage()
		}
		arg22 := flag.Arg(1)
		mbTrans23 := thrift.NewTMemoryBufferLen(len(arg22))
		defer mbTrans23.Close()
		_, err24 := mbTrans23.WriteString(arg22)
		if err24 != nil {
			Usage()
			return
		}
		factory25 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt26 := factory25.GetProtocol(mbTrans23)
		argvalue0 := contact.NewContact()
		err27 := argvalue0.Read(jsProt26)
		if err27 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Update(value0))
		fmt.Print("\n")
		break
	case "destroy":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Destroy requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.Destroy(value0))
		fmt.Print("\n")
		break
	case "fetch":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Fetch requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Fetch())
		fmt.Print("\n")
		break
	case "reset":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Reset requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Reset())
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
