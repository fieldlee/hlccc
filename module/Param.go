package module

type RegisterParam struct {
	ProductId   string `json:"productId"`
	ProductName string `json:"productName"`
	Operation   string `json:"operation"`
	Operator    string `json:"operator"`
	CreateTime  string `json:"createTime"`
	InModule    string `json:"inModule"`    // 入栏批次
	Kind        string `json:"kind"`        // 种类
	Type        string `json:"type"`        // 品种
	MapPosition string `json:"mapPosition"` // 地理位置
}

type ChangeOrgParam struct {
	ProductId  string `json:"productId"`
	ToOrgMsgId string `json:"toOrgMsgId"`
}

type ComfirmChangeParam struct {
	ProductId string `json:"productId"`
}

type DestroyParam struct {
	ProductId string `json:"productId"`
	SerialNum string `json:"serialNum"`
}

type QueryParam struct {
	ProductId string `json:"productId"`
}

type QueryTxParam struct {
	TxId string `json:"txId"`
}

type QueryProductDetailParam struct {
	ProductId string `json:"productId"`
}

type GetProductIdsByInModuleParam struct {
	InModule string `json:"inModule"`
}
