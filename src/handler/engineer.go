package handler
type Engineer struct {

}

func ( a *Engineer) GetData() *Result  {
	var data= map[string]interface{}{"ZZZT":"在职状态",
		"FWDBH":"服务店编号",
		"FWDMC":"服务店名称",
		"GCSBH":"工程师编号11",
		"GCSXM":"工程师姓名11",
		"XJ":"星级",
		"GCSDJ":"工程师等级",
		"SFZH":"身份证号",
		"XB":"性别",
		"XL":"学历",
		"CSRQ":"2018-12-12 12:12:12",
		"RSRQ":"2018-12-12 12:12:12",
		"SJHM":"工程师手机号码",
		"DZYJ":"电子邮件",
		"GZZZT":"工作证状态",
		"GZZFFSJ":"2018-12-12 12:12:12",
		"GZZDQSJ":"2018-12-12 12:12:12",
		"FJSM":"附加说明",
		"HSZXDM":"核算中心代码",
		"HSZX":"核算中心",}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewEngineer()*Engineer  {
	return &Engineer{}
}