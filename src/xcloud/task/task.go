package task

import (
	//"os/exec"
	"time"
	"fmt"
	"errors"
)

func Build(a string) (int, error) {
	for i := 0; i <= 20; i++ {
		fmt.Println(a, "core bbb init")
		time.Sleep(3 * time.Second)
	}
	return -1, errors.New("")
}
