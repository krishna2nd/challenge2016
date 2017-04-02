package regions

import (
	"bufio"
	"log"
	"os"
	"strings"
	//"fmt"
)

type Info struct {
	Code,
	Name string
}
type Country struct {
	Info
}
type City struct {
	Info
}

type Province struct {
	Info
}
type DistributionRegion struct {
	Country  *Country
	City     *City
	Province *Province
}

func AddRecord(tree *Node, data []string) *Node {
	//fmt.Println(tree, data)
	return tree.Add(&DistributionRegion{
		City: &City{
			Info{
				Code: data[0],
				Name: data[3],
			},
		},
		Province: &Province{
			Info{
				Code: data[1],
				Name: data[4],
			},
		},
		Country: &Country{
			Info{
				Code: data[2],
				Name: data[5],
			},
		},
	})
	return tree
}

func GenerateMapFromCSV(csv string) *Node {
	file, err := os.Open(csv)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var parts []string
	var tree = NewTree()
	for scanner.Scan() {
		parts = strings.Split(scanner.Text(), ",")
		AddRecord(tree, parts)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return tree
}

func Root() *Node {
	return parent
}

func Find(names []string) *Node {
	var (
		node *Node = parent
		key  string
	)
	for index := len(names) - 1; index >= 0; index-- {
		key = strings.TrimSpace(names[index])
		if nil != node {
			node = node.Get(key)
		}
	}
	return node
}
