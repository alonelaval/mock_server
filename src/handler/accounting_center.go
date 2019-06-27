package handler


type AccountingCenter struct {

}

func ( a *AccountingCenter) GetData() *Result  {
	var data= map[string]interface{}{
		"HSZXDM":"核算中心代码",
		"HSZX":"核算中心",
	}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewAccountingCenter()*AccountingCenter  {
	return &AccountingCenter{}
}