package handler

//预约
type Appointment struct {
	
}

func ( a *Appointment) GetData() *Result  {
	var data= map[string]string{}
	//rows := []*Row{NewRow(data)}
	rows := []map[string]string{data}
	result := NewResult("111","test",rows)
	return result
}

func NewAppointment()*Appointment  {
	return &Appointment{}
}