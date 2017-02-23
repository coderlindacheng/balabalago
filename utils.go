package utils

import (
	"bufio"
	"os"
	"fmt"
	"github.com/coderlindacheng/balabalago/special_string"
)

func Pause() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请按任意键继续")
	reader.ReadRune()
	os.Exit(0)
}
