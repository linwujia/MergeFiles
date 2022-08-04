package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

func main() {
	dir := "E:/log/500"
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	fileInfos = filter(fileInfos, func(info os.FileInfo) bool {
		return info.IsDir() || !strings.HasSuffix(info.Name(), ".loge")
	})

	sort.Slice(fileInfos, func(i, j int) bool {
		fi := fileInfos[i]
		fj := fileInfos[j]
		return strings.Compare(fi.Name(), fj.Name()) < 0
	})

	for _, info := range fileInfos {
		if !info.IsDir() {
			log.Printf("file name:%s", info.Name())
		}
	}

	var waitGroup = &sync.WaitGroup{}
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		mergeFile(fileInfos, dir, "E:/log/500/all_log.log")
	}()
	waitGroup.Wait()
}

func filter(slice []os.FileInfo, filter func(info os.FileInfo) bool) []os.FileInfo {
	n := 0
	for _, inf := range slice {
		if !filter(inf) {
			slice[n] = inf
			n++
		}
	}

	return slice[:n]
}

func mergeFile(slice []os.FileInfo, dir, fileName string)  {
	destFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer destFile.Close()

	for _, info := range slice {
		file, err1 := os.Open(fmt.Sprintf("%s/%s", dir, info.Name()))
		if err1 != nil {
			log.Printf("open file err %v", err1)
			continue
		}

		written, err2 := io.Copy(destFile, file)
		if err2 != nil {
			log.Printf("copy file err %v", err2)
			continue
		}

		log.Printf("copy file len %d", written)
		file.Close()
	}
}
