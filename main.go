package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
)

var (
	win32_path = "C:\\intelFPGA\\20.1\\modelsim_ase\\win32aloem\\"
	test_path  = "C:/Users/rncb0/Code/VHDL/hw17.vhd"
	log_stats  = true
)

func truthTable() {

}

func check(err error, where string) {
	if err != nil {
		color.Red("Something went wrong at ", where)
		panic("Program stopping!")
	}
}

func examine(cmdIn io.WriteCloser, arch string, v string) {
	cmdIn.Write([]byte("examine /switch/a"))
	cmdIn.Close()
}

func main() {
	color.Green("Starting Orus...")
	//Make the lib at work
	vlib_out, _ := exec.Command("cmd", "/C", "vlib work").Output()
	if log_stats {
		fmt.Println(string(vlib_out))
		color.Green("vlib done.")
	}

	//vcom path
	vcom_out, _ := exec.Command("cmd", "/C", "vcom -work work -2002 -explicit -stats=none", test_path).Output()
	if log_stats {
		fmt.Println(string(vcom_out))
		color.Green("vcom done.")
	}

	//start vsim subprocess
	if log_stats {
		color.Cyan("Starting vsim...")
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		color.Cyan("Closing vsim...")
		os.Exit(1)
	}()
	vsim := exec.Command("cmd", "/C", "vsim -c work.switch")
	vsimIn, _ := vsim.StdinPipe()
	vsimOut, _ := vsim.StdoutPipe()
	buf := bufio.NewReader(vsimOut)
	vsim.Start()
	for {
		line, _, _ := buf.ReadLine()
		fmt.Println(string(line))
		examine(vsimIn, "switch", "a")
		// examine(vsimIn, "switch", "b")
		vsimBytes, _ := ioutil.ReadAll(vsimOut)
		time.Sleep(100 * time.Millisecond)
		fmt.Println(string(vsimBytes))
		vsim.Wait()
	}

}
