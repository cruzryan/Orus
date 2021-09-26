package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

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
	io.WriteString(cmdIn, "examine /switch/a\n")
}

func ReadOutput(output chan string, rc io.ReadCloser) {
	r := bufio.NewReader(rc)
	for {
		x, _ := r.ReadString('\n')
		output <- string(x)
	}
}

func main() {
	color.Green("Starting Orus...")
	//Make the lib at work
	vlib_out, err := exec.Command("cmd", "/C", "vlib work").Output()
	check(err, "vlib work")
	if log_stats {
		fmt.Println(string(vlib_out))
		color.Green("vlib done.")
	}

	//vcom path
	vcom_out, err := exec.Command("cmd", "/C", "vcom -work work -2002 -explicit -stats=none", test_path).Output()
	check(err, "vcom")
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

	vsim.Start()
	go examine(vsimIn, "switch", "a")

	output := make(chan string)
	defer close(output)
	go ReadOutput(output, vsimOut)

	for o := range output {
		fmt.Println(o)
	}
}
