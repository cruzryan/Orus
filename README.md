# Orus
Orus makes debugging vhdl with IntelModel Sim easier, you can do it right from the command line!


## Requirements: 
1. Have Intel ModelSim installed
2. Have Go installed (to run from source)


## How to run:

`go run .`

From executable:

`Orus.exe [file you want to run]`

Example:

`Orus.exe C:/Users/rncb0/Code/VHDL/hw17.vhd`

## Bugs

1. You need to save your file twice for Orus to update the signals
2. Takes ~1s to compile and run. How do we make it run faster?