package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func validIP(ip string) bool {
	if len(ip) == 0 {
		return false
	}
	ips := strings.Split(ip, ".")

	// ip不是四段
	if len(ips) != 4 {
		return false
	}
	for i := 0; i < 4; i++ {
		data, err := strconv.Atoi(ips[i])
		if err != nil {
			return false
		}

		if data < 0 || data > 255 || (data != 0 && ips[0] == "0") {
			return false
		}
	}
	return true
}

func main() {
	var input string
	for {
		_, err := fmt.Scan(&input)
		if err == io.EOF {
			break
		}

		if validIP(input) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
