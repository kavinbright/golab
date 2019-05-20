package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"wfw/third/config"
)

var (
	configFile = flag.String("config", "config.ini", "config file")
)

type jsonConfig struct {
	Host string
	Port string
	Pwd  string
	User userInfo
}

type userInfo struct {
	Name string
	Adr  string
}

func main() {
	fmt.Println("<--------------------------------Beego/config--------------------------------------->")
	// beego config包
	flag.Parse()
	fmt.Println("configFile: ", *configFile)
	//初始化配置文件
	err := config.InitConfig("ini", *configFile)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Config:", config.String("host"))
	fmt.Println("Port:", config.String("port"))
	fmt.Println("Port:", config.String("pwd"))
	fmt.Println("UserName:", config.String("user::name"))

	fmt.Println("<--------------------------------JSON--------------------------------------->")
	// json
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	decoder := json.NewDecoder(file)
	jsonConfigFile := jsonConfig{}
	err = decoder.Decode(&jsonConfigFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("host: ", jsonConfigFile.Host)
	fmt.Println("port: ", jsonConfigFile.Port)
	fmt.Println("pwd: ", jsonConfigFile.Pwd)
	fmt.Println("UserName: ", jsonConfigFile.User.Name)
	fmt.Println("UserAdr: ", jsonConfigFile.User.Adr)
}
