package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/go-p5/p5"
)

var (
	win32_path  = "C:\\intelFPGA\\20.1\\modelsim_ase\\win32aloem\\"
	path        = "C:/Users/rncb0/Code/VHDL/hw17.vhd"
	log_stats   = true
	vsim_writer io.WriteCloser

	examine_in_process = false
	var_to_examine     = ""
	current_entity     = ""
)

var vhdl_vars []VHDL_VAR

type VHDL_VAR struct {
	name   string
	value  string
	entity string
}

func printSig(col color.Attribute, sig, val string) {
	color.Set(col)
	fmt.Println(sig + " | " + val)
	color.Unset()
}

func truthTable() {
	fmt.Println("VLENGTH: ", len(vhdl_vars))
	for i := 0; i < len(vhdl_vars); i++ {
		switch vhdl_vars[i].value {
		case "0":
			printSig(color.FgRed, vhdl_vars[i].name, vhdl_vars[i].value)
			break
		case "1":
			printSig(color.FgGreen, vhdl_vars[i].name, vhdl_vars[i].value)
			break
		default:
			printSig(color.FgYellow, vhdl_vars[i].name, vhdl_vars[i].value)
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
	io.WriteString(vsim_writer, "examine /"+arch+"/"+v+"\n")
}

func examineAll() {
	for i := 0; i < len(vhdl_vars); i++ {
		examine(vhdl_vars[i].entity, vhdl_vars[i].name)
	}
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

			if strings.Contains(o, "examine") {
				examine_in_process = true
				var_to_examine = strings.TrimSuffix(strings.Split(o, "/")[2], "\n")
				// color.Red("EXAMINE")
			}

			if examine_in_process && o[0] == '#' {
				examine_in_process = false
				found := false
				for i := 0; i < len(vhdl_vars); i++ {
					if vhdl_vars[i].name == var_to_examine {
						vhdl_vars[i].value = string(o[2])
						found = true
					}
				}
				if !found {
					fmt.Println("------------ADDED, SZ : ", len(vhdl_vars))
					vhdl_vars = append(vhdl_vars, VHDL_VAR{name: var_to_examine, value: string(o[2])})
				}
			}
			fmt.Println(o)
		}
	}
}

func stopVsim() {
	color.Red("Stopping vsim!")
	io.WriteString(vsim_writer, "exit\n")
}

func restartVsim() {
	color.Red("Restarting vsim!")
	//Reset values
	vhdl_vars = nil
	io.WriteString(vsim_writer, "restart\n")
}

func run() {
	color.Red("Running!")
	// io.WriteString(vsim_writer, "force -freeze sim:/switch/A 0 0\n")
	// io.WriteString(vsim_writer, "force -freeze sim:/switch/B 1 0\n")
	io.WriteString(vsim_writer, "run\n")
}

func compile() {

	//Reset values
	vhdl_vars = nil

	//Make the lib at work
	go vlib()
	//vcom compile path
	go vcom()

}

func main() {

	color.Green("Starting Orus...")

	//Analyze file
	analyze()

	//Hot reload file watching
	go watch()

	//Close program when user hits CTRL+C
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		color.Cyan("Closing Orus...")
		os.Exit(1)
	}()
	compile()
	go p5.Run(setup, draw)
	startVsim()
}
