package pools

import (
	"fmt"
	"github.com/mayankyadavy29/kannspeak/pkg/constants"
)

// PoolMap will contain list of all the pools identified by their constants.PoolEnum
var PoolMap map[constants.PoolEnum][]string

// init PoolMap with all the pools
func init() {
	PoolMap = make(map[constants.PoolEnum][]string)
	PoolMap[constants.NOUNS] = nouns
	PoolMap[constants.PRONOUNS] = pronouns
	PoolMap[constants.NUMBERS] = numbers
	PoolMap[constants.VERBS] = verbs
	PoolMap[constants.RELATIONS] = relations
}

func ConvertStrToPool(poolType string) constants.PoolEnum {
	switch poolType {
	case string(constants.NOUNS):
		return constants.NOUNS
	case string(constants.PRONOUNS):
		return constants.PRONOUNS
	case string(constants.NUMBERS):
		return constants.NUMBERS
	case string(constants.VERBS):
		return constants.VERBS
	case string(constants.RANDOM):
		return constants.RANDOM
	case string(constants.RELATIONS):
		return constants.RELATIONS
	default:
		fmt.Println("You can only change pool type to the pre-defined types: [random, nouns, pronouns, verbs, numbers, relationships]. For more info do `help`.")
	}
	return ""
}
