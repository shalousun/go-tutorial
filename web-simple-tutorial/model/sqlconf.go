package model

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type MysqlConf struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"password"`
	DbName string `yaml:"database"`
	Port   int    `yaml:"port"`
}

func (m *MysqlConf) GetConf() *MysqlConf {
	basePath, err := os.Getwd()
	if err != nil {
		fmt.Println("base path error")
	}
	fileName := filepath.Join(basePath, "conf", "conf.yaml")
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("load conf error")
	}
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		fmt.Println(err.Error())
	}
	return m
}
