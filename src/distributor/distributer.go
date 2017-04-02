package distributor

import (
	"regions"
	"strings"
	//"fmt"
)

type Permissions struct {
	include,
	exclude []*regions.Node
}

/*
func (p Permissions) String() string {
	for _, r := range p.include {
		fmt.Print(*r)
	}

	for _, r := range p.exclude {
		fmt.Print(*r)
	}
	return  "test"
}
*/

var distributers = make(map[string]*Distributer, 0)

type Distributer struct {
	name        string
	permissions Permissions
	parent, sub *Distributer
}

func (d *Distributer) include(region *regions.Node) {
	if nil != d.permissions.include {
		d.permissions.include = append(d.permissions.include, region)
	} else {
		d.permissions.include = []*regions.Node{region}
	}
}

func (d *Distributer) exclude(region *regions.Node) {
	if nil != d.permissions.exclude {
		d.permissions.exclude = append(d.permissions.exclude, region)
	} else {
		d.permissions.exclude = []*regions.Node{region}
	}
}

func (d *Distributer) permission(permissions []string) (*Distributer, error) {
	var (
		parts  []string
		region *regions.Node
	)

	for _, permission := range permissions {
		parts = strings.SplitN(permission, ":", 2)
		if len(parts) == 2 {
			region = regions.Find(strings.Split(parts[1], "-"))
		}
		if nil == region {
			region = regions.Root()
		}
		switch parts[0] {
		case "INCLUDE":
			d.include(region)
			break
		case "EXCLUDE":
			d.exclude(region)
			break
		}
	}

	return d, nil
}

func (d *Distributer) UpdatePermissions(permissions []string) (*Distributer, error) {
	d.permissions = Permissions{
		include: []*regions.Node{},
		exclude: []*regions.Node{},
	}
	return d.permission(permissions)
}

func (d *Distributer) GetPermissions() Permissions {
	return d.permissions
}

func New(name string, permission []string, parent *Distributer) (*Distributer, error) {
	var d = &Distributer{
		name: name,
	}
	_, err := d.permission(permission)
	if nil != err {
		return nil, err
	}
	d.parent = parent
	add(d)
	return d, nil
}

func add(newDistributer *Distributer) {
	distributers[newDistributer.name] = newDistributer
}

func GetAll() map[string]*Distributer {
	return distributers
}

func Get(name string) *Distributer {
	return distributers[name]
}
