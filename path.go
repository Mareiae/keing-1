package keing

import (
	"fmt"
	"strconv"
)

type Path struct {
	method string
	path   string
}

type Paths struct {
	maxNum int
	paths  map[int]*Path
}

func (p *Paths) AddPath(method, path string, n int) {
	p.paths[n] = &Path{
		method: method,
		path:   path,
	}
	p.maxNum++
}

func (p *Paths) getPathString(n int) string {
	str := fmt.Sprintf("--> func%d (%d handlers)", n, p.maxNum)
	s := fmt.Sprintf("[Keing] %-"+strconv.Itoa(20)+"s%-"+strconv.Itoa(20)+"s%-"+strconv.Itoa(20)+"s", p.paths[n].method, p.paths[n].path, str)
	return s
}
func (p *Paths) ShowAllPathString() {
	for k, _ := range p.paths {
		fmt.Println(p.getPathString(k))
	}
}
