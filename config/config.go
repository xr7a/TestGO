package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	ConnectionString string `json:"connectionString"`
	Traces           bool   `json:"traces"`
}

func NewConfig() (Config, error) {
	env := GetEnv()

	var result Config
	currentDir, err := os.Getwd()

	fmt.Printf("Текущая рабочая директория: %s\n", currentDir)
	filePath := fmt.Sprintf(".config/appsettings.%s.json", env)
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err = json.Unmarshal(fileBytes, &result); err != nil {
		log.Fatalf("Unmarshalling error: %s", err)
	}

	return result, nil
}
