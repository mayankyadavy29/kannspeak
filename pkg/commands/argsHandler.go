package commands

import (
	"fmt"
	"strconv"
)

func Args1(arg string) {
	switch arg {
	case "run":
		run()
	case "help":
		help()
	case "cfg":
		cfg()
	case "full":
		full()
	default:
		fmt.Println("Incorrect argument provided. For more info do `help`.")
	}
}

func Args2(arg1, arg2 string) {
	switch arg1 {
	case "pool":
		pool(arg2)
	case "len":
		l, err := strconv.Atoi(arg2)
		if err != nil {
			fmt.Println("You can only pass integer as a length.")
			return
		}
		leng(l)
	case "ask":
		ask(arg2)
	case "full":
		fullWithType(arg2)
	default:
		fmt.Println("Incorrect argument provided. For more info do `help`.")
	}
}
