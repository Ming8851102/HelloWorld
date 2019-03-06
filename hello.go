package main

import (
	"encoding/json"
	"fmt"
	"helloworld/printdata"
	"os"
	"os/exec"
)

type School struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Call shell command wget to download. The reason to use wget is that wget
// supports automatically resume download. So this package only runs on Linux
// systems.
func main() {
	fmt.Println("Hello, world.")
	//var s *School
	params := School{ID: 123, Name: "Primary School"}
	//fmt.Println(params)
	//printdata.Rec(params)

	d, _ := json.Marshal(params)
	printdata.RecByte(d)
	//printdata.Rec(params)
	//fmt.Println(d)

}


func wget(url, filepath string) error {
	// run shell `wget URL -O filepath`
	cmd := exec.Command("wget", url, "-O", filepath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
