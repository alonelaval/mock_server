package handler
type ProductPrice struct {

}

func ( a *ProductPrice) GetData() *Result  {
	var data= map[string]interface{}{
		"cppldm":"产品品类代码",
		"cppl":"产品品类",
		"gmqd":"购买渠道",
		"zffs":"支付方式",
		"cplx":"产品类型",
		"ybcpmc":"延保产品名称",
		"jg":"价格",
		"dxyj":"电销佣金",
		"gcsyj":"工程师佣金",
		"fwdyj":"服务店佣金",
	}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewProductPrice()*ProductPrice  {
	return &ProductPrice{}
}