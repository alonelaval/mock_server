package handler


type Order struct {
	
}

func (c *Order) GetData()*Result {

	var data= map[string]string{"PGDH":"1",
		"PGLY":"派工来源",
		"LYID":"来源ID",
		"TMMJID":"天猫买家ID",
		"TMSX":"天猫属性",
		"FWLX":"服务类型",
		"HSZXDM":"核算中心代码",
		"HSZX":"核算中心",
		"HSDQDM":"核算大区代码",
		"HSDQ":"核算大区",
		"HSPQDM":"核算片区代码",
		"HSPQ":"核算片区",
		"FWDBH":"服务店编号",
		"FWDMC":"服务店名称",
		"KHBH":"客户编号",
		"KHXM":"客户姓名",
		"KHXB":"客户性别",
		"SJHM":"手机号码",
		"GDDH":"固定电话",
		"BYDH":"备用电话",
		"KHBZ":"客户备注",
		"SFDM":"省份代码",
		"SFMC":"省_直辖市_自治区",
		"SDQDM":"地区_市代码",
		"SDQMC":"地区_市",
		"QXDM":"区_县代码",
		"QXMC":"区_县",
		"ZJDDM":"镇_街道代码",
		"ZJDMC":"镇_街道",
		"XXDZ":"详细地址",
		"YCBZ":"远程标志",
		"FWSX":"服务时限",
		"CPPLDM":"产品品类代码",
		"CPPL":"产品品类",
		"CPLXDM":"产品类型代码",
		"CPLX":"产品类型",
		"CPXLDM":"产品系列代码",
		"CPXL":"产品系列",
		"CPXH":"产品型号",
		"GCSBH":"工程师编号",
		"GCSXM":"工程师姓名",
		"KHKSS":"J考核开始时间",
		"YYSJ":"预约时间",
	}
	rows := []map[string]string{data}
	result := NewResult("111","test",rows)

	//jsonBytes, err := json.Marshal(result)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//str := string(jsonBytes)
	//
	return  result
}

func NewOrder()*Order  {
	return &Order{}
}

