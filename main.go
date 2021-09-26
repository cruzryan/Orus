package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

var (
	win32_path = "C:\\intelFPGA\\20.1\\modelsim_ase\\win32aloem\\"
	test_path  = "C:/Users/rncb0/Code/VHDL/hw17.vhd"
)

func main() {
	fmt.Println("sup boys")

	//Make the lib at work
	vlib_out, _ := exec.Command("cmd", "/C", "vlib work").Output()
	fmt.Println(string(vlib_out))

	//vcom path
	vcom_out, _ := exec.Command("cmd", "/C", "vcom -work work -2002 -explicit -stats=none", test_path).Output()
	fmt.Println(string(vcom_out))

	grepCmd := exec.Command("cmd", "/C", "vsim -c work.switch")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("examine /switch/a"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println(string(grepBytes))

}
