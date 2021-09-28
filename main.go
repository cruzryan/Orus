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
	win32_path  = "C:\\intelFPGA\\20.1\\modelsim_ase\\win32aloem\\"
	path        = "C:/Users/rncb0/Code/VHDL/hw17.vhd"
	log_stats   = true
	vsim_writer io.WriteCloser
)

func printSig(col color.Attribute, sig, val string) {
	color.Set(col)
	fmt.Println(sig + " | " + val)
	color.Unset()
}

func truthTable(vals, sigs []string) {
	// sigs := []string{"A", "B", "C"}
	// vals := []string{"0", "1", "U"}

	for i := 0; i < len(sigs); i++ {
		switch vals[i] {
		case "0":
			printSig(color.FgRed, sigs[i], vals[i])
			break
		case "1":
			printSig(color.FgGreen, sigs[i], vals[i])
			break
		case "U":
			printSig(color.FgYellow, sigs[i], vals[i])
			break
		}
	}
}

func check(err error, where string) {
	if err != nil {
		color.Red("Something went wrong at ", where)
		panic("Program stopping!")
	}
}

func examine(arch string, v string) {
	io.WriteString(vsim_writer, "examine /switch/a\n")
}

func ReadOutput(output chan string, rc io.ReadCloser) {
	r := bufio.NewReader(rc)
	for {
		x, _ := r.ReadString('\n')
		output <- string(x)
	}
}

func vlib() {
	vlib_out, err := exec.Command("cmd", "/C", "vlib work").Output()
	check(err, "vlib work")
	if log_stats {
		fmt.Println(string(vlib_out))
		color.Green("vlib done.")
	}
}

func vcom() {
	vcom_out, err := exec.Command("cmd", "/C", "vcom -work work -2002 -explicit -stats=none", path).Output()
	check(err, "vcom")
	if log_stats {
		fmt.Println(string(vcom_out))
		color.Green("vcom done.")
	}
}

func startVsim() {
	if log_stats {
		color.Cyan("Starting vsim...")
	}

	vsim := exec.Command("cmd", "/C", "vsim -c work.switch")
	vsimIn, _ := vsim.StdinPipe()
	vsimOut, _ := vsim.StdoutPipe()
	vsim_writer = vsimIn

	vsim.Start()

	output := make(chan string)
	defer close(output)
	go ReadOutput(output, vsimOut)
	for o := range output {
		if len(o) > 0 {
			fmt.Println(o)
		}
	}
}

func stopVsim() {
	io.WriteString(vsim_writer, "exit\n")
}

func compileAndRun() {
	//Make the lib at work
	vlib()
	//vcom compile path
	vcom()
	//start vsim
	// startVsim()
}

func main() {

	color.Green("Starting Orus...")

	//Hot reload file watching
	go watch()

	//Close program when user hits CTRL+C
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		color.Cyan("Closing vsim...")
		os.Exit(1)
	}()

	compileAndRun()
	startVsim()
}
