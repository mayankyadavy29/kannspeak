package constants

import (
	"os"
	"path/filepath"
)

// PoolEnum defines all types of available pools to learn from
type PoolEnum string

const (
	RANDOM    PoolEnum = "random"
	PRONOUNS  PoolEnum = "pronouns"
	NOUNS     PoolEnum = "nouns"
	NUMBERS   PoolEnum = "numbers"
	VERBS     PoolEnum = "verbs"
	RELATIONS PoolEnum = "relations"
)

var (
	user, _     = os.UserHomeDir()
	DefWinCfg   = filepath.Join(user, "Documents", "KannSpeak", "kannspeak.json")
	DefLinuxCfg = filepath.Join("usr", "local", "etc")
)
