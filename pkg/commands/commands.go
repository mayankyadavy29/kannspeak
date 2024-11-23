package commands

import (
	"encoding/json"
	"fmt"
	"github.com/mayankyadavy29/kannspeak/pkg/config"
	"github.com/mayankyadavy29/kannspeak/pkg/constants"
	"github.com/mayankyadavy29/kannspeak/pkg/pools"
	"math/rand"
	"strings"
	"time"
)

func prettyEntry(entry string) []string {
	return strings.Split(entry, ":")
}

func random(len int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(len)
}

func run() {
	cfg := config.GetConfig()
	curPool := cfg.Pool
	if cfg.Pool == constants.RANDOM {
		idx := random(len(pools.PoolMap))
		keys := make([]constants.PoolEnum, 0, len(pools.PoolMap))
		for key := range pools.PoolMap {
			keys = append(keys, key)
		}
		curPool = keys[idx]
	}
	rd := random(cfg.MaxPoolLen)
	entry := pools.PoolMap[curPool][rd%len(pools.PoolMap[curPool])]
	pEntry := prettyEntry(entry)
	fmt.Println("Your word is :", pEntry[0])
	fmt.Println("What is its kannada translation?")
	var kannWord string
	_, err := fmt.Scanln(&kannWord)
	if err != nil {
		fmt.Println("error while reading the input")
		return
	}
	fmt.Println("Correct translation is:", pEntry[1])
}

func pool(poolType string) {
	cfg := config.GetConfig()
	p := pools.ConvertStrToPool(poolType)
	if p == "" {
		return
	}
	cfg.Pool = p
	config.SetConfig(cfg)
	fmt.Println("Successfully updated the pool type to", poolType)
}

func leng(poolLen int) {
	cfg := config.GetConfig()
	curLen := len(pools.PoolMap[cfg.Pool])
	if poolLen > curLen {
		fmt.Println("Maximum pool length is ", curLen, ". You cannot exceed it!!")
	}
}

func ask(engWord string) {
	for _, v := range pools.PoolMap {
		for _, e := range v {
			pEntry := prettyEntry(e)
			if pEntry[0] == engWord {
				fmt.Println("The correct translation for your english word is:", pEntry[1])
				return
			}
		}
	}
	fmt.Println("Could not find any translation for your english word. Please check the word again!")
}

func help() {
	s := `run : Give user an english word. Get translation from user and return correct kannada word.
cfg : Return the current config
len : Update the length of current pool. This tells how many words would be used to get a random word.
ask : You can ask for kannada translation of an english word. It can be a small number initially. When you become comfortable then you can increase the limit.
pool : Current pool of words from which a random word would be picked up. Available pool types : [random, nouns, pronouns, verbs, numbers, relationships].
full : Return the full list of translations in pretty format. If pool type is given, then return the full list of only that specific pool else return global list.`
	fmt.Println(s)
}

func cfg() {
	cfg := config.GetConfig()
	cfgBytes, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		fmt.Println("Could not get config. Restart the app")
		return
	}
	fmt.Println(string(cfgBytes))
}

func full() {
	fmt.Println("English word          ||          Kannada word")
	fmt.Println("-----------------------------------------")
	for k := range pools.PoolMap {
		fullWithType(string(k))
	}
}

func fullWithType(poolType string) {
	p := pools.ConvertStrToPool(poolType)
	if p == "" {
		return
	}
	for _, e := range pools.PoolMap[p] {
		pEntry := prettyEntry(e)
		fmt.Println("    ", pEntry[0], "<----------------->", pEntry[1])
	}
}
