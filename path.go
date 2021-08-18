package keing

import (
	"fmt"
	"strconv"
)

//Path 单个路径结构
type Path struct {
	method string
	path   string
}

//Paths 全部路径结构
type Paths struct {
	maxNum int
	paths  map[int]*Path
}

//AddPath 将路径添加到路径表里
func (p *Paths) AddPath(method, path string, n int) {
	p.paths[n] = &Path{
		method: method,
		path:   path,
	}
	p.maxNum++
}

//获取单个路径
func (p *Paths) getPathString(n int) string {
	str := fmt.Sprintf("--> func%d (%d handlers)", n, p.maxNum)
	s := fmt.Sprintf("[Keing] %-"+strconv.Itoa(20)+"s%-"+strconv.Itoa(20)+"s%-"+strconv.Itoa(20)+"s", p.paths[n].method, p.paths[n].path, str)
	return s
}

//ShowAllPathString 获取当前所里路径并格式化输出
func (p *Paths) ShowAllPathString() {
	for k := range p.paths {
		fmt.Println(p.getPathString(k))
	}
}
