package main

import (
	csv "github.com/billyct/split-csv"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(basePath())
	checkError(err)

	pathSetting := filepath.Join(basePath(), "setting.csv")

	var pathRecord string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".csv") && file.Name() != "setting.csv" {
			pathRecord = filepath.Join(basePath(), file.Name())
		}
	}

	err = csv.Split(pathRecord, pathSetting)
	checkError(err)
}

// https://blog.csdn.net/skh2015java/article/details/78515002
func basePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	checkError(err)
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

// simple panic error
func checkError(err error) {
	if err != nil {
		red := color.New(color.FgHiWhite).Add(color.Bold).Add(color.BgRed)
		red.Println(err.Error())
		os.Exit(1)
	}
}
