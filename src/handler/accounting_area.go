package handler



type AccountingArea struct {

}

func ( a *AccountingArea) GetData() *Result  {
	var data= map[string]interface{}{
		"hsdqdm":"核算大区代码",
		"hsdq":"核算大区",
		"hszxdm":"核算中心代码",
		"hszx":"核算中心",
		"hspqdm":"核算片区代码",
		"hspq":"核算片区",
		"glybh":"管理员编号",
		"glyxm":"管理员姓名",
	}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewAccountingArea()*AccountingArea  {
	return &AccountingArea{}
}