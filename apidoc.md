# 系统设计报告书 #


## 智能合约接口设计文档 ##


### 版本：V1.1 ###
### 状态：初稿 ###
#### 日期：2018-3-7 ####




## 1 需求概述 


- 1.1 背景及目的
##### 为了明确区块链与核心系统的调用规则，特此编此文档。 #####


- 1.2 适用范围
##### 本文档的的适用对象为业务系统的技术开发人员，日常维护人员。他们需要熟悉以下基本知识：#####
##### 熟悉Restful接口的调用、HttpPost相关基础知识 #####




## 2 集成方案


- 2.1 集成概述
##### 2.1.1 接口列表 #####


### login


接口名称 | 方法代码 | 功能描述
-|:-:|-
用户认证接口 | -- | 获取用户交互凭证进行用户权限认证


### query


接口名称 | 方法代码 | 功能描述
-|:-:|-
查询产品交易历史 | QueryProduct | 功能描述: 通过产品ID，查询该产品所有的交易基本信息
查询交易详情 | QueryTx | 根据交易ID，查询交易详情


### invoke


接口名称 | 方法代码 | 功能描述
-|:-:|-
产品注册 | Register | 注册
产品属性更变 | ChangeProduct | 修改产品的基本属性
产品权属更变 | ChangeOwner | 修改产品的所属权
权属更变确认 | ConfirmChangeOwner | 产品权属变更交易，需要对手方确认
产品销毁 | DestoryProduct | 产品销售后，所属权不属于任何组织，无法再更改产品或变更权属，可视为销毁


- 2.1.2 接入说明
* 系统分为三种不同类型的接口服务，其中“用户认证接口”是用户使用账号以及密码信息获取的有效凭证（token）。调用具体业务接口时需要使用此凭证，此凭证在一定时间内有效。重复调用会导致之前的凭证失效。

- 2.1.3 公共参数说明
* 注：
本章节的说明是针对业务接口（invoke）的说明，获取token的公共接口（login）不在此范围内，请注意。

#### 2.1.3.1 参数结构


- 输入参数结构:

```
{
	"peers": ["节点名","节点名"],
	"channelName": "渠道名称",		// 必填
	"chaincodeName": "智能合约名称",	// 必填 
	"fcn": "方法名称",			// 必填
	"args": ["方法名","业务报文"],
}
```


- 输出参数结构:

```
{
	"success": "调用结果",
	"message": "调用结果数据"
}
```


#### 2.1.3.2 输入参数说明:


数据项名称 | 标签名 | 数据属性 | 限制 | 必填 | 备注
-|:-:|:-:|:-:|:-:|-
调用的节点地址 | peers | []string | -- | T | 指定了调用的区块链节点的信息，待区块链系统提供
渠道名称 | channelName | string | 长度(0 - 50) | T | 指定了渠道的名称，待区块链系统提供
智能合约名称 | chaincodeName | string | 长度(0 - 50) | T | 智能合约名称，待区块链系统提供
方法名称 | fcn | string | 长度(0 - 10) | T | 指定了调用的方法名，invoke或者query
业务参数 | args | []string | -- | T | 数组包括两个参数，第一个参数为接口方法代码（参考2.1.1章节），<br />第二个参数为调用的业务报文json串，<br />具体可参考3章节接口说明(除去3.1章节以外)


#### 2.1.3.3 输出参数说明:


数据项名称 | 标签名 | 数据属性 | 备注
-|:-:|:-:|-
调用结果 | success | string | true: 成功<br />false: 失败
调用结果数据 | message | string | query方法: 调用的返回结果<br />invoke方法: 交易的txId


<br /><br /><br /><br />
## 3 接口说明


### 3.1 获取调用凭证
#### 3.1.1 参数结构 ####


###### 输入参数


- 输入参数结构:
```
{
	"ursername": "用户名",	// 必填
	"password": "密码",	// 必填
	"orgName": "机构名称"	// 必填
}
```
- 输入参数示例:
```
{
	"username": "admin",	// 必填
	"password": "cde34rfv",	// 必填
	"orgName": "org1"	// 必填
}
```


###### 输出参数


- 输出参数结构
```
{
	"success": "返回结果",
	"secret": "密码",
	"message": "返回结果描述",
	"token": "返回的用户凭证",
	"exp": 过期时间点,
	"usefulTime": 有效时长(单位/秒)
}
```
- 输出参数示例
```
{
	"success": "true",
	"secret": "",
	"message": "admin enrolled Successfully",
	"token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0OTg0ODM0NjcsInVzZXJuYW1lIjoiYWRtaW4xIiwib3JnTmFtZSI6Im9yZzEiLCJpYXQiOjE0OTg0NDc0Njd9.YJVFfcfii1E8mYpzJ5Ac2mG0n_PLRCd97w429WD_A8A",
	"exp": 1504528295,
	"usefulTime": 36000
}
```


#### 3.1.2 请求报文参数说明


数据项名称 | 标签名 | 数据属性 | 限制 | 必填 | 备注
-|:-:|:-:|:-:|:-:|-
用户名 | username | string | length(1:30) | T | --
密码 | password | string | length(1:30) | T | --
机构名称 | orgName | string | length(1:30) | T | --


#### 3.1.3 返回报文参数说明


数据项名称 | 标签名 | 数据属性 | 备注
-|:-:|:-:|-
调用结果 | success | string | true: 成功<br />false: 失败
密码 | secret | string | --
调用结果描述 | message | string | --
用户凭证 | token | string | --
过期时间点 | exp | long int | --
过期时长 | usefulTime | long int | 单位/秒


<br /><br /><br /><br />
### 3.2 产品注册 Register


#### 3.2.1 参数结构


###### 请求参数


- 请求参数结构
```
{
	"productId": "产品编号",	// 必填
	"productName": "产品名称",	// 必填
	"inModule": "入栏批次",	
	"kind": "种类",			// 必填
	"type": "品种",			// 必填
	"mapPosition": "地理位置",	// 必填
	"iSerial": "入栏终端编号",	// 必填
	"kSerial": "待宰终端编号",	// 必填
	"lairage": "入栏时间",		// 必填
	"days": "生长天数",		// 必填
	"condition": "健康状况"		// 必填
	"comment": "行为"		// 必填
	"penNum": "圈号"		// 必填
}
```
- 请求参数示例
```
{
	"productId": "800010001",		// 必填
	"productName": "第八系",		// 必填
	"inModule": "xxxx1"	
	"kind": "羊",				// 必填
	"type": "小肥羊",			// 必填
	"mapPosition": "北纬17°，东经143°",	// 必填
	"iserial": "AX00010001",		
	"kserial": "AY00010001",		
	"lairage": "2001-01-01",		// 必填
	"days": "50"				// 必填
	"condition": "良好"			// 必填
	"comment": "注册"			// 必填
	"penNum": "A1"				// 必填
}
```


#### 3.2.2 请求参数说明


数据项名称 | 标签名 | 数据属性 | 必填 | 备注
-|:-:|:-:|:-:|:-:|-
产品编号 | productId | string | T | --
产品名称 | productName | string | T | --
入栏批次 | inModule | string | T | --
待宰批次 | killModule | string | F | --
种类 | kind | string | T | --
品种 | type | string | T | --
地理位置 | mapPosition | string | T | --
入栏终端编号 | iSerial | string | T | --
待宰终端编号 | kSerail | string | T | --
入栏时间 | lairage | string | T | --
生长天数 | days | string | T | --
健康状况 | condition | string | T | --
行为 | comment | string | T | --
圈号 | penNum | string | T | --


<br /><br /><br /><br />
### 3.3 产品属性变更 ChangeProduct


#### 3.3.1 参数结构


###### 输入参数


- 输入参数结构
```
{
	"productId": "产品编号",	// 必填
	"productName": "产品名称"
}
```
- 输入参数示例
```
{
	"productId": "800010001",	// 必填
	"productName": "第8系列一号"
}
```


#### 3.3.2 参数说明


数据项名称 | 标签名 | 数据属性 | 限制 | 必填 | 备注
-|:-:|:-:|:-:|:-:|-
产品编号 | productId | string | length(1:100) | T | --
产品名称 | productName | string | length(1:200) | F | --


<br /><br /><br /><br />
### 3.4 产品权属变更 ChangeOwner


#### 3.4.1 参数结构


###### 参数输入


- 输入参数结构:
```
{
	"productId": "产品编号",		// 必填
	"toOrgMsgId": "变更方的组织机构ID"	// 必填
}
```
- 输入参数示例:
```
{
	"productId": "800010001",		// 必填
	"toOrgMsgId": "pinganmsp"		// 必填
}
```


#### 3.4.2 参数说明


数据项名称 | 标签名 | 数据属性 | 限制 | 必填 | 备注
-|:-:|:-:|:-:|:-:|-
产品编号 | productId | string | length(1:100) | T | --
变更方的组织机构ID | toOrgMsgId | string | length(1:100) | T | --


<br /><br /><br /><br />
### 3.5 确认权属变更 ConfirmChangeOwner


#### 3.5.1 参数结构


###### 参数输入


- 输入参数结构:
```
{
	"productId": "产品编号",		// 必填
}
```
- 输入参数示例:
```
{
	"productId": "800010001",		// 必填
}
```
- 备注:
```
交易发起时会校验交易发起方的权限，是否能确认该笔交易，以及是否已经被确认
```


#### 3.5.2 参数说明


数据项名称 | 标签名 | 数据属性 | 限制 | 必填 | 备注
-|:-:|:-:|:-:|:-:|-
产品编号 | productId | string | length(1:100) | T | --


<br /><br /><br /><br />
### 3.6 产品销毁 DestroyProduct


#### 3.6.1 参数结构


###### 参数输入


- 输入参数结构:
```
{
	"productId": "产品编号",	// 必填
	"serialNum": "销售订单流水号"	// 必填
}
```
- 输入参数示例:
```
{
	"productId": "800010001",	// 必填
	"serialNum": "2018021111190001"	// 必填
}
```
- 备注:
```
交易发起时会校验交易发起方的权限 
```


#### 3.6.2 输入参数说明


数据项名称 | 标签名 | 数据属性 | 限制 | 必填 | 备注
-|:-:|:-:|:-:|:-:|-
产品编号 | productId | string | length(1:100) | T | --
销售订单流水号 | serialNum | string | length(1:100) | T | --


<br /><br /><br /><br />
### 3.7 查询产品交易历史 QueryProduct


#### 3.7.1 参数结构


- 输入参数结构:
```
{
	"productId": "产品编号",	// 必填
}
```
- 输入参数示例:
```
{
	"productId": "800010001",	// 必填
}
```


- 输出参数结构:
```
[{
	"txId": "交易ID",
	"txType": "交易类型",
	"preOwner": "上一次所属",
	"currentOwner": "当前所属",
	"txTime": "交易时间"
}]
```
- 输出参数示例:
```
[{
	"txId": "6e29bda8b55e3d92d785536ed00e8a2b5812e02e2d8339252c02380624ec8f45",
	"txType": "01",
	"preOwner": "SYSTEM",
	"currentOwner": "shiemsp",
	"txTime": "1518145694"
}, {
	"txId": "0725288e8d4ed16e1b5e5c3d39d7b3f2971e99304c41495a935478ceb2aa1c8e",
	"txType": "04",
	"preOwner": "SYSTEM",
	"currentOwner": "shiemsp",
	"txTime": "1518146300"
}, {
	"txId": "1a51615b6b98e3ab0f4b474c1421e79ab72aefc7d863bfc2106ad93939c7620c",
	"txType": "02",
	"preOwner": "shiemsp",
	"currentOwner": "UNCONFIRM_taipingmsp",
	"txTime": "1518146348"
}, {
	"txId": "8ed5fbcaf7bddb7fe5fd72d478920f4984babfa617b03bd2825f0f9885de4fe3",
	"txType": "03",
	"preOwner": "UNCONFIRM_taipingmsp",
	"currentOwner": "taipingmsp",
	"txTime": "1518146435"
}, {
	"txId": "444ae96db7090762979b522a8b7c421dea8c34a778aefed8ea127310674bacaa",
	"txType": "99",
	"preOwner": "taipingmsp",
	"currentOwner": "2018020914420001",
	"txTime": "1518146458"
}]
```


#### 3.7.2 输入参数说明


数据项名称 | 标签名 | 数据属性 | 限制 | 必填 | 备注
-|:-:|:-:|:-:|:-:|-
产品编号 | productId | string | length(1:100) | T | --


#### 3.7.2 输出参数说明


数据项名称 | 标签名 | 数据属性 | 备注
-|:-:|:-:|-
交易ID | txId | string | --
交易类型 | txType | [2]char | 参考4.1章节代码定义
前一次所属 | preOwner | string | --
当前所属 | currentOwner | string | --
交易时间 | txTime | long int | --


<br /><br /><br /><br />
## 3.8 查询交易详情


#### 3.8.1 参数结构 QueryTx


- 输入参数结构
```
{
	"txId": "交易ID"								// 必填
}
```
- 输入参数示例
```
{
	"txId": "8ed5fbcaf7bddb7fe5fd72d478920f4984babfa617b03bd2825f0f9885de4fe3"	// 必填
}
```
- 备注
```
不同的交易类型，返回的交易详情信息不同，具体格式参考下文。
```


	- 01-产品注册交易


	###### 报文结构
	```
	{
		"productId": "产品编号",
		"productName": "产品名称"
		"kind": "种类",
		"type": "品种",
		"earCode": "耳标编号",
		"nLatitude": "北纬",
		"eLongitude": "东经",
		"serial": "终端编号",
		"eName": "操作人员名字",
		"birth": "注册时间戳",
		"condition": "健康状况"
	}
	```
	###### 报文示例
	```
	{
		"productId": "800010001",
		"productName": "第8系列一号"
		"kind": 1,
		"type": 20,
		"earCode": "ER20180302",
		"nLatitude": 35.14,
		"eLongitude": 104.17,
		"serial": "AX00010001",
		"name": "穆阳人"
		"birth": "1520481916",
		"condition": 0	
	}
	```

	
	- 02-产品权属变更


	###### 报文结构
	```
	{
		"before": {				// 变更前信息
			"preOwner": "前一次所属",
			"currentOwner": "当前所属"
		},
		"after": {				// 变更后信息
			"preOwner": "前一次所属",
			"currentOwner": "当前所属"
		}
	}
	```
	###### 报文示例
	```
	{
		"before": {				// 变更前信息
			"preOwner": "pinganmsp",
			"currentOwner": "renbaomsp"
		},
		"after": {				// 变更后信息
			"preOwner": "renbaomsp",
			"currentOwner": "UNCONFIRM_taipingmsp"
		}
	}
	```


	- 03-确认权属变更


	###### 报文结构
	```
	{
		"before": {				// 变更前信息
			"preOwner": "前一次所属",
			"currentOwner": "当前所属"
		},
		"after": {				// 变更后信息
			"preOwner": "前一次所属",
			"currentOwner": "当前所属"
		}
	}
	```
	###### 报文示例
	```
	{
		"before": {				// 变更前信息
			"preOwner": "pinganmsp",
			"currentOwner": "renbaomsp"
		},
		"after": {				// 变更后信息
			"preOwner": "renbaomsp",
			"currentOwner": "UNCONFIRM_taipingmsp"
		}
	}
	```


	- 04-产品属性变更


	###### 报文结构
	```
	{
		"before": {				// 变更前信息
			"productId": "产品编号",
			"productName": "产品名称"
		},
		"after": {				// 变更后信息
			"productId": "产品编号",
			"productName": "产品名称"
		}
	}
	```
	###### 报文示例
	```
	{
		"before": {				// 变更前信息
			"productId": "800010001",
			"productName": "变更前名称"
		},
		"after": {				// 变更后信息
			"productId": "800010001",
			"productName": "变更后名称"
		}
	}
	```


	- 99-产品销毁
	###### 报文结构
	```
	{
		"before": {				// 变更前信息
			"preOwner": "前一次所属",
			"currentOwner": "当前所属"
		},
		"after": {				// 变更后信息
			"preOwner": "前一次所属",
			"currentOwner": "销售时的订单流水号"
		}
	}
	```
	###### 报文示例
	```
	{
		"before": {				// 变更前信息
			"preOwner": "UNCONFIRM_taipingmsp",
			"currentOwner": "taipingmsp"
		},
		"after": {				// 变更后信息
			"preOwner": "taipingmsp",
			"currentOwner": "2018021114390001"
		}
	}
	```


#### 3.8.2 输入参数说明


数据项名称 | 标签名 | 数据属性 | 限制 | 必填 | 备注
-|:-:|:-:|:-:|:-:|-
交易 ID | txId | string | length(1:500) | T | --


#### 3.8.2 输出参数说明


数据项名称 | 标签名 | 数据属性 | 备注
-|:-:|:-:|-
产品编号 | productId | string | --
产品名称 | productName | string | --
前一次所属 | preOwner | string | --
当前所属 | currentOwner | string | 产品销毁时，为销售时的订单流水号


<br /><br /><br /><br />
## 4 代码定义


### 4.1 交易类型(txType)


代码 | 名称 | 说明
-|:-:|-
01 | 产品注册 | --
02 | 产品权属变更交易 | --
03 | 确认产品权属变更交易 | --
04 | 产品属性更变 | --
99 | 产品销售销毁 | --
