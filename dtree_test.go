package dtree

import "testing"

func TestDTree(t *testing.T) {
	dt := NewDTree()
	dt.Set("www.baidu.com", "hello")
	t.Log(dt.Get("www.baidu.com"))
}
