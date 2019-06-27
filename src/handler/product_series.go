package handler
type ProductSeries struct {

}

func ( a *ProductSeries) GetData() *Result  {
	var data= map[string]interface{}{
		"CPLXDM":"产品类型代码",
		"CPLX":"产品类型",
		"CPXLDM":"产品系列代码",
		"CPXL":"产品系列",
		"ESSNWCJLF":"24H内完成奖励费",
		"SSBNWCJLF":"48H内完成奖励费",
		"CLBZ":"测量标志",
		"CD":"长度",
		"GD":"高度",
		"SD":"深度",
	}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewProductSeries()*ProductSeries  {
	return &ProductSeries{}
}