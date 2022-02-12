package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {

	wm := emoji.Sprint(":world_map:!")
	fmt.Println("Hello", wm)

}
