package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/mayankyadavy29/kannspeak/pkg/commands"
	"os"
	"os/signal"
	"strings"
)

/*
	Accept all instructions

• run: Return the english word to translate. Get the user input. Return the correct kannada word.

• pool pool_type : Change the pool type. You can only change it to the pre-defined types. For more info do `help`.

• len pool_len : Change the number of words in random pool. (It can be a small number initially. When you become comfortable then you can increase the limit).

• ask eng_word : Return the kannada word for this eng word

• full pool_type : Return the full list of that pool type. If no pool_type, then return global list

• cfg : Return the current config

• help : List down all the methods available from `help`.
run : Give user an english word. Get translation from user and return correct kannada word.
cfg : Return the current config
len : Update the length of current pool. This tells how many words would be used to get a random word.
ask : You can ask for kannada translation of an english word.
pool : Current pool of words from which a random word would be picked up. List down all the available pool types.
full : Return the full list of translations in pretty format. If pool type is given, then return the full list of only that specific pool else return global list.
*/
func main() {
	fmt.Println("From Can't Speak to Kann Speak!")
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
			return
		default:
			fmt.Println("\nSelect command: run | cfg | pool pool_type | len pool_len | ask eng_word | full pool_type | help")
			scanner.Scan()
			line := scanner.Text()
			inpList := strings.Split(line, " ")
			switch len(inpList) {
			case 1:
				commands.Args1(inpList[0])
			case 2:
				commands.Args2(inpList[0], inpList[1])
			}
		}
	}
}
