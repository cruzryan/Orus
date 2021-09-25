import os 
import sys

win32_path = "C:\\intelFPGA\\20.1\\modelsim_ase\\win32aloem\\"
test_path = "C:/Users/rncb0/Code/VHDL/hw17.vhd"

def compile():

	#Go to folder
	os.system("cd " + test_path)

	#Work libs
	os.system(win32_path + "vlib work")

	#Compile
	os.system(win32_path + "vcom -work work -2002 -explicit -stats=none " + test_path)

	#Simulate
	os.system("vsim -c work.switch");


if __name__ == "__main__":
	print("Starting Orus...")
	compile()