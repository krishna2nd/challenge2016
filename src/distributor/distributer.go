package distributor

import (
	"regions"
	"strings"
	//"fmt"
)

type Permissions struct{
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

type Distributer struct {
	name string
	permissions Permissions
	parent, sub *Distributer
}

func (d *Distributer) include(region *regions.Node)  {
	if ( nil != d.permissions.include) {
		d.permissions.include = append(d.permissions.include, region)
	} else {
		d.permissions.include = []*regions.Node {region}
	}
}

func (d *Distributer) exclude(region *regions.Node)  {
	if ( nil != d.permissions.exclude) {
		d.permissions.exclude = append(d.permissions.exclude, region)
	} else {
		d.permissions.exclude = []*regions.Node {region}
	}
}

func (d *Distributer) permission(permissions []string) *Distributer {
	var (
		parts []string
		region *regions.Node
	)
	for _, permission := range permissions {
		parts = strings.SplitN(permission, ":", 2)
		if (len(parts) == 2) {
			region = regions.Find(strings.Split(parts[1], "-"))
		}
		if nil == region {
			region = regions.Root()
		}
		switch parts[0] {
		case "INCLUDE": d.include(region); break;
		case "EXCLUDE": d.exclude(region); break;
		}
	}

	return d;
}

func (d *Distributer) GetPermissions() Permissions {
	return d.permissions;
}

func New(name string, permission []string) *Distributer {
	var d = &Distributer{
		name:name,
	}
	d.permission(permission)
	return d;
}


