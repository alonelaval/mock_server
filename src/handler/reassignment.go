package handler


//预约改派
type Reassignment struct {

}

func ( a *Reassignment) GetData() *Result  {
	var data= map[string]string{}
	//rows := []*Row{NewRow(data)}
	rows := []map[string]string{data}
	result := NewResult("111","test",rows)
	return result
}

func NewReassignment()*Reassignment  {
	return &Reassignment{}
}