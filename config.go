package goconfig

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

//const middle = "========="
const SEP = "=" // key 和 value 分隔符
const NOTE = "#"        // #开头的为注释
const MODEL_START = "[" // [开头的为注释
const MODEL_END = "]"   // [开头的为注释
// 读取配置文件

type node struct {
	key string
	value []byte
	note [][]byte
}

type groupLine struct {
	group []*node  // 组的行
	note [][]byte  // 组注释
	name []byte   // 组名
}

type config struct {
	Groups []*groupLine  // 组
	Lines []*node    // 单key
	Read []byte  // 文件读出来的所有内容
	Write []byte  // 文件写的所有内容
	KeyValue map[string][]byte   // 键值缓存， key的值  key or group.key
	Filepath string  // 配置文件路径
}


var fl *config

func InitConf(configpath string) {

	fptmp := filepath.Clean(configpath)
	//判断文件目录是否存在
	_, err := os.Stat(filepath.Dir(fptmp))
	if err != nil {
		// 不存在就先创建目录
		if err := os.MkdirAll(filepath.Dir(fptmp), 0755); err != nil {
			panic(err)
		}
	}
	fl = &config{
		Filepath: configpath,
		Lines: make([]*node, 0),
		KeyValue: make(map[string][]byte),
	}
	fl.Read, err = ioutil.ReadFile(fptmp)
	if err != nil {
		panic(err)
	}

	fl.readlines()
}


// 读取配置文件到全局变量，并检查重复项, 重载配置文件执行这个函数
