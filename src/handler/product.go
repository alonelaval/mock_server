package handler
type Product struct {

}

func ( a *Product) GetData() *Result  {
	var data= map[string]interface{}{
		"CPPLDM":"产品品类代码",
		"CPPL":"产品品类",
		"CPLXDM":"产品类型代码",
		"CPLX":"产品类型",
		"CPXLDM":"产品系列代码",
		"CPXL":"产品系列",
		"CPTBSM":"产品特别说明",
		"CPXH":"产品型号",
		"TMDYZ":"条码对应值",
		"GYSDM":"供应商代码",
		"GYSMC":"供应商名称",
		"SSRQ":"2018-12-12 12:12:12",
		"LCRQ":"2018-12-12 12:12:12",
		"ZZRQ":"2018-12-12 12:12:12",
		"ZZSCHL":"终止时出货量",
		"QNTCBS":"七年停产标识",
		"HXLX":"换修类型",
		"YXLX":"允许拉修",
		"LXFY":"拉修费用",
		"BXQNX":"保修期年限",
		"JXPBZ":"寄修品标志",
		"JXZYF":"寄修支援费",
		"JXCZF":"寄修操作费",
		"PBFXXH":"品保分析型号",
		"PZFK":"品质反馈",
		"SZYTBZ":"送装一体标志",
		"CBBZ":"超保标志",
		"D1BCSX":"D1补偿上限",
		"D2BCSX":"D2补偿上限",
		"D3BCSX":"D3补偿上限",
		"RXGZBCSX":"软性故障补偿上限",
		"YXGZBCSX":"硬性故障补偿上限",
		"LKHBZFY":"L库换包装费用",
		"JSLX":"结算类型",
		"ZXAZF":"专修安装费",
		"ZDAZF":"重点安装费",
		"PTAZF":"普通安装费",
		"YBAZF":"一般安装费",
		"LKJSLT":"6口径水龙头",
		"DBAZ":"顶板安装",
		"GLQAZ":"过滤器安装",
		"AZDTJL":"安装单台奖励",
		"WXDTJL":"维修单台奖励",
		"SFGMYB":"是否购买延保",
	}
	rows := []map[string]interface{}{data}
	result := NewResult("111","test",rows)
	return result
}

func NewProduct()*Product  {
	return &Product{}
}