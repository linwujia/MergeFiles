package main

import (
	"os"
	"strings"
)

type Files []os.FileInfo


// 实现Sort接口中的三个方法
func (f Files) Len() int {
	return len(f)
}

func (f Files) Less(i, j int) bool {
	fi := f[i]
	fj := f[j]
	return strings.Compare(fi.Name(), fj.Name()) > 0
}


func (f Files) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
