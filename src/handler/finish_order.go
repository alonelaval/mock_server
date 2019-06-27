package handler


//预约改派
type FinishOrder struct {

}

func ( a *FinishOrder) GetData() *Result  {
	var data= map[string]interface{}{}
	//rows := []*Row{NewRow(data)}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewFinishOrder()*FinishOrder  {
	return &FinishOrder{}
}