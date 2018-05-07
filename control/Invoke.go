package control

import (
	"encoding/json"
	"hlccc/log"
	"hlccc/module"
	"hlccc/service"
	"strconv"

	"hlccc/common"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Invoke is called by fabric to execute a transaction
func (t *ProductTrace) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	log.Logger.Info("###########调用invoke开始############")
	function, args := stub.GetFunctionAndParameters()
	log.Logger.Info("###########function:", function)
	if len(args) <= 0 {
		return shim.Error("Invoke args error. " + strconv.Itoa(len(args)))
	}
	if args[0] == "Register" {
		return t.Register(stub, args[1:])
	} else if args[0] == "BatchRegister" {
		return t.BatchRegister(stub, args[1:])
	} else if args[0] == "QueryProductChange" {
		return t.QueryProductChange(stub, args[1:])
	} else if args[0] == "GetHistoryByProduct" {
		return t.GetHistoryByProduct(stub, args[1:])
	} else if args[0] == "ChangeOwner" {
		return t.ChangeOwner(stub, args[1:])
	} else if args[0] == "ConfirmChangeOwner" {
		return t.ConfirmChangeOwner(stub, args[1:])
	} else if args[0] == "DestroyProduct" {
		return t.DestroyProduct(stub, args[1:])
	} else if args[0] == "QueryTx" {
		return t.QueryTx(stub, args[1:])
	} else if args[0] == "ChangeProduct" {
		return t.ChangeProductInfo(stub, args[1:])
	} else if args[0] == "QueryProductDetail" {
		return t.QueryProductDetail(stub, args[1:])
	} else if args[0] == "GetProductIdsByInModule" {
		return t.GetProductIdsByInModule(stub, args[1:])
	} else if args[0] == "GetAllInModule" {
		return t.GetAllInModule(stub)
	}
	return shim.Error("Invalid invoke function name. " + args[0])
}

/**
注册产品信息入口
**/
func (t *ProductTrace) Register(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Logger.Info("##############调用产品注册接口开始###############")
	if len(args) < 1 {
		return shim.Error("Register:Incorrect number of arguments. Incorrect number is : " + strconv.Itoa(len(args)))
	}
	var Param module.RegisterParam
	err := json.Unmarshal([]byte(args[0]), &Param)
	if err != nil {
		log.Logger.Error("######解析传入报文参数报错", err)
		return shim.Error("Unmarshal Register args[0] Error," + err.Error())
	}
	return service.Register(stub, Param)
}

/**
注册产品信息入口
**/
func (t *ProductTrace) BatchRegister(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Logger.Info("##############调用批量产品注册接口开始###############")
	if len(args) < 1 {
		return shim.Error("Register:Incorrect number of arguments. Incorrect number is : " + strconv.Itoa(len(args)))
	}
	var Params []module.RegisterParam

	err := json.Unmarshal([]byte(args[0]), &Params)
	if err != nil {
		log.Logger.Error("######解析批量传入报文参数报错", err)
		return shim.Error("Unmarshal Register args[0] Error," + err.Error())
	}
	return service.BatchRegister(stub, Params)
}

/**
变更产品所属信息入口
**/
func (t *ProductTrace) ChangeOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Logger.Info("##############调用变更产品所属接口开始###############")
	if len(args) < 1 {
		return shim.Error("ChangeOwner:Incorrect number of arguments. Incorrect number is : " + strconv.Itoa(len(args)))
	}
	var Param module.ChangeOrgParam
	err := json.Unmarshal([]byte(args[0]), &Param)
	if err != nil {
		log.Logger.Error("######解析传入报文参数报错", err)
		return shim.Error("Unmarshal ChangeOwner args[0] Error," + err.Error())
	}
	return service.ChangeOwner(stub, Param)
}

/**
确认产品权属变更接口入口
**/
func (t *ProductTrace) ConfirmChangeOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Logger.Info("##############调用确认产品权属变更接口开始###############")
	if len(args) < 1 {
		return shim.Error("ConfirmChangeOwner:Incorrect number of arguments. Incorrect number is : " + strconv.Itoa(len(args)))
	}
	var Param module.ComfirmChangeParam
	err := json.Unmarshal([]byte(args[0]), &Param)
	if err != nil {
		log.Logger.Error("######解析传入报文参数报错", err)
		return shim.Error("Unmarshal ConfirmChangeOwner args[0] Error," + err.Error())
	}
	return service.ConfirmChangeOwner(stub, Param)
}

/**
产品销毁入口
**/
func (t *ProductTrace) DestroyProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Logger.Info("##############调用产品销毁接口开始###############")
	if len(args) < 1 {
		return shim.Error("DestroyProduct:Incorrect number of arguments. Incorrect number is : " + strconv.Itoa(len(args)))
	}
	var Param module.DestroyParam
	err := json.Unmarshal([]byte(args[0]), &Param)
	if err != nil {
		log.Logger.Error("######解析传入报文参数报错", err)
		return shim.Error("Unmarshal DestroyProduct args[0] Error," + err.Error())
	}
	return service.DestroyProduct(stub, Param)
}

/**
查询产品流转入口
**/
func (t *ProductTrace) QueryProductChange(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Logger.Info("##############调用查询产品信息接口开始###############")
	if len(args) < 1 {
		return shim.Error("QueryProductChange:Incorrect number of arguments. Incorrect number is : " + strconv.Itoa(len(args)))
	}
	var Param module.QueryParam
	err := json.Unmarshal([]byte(args[0]), &Param)
	if err != nil {
		log.Logger.Error("######解析传入报文参数报错", err)
		return shim.Error("Unmarshal QueryProduct args[0] Error," + err.Error())
	}
	return service.QueryProductChange(stub, Param)
}

/**
查询产品详情入口
**/
func (t *ProductTrace) QueryProductDetail(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Logger.Info("##############调用查询产品详情接口开始###############")

	if len(args) < 1 {
		return shim.Error("QueryProductDetail:Incorrect number of arguments. Incorrect number is : " + strconv.Itoa(len(args)))
	}

	var Param module.QueryProductDetailParam

	err := json.Unmarshal([]byte(args[0]), &Param)
	if err != nil {
		log.Logger.Error("######解析传入报文参数报错", err)
		return shim.Error("Unmarshal QueryProduct args[0] Error," + err.Error())
	}

	return service.QueryProductDetail(stub, Param)
}

/**
查询交易详情入口
**/
func (t *ProductTrace) QueryTx(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Logger.Info("##############调用查询交易信息接口开始###############")
	if len(args) < 1 {
		return shim.Error("QueryTx:Incorrect number of arguments. Incorrect number is : " + strconv.Itoa(len(args)))
	}
	var Param module.QueryTxParam
	err := json.Unmarshal([]byte(args[0]), &Param)
	if err != nil {
		log.Logger.Error("######解析传入报文参数报错", err)
		return shim.Error("Unmarshal QueryTx args[0] Error," + err.Error())
	}
	return service.QueryTx(stub, Param)
}

/**
变更产品属性入口
**/
func (t *ProductTrace) ChangeProductInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Logger.Info("##############调用变更产品属性接口开始###############")
	if len(args) < 1 {
		return shim.Error("ChangeProductInfo:Incorrect number of arguments. Incorrect number is : " + strconv.Itoa(len(args)))
	}
	log.Logger.Info("args:", args[0])
	param, err := common.Json2map(args[0])
	if err != nil {
		log.Logger.Error("######解析传入报文参数报错", err)
		return shim.Error("Convert args[0] to map error," + err.Error())
	}
	log.Logger.Info(param, "xxxxxxx")
	return service.ChangeProductInfo(stub, param)
}

/**
按批次查询耳标号
**/
func (t *ProductTrace) GetProductIdsByInModule(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Logger.Info("##############调用按查询耳标号接口开始###############")

	if len(args) < 1 {
		return shim.Error("GetProductIdsByInModule:Incorrect number of arguments. Incorrect number is : " + strconv.Itoa(len(args)))
	}

	var Param module.GetProductIdsByInModuleParam

	err := json.Unmarshal([]byte(args[0]), &Param)
	if err != nil {
		log.Logger.Error("######解析传入报文参数报错", err)
		return shim.Error("Unmarshal GetProductIdsByInModule args[0] Error," + err.Error())
	}

	return service.GetProductIdsByInModule(stub, Param)
}

/**
查询所有批次
**/
func (t *ProductTrace) GetAllInModule(stub shim.ChaincodeStubInterface) pb.Response {
	log.Logger.Info("##############调用查询所有批次开始###############")

	return service.GetAllInModule(stub)
}

/**
查询所有批次
**/
func (t *ProductTrace) GetHistoryByProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Logger.Info("##############调用查询资产历史数据###############")

	if len(args) < 1 {
		return shim.Error("GetHistoryByProduct:Incorrect number of arguments. Incorrect number is : " + strconv.Itoa(len(args)))
	}
	var Param module.QueryProductDetailParam
	err := json.Unmarshal([]byte(args[0]), &Param)
	if err != nil {
		log.Logger.Error("######解析传入报文参数报错", err)
		return shim.Error("Unmarshal GetHistoryByProduct args[0] Error," + err.Error())
	}
	return service.GetHistoryByProduct(stub, Param)
}
