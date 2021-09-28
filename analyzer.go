package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func analyze() {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "entity") {
			splt := strings.Split(line, " ")
			current_entity = strings.ToLower(splt[1])
		}

		if (strings.Contains(line, "std_logic") || strings.Contains(line, "STD_LOGIC_VECTOR")) && !strings.Contains(line, "use") {
			splt := strings.Split(line, ":")
			name := strings.ReplaceAll(splt[0], "\t", "")
			val := "U"

			if len(splt) == 3 {
				p1 := string(splt[2][len(splt[2])-4])
				p2 := string(splt[2][len(splt[2])-3])

				//Very bad fix
				if p1 == "0" || p1 == "1" || p1 == "U" {
					val = p1
				}

				if p2 == "0" || p2 == "1" || p2 == "U" {
					val = p2
				}
			}

			found := false

			for i := 0; i < len(vhdl_vars); i++ {
				if vhdl_vars[i].name == name {
					vhdl_vars[i].value = string(val)
					found = true
				}
			}

			if current_entity == "" {
				panic("No entity found in file!")
			}

			if !found {
				vhdl_vars = append(vhdl_vars, VHDL_VAR{name: name, value: string(val), entity: current_entity})
			}
		}

		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
