package dtree

import (
	"sync"
)

// dNode domain node
type dNode struct {
	// levelDomain top-level-domain, first-level-domain, not include '.'
	levelDomain string
	// extra extra info store in this node
	data interface{}
	// pre previous dNode
	pre *dNode
	// lock lock for next map
	lock sync.RWMutex
	// next next dNode info
	next map[string]*dNode
}

// newDNode return a new dNode pointer
func newDNode(ld string, data interface{}) *dNode {
	return &dNode{
		levelDomain: ld,
		data:        data,
		next:        make(map[string]*dNode),
	}
}

// insert insert data to domain tree
func (dn *dNode) insert(remain string, data interface{}) {
	index := IndexRight(String2Bytes(remain), '.')
	if index < 0 {
		if next, ok := dn.next[remain]; ok {
			next.data = data
		} else {
			dn.next[remain] = newDNode(remain, data)
		}
		return
	}
	lvlDomain := remain[index+1:]
	if next, ok := dn.next[lvlDomain]; ok {
		next.insert(remain[:index], data)
	} else {
		dn.next[lvlDomain] = newDNode(lvlDomain, nil)
		dn.next[lvlDomain].insert(remain[:index], data)
	}
}

func (dn *dNode) get(remain string) interface{} {
	index := IndexRight(String2Bytes(remain), '.')
	if index < 0 {
		if next, ok := dn.next[remain]; ok {
			return next.data
		}
		return nil
	}
	lvlDomain := remain[index+1:]
	if next, ok := dn.next[lvlDomain]; ok {
		return next.get(remain[:index])
	}
	return nil
}
