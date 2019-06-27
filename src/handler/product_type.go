package handler
type ProductType struct {

}

func ( a *ProductType) GetData() *Result  {
	var data= map[string]interface{}{
		"CPPLDM":"产品品类代码",
		"CPPL":"产品品类",
		"CPLXDM":"产品类型代码",
		"CPLX":"产品类型",
		"JHWYBZ":"机号唯一标志",
		"BUDM":"BU代码",
		"BUMC":"BU名称",
		"GCDM":"工厂代码",
		"GCMC":"工厂名称",
		"ZZGZGLT":"制造规则管理套",
		"ZZGZGLTDM":"制造规则管理套代码",
		"PBDSFLAG":"平板电视FLAG",
		"ZJBZ":"自接标志",
		"SXPBZ":"送修品标志",
		"WXKXBZ":"微信可选标志",
		"TGAZFWBZ":"提供安装服务标志",
		"TSSM":"提示说明",
		"BZ":"备注",
		"HSZXBH":"核算中心编号",
		"HSZX":"核算中心",
	}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewProductType()*ProductType  {
	return &ProductType{}
}