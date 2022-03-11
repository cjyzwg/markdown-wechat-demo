package main

import (
	"fmt"

	"github.com/cjyzwg/markdown-wechat/draft"
	"github.com/fastwego/offiaccount"
	"github.com/spf13/viper"
)

var App *offiaccount.OffiAccount

func init() {
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	config := offiaccount.Config{
		Appid:  viper.GetString("APPID"),
		Secret: viper.GetString("SECRET"),
	}
	App = offiaccount.New(config)
}

func main() {
	config_file := &draft.ConfigFile{
		MarkdownFilePath: "assets/test.md",
		CssFilePath:      "assets/xzdnf-wechat.css",
		ImagePath:        "assets/img.jpg",
		Title:            "test",
		Author:           "c",
		Digest:           "c",
	}

	res, err := draft.DraftRun(config_file, App)
	fmt.Println(res, err)

}
