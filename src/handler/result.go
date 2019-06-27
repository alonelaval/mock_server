package handler

type Result struct {
	RESULT string
	ERRMSG string
	TableList [] map[string]interface{}
}

//type Row struct {
//	Data  map[string]string
//}



func NewResult(result string,errMsg string,data[] map[string]interface{})*Result{
 	return &Result{RESULT:result,ERRMSG:errMsg,TableList:data}
}

//func NewRow(data map[string]string) *Row{
//	return &Row{Data:data}
//}



