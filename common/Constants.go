package common

const (
	//下划线
	ULINE = "_"
	//未确认
	UNCOMFIRM = "UNCONFIRM"
	//产品注册所属
	SYSTEM = "SYSTEM"
	//产品所属KEY
	PRODUCT_OWNER = "PRODUCT_OWNER"
	//交易类型（产品注册）
	REGISTER = "01"
	//交易类型(权属变更)
	CHANGE_OWNER = "02"
	//交易类型(确认权属变更)
	CONFIRM_CHANGE_OWNER = "03"
	//交易类型(产品信息变更)
	CHANGE_PRODUCT = "04"
	//交易类型(产品销毁)
	DESTROY = "99"
	//产品信息KEY
	PRODUCT_INFO = "PRODUCT_INFO"
	//产品ID
	PRODUCT_ID = "productId"

	// error 类型和描述
	ERROR_CODE_EXIST    = "001"
	ERROR_MSG_EXIST     = "Asset has registered,please check the asset No."
	ERROR_CODE_Mashal   = "002"
	ERROR_MSG_Mashal    = "Asset information mashal error,please check post data information."
	ERROR_CODE_REGISTER = "003"
	ERROR_MSG_REGISTER  = "Asset register error,please try again."

	ERROR_CODE_OWNER = "004"
	ERROR_MSG_OWNER  = "Save asset owner info error."

	ERROR_CODE_TX = "005"
	ERROR_MSG_TX  = "Update transcation information errors"
)
