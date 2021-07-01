package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	a := make([]int, 1000)
	for i := 0; i < 100; i++ {
		a[i] = i
	}
	b, _ := json.Marshal(a[:10])
	//b, _ := json.Marshal(a[:100])

	var cmd0 *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd0 = exec.Command("cmd", "/c", "echo", string(b))
	} else {
		//cmd0 := exec.Command("echo", "-n", "My first command comes from golang.")
		cmd0 = exec.Command("echo", "-n", string(b))
	}

	// 开启管道
	stdout, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Couldn't obtain the stdout pipe for command No.0: %s\n", err)
		return
	}

	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The command No.0 can not be startup: %s\n", err)
		return
	}

	output0 := make([]byte, 30)
	n, err := stdout.Read(output0)
	fmt.Println(n)
	if err != nil {
		fmt.Printf("Error: Couldn't  read data from the pipe: %s\n", err)
	}
	fmt.Printf("%s\n", output0[:n])

}
