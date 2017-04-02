package distributor

import "msg"

// SetLevel1Permission to execute the command and provide result
func SetLevelPermission(disLevel1, disLevel2 string, subCmds []string) (string, error) {
	var err error
	// Check parent level1 distributer exists or not
	dis1 := Get(disLevel1)
	if nil == dis1 {
		return msg.Empty, msg.ErrDistributerNotFound(disLevel1)
	}

	// Check parent level1 distributer exists or not
	dis2 := Get(disLevel2)
	if nil == dis2 {
		_, err = New(disLevel2, subCmds, dis1)
	} else {
		_, err = dis2.UpdatePermissions(subCmds)
	}
	if nil != err {
		return msg.Empty, msg.ErrInheritPermissions(disLevel2, disLevel1)
	}

	return msg.Success, nil
}

// SetLevel2Permission to execute the command and provide result
func SetLevel2Permission(disL1, disL2, disL3 string, subCmds []string) (string, error) {
	dis1 := Get(disL1)
	if nil == dis1 {
		return msg.Empty, msg.ErrDistributerNotFound(disL1)
	}

	dis2 := Get(disL2)
	if nil == dis2 {
		return msg.Empty, msg.ErrDistributerNotFound(disL2)
	}

	return SetLevelPermission(disL2, disL3, subCmds)
}
