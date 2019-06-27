package handler
type AreaShop struct {

}

func ( a *AreaShop) GetData() *Result  {
	var data= map[string]interface{}{
		"CPPLDM":"产品品类代码",
		"CPPL":"产品品类",
		"QYDM":"区域代码",
		"QYMC":"区域名称",
		"FWDBH":"服务店编号",
		"FWDMC":"服务店名称",
		"YCBZ":"远程标志",
		"FWSX":"服务时限",
		"PDBZ":"派单标志",
		"HSZXDM":"核算中心代码",
		"HSZX":"核算中心",
	}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewAreaShop()*AreaShop  {
	return &AreaShop{}
}