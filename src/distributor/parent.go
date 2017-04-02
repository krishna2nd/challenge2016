package distributor

import "msg"

// setIndividualPermission to execute the command and provide result
func SetIndividualPermission(disName string, subCmds []string) (string, error) {
	var err error
	dis := Get(disName)
	if nil == dis {
		_, err = New(disName, subCmds, nil)
	} else {
		_, err = dis.UpdatePermissions(subCmds)
	}
	if nil != err {
		return msg.Empty, err
	}
	return msg.Success, nil
}

