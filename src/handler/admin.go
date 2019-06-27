package handler
type Admin struct {

}

func ( a *Admin) GetData() *Result  {
	var data= map[string]interface{}{
		"GLYBH":"管理员编号11",
		"GLYXM":"管理员姓名11",
		"SFZHM":"身份证号码",
		"HSZX":"核算中心",
		"SJHM":"管理员手机号码",
		"DZYJ":"电子邮件",
		"OAID":"OAID",
	}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewAdmin()*Admin  {
	return &Admin{}
}