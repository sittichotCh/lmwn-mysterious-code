package main

import (
	"fmt"
	"github.com/ongSitti/lmwn-mysterious-code/util"
)

func main() {
	const secret = "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	message, err := util.DecodeBase64String(secret)
	if err != nil {
		panic(err)
	}

	fmt.Printf("message : %s", message)
}
