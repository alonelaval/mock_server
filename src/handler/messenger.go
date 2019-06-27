package handler
type Messenger struct {

}

func ( a *Messenger) GetData() *Result  {
	var data= map[string]interface{}{"FWDBH":"服务店编号",
		"XXYBH":"信息员编号11",
		"XXYXM":"信息员姓名11",
		"SFZH":"身份证号",
		"XB":"性别",
		"YDDH":"信息员手机号码",
		"BGDH":"办公电话",
		"DZYJ":"电子邮件",
		"CSRQ":"2018-12-12 12:12:12",
		"RSRQ":"2018-12-12 12:12:12",
		"XL":"学历",
		"ZT":"状态",
		"ZHSX":"帐号属性",
		"FJSM":"附加说明",
		"HSZXDM":"核算中心代码",
		"HSZX":"核算中心",}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewMessenger()*Messenger  {
	return &Messenger{}
}