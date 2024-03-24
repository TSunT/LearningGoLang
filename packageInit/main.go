package main

/**
报错 main.go:4:2: package helloworld/packageInit/lib1 is not in std
配置 go env -w GO111MODULE=off 即可
因为在配置的GOPATH中找相应的非std包
如果配置空字符串只能在GOPATH里面找std包
*/
import (
	"helloworld/packageInit/lib1"
	"helloworld/packageInit/lib2"
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}
