package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv(envPath string) {
	var err error
	if envPath == "" {
		envPath = ".env"
	}
	err = godotenv.Load(envPath)
	if err != nil {
		fmt.Println("load env file failed", envPath, err)
	}
}

// get env
func Get(k string) string {
	v, ok := os.LookupEnv(k)
	if ok == false {
		fmt.Println("******* ENV ERROR ****** CANNOT FIND ENV KEY :" + k)
		return ""
	}
	return v
}

func InterfaceToMap(i interface{}) map[string]interface{} {
	b, _ := json.Marshal(i)
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m
}
