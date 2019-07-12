package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type BlogConfig struct {
	ServerName        string `toml:"server_name"`
	LogEnable         bool   `toml:"log_enable"`
	LogPath           string `toml:"log_path"`
	ImgPath           string `toml:"img_path"`
	ServerPort        int    `toml:"server_port"`
	EnvProd           bool   `toml:"environment_prod"`
	SecretKey         string `toml:"secret_key"`
	DefaultClientUser string `toml:"default_client_user"`
}

func main() {
	path := "./config.toml"

	config := new(BlogConfig)
	toml.DecodeFile(path, config)
	fmt.Println(config.ServerName)
}
