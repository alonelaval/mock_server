package handler
type ProductCategory struct {

}

func ( a *ProductCategory) GetData() *Result  {
	var data= map[string]interface{}{
		"CPPLDM":"产品品类代码",
		"CPPL":"产品品类",
		"KPFSQ":"空跑费申请",
		"YCF":"远程费",
		"HSZXDM":"核算中心代码",
		"HSZX":"核算中心",
	}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewProductCategory()*ProductCategory  {
	return &ProductCategory{}
}