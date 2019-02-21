package dtree

import "testing"

func TestIndexRight(t *testing.T) {
	t.Log(IndexRight([]byte("www.baidu.com"), '.'))
	t.Log(IndexRight([]byte("com"), '.'))
}
