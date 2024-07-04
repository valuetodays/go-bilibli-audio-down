package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func Test1(t *testing.T) {
	var commandArgs []string
	commandArgs = append(commandArgs, "https://www.bilibili.com/audio/music-service-c/web/url?sid=4401746")
	var commandArgsForFileLength = append(commandArgs, " -I ")
	//var cmdForFileLength = commonCmd + " -I "
	command := exec.Command("curl", commandArgsForFileLength...)
	output, err := command.Output()
	fmt.Println("err", err)
	fmt.Println("output", string(output))

}
