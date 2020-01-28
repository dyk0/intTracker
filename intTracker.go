package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/dyk0/intTracker/stringstuff"
)

//Setup constant variables for the server connection settings and intTracker logic

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8888"
	CONN_TYPE = "tcp"
	NUM_MIN   = 0
	NUM_MAX   = 4294967295
)

// Let's create the mapping, each bucket number will have an array of names, called record
// First, we create an empty slice of max numbersize
type record struct {
	name string
	min  uint32
	max  uint32
}

var collection = make([][]string, 100)

func main() {
	// Incomming connections
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		//Listen for an incomming Connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		//Handle connection in a new goroutine
		go handleRequest(conn)
	}

}

// Handles incoming requests
// func addCommand(s []string) map[string]*record {
func addCommand(s []string, conn net.Conn) {
	if len(s) == 4 {
		if !stringstuff.Parse(s[3]) {
			e := newRecord(s)
			for i := e.min; i <= e.max; i++ {
				old_names := collection[i]
				names := append(old_names, e.name)
				collection[i] = names
			}
			conn.Write([]byte("OK" + "\n"))
		} else {
			fmt.Println(stringstuff.Parse(s[3]))
			fmt.Println("Incorrect Name " + s[3])
		}
	} else {
		fmt.Println("ERROR" + "\n")
		conn.Write([]byte("ERROR" + "\n"))
	}
}

func delCommand(s []string, conn net.Conn) {
	// create a new record
	e := newRecord(s)
	// if the name is empty, then we merc everything
	if e.name != "" {
		for i := e.min; i <= e.max; i++ {
			record := collection[i]
			// find and replace the name
			for j, v := range record {
				if v == e.name {
					record = append(record[:j], record[j+1:]...)
					collection[i] = record
					break
				}
			}
		}
	} else {

		for i := e.min; i <= e.max; i++ {
			record := collection[i]
			// merc everything`
			collection[i] = record[:0]
		}
	}
	conn.Write([]byte("OK" + "\n"))
}
func findCommand([]string) {}

func newRecord(s []string) record {
	r := record{}
	if len(s) == 4 {
		entry := strings.TrimSuffix(s[3], "\n")
		r.name = entry

	} else {

		r.name = ""
	}

	min, err := strconv.Atoi(s[1])
	if err != nil {
		fmt.Println("ERROR" + "\n")
	}
	r.min = uint32(min)
	max, err := strconv.Atoi(strings.TrimSuffix(s[2], "\n"))
	if err != nil {
		fmt.Println("ERROR" + "\n")
	}

	r.max = uint32(max)
	return r
}

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		fmt.Println("...client left...")
		conn.Close()
		return
	}
	message := strings.Split(string(buffer), " ")
	switch message[0] {
	case "ADD":
		fmt.Println("Add command received")
		addCommand(message, conn)
		fmt.Println(collection)
	case "DEL":
		fmt.Println("Del command received")
		delCommand(message, conn)
		fmt.Println(collection)
	case "FIND":
		fmt.Println("Find command received")
		findCommand(message)
	default:
		conn.Write([]byte("ERROR invailid arguments" + "\n"))
		conn.Close()
	}
	handleRequest(conn)
}
