package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var serifs = map[string][]string{
	"nobue": []string{
		"えーはじめまして。あたしは伊藤伸恵。ご覧のとおりごく普通の16歳です",
		"長い付き合いだけど、こいつの頭ん中はよく分かりません",
		"んじゃ、失礼",
	},
	"chika": []string{
		"いい加減、玄関から入りなよみっちゃん",
		"禁煙しろ",
		"おまえ、アホだろ",
	},
	"miu": []string{
		"君らここでなにしてんの？",
		"灰皿を作りませんか？",
		"よし、それでは今回作る灰皿の見本だが・・・",
	},
}

func main() {
	cli.AppHelpTemplate = `Name:
	ichigo-cli - display ichigo-mashimaro serifs

Usage:
	ichigo-cli [options] [name]

Example:
	ichigo-cli nobue

Options:
	-h --help
	-v --version

Names:
	nobue: 伸恵
	chika: 千佳
	miu: 美羽`

	app := &cli.App{
		Name:  "ichigo-cli",
		Usage: "display ichigo-mashimaro serifs",
		Action: func(c *cli.Context) error {
			return run(c)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	name := c.Args().Get(0)
	if name == "" {
		return errors.New("name is required")
	}
	if serifs[name] == nil {
		return fmt.Errorf("%s is invalid", name)
	}
	size := len(serifs[name])

	rand.Seed(time.Now().UnixNano())
	fmt.Println(serifs[name][rand.Intn(size)])
	return nil
}
