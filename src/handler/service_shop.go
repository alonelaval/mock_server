package handler
type ServiceShop struct {

}

func ( a *ServiceShop) GetData() *Result  {
	var data= map[string]interface{}{
		"FWDBH":"服务店编号",
		"FWDMC":"服务店名称",
		"FWDBM":"服务店别名",
		"FWDJB":"服务店级别",
		"SFDM":"省份代码",
		"SFMC":"省_直辖市_自治区",
		"SDQDM":"地区_市代码",
		"SDQMC":"地区_市",
		"QXDM":"区_县代码",
		"QXMC":"区_县",
		"ZJDDM":"镇_街道代码",
		"ZJDMC":"镇_街道",
		"XXDZ":"详细地址",
		"JD":"经度",
		"WD":"纬度",
		"DYWD":"地域维度",
		"CLDH":"彩铃电话",
		"CZHM":"传真号码",
		"YZBM":"邮政编码",
		"SHR":"收货人",
		"SHDH":"收货电话",
		"SHDZ":"收货地址",
		"FRXM":"法人姓名",
		"FRDH":"法人电话",
		"FRYJ":"法人邮件",
		"FZRXM":"负责人姓名",
		"FZRDH":"负责人电话",
		"FZRYJ":"负责人邮件",
		"NSRZZ":"纳税人资质",
		"YYZZDM":"营业执照代码",
		"KHYX":"开户银行",
		"YXZH":"银行帐号",
		"CWXM":"财务姓名",
		"CWZZ":"财务资质",
		"CWDH":"财务电话",
		"CWDZ":"财务地址",
		"QYXZ":"企业性质",
		"CLRQ":"2018-12-12 12:12:12",
		"ZCZJ":"注册资金",
		"YGRS":"1212",
		"YYKSSJ":"2018-12-12 12:12:12",
		"YYJSSJ":"2018-12-12 12:12:12",
		"LJXSQD":"零件销售清单",
		"LJFPLX":"零件发票类型",
		"LJFPJSDZ":"零件发票寄送地址",
		"ZDFWXT":"自店服务系统",
		"ZXRQ":"2018-12-12 12:12:12",
		"MTCL":"门头材料",
		"MTDL":"门头独立",
		"JDQDL":"接待区独立",
		"JYMJ":"经营面积",
		"LJKFMJ":"零件库房面积",
		"ZYQYPP":"主要签约品牌",
		"SHG":"生活馆",
		"BZ":"备注",
		"WLDW":"物流单位",
		"JYFW":"经营范围",
		"KHDJ":"考核等级",
		"FCXQ":"返厂限期",
		"XCDZTSL":"现场打折贴税率",
		"SGQZXS":"扫关权重系数",
		"JCFWQZXS":"基础服务权重系数",
		"ZZFWQZXS":"增值服务权重系数",
	}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewServiceShop()*ServiceShop  {
	return &ServiceShop{}
}