package dtree

// DTree domain tree
type DTree interface {
	// Set set domain into domain tree with data
	Set(domain string, data interface{})
	// Get get data from domain tree
	Get(domain string) (data interface{})
}

type dtree struct {
	root *dNode
}

// Set set data with domain
func (dt *dtree) Set(domain string, data interface{}) {
	dt.root.insert(domain, data)
}

func (dt *dtree) Get(domain string) interface{} {
	return dt.root.get(domain)
}

func NewDTree() DTree {
	return &dtree{root: newDNode("", nil)}
}
