package service

import (
	"encoding/json"
	"hlccc/common"
	"hlccc/log"
	"hlccc/module"
	"time"

	"reflect"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

/** 产品注册 **/
func Register(stub shim.ChaincodeStubInterface, param module.RegisterParam) pb.Response {
	//校验产品是否已注册
	jsonByte, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	if err != nil {
		return shim.Error("get product txList error" + err.Error())
	}
	if jsonByte != nil {
		return shim.Error("product has been register" + err.Error())
	}

	var txId = stub.GetTxID()
	product := module.ProductInfo{}
	product.ProductId = param.ProductId
	product.InModule = param.InModule
	product.Operation = param.Operation
	product.Operator = param.Operator
	product.Kind = param.Kind
	product.Type = param.Type
	product.MapPosition = param.MapPosition
	product.CreateTime = param.CreateTime

	jsonByte, err = json.Marshal(product)
	if err != nil {
		return shim.Error("Mashal productInfo error" + err.Error())
	}

	//保存交易详细信息
	err = stub.PutState(txId, jsonByte)
	if err != nil {
		return shim.Error("Put productInfo error" + err.Error())
	}

	//保存产品详细信息
	err = stub.PutState(common.PRODUCT_INFO+common.ULINE+param.ProductId, jsonByte)
	if err != nil {
		return shim.Error("Put productInfo error" + err.Error())
	}

	//保存产品所属信息
	var productOwner = module.ProductOwner{}
	productOwner.PreOwner = common.SYSTEM
	productOwner.CurrentOwner = getMspid(stub)
	jsonByte, err = json.Marshal(productOwner)
	if err != nil {
		return shim.Error("Mashal productOwner error" + err.Error())
	}
	err = stub.PutState(common.PRODUCT_OWNER+common.ULINE+param.ProductId, jsonByte)
	if err != nil {
		return shim.Error("Put productOwner error" + err.Error())
	}

	//更新产品的交易基本信息列表
	var txInfoAdd = module.TxInfoAdd{}
	txInfoAdd.MapPosition = product.MapPosition
	txInfoAdd.Operation = product.Operation
	txInfoAdd.Operator = product.Operator

	txInfoAdd.OperateTime = time.Now().Format("2006-01-02T15:04:05Z07:00")

	err = putTxId(stub, param.ProductId, productOwner, common.REGISTER, txInfoAdd)
	if err != nil {
		return shim.Error("Put TxList error" + err.Error())
	}
	return shim.Success(nil)
}

/** 批量产品注册 **/
func BatchRegister(stub shim.ChaincodeStubInterface, param []module.RegisterParam) pb.Response {
	//校验产品是否已注册
	errlist := make([]module.RegisterErr, 0)

	for index := 0; index < len(param); index++ {
		product := param[index]
		// 保存资产
		searchJsonByte, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + product.ProductId)
		if searchJsonByte != nil {
			errinfo := module.RegisterErr{}
			errinfo.ProductId = product.ProductId
			errinfo.ErrorCode = common.ERROR_CODE_EXIST
			errinfo.ErrorInfo = common.ERROR_MSG_EXIST
			errlist = append(errlist, errinfo)
			continue
		}
		jsonByte, err := json.Marshal(product)
		if err != nil {
			errinfo := module.RegisterErr{}
			errinfo.ProductId = product.ProductId
			errinfo.ErrorCode = common.ERROR_CODE_Mashal
			errinfo.ErrorInfo = common.ERROR_MSG_Mashal
			errlist = append(errlist, errinfo)

		} else {
			//保存产品详细信息
			err = stub.PutState(common.PRODUCT_INFO+common.ULINE+product.ProductId, jsonByte)
			if err != nil {
				errinfo := module.RegisterErr{}
				errinfo.ProductId = product.ProductId
				errinfo.ErrorCode = common.ERROR_CODE_REGISTER
				errinfo.ErrorInfo = common.ERROR_MSG_REGISTER
				errlist = append(errlist, errinfo)

			}
		}

		//保存产品所属信息
		var productOwner = module.ProductOwner{}
		productOwner.PreOwner = common.SYSTEM
		productOwner.CurrentOwner = getMspid(stub)
		jsonByte, err = json.Marshal(productOwner)
		if err != nil {
			errinfo := module.RegisterErr{}
			errinfo.ProductId = product.ProductId
			errinfo.ErrorCode = common.ERROR_CODE_Mashal
			errinfo.ErrorInfo = common.ERROR_MSG_Mashal
			errlist = append(errlist, errinfo)

		}
		err = stub.PutState(common.PRODUCT_OWNER+common.ULINE+product.ProductId, jsonByte)
		if err != nil {
			errinfo := module.RegisterErr{}
			errinfo.ProductId = product.ProductId
			errinfo.ErrorCode = common.ERROR_CODE_OWNER
			errinfo.ErrorInfo = common.ERROR_MSG_OWNER
			errlist = append(errlist, errinfo)

		}

		//更新产品的交易基本信息列表
		var txInfoAdd = module.TxInfoAdd{}
		txInfoAdd.MapPosition = product.MapPosition
		txInfoAdd.Operation = product.Operation
		txInfoAdd.Operator = product.Operator

		txInfoAdd.OperateTime = time.Now().Format("2006-01-02T15:04:05Z07:00")
		log.Logger.Info("#################批量入拦操作###############")
		log.Logger.Info(product)
		log.Logger.Info(product.ProductId)
		err = putTxId(stub, product.ProductId, productOwner, common.REGISTER, txInfoAdd)
		if err != nil {
			errinfo := module.RegisterErr{}
			errinfo.ProductId = product.ProductId
			errinfo.ErrorCode = common.ERROR_CODE_TX
			errinfo.ErrorInfo = common.ERROR_MSG_TX
			errlist = append(errlist, errinfo)
		}
	}

	var txId = stub.GetTxID()
	jsonByteParam, err := json.Marshal(param)
	if err != nil {
		errinfo := module.RegisterErr{}
		errinfo.ProductId = ""
		errinfo.ErrorCode = common.ERROR_CODE_TX
		errinfo.ErrorInfo = common.ERROR_MSG_TX
		errlist = append(errlist, errinfo)

	}
	//保存交易详细信息
	err = stub.PutState(txId, jsonByteParam)
	if err != nil {
		errinfo := module.RegisterErr{}
		errinfo.ProductId = ""
		errinfo.ErrorCode = common.ERROR_CODE_TX
		errinfo.ErrorInfo = common.ERROR_MSG_TX
		errlist = append(errlist, errinfo)
	}
	// if len(errlist) == 0 {
	// 	return shim.Success(nil)
	// } else {
	// 	errjsonbyte, _ := json.Marshal(errlist)
	// 	return shim.Error(errjsonbyte)
	// }
	if len(errlist) == 0 {
		return shim.Success(nil)
	} else {
		errjsonbyte, _ := json.Marshal(errlist)
		// return shim.Error(string(errjsonbyte))
		return shim.Success(errjsonbyte)
	}
}

/** 查询产品详细信息 **/
func QueryProductDetail(stub shim.ChaincodeStubInterface, param module.QueryProductDetailParam) pb.Response {
	jsonByte, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	if err != nil {
		return shim.Error("Get Product Detail error" + err.Error())
	}

	return shim.Success(jsonByte)
}

/** 查询产品流转信息**/
func QueryProductChange(stub shim.ChaincodeStubInterface, param module.QueryParam) pb.Response {
	//查询产品现有交易列表
	jsonByte, err := stub.GetState(param.ProductId)
	if err != nil {
		return shim.Error("get product txList error" + err.Error())
	}
	return shim.Success(jsonByte)
}

/** 查询交易信息**/
func QueryTx(stub shim.ChaincodeStubInterface, param module.QueryTxParam) pb.Response {
	//查询交易详情
	jsonByte, err := stub.GetState(param.TxId)
	if err != nil {
		return shim.Error("get tx info error" + err.Error())
	}
	return shim.Success(jsonByte)
}

/** 权属变更**/
func ChangeOwner(stub shim.ChaincodeStubInterface, param module.ChangeOrgParam) pb.Response {
	//查询产品当前所属
	productOwner, err := queryProductOwner(stub, param.ProductId)
	if err != nil {
		return shim.Error("get productOwner error" + err.Error())
	}
	//验证交易发起方是否有权限
	if getMspid(stub) != productOwner.CurrentOwner {
		return shim.Error("tx sender has no auth to change owner")
	}
	//更改产品权属信息&记录交易详情
	var changeOwner = module.ChangeOwner{}
	changeOwner.Before.PreOwner = productOwner.PreOwner
	changeOwner.Before.CurrentOwner = productOwner.CurrentOwner
	changeOwner.After.PreOwner = productOwner.CurrentOwner
	changeOwner.After.CurrentOwner = common.UNCOMFIRM + common.ULINE + strings.Replace(param.ToOrgMsgId, " ", "", -1)
	err = changeProductOwner(stub, changeOwner.Before, changeOwner.After, param.ProductId)
	if err != nil {
		return shim.Error("change product owner error" + err.Error())
	}
	//更新产品交易列表信息
	var txInfoAdd = module.TxInfoAdd{}
	txInfoAdd.MapPosition = productOwner.CurrentOwner
	txInfoAdd.Operation = "ChangeOwner"
	txInfoAdd.Operator = getMspid(stub)

	txInfoAdd.OperateTime = time.Now().Format("2006-01-02T15:04:05Z07:00")

	err = putTxId(stub, param.ProductId, changeOwner.After, common.CHANGE_OWNER, txInfoAdd)
	if err != nil {
		return shim.Error("Put TxList error" + err.Error())
	}
	return shim.Success(nil)
}

/** 确认权属变更**/
func ConfirmChangeOwner(stub shim.ChaincodeStubInterface, param module.ComfirmChangeParam) pb.Response {
	//查询产品当前所属
	productOwner, err := queryProductOwner(stub, param.ProductId)
	if err != nil {
		return shim.Error("get productOwner error" + err.Error())
	}
	//验证交易发起方是否有权限
	currentOwner := productOwner.CurrentOwner
	if !strings.Contains(currentOwner, common.UNCOMFIRM) {
		return shim.Error("change tx has been confirmed")
	}
	if getMspid(stub) != currentOwner[10:] {
		return shim.Error("tx sender has no auth to confirm change owner")
	}
	//更改产品权属信息&记录交易详情
	var changeOwner = module.ChangeOwner{}
	changeOwner.Before.PreOwner = productOwner.PreOwner
	changeOwner.Before.CurrentOwner = productOwner.CurrentOwner
	changeOwner.After.PreOwner = productOwner.CurrentOwner
	changeOwner.After.CurrentOwner = currentOwner[10:]
	err = changeProductOwner(stub, changeOwner.Before, changeOwner.After, param.ProductId)
	if err != nil {
		return shim.Error("change product owner error" + err.Error())
	}
	//更新产品交易列表信息
	var txInfoAdd = module.TxInfoAdd{}
	txInfoAdd.MapPosition = changeOwner.After.CurrentOwner
	txInfoAdd.Operation = "ConfirmChange"
	txInfoAdd.Operator = getMspid(stub)

	txInfoAdd.OperateTime = time.Now().Format("2006-01-02T15:04:05Z07:00")

	err = putTxId(stub, param.ProductId, changeOwner.After, common.CONFIRM_CHANGE_OWNER, txInfoAdd)
	if err != nil {
		return shim.Error("Put TxList error" + err.Error())
	}

	/** 更改前后所属 **/
	product := module.ProductInfo{}

	jsonByte, err := stub.GetState(common.PRODUCT_INFO + common.ULINE + param.ProductId)
	if err != nil {
		return shim.Error("Put productInfo error" + err.Error())
	}

	err = json.Unmarshal(jsonByte, &product)
	if err != nil {
		return shim.Error("Unmarshal JSON error" + err.Error())
	}

	product.PreOwner = product.CurrentOwner
	product.CurrentOwner = getMspid(stub)

	jsonByte, err = json.Marshal(product)
	if err != nil {
		return shim.Error("Unmarshal JSON error" + err.Error())
	}

	err = stub.PutState(common.PRODUCT_INFO+common.ULINE+param.ProductId, jsonByte)
	if err != nil {
		return shim.Error("Marshal JSON error" + err.Error())
	}

	return shim.Success(nil)
}

/** 产品售出销毁**/
func DestroyProduct(stub shim.ChaincodeStubInterface, param module.DestroyParam) pb.Response {
	//查询产品当前所属
	productOwner, err := queryProductOwner(stub, param.ProductId)
	if err != nil {
		return shim.Error("get productOwner error" + err.Error())
	}
	//验证交易发起方是否有权限
	if getMspid(stub) != productOwner.CurrentOwner {
		return shim.Error("tx sender has no auth to confirm change owner")
	}
	//销毁产品&记录交易详情
	var changeOwner = module.ChangeOwner{}
	changeOwner.Before.PreOwner = productOwner.PreOwner
	changeOwner.Before.CurrentOwner = productOwner.CurrentOwner
	changeOwner.After.PreOwner = productOwner.CurrentOwner
	changeOwner.After.CurrentOwner = param.SerialNum
	err = changeProductOwner(stub, changeOwner.Before, changeOwner.After, param.ProductId)
	if err != nil {
		return shim.Error("change product owner error" + err.Error())
	}
	//更新产品交易列表信息
	var txInfoAdd = module.TxInfoAdd{}
	txInfoAdd.MapPosition = productOwner.CurrentOwner
	txInfoAdd.Operation = "Destroy"
	txInfoAdd.Operator = getMspid(stub)

	txInfoAdd.OperateTime = time.Now().Format("2006-01-02T15:04:05Z07:00")

	err = putTxId(stub, param.ProductId, changeOwner.After, common.DESTROY, txInfoAdd)
	if err != nil {
		return shim.Error("Put TxList error" + err.Error())
	}
	return shim.Success(nil)
}

/** 产品属性变更 **/
func ChangeProductInfo(stub shim.ChaincodeStubInterface, param map[string]interface{}) pb.Response {
	//查询产品当前所属
	log.Logger.Info("#################param###############")
	log.Logger.Info(param)
	productId := reflect.ValueOf(param[common.PRODUCT_ID]).String()
	log.Logger.Info(productId)
	productOwner, err := queryProductOwner(stub, productId)
	if err != nil {
		return shim.Error("get productOwner error" + err.Error())
	}
	// 验证交易发起方是否有权限
	if getMspid(stub) != productOwner.CurrentOwner {
		return shim.Error("tx sender has no auth to change productInfo")
	}
	// 更改产品详细信息&记录交易详情
	err = changeProductInfo(stub, param)
	if err != nil {
		return shim.Error("change productInfo error" + err.Error())
	}
	// 更新产品交易列表信息
	var productInfo = module.ProductInfo{}
	common.SetStructByJsonName(&productInfo, param)

	var txInfoAdd = module.TxInfoAdd{}
	txInfoAdd.MapPosition = productInfo.MapPosition
	txInfoAdd.Operation = productInfo.Operation

	txInfoAdd.Operator = getMspid(stub)
	txInfoAdd.OperateTime = time.Now().Format("2006-01-02T15:04:05Z07:00")

	err = putTxId(stub, productId, productOwner, common.CHANGE_PRODUCT, txInfoAdd)
	if err != nil {
		return shim.Error("Put TxList error" + err.Error())
	}
	return shim.Success(nil)
}

/** 按批次查询耳标 **/
func GetProductIdsByInModule(stub shim.ChaincodeStubInterface, param module.GetProductIdsByInModuleParam) pb.Response {
	inModule := param.InModule // 需要查询的名字
	queryString := `{"selector":{"_id": {"$regex": "PRODUCT_INFO"},"data.inModule":"` + inModule + `"}}`

	resultsIterator, err := stub.GetQueryResult(queryString) // 必须是CouchDB才行
	if err != nil {

		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	results := make([]module.ProductInfo, 0)
	for resultsIterator.HasNext() {
		result, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		product := module.ProductInfo{}
		// log.Println(result.Namespace, result.Key, result.Value)
		err = json.Unmarshal(result.Value, &product)
		if err != nil {

		} else {
			results = append(results, product)
		}
	}
	resultsJSON, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsJSON)
}

/** 查询所有批次 **/
func GetAllInModule(stub shim.ChaincodeStubInterface) pb.Response {
	queryString := `{"selector":{"_id": {"$regex": "PRODUCT_INFO"},"data.inModule":{"$gt":null}}}`
	//
	resultsIterator, err := stub.GetQueryResult(queryString) //必须是CouchDB才行
	if err != nil {

		return shim.Error(err.Error())
	}

	results := make([]string, 0)

	for resultsIterator.HasNext() {
		result, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		var product = module.ProductInfo{}

		err = json.Unmarshal(result.Value, &product)
		if err != nil {

		} else {
			results = append(results, product.InModule)
		}
	}
	resultsList := common.RemoveDuplicatesAndEmpty(results)
	// log.Println(resultsList)
	resultsJSON, err := json.Marshal(resultsList)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsJSON)
}

/** 查询历史信息**/
func GetHistoryByProduct(stub shim.ChaincodeStubInterface, param module.QueryProductDetailParam) pb.Response {
	historyProducts, err := stub.GetHistoryForKey(common.PRODUCT_INFO + common.ULINE + param.ProductId)

	if err != nil {
		return shim.Error("get history err:" + err.Error())
	}
	defer historyProducts.Close()
	results := make([]module.ProductInfo, 0)
	for historyProducts.HasNext() {
		result, err := historyProducts.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		product := module.ProductInfo{}

		err = json.Unmarshal(result.Value, &product)
		if err != nil {
			return shim.Error(err.Error())
		} else {
			results = append(results, product)
		}
	}
	resultsJSON, err := json.Marshal(results)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(resultsJSON)
}
