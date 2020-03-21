package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	//request line
	method := strings.Fields(ln)[0]
	uri := strings.Fields(ln)[1]
	fmt.Println("METHOD -> ", method)
	fmt.Println("URI -> ", uri)

	//multiplexer
	if method == "GET" && uri == "/" {
		index(conn)
	}
	if method == "GET" && uri == "/post" {
		post(conn)
	}
	if method == "POST" && uri == "/post" {
		postProcess(conn)
	}
}

func index(conn net.Conn) {
	body := `
		<!DOCTYPE html>
			<html lang="en">
			<head><meta charset="utf-8">
				<title>The Go Programming Language</title>
			</head>
				<body>
					<strong>Index</strong><br>
					<a href="/">index</a><br>
					<a href="/post">post</a><br>
				</body>
			</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func post(conn net.Conn) {
	body := `
		<!DOCTYPE html>
			<html lang="en">
			<head><meta charset="utf-8">
				<title>The Go Programming Language</title>
			</head>
				<body>
					<strong>Post</strong><br>
					<a href="/">index</a><br>
					<a href="/post">post</a><br>
					<form method="post" action="/post">
						<input type="submit" value="post">
					</form>
				</body>
			</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func postProcess(conn net.Conn) {
	body := `
		<!DOCTYPE html>
			<html lang="en">
			<head><meta charset="utf-8">
				<title>The Go Programming Language</title>
			</head>
				<body>
					<strong>Post Process</strong><br>
					<a href="/">index</a><br>
					<a href="/post">post</a><br>
				</body>
			</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
