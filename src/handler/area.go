package handler
type Area struct {

}

func ( a *Area) GetData() *Result  {
	var data= map[string]interface{}{
		"QYCJ":"区域层级",
		"QYDM":"区域代码",
		"QYMC":"区域名称",
		"BZJL":"标准距离",
	}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewArea()*Area  {
	return &Area{}
}