package module

//交易基本信息
type TxInfo struct {
	TxId         string `json:"txId"`
	TxType       string `json:"txType"`
	PreOwner     string `json:"preOwner"`
	CurrentOwner string `json:"currentOwner"`
	Operation    string `json:"operation"`   // 具体操作行为，对应Comment
	MapPosition  string `json:"mapPosition"` // 发生交易的地理位置
	Operator     string `json:"operator"`    // 对应任何交易的操作员
	TxTime       string `json:"txTime"`
}

//交易附加信息
type TxInfoAdd struct {
	Operation   string `json:"operation"`   // 具体操作行为，对应Comment
	MapPosition string `json:"mapPosition"` // 发生交易的地理位置
	Operator    string `json:"operator"`    // 对应任何交易的操作员
	OperateTime string `json:"operateTime"` // 操作时间
}

//产品详情
type ProductInfo struct {
	// 流转

	MapPosition  string `json:"mapPosition"`  // 地理位置
	Operation    string `json:"operation"`    // 行为
	Operator     string `json:"operator"`     // 操作员
	PreOwner     string `json:"preOwner"`     // 上游
	CurrentOwner string `json:"currentOwner"` // 当前

	// 入栏
	ProductId string `json:"productId"` // 耳标编号
	InModule  string `json:"inModule"`  // 入栏批次
	Kind      string `json:"kind"`      // 种类
	Type      string `json:"type"`      // 品种

	// 所有时间
	CreateTime     string `json:"createTime"`     // 入栏时间
	FeedTime       string `json:"feedTime"`       // 喂养时间
	MedicationTime string `json:"medicationTime"` // 防疫时间
	PreventionTime string `json:"preventionTime"` // 检疫时间
	SaveTime       string `json:"saveTime"`       // 救治时间
	LostTime       string `json:"lostTime"`       // 灭失时间
	FattenedTime   string `json:"fattenedTime"`   // 出栏时间

	// 喂养
	Feeder string `json:"feeder"` // 喂养人
	FeedId string `json:"feedId"` // 饲料ID

	// 防疫
	Medder string `json:"medder"` // 防疫人
	MedId  string `json:"medId"`  // 药品ID

	// 检疫
	Preventer     string `json:"preventer"`     // 检疫人
	PreventResult string `json:"preventResult"` // 检疫结果
	PreventName   string `json:"preventName"`   // 病情

	// 救治
	Inspector     string `json:"inspector"`     // 救治人
	InspectResult string `json:"inspectResult"` // 救治结果
	Treatment     string `json:"treatment"`     // 救治措施
	InspectId     string `json:"inspectId"`     // 药品ID

	// 灭失
	Loser     string `json:"loser"`     // 灭失人
	LostCause string `json:"lostCause"` // 灭失原因
	LostTreat string `json:"lostTreat"` // 处理方式

	// 出栏
	OutModule string `json:"outModule"` // 出栏编号
	Name      string `json:"name"`      // 操作员
}

//变更所属交易详情
type ChangeOwner struct {
	Before ProductOwner `json:"before"`
	After  ProductOwner `json:"after"`
}

//确认变更所属交易详情
type ConfirmChangeOwner struct {
	Before ProductOwner `json:"before"`
	After  ProductOwner `json:"after"`
}

//产品属性变更交易详情
type ChangeProduct struct {
	Before ProductInfo `json:"before"`
	After  ProductInfo `json:"after"`
}

//产品销毁交易详情
type DestoryProduct struct {
	Before ProductOwner `json:"before"`
	After  ProductOwner `json:"after"`
}

//产品当前所属信息
type ProductOwner struct {
	PreOwner     string `json:"preOwner"`
	CurrentOwner string `json:"currentOwner"`
}

// error return
type RegisterErr struct {
	ProductId string `json:"productid"`
	ErrorCode string `json:"errcode"`
	ErrorInfo string `json:"errmsg"`
}
