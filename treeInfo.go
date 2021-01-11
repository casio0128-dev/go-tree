package main

import "strings"

type TreeInfo struct {
	fileInfo fileInfo
	tabInfo tabInfo
}

type fileInfo map[string]*[]string
type tabInfo map[string]int

/*
	treeInfoの初期化
*/
func(tInfo *TreeInfo) initTreeInfo() {
	tInfo.fileInfo = make(map[string]*[]string)
	tInfo.tabInfo = make(map[string]int)
}

/*
	指定されたキーが存在するかチェック
*/
func (tInfo TreeInfo) isExists(key string) bool {
	if _, ok := tInfo.fileInfo[key]; ok {
		return true
	}
	return false
}

/*
	file情報を取得
*/
func (tInfo TreeInfo) getFiles(key string) []string {
	if tInfo.isExists(key) {
		return *tInfo.fileInfo[key]
	}
	return nil
}

/*
	タブの個数を取得
*/
func (tInfo TreeInfo) tabNumber(key string) int {
	if tInfo.isExists(key) {
		return tInfo.tabInfo[key]
	}
	return 0
}

/*
	指定された複数のパスから、map[string]*[]stringのマップを作成する
*/
func (tInfo TreeInfo) makeMap(paths []string) {
	for _, path := range paths {
		var values []string
		files := getFileInfos(path)
		for _, file := range files {
			values = append(values, file.Name())
		}
		tInfo.fileInfo[path] = &values
	}
}

/*
	baseパスのparentを取得する
*/
func (tInfo TreeInfo) getParentPath(key string) string {
	if tInfo.isExists(key) {
		separatedKeys := strings.Split(key, separator())
		lastElement := separatedKeys[0:len(separatedKeys)-1]
		return strings.Join(lastElement, separator())
	}
	return ""
}