package main

import (
	"flag"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net"
)

var (
	listenAddr string
	dialAddr   string
	seed       = flag.Uint("seed", 0, "Seed for mask sequence")
)

func init() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalln("need 2 arguments: <listenAddr> <dialAddr>")
	}
	listenAddr = flag.Arg(0)
	dialAddr = flag.Arg(1)
}

func main() {
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept() failed.")
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(connIn net.Conn) {
	defer connIn.Close()
	connOut, err := net.Dial("tcp", dialAddr)
	if err != nil {
		log.Println("Dial() failed.")
		return
	}
	defer connOut.Close()
	tunnel := connIn.RemoteAddr().String() + "<->" +
		connOut.RemoteAddr().String()
	log.Println(tunnel, "established.")
	defer log.Println(tunnel, "closed.")
	ch := make(chan struct{}, 1)
	transfer := func(dst, src net.Conn) {
		io.Copy(dst, transform.NewReader(src, new(Transformer)))
		ch <- struct{}{}
	}
	go transfer(connIn, connOut)
	go transfer(connOut, connIn)
	<-ch
}

type Transformer struct {
	seed uint32
}

func (t *Transformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	if len(dst) < len(src) {
		err = transform.ErrShortDst
		return
	}
	for i, ch := range src {
		t.seed++
		dst[i] = ch ^ byte(t.seed)
	}
	nDst = len(src)
	nSrc = len(src)
	return
}

func (t *Transformer) Reset() {
	t.seed = uint32(*seed)
}
