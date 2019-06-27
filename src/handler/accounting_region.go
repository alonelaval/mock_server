package handler
type AccountingRegion struct {

}

func ( a *AccountingRegion) GetData() *Result  {
	var data= map[string]interface{}{
		"HSDQDM":"核算大区代码",
		"HSDQ":"核算大区",
		"HSZXDM":"核算中心代码",
		"HSZX":"核算中心",
	}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewAccountingRegion()*AccountingRegion  {
	return &AccountingRegion{}
}