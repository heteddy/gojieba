package gojieba

import (
	"github.com/heteddy/gojieba/deps/cppjieba"
	"github.com/heteddy/gojieba/deps/limonp"
	"github.com/heteddy/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
