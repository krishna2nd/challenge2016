package regions

import (
	"strings"
)

type IAddressNode interface {
	Add(*Node) *Node
	//GetChildren(Data string) *Node
	//GetChildrens() []*Node
	//SetChildrens() []*Node

	isCoutry() bool
	isCity() bool
	isProvince() bool
	findCountry() *Node
	findCity() *Node
	findProvince() *Node
}
type Node struct {
	data      *Info
	childrens map[string]*Node
	parent    *Node
}

var parent *Node

func NewTree() *Node {
	if nil == parent {
		parent = NewNode(nil, nil)
	}
	return parent
}

func NewNode(data *Info, parent *Node) *Node {
	return &Node{
		data:      data,
		parent:    parent,
		childrens: make(map[string]*Node),
	}
}

func (n *Node) get(name string) *Node {
	node, ok := n.childrens[name]
	if ok {
		return node
	}
	return nil
}

func (n *Node) Get(name string) *Node {
	return n.childrens[name]
}

func (n *Node) Add(info *DistributionRegion) *Node {
	var key string
	genKey := func(key string) string {
		key = strings.ToUpper(
			strings.Replace(
				strings.TrimSpace(key),
				" ", "", -1))
		return key
	}
	key = genKey(info.Country.Info.Name)
	countryNode := parent.get(key)
	if nil == countryNode {
		countryNode = NewNode(&info.Country.Info, n)
		parent.childrens[key] = countryNode
	}
	key = genKey(info.Province.Info.Name)
	provinceNode := countryNode.get(key)
	if nil == provinceNode {
		provinceNode = NewNode(&info.Province.Info, countryNode)
		countryNode.childrens[key] = provinceNode
	}
	key = genKey(info.City.Info.Name)
	cityNode := provinceNode.get(key)
	if nil == cityNode {
		cityNode = NewNode(&info.City.Info, provinceNode)
		provinceNode.childrens[key] = cityNode
	}
	return n
}
