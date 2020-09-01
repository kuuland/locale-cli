package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

var (
	extMap        = make(map[string]bool)
	reg           = regexp.MustCompile(`L\(['"]([^'"]*)['"],\s*['"]([^'"]*)['"]\)`)
	totalFileKeys = make(map[string]map[string]string)
)

func env(keys ...string) string {
	for _, key := range keys {
		if v := os.Getenv(key); v != "" {
			return v
		}
	}
	return ""
}

func init() {
	exts := env("EXTS", "PLUGIN_EXTS")
	if exts == "" {
		exts = ".go,.js,.jsx,.ts,.tsx,.vue"
	}
	split := strings.Split(exts, ",")
	for _, n := range split {
		extMap[n] = true
	}
}

func dirs() []string {
	var paths []string
	if s := env("DIRS", "PLUGIN_DIRS"); s != "" {
		paths = strings.Split(s, ",")
	} else if len(os.Args) > 1 {
		paths = os.Args[1:]
	}
	return paths
}

func getFiles(pathname string, s []string) ([]string, error) {
	fromSlash := filepath.FromSlash(pathname)
	rd, err := ioutil.ReadDir(fromSlash)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := filepath.Join(fromSlash, fi.Name())
			s, err = getFiles(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := filepath.Join(fromSlash, fi.Name())
			fileExt := path.Ext(fullName)
			if extMap[fileExt] {

				buf, err := ioutil.ReadFile(fullName)
				if err != nil {
					panic(err)
				}
				content := string(buf)
				result := reg.FindAllStringSubmatch(content, -1)
				if len(result) > 0 {
					keys := totalFileKeys[fullName]
					if keys == nil {
						keys = make(map[string]string)
					}

					for _, sub := range result {
						if len(sub) < 3 {
							continue
						}
						key := sub[1]
						defaultValue := sub[2]
						keys[key] = defaultValue
					}

					totalFileKeys[fullName] = keys
				}
				s = append(s, fullName)
			}
		}
	}
	return s, nil
}

func main() {
	var (
		totalPaths []string
		startTime  = time.Now()
	)
	paths := dirs()
	if len(paths) == 0 {
		fmt.Println("请指定扫描目录，支持以下两种用法：")
		fmt.Println("\t1. kuu-locale dir1 dir2 ...")
		fmt.Println("\t2. DIRS=dir1,dir2 kuu-locale")
		return
	}

	for _, p := range paths {
		totalPaths, _ = getFiles(p, totalPaths)
	}
	var (
		logBuf bytes.Buffer
		csvBuf bytes.Buffer

		entries   []string
		existsKey = make(map[string]string)
	)

	for i, p := range totalPaths {
		var (
			keys []string
			vs   = totalFileKeys[p]
		)
		for k, v := range vs {
			if _, has := existsKey[k]; has {
				continue
			}
			existsKey[k] = v
			keys = append(keys, k)
		}
		logBuf.WriteString(fmt.Sprintf("%d.%s(%d):\n", i+1, p, len(keys)))
		entries = append(entries, keys...)
		sort.Strings(keys)
		if len(keys) > 0 {
			for _, k := range keys {
				v := vs[k]
				logBuf.WriteString(fmt.Sprintf("%s=%s\n", k, v))
			}
		} else {
			logBuf.WriteString("No key found.\n")
		}
		logBuf.WriteString("\n")
	}

	sort.Strings(entries)
	csvBuf.WriteString("国际化键,用途描述,English,简体中文,繁體中文\n")
	for _, entry := range entries {
		value := existsKey[entry]
		csvBuf.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s\n", entry, "", value, value, value))
	}
	totalCost := fmt.Sprintf("总耗时：%v\n", time.Since(startTime))
	totalFilesCount := fmt.Sprintf("扫描文件数：%d\n", len(totalFileKeys))
	totalStmtsCount := fmt.Sprintf("命中声明数：%d\n", len(existsKey))

	logBuf.WriteString("\n")
	logBuf.WriteString(totalCost)
	logBuf.WriteString(totalFilesCount)
	logBuf.WriteString(totalStmtsCount)

	fmt.Print(totalCost)
	fmt.Print(totalFilesCount)
	fmt.Print(totalStmtsCount)

	logFileName := "locale.log"
	csvFileName := "locale.csv"

	fmt.Printf("详细日志文件：%s\n", logFileName)
	fmt.Printf("翻译参考文件：%s\n", csvFileName)
	if err := ioutil.WriteFile(logFileName, []byte(logBuf.String()), os.ModePerm); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(csvFileName, []byte(csvBuf.String()), os.ModePerm); err != nil {
		panic(err)
	}
}
