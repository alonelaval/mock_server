package handler
type PointInspection struct {

}

func (c *PointInspection) GetData()*Result {

	var subData = map[string]interface{}{
		"xmbh":"项目编号",
		"xmmc":"项目名称",
		"xmjz":"项目基准",
		"djff":"点检方法",
	}
	var subList= []map[string]interface{}{subData}

	var data= map[string]interface{}{
		"djfadm":"点检方案代码",
		"djfa":"点检方案",
		"sm":"说明",
		"SubTableList":subList,
	}

	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)

	//jsonBytes, err := json.Marshal(result)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//str := string(jsonBytes)
	//
	return  result
}

func NewPointInspection()*PointInspection  {
	return &PointInspection{}
}


