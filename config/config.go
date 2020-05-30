package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	Rows      int     `json:"rows"`
	Columns   int     `json:"columns"`
	Threshold float64 `json:"threshold"`
	Fps       int     `json:"fps"`
}

func ReadFile(cfgFile string) *Config {
	cnf := Config{}

	jsonFile, err := os.Open(cfgFile)
	if err != nil {
		fmt.Println(err)
	}

	fileContents, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	err =json.Unmarshal(fileContents, &cnf)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened ", cfgFile)

	defer jsonFile.Close()

	return &cnf
}