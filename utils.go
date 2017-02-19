package utils

import (
	"bufio"
	"os"
	"fmt"
)

func Pause() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请按任意键继续")
	reader.ReadRune()
	os.Exit(0)
}