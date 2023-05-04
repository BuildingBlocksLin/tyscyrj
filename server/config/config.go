package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var cfg Config

func init() {
	readFile, err := os.ReadFile("./config/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(readFile, &cfg)
	if err != nil {
		return
	}
	fmt.Println(cfg)
}

func GetItemByName(name string) *Item {
	for _, ws := range cfg.WorkshopS {
		for _, item := range ws.Product {
			if item.Name == name {
				return item
			}
		}
	}
	return nil
}
