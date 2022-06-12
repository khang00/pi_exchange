package core

import "time"

type PlaceHolderOperation = func() string
type OperationPlaceHolder struct {
	operationTable map[string]PlaceHolderOperation
}

func NewOperationPlaceHolder() OperationPlaceHolder {
	table := make(map[string]PlaceHolderOperation, 0)
	table["TODAY"] = func() string {
		layout := "31 DEC 2020"
		return time.Now().Format(layout)
	}

	return OperationPlaceHolder{
		operationTable: table,
	}
}

func (o OperationPlaceHolder) getValueForPlaceHolder(placeHolderName string) (string, bool) {
	value, ok := o.operationTable[placeHolderName]
	if !ok {
		return "", ok
	}

	return value(), ok
}
