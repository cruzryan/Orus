import os 
import sys

from colorama import init, Fore, Back, Style
init()

#a = 1 n
#a = 0 
#a = 0, b = 1, c = 1, d = 3 
#a = 0010 

win32_path = "C:\\intelFPGA\\20.1\\modelsim_ase\\win32aloem\\"
test_path = "C:/Users/rncb0/Code/VHDL/hw17.vhd"

#On user toggle
def showTable():
	pass


#move to utils or something
def log(txt):
	print(txt)
	print(Style.RESET_ALL)

#Yup yup
def hotreload():
	pass

def compile():

	#Work libs
	os.system(win32_path + "vlib work")

	#Compile
	os.system(win32_path + "vcom -work work -2002 -explicit -stats=none " + test_path)

	#Simulate
	os.system("vsim -c work.switch");


if __name__ == "__main__":
	log(Fore.GREEN + "Starting Orus...")
	compile()