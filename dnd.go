package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var debug = flag.Bool("debug", false, "Enable Debug Mode")

func dprint(s interface{}) {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *debug {
		fmt.Println(s)
	}
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	stdin := bufio.NewScanner(os.Stdin)
	var stinput = func(desc string) string {
		fmt.Printf("%s> ", desc)
		stdin.Scan()
		return strings.TrimSpace(stdin.Text())
	}
	var itinput = func(desc string) int64 {
		val, _ := strconv.ParseInt(stinput(desc), 10, 64)
		return val
	}
	cmdLoop(stinput, itinput)
} // Main
