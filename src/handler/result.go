package handler

type Result struct {
	Result string
	ErrMsg string
	TableList [] map[string]string
}

//type Row struct {
//	Data  map[string]string
//}



func NewResult(result string,errMsg string,data[] map[string]string)*Result{
 	return &Result{Result:result,ErrMsg:errMsg,TableList:data}
}

//func NewRow(data map[string]string) *Row{
//	return &Row{Data:data}
//}