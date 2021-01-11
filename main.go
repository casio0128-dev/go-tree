package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
)

// Windowsのファイルパスに使われるセパレータ
const WIN_SEPARATOR = "\\"
// Linuxのファイルパスに使われるセパレータ
const LIN_SEPARATOR = "/"
// Macのファイルパスに使われるセパレータ（LinuxベースのためLinuxと同様）
const MAC_SEPARATOR = LIN_SEPARATOR

const containsHiddenFile = false

func main() {
	// 隠しファイルをTreeに含めるか
	//var containsHiddenFiles = flag.Bool("h", false, "contains hidden files for tree.")
	//flag.Parse()

	var directories = getDirectories(".")
	var tInfo TreeInfo
	tInfo.initTreeInfo()
	tInfo.makeMap(directories)

	fmt.Println(tInfo)

	for _, dir := range directories {
		fmt.Println(dir, countSeparator(dir))

		for _, fileName := range tInfo.getFiles(dir) {
			fmt.Println("\t" + fileName)
		}
	}
}

func countSeparator(path string) int {
	return len(strings.Split(path, separator())) - 1
}

/*
	指定されたディレクトリは以下のファイル情報を取得（ディレクトリ）
*/
func getDirectories(p string) []string {
	var path = p
	var dirs = getFileInfos(path)
	var result []string

	if strings.EqualFold(path, ".") {
		result = append(result, path)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			if !containsHiddenFile && !strings.EqualFold(dir.Name(), ".") && strings.HasPrefix(dir.Name(), ".") {
				continue
			}

			if path[len(path)-1:] != separator() {
				suffixAddSeparator(&path)
			}
			result = append(result, path + dir.Name())
			result = append(result, getDirectories(path + dir.Name())...)
		}
	}

	return result
}

/*
	指定されたディレクトリ配下のファイル情報を取得（ファイル）
*/
func getFiles(p string) []string {
	var path = p
	var dirs = getFileInfos(path)
	var result []string

	for _, dir := range dirs {
		if !dir.IsDir() {
			result = append(result, path + dir.Name())
		}
	}

	return result
}

/*
	指定されたディレクトリ配下のファイル情報を取得（ファイル、ディレクトリ含む）
 */
func getFileInfos(path string) []os.FileInfo {
	if path == "" {
		fmt.Println("cannot input path for blank.")
		return nil
	}

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

/*
	指定されたパスにあるファイルの個数を取得
 */
func countDirectories(path string) uint8 {
	dirs := getFileInfos(path)
	return uint8(len(dirs))
}

/*
	指定されたパス以下のファイル情報を取得
 */
func search(path string) []string {
	var dir = getFileInfos(".")
	var result []string

	for _, file := range dir {
		var filePath = path
		if filePath[len(filePath)-1:] != separator() {
			suffixAddSeparator(&filePath)
		}

		if file.IsDir() {
			result = append(result, search(filePath + file.Name())...)
		} else {
			result = append(result, filePath + file.Name())
		}
	}
	return result
}

/*
	OS別のセパレータを返す関数
 */
func separator() string {
	switch runtime.GOOS {
	case "windows":
		return WIN_SEPARATOR
	case "linux":
		return LIN_SEPARATOR
	case "darwin":
		return MAC_SEPARATOR
	}
	return ""
}

/*
	末尾にファイルのセパレータを付与する関数
	末尾にセパレータが付与済みかどうかは確認しないため、独自に実施する必要有
 */
func suffixAddSeparator(path *string) {
	*path += separator()
}