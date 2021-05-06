package model

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Status_Code int32

const (
	Status_CODE_UNSPECIFIED            Status_Code = 0        // 正常
	Status_CODE_TASK_CONCURRENCY_LIMIT Status_Code = 20050001 // 任务达到并发数限制
)

// Enum value maps for Status_Code.
var (
	Status_Code_name = map[int32]string{
		0:        "CODE_UNSPECIFIED",
		20050001: "CODE_TASK_CONCURRENCY_LIMIT",
	}
	Status_Code_value = map[string]int32{
		"CODE_UNSPECIFIED":            0,
		"CODE_TASK_CONCURRENCY_LIMIT": 20050001,
	}
)

func (x Status_Code) Enum() *Status_Code {
	p := new(Status_Code)
	*p = x
	return p
}

func (x Status_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_asynctask_proto_asynctask_proto_enumTypes[0].Descriptor()
}

func (Status_Code) Type() protoreflect.EnumType {
	return &file_asynctask_proto_asynctask_proto_enumTypes[0]
}

func (x Status_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status_Code.Descriptor instead.
func (Status_Code) EnumDescriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{0, 0}
}

// 任务分类
type Task_Category int32

const (
	Task_CATEGORY_UNSPECIFIED Task_Category = 0
	Task_CATEGORY_EXPORT      Task_Category = 1   // 导入导出
	Task_CATEGORY_OTHER       Task_Category = 999 // 未分类/其他分类
)

// Enum value maps for Task_Category.
var (
	Task_Category_name = map[int32]string{
		0:   "CATEGORY_UNSPECIFIED",
		1:   "CATEGORY_EXPORT",
		999: "CATEGORY_OTHER",
	}
	Task_Category_value = map[string]int32{
		"CATEGORY_UNSPECIFIED": 0,
		"CATEGORY_EXPORT":      1,
		"CATEGORY_OTHER":       999,
	}
)

func (x Task_Category) Enum() *Task_Category {
	p := new(Task_Category)
	*p = x
	return p
}

func (x Task_Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_Category) Descriptor() protoreflect.EnumDescriptor {
	return file_asynctask_proto_asynctask_proto_enumTypes[1].Descriptor()
}

func (Task_Category) Type() protoreflect.EnumType {
	return &file_asynctask_proto_asynctask_proto_enumTypes[1]
}

func (x Task_Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_Category.Descriptor instead.
func (Task_Category) EnumDescriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{2, 0}
}

// 任务类型，形如TYPE_模块_数据模型_操作
// 前端展示的任务名称为“任务类型名称-时间”，任务类型名称请填写在枚举的注释中。请认真填写任务类型名称，以“数据类型名称操作名称”的格式定义
type Task_Type int32

const (
	Task_TYPE_UNSPECIFIED                                  Task_Type = 0
	Task_TYPE_ASYNCTASK_FOO_IMPORT                         Task_Type = 1    // 正常测试
	Task_TYPE_ASYNCTASK_FOO_EXPORT                         Task_Type = 2    // 失败测试
	Task_TYPE_ORDER_ORDER_IMPORT                           Task_Type = 1000 // 订单导入
	Task_TYPE_ORDER_ORDERSTATE_UPDATE                      Task_Type = 1001 // 订单状态修改
	Task_TYPE_ORDER_ORDEROUTSTOCK_UPDATE                   Task_Type = 1002 // 订单缺货修改
	Task_TYPE_ORDER_ORDERSSUPRICE_UPDATE                   Task_Type = 1003 // 订单ssu价格修改
	Task_TYPE_ORDER_SSUPRICE_SYNC                          Task_Type = 1004 // 报价单最新单价同步到订单
	Task_TYPE_ORDER_ORDER_CREATE                           Task_Type = 1005 // 订单创建
	Task_TYPE_ORDER_ORDER_EXPORT                           Task_Type = 1006 // 订单导出
	Task_TYPE_ORDER_ORDERDRIVER_UPDATE                     Task_Type = 1007 // 订单司机修改
	Task_TYPE_ORDER_ORDERDRIVER_AUTOUPDATE                 Task_Type = 1008 // 订单司机自动修改
	Task_TYPE_ORDER_ORDER_SUBMIT_PURCHASE_TASK             Task_Type = 1009 // 订单发布采购计划
	Task_TYPE_ORDER_SCAN_RECEIPT                           Task_Type = 1010 // 扫码回单
	Task_TYPE_ORDER_COMBINESSU_SYNC                        Task_Type = 1011 // 订单批量同步组合商品信息到订单
	Task_TYPE_ORDER_ORDERDETAILSSU_UPDATE                  Task_Type = 1012 // 订单批量修改订单商品
	Task_TYPE_ORDER_SORTING_COMPANY_SUMMARY_EXPORT         Task_Type = 1013 // 按公司汇总导出分拣单
	Task_TYPE_ORDER_ORDER_BATCH_CREATE_BY_MENU             Task_Type = 1014 // 根据菜谱批量创建订单
	Task_TYPE_PRODUCTION_TASK_CREATE                       Task_Type = 1100 // 生产计划创建
	Task_TYPE_PRODUCTION_PURCHASE_SKU_EXPORT               Task_Type = 1101 // 物料需求汇总
	Task_TYPE_PRODUCTION_TASK_EXPORT                       Task_Type = 1102 // 生产计划导出
	Task_TYPE_PRODUCTION_BOM_EXPORT                        Task_Type = 1103 // bom导出
	Task_TYPE_PRODUCTION_TASK_DATA_EXPORT                  Task_Type = 1104 // 生产计划报表导出
	Task_TYPE_PRODUCTION_BOM_IMPORT                        Task_Type = 1105 // bom导入
	Task_TYPE_PRODUCTION_PROCESS_TEMPLATE_EXPORT           Task_Type = 1106 // 工序导出
	Task_TYPE_PRODUCTION_TASK_RELEASE                      Task_Type = 1107 // 计划发布任务
	Task_TYPE_PRODUCTION_PROCESS_TASK_EXPORT               Task_Type = 1108 // 生产任务导出
	Task_TYPE_MERCHANDISE_SKU_EXPORT                       Task_Type = 1200 // 商品导出（task_data.parameters.values[0] 是 excel 文件的下载链接）
	Task_TYPE_MERCHANDISE_SSU_EXPORT                       Task_Type = 1201 // 商品规格导出
	Task_TYPE_MERCHANDISE_SKU_SSU_UPDATE                   Task_Type = 1202 // sku和ssu批量修改
	Task_TYPE_MERCHANDISE_CREATE_ALL                       Task_Type = 1203 // 分类,spu,sku新建
	Task_TYPE_MERCHANDISE_CREATE_SSU                       Task_Type = 1204 // ssu新建
	Task_TYPE_MERCHANDISE_MENU_DETAIL_EXPORT               Task_Type = 1205 // 菜谱导出
	Task_TYPE_MERCHANDISE_MENU_DETAIL_IMPORT               Task_Type = 1206 // 菜谱导入
	Task_TYPE_MERCHANDISE_COMBINE_SSU_IMPORT               Task_Type = 1207 // 组商品导入
	Task_TYPE_MERCHANDISE_COMBINE_SSU_EXPORT               Task_Type = 1208 // 组商品导出
	Task_TYPE_MERCHANDISE_COMBINE_SSU_UPDATE               Task_Type = 1209 // 组商品导出修改
	Task_TYPE_MERCHANDISE_QUOTATION_EXPORT                 Task_Type = 1210 // 报价协议单导出
	Task_TYPE_MERCHANDISE_QUOTATION_EXPORT_BY_SSU          Task_Type = 1211 // 报价协议单按商品导出
	Task_TYPE_MERCHANDISE_QUOTATION_IMPORT                 Task_Type = 1212 // 报价协议单导入
	Task_TYPE_MERCHANDISE_QUOTATION_IMPORT_CREATE_BY_SSU   Task_Type = 1213 // 按商品导入协议单并创建
	Task_TYPE_MERCHANDISE_QUOTATION_IMPORT_UPDATE_BY_SSU   Task_Type = 1214 // 报价协议单导入并更新
	Task_TYPE_ENTERPRISE_CUSTOMER_EXPORT                   Task_Type = 1300 // 商户导出
	Task_TYPE_ENTERPRISE_CUSTOMER_IMPORT                   Task_Type = 1301 // 商户导入
	Task_TYPE_ENTERPRISE_SUPPLIER_EXPORT                   Task_Type = 1302 // 供应商导出
	Task_TYPE_ENTERPRISE_SUPPLIER_IMPORT                   Task_Type = 1303 // 供应商导入
	Task_TYPE_INVENTORY_STOCK_SHEET_LIST_EXPORT            Task_Type = 1400 // 库存单据导出
	Task_TYPE_INVENTORY_BATCH_LOG_EXPORT                   Task_Type = 1402 // 批次流水导出
	Task_TYPE_INVENTORY_COMMIT_ADJUST_SHEET                Task_Type = 1403 // 库存调整单提交
	Task_TYPE_INVENTORY_BATCH_SYNC_STOCK_SHEET             Task_Type = 1404 // 批量同步单据
	Task_TYPE_INVENTORY_CHANGE_LOG_EXPORT                  Task_Type = 1410 // 商品台账导出
	Task_TYPE_INVENTORY_PURCHASE_IN_LOG_EXPORT             Task_Type = 1411 // 采购入库明细导出
	Task_TYPE_INVENTORY_SALE_OUT_LOG_EXPORT                Task_Type = 1412 // 销售出库明细导出
	Task_TYPE_INVENTORY_MATERIAL_OUT_LOG_EXPORT            Task_Type = 1413 // 领料出库明细导出
	Task_TYPE_INVENTORY_MATERIAL_IN_LOG_EXPORT             Task_Type = 1414 // 退料入库明细导出
	Task_TYPE_INVENTORY_PRODUCT_IN_LOG_EXPORT              Task_Type = 1415 // 生产入库明细导出
	Task_TYPE_INVENTORY_REFOUND_OUT_LOG_EXPORT             Task_Type = 1416 // 采购退货出库明细导出
	Task_TYPE_INVENTORY_OTHER_OUT_LOG_EXPORT               Task_Type = 1417 // 其他出库明细导出
	Task_TYPE_INVENTORY_OTHER_IN_LOG_EXPORT                Task_Type = 1418 // 其他入库明细导出
	Task_TYPE_INVENTORY_INCREASE_IN_LOG_EXPORT             Task_Type = 1419 // 盘盈入库明细导出
	Task_TYPE_INVENTORY_LOSS_OUT_LOG_EXPORT                Task_Type = 1420 // 盘亏出库明细导出
	Task_TYPE_INVENTORY_CUSTOMER_TURNOVER_EXPORT           Task_Type = 1421 // 周转物商户借出汇总导出
	Task_TYPE_INVENTORY_LOAN_EXPORT                        Task_Type = 1422 // 周转物借出记录导出
	Task_TYPE_INVENTORY_REVERT_EXPORT                      Task_Type = 1423 // 周转物归还记录导出
	Task_TYPE_INVENTORY_CUSTOMER_TURNOVER_LOG_EXPORT       Task_Type = 1424 // 商户周转物借出归还记录
	Task_TYPE_INVENTORY_REFOUND_IN_LOG_EXPORT              Task_Type = 1425 // 销售退货入库记录导出
	Task_TYPE_INVENTORY_TRANSFER_IN_LOG_EXPORT             Task_Type = 1426 // 调拨入库记录导出
	Task_TYPE_INVENTORY_TRANSFER_OUT_LOG_EXPORT            Task_Type = 1427 // 调拨出库记录导出
	Task_TYPE_INVENTORY_SSU_STOCK_VALUE_EXPORT             Task_Type = 1430 // ssu货值成本表导出
	Task_TYPE_INVENTORY_SHEET_EXPORT                       Task_Type = 1450 // 单据导出
	Task_TYPE_INVENTORY_BATCH_UPDATE_STOCK_SHEET           Task_Type = 1498 // 批量修改单据
	Task_TYPE_INVENTORY_SKU_STOCK_EXPORT                   Task_Type = 1499 // 库存总览导出
	Task_TYPE_SORTING_SORTING_PERFORMANCE_EXPORT           Task_Type = 1500 // 分拣绩效导出
	Task_TYPE_PURCHASE_PURCHASE_ORDER_EXPORT               Task_Type = 1601 //采购单导出
	Task_TYPE_PURCHASE_INQUIRY_PRICE_IMPORT                Task_Type = 1602 //询价导入
	Task_TYPE_PURCHASE_PURCHASE_TASK_RELEASE               Task_Type = 1603 //下达采购计划
	Task_TYPE_PURCHASE_PURCHASE_TASK_SWITCH_SUPPLIER       Task_Type = 1604 //批量修改供应商
	Task_TYPE_PURCHASE_PURCHASE_TASK_SWITCH_PURCHASER      Task_Type = 1605 //批量修改采购员
	Task_TYPE_PURCHASE_PURCHASE_TASK_CREATE_PURCHASE_ORDER Task_Type = 1606 //批量生成采购单
	Task_TYPE_PURCHASE_PURCHASE_TASK_EXPORT                Task_Type = 1607 //采购计划导出
	Task_TYPE_PURCHASE_INQUIRY_PRICE_IMPORT_TMPL           Task_Type = 1608 //询价导入模版
	Task_TYPE_PURCHASE_PURCHASE_TASK_EXPORT_DIMENSION      Task_Type = 1609 //采购计划导出二维表
	Task_TYPE_PURCHASE_PURCHASE_TASK_BATCH_REDUCE_STOCK    Task_Type = 1610 // 批量扣减库存
	Task_TYPE_ESHOP_ESHOP_ORDER_EXPORT                     Task_Type = 1706 // eshop订单导出
	Task_TYPE_ESHOP_ESHOP_ORDER_UPDATE                     Task_Type = 1707 // eshop订单批量修改
	Task_TYPE_ESHOP_LEAVEAPPLICATION_EXPORT                Task_Type = 1801 // eshop请假导出
	Task_TYPE_AFTER_SALE_EXPORT_ORDER                      Task_Type = 1900 // 导出售后订单
	Task_TYPE_AFTER_SALE_EXPORT_TASK                       Task_Type = 1901 // 导出取货任务
	Task_TYPE_AFTER_SALE_EXPORT_ORDER_SHEET                Task_Type = 1902 // 导出售后分析报表
	Task_TYPE_AFTER_SALE_EXPORT_ORDER_DETAIL               Task_Type = 1903 // 导出订单售后明细报表
	// INITIALIZATION
	Task_TYPE_INITIALIZATION_CLEAN_GROUP_TICKET  Task_Type = 2000 // 清空 Group 单据
	Task_TYPE_INITIALIZATION_CLEAN_GROUP_DATA    Task_Type = 2001 // 清空 Group 数据
	Task_TYPE_FINANCE_SETTLE_EXPORT              Task_Type = 2100 // 商户结算信息导出
	Task_TYPE_FINANCE_BATCH_PAY_SETTLE           Task_Type = 2101 // 结款单批量结算
	Task_TYPE_FINANCE_FLOW_EXPORT                Task_Type = 2102 // 流水导出
	Task_TYPE_FINANCE_ACCOUNT_BALANCE_EXPORT     Task_Type = 2103 // 商户余额导出
	Task_TYPE_FINANCE_BATCH_SUBMIT_SETTLE_SHEET  Task_Type = 2104 // 结款单批量提交
	Task_TYPE_FINANCE_MERCHANT_SETTLEMENT_EXPORT Task_Type = 2105 // 对账单详情导出
)

// Enum value maps for Task_Type.
var (
	Task_Type_name = map[int32]string{
		0:    "TYPE_UNSPECIFIED",
		1:    "TYPE_ASYNCTASK_FOO_IMPORT",
		2:    "TYPE_ASYNCTASK_FOO_EXPORT",
		1000: "TYPE_ORDER_ORDER_IMPORT",
		1001: "TYPE_ORDER_ORDERSTATE_UPDATE",
		1002: "TYPE_ORDER_ORDEROUTSTOCK_UPDATE",
		1003: "TYPE_ORDER_ORDERSSUPRICE_UPDATE",
		1004: "TYPE_ORDER_SSUPRICE_SYNC",
		1005: "TYPE_ORDER_ORDER_CREATE",
		1006: "TYPE_ORDER_ORDER_EXPORT",
		1007: "TYPE_ORDER_ORDERDRIVER_UPDATE",
		1008: "TYPE_ORDER_ORDERDRIVER_AUTOUPDATE",
		1009: "TYPE_ORDER_ORDER_SUBMIT_PURCHASE_TASK",
		1010: "TYPE_ORDER_SCAN_RECEIPT",
		1011: "TYPE_ORDER_COMBINESSU_SYNC",
		1012: "TYPE_ORDER_ORDERDETAILSSU_UPDATE",
		1013: "TYPE_ORDER_SORTING_COMPANY_SUMMARY_EXPORT",
		1014: "TYPE_ORDER_ORDER_BATCH_CREATE_BY_MENU",
		1100: "TYPE_PRODUCTION_TASK_CREATE",
		1101: "TYPE_PRODUCTION_PURCHASE_SKU_EXPORT",
		1102: "TYPE_PRODUCTION_TASK_EXPORT",
		1103: "TYPE_PRODUCTION_BOM_EXPORT",
		1104: "TYPE_PRODUCTION_TASK_DATA_EXPORT",
		1105: "TYPE_PRODUCTION_BOM_IMPORT",
		1106: "TYPE_PRODUCTION_PROCESS_TEMPLATE_EXPORT",
		1107: "TYPE_PRODUCTION_TASK_RELEASE",
		1108: "TYPE_PRODUCTION_PROCESS_TASK_EXPORT",
		1200: "TYPE_MERCHANDISE_SKU_EXPORT",
		1201: "TYPE_MERCHANDISE_SSU_EXPORT",
		1202: "TYPE_MERCHANDISE_SKU_SSU_UPDATE",
		1203: "TYPE_MERCHANDISE_CREATE_ALL",
		1204: "TYPE_MERCHANDISE_CREATE_SSU",
		1205: "TYPE_MERCHANDISE_MENU_DETAIL_EXPORT",
		1206: "TYPE_MERCHANDISE_MENU_DETAIL_IMPORT",
		1207: "TYPE_MERCHANDISE_COMBINE_SSU_IMPORT",
		1208: "TYPE_MERCHANDISE_COMBINE_SSU_EXPORT",
		1209: "TYPE_MERCHANDISE_COMBINE_SSU_UPDATE",
		1210: "TYPE_MERCHANDISE_QUOTATION_EXPORT",
		1211: "TYPE_MERCHANDISE_QUOTATION_EXPORT_BY_SSU",
		1212: "TYPE_MERCHANDISE_QUOTATION_IMPORT",
		1213: "TYPE_MERCHANDISE_QUOTATION_IMPORT_CREATE_BY_SSU",
		1214: "TYPE_MERCHANDISE_QUOTATION_IMPORT_UPDATE_BY_SSU",
		1300: "TYPE_ENTERPRISE_CUSTOMER_EXPORT",
		1301: "TYPE_ENTERPRISE_CUSTOMER_IMPORT",
		1302: "TYPE_ENTERPRISE_SUPPLIER_EXPORT",
		1303: "TYPE_ENTERPRISE_SUPPLIER_IMPORT",
		1400: "TYPE_INVENTORY_STOCK_SHEET_LIST_EXPORT",
		1402: "TYPE_INVENTORY_BATCH_LOG_EXPORT",
		1403: "TYPE_INVENTORY_COMMIT_ADJUST_SHEET",
		1404: "TYPE_INVENTORY_BATCH_SYNC_STOCK_SHEET",
		1410: "TYPE_INVENTORY_CHANGE_LOG_EXPORT",
		1411: "TYPE_INVENTORY_PURCHASE_IN_LOG_EXPORT",
		1412: "TYPE_INVENTORY_SALE_OUT_LOG_EXPORT",
		1413: "TYPE_INVENTORY_MATERIAL_OUT_LOG_EXPORT",
		1414: "TYPE_INVENTORY_MATERIAL_IN_LOG_EXPORT",
		1415: "TYPE_INVENTORY_PRODUCT_IN_LOG_EXPORT",
		1416: "TYPE_INVENTORY_REFOUND_OUT_LOG_EXPORT",
		1417: "TYPE_INVENTORY_OTHER_OUT_LOG_EXPORT",
		1418: "TYPE_INVENTORY_OTHER_IN_LOG_EXPORT",
		1419: "TYPE_INVENTORY_INCREASE_IN_LOG_EXPORT",
		1420: "TYPE_INVENTORY_LOSS_OUT_LOG_EXPORT",
		1421: "TYPE_INVENTORY_CUSTOMER_TURNOVER_EXPORT",
		1422: "TYPE_INVENTORY_LOAN_EXPORT",
		1423: "TYPE_INVENTORY_REVERT_EXPORT",
		1424: "TYPE_INVENTORY_CUSTOMER_TURNOVER_LOG_EXPORT",
		1425: "TYPE_INVENTORY_REFOUND_IN_LOG_EXPORT",
		1426: "TYPE_INVENTORY_TRANSFER_IN_LOG_EXPORT",
		1427: "TYPE_INVENTORY_TRANSFER_OUT_LOG_EXPORT",
		1430: "TYPE_INVENTORY_SSU_STOCK_VALUE_EXPORT",
		1450: "TYPE_INVENTORY_SHEET_EXPORT",
		1498: "TYPE_INVENTORY_BATCH_UPDATE_STOCK_SHEET",
		1499: "TYPE_INVENTORY_SKU_STOCK_EXPORT",
		1500: "TYPE_SORTING_SORTING_PERFORMANCE_EXPORT",
		1601: "TYPE_PURCHASE_PURCHASE_ORDER_EXPORT",
		1602: "TYPE_PURCHASE_INQUIRY_PRICE_IMPORT",
		1603: "TYPE_PURCHASE_PURCHASE_TASK_RELEASE",
		1604: "TYPE_PURCHASE_PURCHASE_TASK_SWITCH_SUPPLIER",
		1605: "TYPE_PURCHASE_PURCHASE_TASK_SWITCH_PURCHASER",
		1606: "TYPE_PURCHASE_PURCHASE_TASK_CREATE_PURCHASE_ORDER",
		1607: "TYPE_PURCHASE_PURCHASE_TASK_EXPORT",
		1608: "TYPE_PURCHASE_INQUIRY_PRICE_IMPORT_TMPL",
		1609: "TYPE_PURCHASE_PURCHASE_TASK_EXPORT_DIMENSION",
		1610: "TYPE_PURCHASE_PURCHASE_TASK_BATCH_REDUCE_STOCK",
		1706: "TYPE_ESHOP_ESHOP_ORDER_EXPORT",
		1707: "TYPE_ESHOP_ESHOP_ORDER_UPDATE",
		1801: "TYPE_ESHOP_LEAVEAPPLICATION_EXPORT",
		1900: "TYPE_AFTER_SALE_EXPORT_ORDER",
		1901: "TYPE_AFTER_SALE_EXPORT_TASK",
		1902: "TYPE_AFTER_SALE_EXPORT_ORDER_SHEET",
		1903: "TYPE_AFTER_SALE_EXPORT_ORDER_DETAIL",
		2000: "TYPE_INITIALIZATION_CLEAN_GROUP_TICKET",
		2001: "TYPE_INITIALIZATION_CLEAN_GROUP_DATA",
		2100: "TYPE_FINANCE_SETTLE_EXPORT",
		2101: "TYPE_FINANCE_BATCH_PAY_SETTLE",
		2102: "TYPE_FINANCE_FLOW_EXPORT",
		2103: "TYPE_FINANCE_ACCOUNT_BALANCE_EXPORT",
		2104: "TYPE_FINANCE_BATCH_SUBMIT_SETTLE_SHEET",
		2105: "TYPE_FINANCE_MERCHANT_SETTLEMENT_EXPORT",
	}
	Task_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":                                  0,
		"TYPE_ASYNCTASK_FOO_IMPORT":                         1,
		"TYPE_ASYNCTASK_FOO_EXPORT":                         2,
		"TYPE_ORDER_ORDER_IMPORT":                           1000,
		"TYPE_ORDER_ORDERSTATE_UPDATE":                      1001,
		"TYPE_ORDER_ORDEROUTSTOCK_UPDATE":                   1002,
		"TYPE_ORDER_ORDERSSUPRICE_UPDATE":                   1003,
		"TYPE_ORDER_SSUPRICE_SYNC":                          1004,
		"TYPE_ORDER_ORDER_CREATE":                           1005,
		"TYPE_ORDER_ORDER_EXPORT":                           1006,
		"TYPE_ORDER_ORDERDRIVER_UPDATE":                     1007,
		"TYPE_ORDER_ORDERDRIVER_AUTOUPDATE":                 1008,
		"TYPE_ORDER_ORDER_SUBMIT_PURCHASE_TASK":             1009,
		"TYPE_ORDER_SCAN_RECEIPT":                           1010,
		"TYPE_ORDER_COMBINESSU_SYNC":                        1011,
		"TYPE_ORDER_ORDERDETAILSSU_UPDATE":                  1012,
		"TYPE_ORDER_SORTING_COMPANY_SUMMARY_EXPORT":         1013,
		"TYPE_ORDER_ORDER_BATCH_CREATE_BY_MENU":             1014,
		"TYPE_PRODUCTION_TASK_CREATE":                       1100,
		"TYPE_PRODUCTION_PURCHASE_SKU_EXPORT":               1101,
		"TYPE_PRODUCTION_TASK_EXPORT":                       1102,
		"TYPE_PRODUCTION_BOM_EXPORT":                        1103,
		"TYPE_PRODUCTION_TASK_DATA_EXPORT":                  1104,
		"TYPE_PRODUCTION_BOM_IMPORT":                        1105,
		"TYPE_PRODUCTION_PROCESS_TEMPLATE_EXPORT":           1106,
		"TYPE_PRODUCTION_TASK_RELEASE":                      1107,
		"TYPE_PRODUCTION_PROCESS_TASK_EXPORT":               1108,
		"TYPE_MERCHANDISE_SKU_EXPORT":                       1200,
		"TYPE_MERCHANDISE_SSU_EXPORT":                       1201,
		"TYPE_MERCHANDISE_SKU_SSU_UPDATE":                   1202,
		"TYPE_MERCHANDISE_CREATE_ALL":                       1203,
		"TYPE_MERCHANDISE_CREATE_SSU":                       1204,
		"TYPE_MERCHANDISE_MENU_DETAIL_EXPORT":               1205,
		"TYPE_MERCHANDISE_MENU_DETAIL_IMPORT":               1206,
		"TYPE_MERCHANDISE_COMBINE_SSU_IMPORT":               1207,
		"TYPE_MERCHANDISE_COMBINE_SSU_EXPORT":               1208,
		"TYPE_MERCHANDISE_COMBINE_SSU_UPDATE":               1209,
		"TYPE_MERCHANDISE_QUOTATION_EXPORT":                 1210,
		"TYPE_MERCHANDISE_QUOTATION_EXPORT_BY_SSU":          1211,
		"TYPE_MERCHANDISE_QUOTATION_IMPORT":                 1212,
		"TYPE_MERCHANDISE_QUOTATION_IMPORT_CREATE_BY_SSU":   1213,
		"TYPE_MERCHANDISE_QUOTATION_IMPORT_UPDATE_BY_SSU":   1214,
		"TYPE_ENTERPRISE_CUSTOMER_EXPORT":                   1300,
		"TYPE_ENTERPRISE_CUSTOMER_IMPORT":                   1301,
		"TYPE_ENTERPRISE_SUPPLIER_EXPORT":                   1302,
		"TYPE_ENTERPRISE_SUPPLIER_IMPORT":                   1303,
		"TYPE_INVENTORY_STOCK_SHEET_LIST_EXPORT":            1400,
		"TYPE_INVENTORY_BATCH_LOG_EXPORT":                   1402,
		"TYPE_INVENTORY_COMMIT_ADJUST_SHEET":                1403,
		"TYPE_INVENTORY_BATCH_SYNC_STOCK_SHEET":             1404,
		"TYPE_INVENTORY_CHANGE_LOG_EXPORT":                  1410,
		"TYPE_INVENTORY_PURCHASE_IN_LOG_EXPORT":             1411,
		"TYPE_INVENTORY_SALE_OUT_LOG_EXPORT":                1412,
		"TYPE_INVENTORY_MATERIAL_OUT_LOG_EXPORT":            1413,
		"TYPE_INVENTORY_MATERIAL_IN_LOG_EXPORT":             1414,
		"TYPE_INVENTORY_PRODUCT_IN_LOG_EXPORT":              1415,
		"TYPE_INVENTORY_REFOUND_OUT_LOG_EXPORT":             1416,
		"TYPE_INVENTORY_OTHER_OUT_LOG_EXPORT":               1417,
		"TYPE_INVENTORY_OTHER_IN_LOG_EXPORT":                1418,
		"TYPE_INVENTORY_INCREASE_IN_LOG_EXPORT":             1419,
		"TYPE_INVENTORY_LOSS_OUT_LOG_EXPORT":                1420,
		"TYPE_INVENTORY_CUSTOMER_TURNOVER_EXPORT":           1421,
		"TYPE_INVENTORY_LOAN_EXPORT":                        1422,
		"TYPE_INVENTORY_REVERT_EXPORT":                      1423,
		"TYPE_INVENTORY_CUSTOMER_TURNOVER_LOG_EXPORT":       1424,
		"TYPE_INVENTORY_REFOUND_IN_LOG_EXPORT":              1425,
		"TYPE_INVENTORY_TRANSFER_IN_LOG_EXPORT":             1426,
		"TYPE_INVENTORY_TRANSFER_OUT_LOG_EXPORT":            1427,
		"TYPE_INVENTORY_SSU_STOCK_VALUE_EXPORT":             1430,
		"TYPE_INVENTORY_SHEET_EXPORT":                       1450,
		"TYPE_INVENTORY_BATCH_UPDATE_STOCK_SHEET":           1498,
		"TYPE_INVENTORY_SKU_STOCK_EXPORT":                   1499,
		"TYPE_SORTING_SORTING_PERFORMANCE_EXPORT":           1500,
		"TYPE_PURCHASE_PURCHASE_ORDER_EXPORT":               1601,
		"TYPE_PURCHASE_INQUIRY_PRICE_IMPORT":                1602,
		"TYPE_PURCHASE_PURCHASE_TASK_RELEASE":               1603,
		"TYPE_PURCHASE_PURCHASE_TASK_SWITCH_SUPPLIER":       1604,
		"TYPE_PURCHASE_PURCHASE_TASK_SWITCH_PURCHASER":      1605,
		"TYPE_PURCHASE_PURCHASE_TASK_CREATE_PURCHASE_ORDER": 1606,
		"TYPE_PURCHASE_PURCHASE_TASK_EXPORT":                1607,
		"TYPE_PURCHASE_INQUIRY_PRICE_IMPORT_TMPL":           1608,
		"TYPE_PURCHASE_PURCHASE_TASK_EXPORT_DIMENSION":      1609,
		"TYPE_PURCHASE_PURCHASE_TASK_BATCH_REDUCE_STOCK":    1610,
		"TYPE_ESHOP_ESHOP_ORDER_EXPORT":                     1706,
		"TYPE_ESHOP_ESHOP_ORDER_UPDATE":                     1707,
		"TYPE_ESHOP_LEAVEAPPLICATION_EXPORT":                1801,
		"TYPE_AFTER_SALE_EXPORT_ORDER":                      1900,
		"TYPE_AFTER_SALE_EXPORT_TASK":                       1901,
		"TYPE_AFTER_SALE_EXPORT_ORDER_SHEET":                1902,
		"TYPE_AFTER_SALE_EXPORT_ORDER_DETAIL":               1903,
		"TYPE_INITIALIZATION_CLEAN_GROUP_TICKET":            2000,
		"TYPE_INITIALIZATION_CLEAN_GROUP_DATA":              2001,
		"TYPE_FINANCE_SETTLE_EXPORT":                        2100,
		"TYPE_FINANCE_BATCH_PAY_SETTLE":                     2101,
		"TYPE_FINANCE_FLOW_EXPORT":                          2102,
		"TYPE_FINANCE_ACCOUNT_BALANCE_EXPORT":               2103,
		"TYPE_FINANCE_BATCH_SUBMIT_SETTLE_SHEET":            2104,
		"TYPE_FINANCE_MERCHANT_SETTLEMENT_EXPORT":           2105,
	}
)

func (x Task_Type) Enum() *Task_Type {
	p := new(Task_Type)
	*p = x
	return p
}

func (x Task_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_asynctask_proto_asynctask_proto_enumTypes[2].Descriptor()
}

func (Task_Type) Type() protoreflect.EnumType {
	return &file_asynctask_proto_asynctask_proto_enumTypes[2]
}

func (x Task_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_Type.Descriptor instead.
func (Task_Type) EnumDescriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{2, 1}
}

type Task_State int32

const (
	Task_STATE_UNSPECIFIED Task_State = 0
	Task_STATE_CREATED     Task_State = 1 // 已创建，开始状态
	Task_STATE_READY       Task_State = 2 // 已就绪，等待被第一次调度Execute
	Task_STATE_RUNNING     Task_State = 3 // 执行中，循环调用Execute
	Task_STATE_CANCELED    Task_State = 4 // 已取消，结束状态
	Task_STATE_FAULTED     Task_State = 5 // 已失败，结束状态
	Task_STATE_COMPLETED   Task_State = 6 // 已完成，结束状态
)

// Enum value maps for Task_State.
var (
	Task_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_CREATED",
		2: "STATE_READY",
		3: "STATE_RUNNING",
		4: "STATE_CANCELED",
		5: "STATE_FAULTED",
		6: "STATE_COMPLETED",
	}
	Task_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STATE_CREATED":     1,
		"STATE_READY":       2,
		"STATE_RUNNING":     3,
		"STATE_CANCELED":    4,
		"STATE_FAULTED":     5,
		"STATE_COMPLETED":   6,
	}
)

func (x Task_State) Enum() *Task_State {
	p := new(Task_State)
	*p = x
	return p
}

func (x Task_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_State) Descriptor() protoreflect.EnumDescriptor {
	return file_asynctask_proto_asynctask_proto_enumTypes[3].Descriptor()
}

func (Task_State) Type() protoreflect.EnumType {
	return &file_asynctask_proto_asynctask_proto_enumTypes[3]
}

func (x Task_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_State.Descriptor instead.
func (Task_State) EnumDescriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{2, 2}
}

// gRPC status code
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{0}
}

// 任务委托
type TaskDelegate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                Task_Type `protobuf:"varint,1,opt,name=type,proto3,enum=ceres.asynctask.Task_Type" json:"type,omitempty"`
	Endpoint            string    `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`                                                       // 代理节点
	MaxConcurrentTask   uint32    `protobuf:"varint,3,opt,name=max_concurrent_task,json=maxConcurrentTask,proto3" json:"max_concurrent_task,omitempty"`         // 任务并发数
	TaskMaxSchduleCount uint32    `protobuf:"varint,4,opt,name=task_max_schdule_count,json=taskMaxSchduleCount,proto3" json:"task_max_schdule_count,omitempty"` // 任务最大调度次数
	TaskMaxFailureCount uint32    `protobuf:"varint,5,opt,name=task_max_failure_count,json=taskMaxFailureCount,proto3" json:"task_max_failure_count,omitempty"` // 任务最大失败次数
}

func (x *TaskDelegate) Reset() {
	*x = TaskDelegate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskDelegate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskDelegate) ProtoMessage() {}

func (x *TaskDelegate) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskDelegate.ProtoReflect.Descriptor instead.
func (*TaskDelegate) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{1}
}

func (x *TaskDelegate) GetType() Task_Type {
	if x != nil {
		return x.Type
	}
	return Task_TYPE_UNSPECIFIED
}

func (x *TaskDelegate) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *TaskDelegate) GetMaxConcurrentTask() uint32 {
	if x != nil {
		return x.MaxConcurrentTask
	}
	return 0
}

func (x *TaskDelegate) GetTaskMaxSchduleCount() uint32 {
	if x != nil {
		return x.TaskMaxSchduleCount
	}
	return 0
}

func (x *TaskDelegate) GetTaskMaxFailureCount() uint32 {
	if x != nil {
		return x.TaskMaxFailureCount
	}
	return 0
}

// 任务
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"PRIMARY_KEY"
	TaskId     uint64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty" gorm:"PRIMARY_KEY"`
	Status     uint64 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Revision   uint64 `protobuf:"varint,3,opt,name=revision,proto3" json:"revision,omitempty"`
	CreateTime uint64 `protobuf:"varint,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime uint64 `protobuf:"varint,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"` // update_time - create_time即为任务调度时间（包含等待时间）
	DeleteTime uint64 `protobuf:"varint,6,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// @inject_tag: gorm:"INDEX:group_id"
	GroupId   uint64        `protobuf:"varint,10,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty" gorm:"INDEX:group_id"`
	CreatorId uint64        `protobuf:"varint,11,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`                 // 任务创建者
	Category  Task_Category `protobuf:"varint,12,opt,name=category,proto3,enum=ceres.asynctask.Task_Category" json:"category,omitempty"` // 任务类型
	// @inject_tag: gorm:"INDEX:type_state"
	Type        Task_Type `protobuf:"varint,20,opt,name=type,proto3,enum=ceres.asynctask.Task_Type" json:"type,omitempty" gorm:"INDEX:type_state"` // 任务类型
	Priority    uint32    `protobuf:"varint,21,opt,name=priority,proto3" json:"priority,omitempty"`                                                // 优先级，越大越高
	SingletonId uint64    `protobuf:"varint,22,opt,name=singleton_id,json=singletonId,proto3" json:"singleton_id,omitempty"`                       //单例模式，非0时，此Type下具有相同singleton_id的任务只能并发调度一个.譬如设置singleton_id为group_id可实现group唯一；搭配max_concurrent_task = 1，并指定singleton_id为固定值，可以实现全局唯一
	// @inject_tag: gorm:"type:mediumblob"
	Data     []byte           `protobuf:"bytes,23,opt,name=data,proto3" json:"data,omitempty" gorm:"type:mediumblob"`  // 任务创建参数
	Name     string           `protobuf:"bytes,24,opt,name=name,proto3" json:"name,omitempty"`                         // 任务名字，覆盖默认类型名字；业务自行处理i18n问题
	UserInfo *proto1.UserInfo `protobuf:"bytes,25,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"` // 任务创建时的用户信息，asynctask.从ctx中获取并存储 TODO 清理过期数据
	// 调度数据，调度器负责维护
	// @inject_tag: gorm:"INDEX:type_state"
	State            Task_State `protobuf:"varint,40,opt,name=state,proto3,enum=ceres.asynctask.Task_State" json:"state,omitempty" gorm:"INDEX:type_state"` // 任务状态
	IsBusy           bool       `protobuf:"varint,41,opt,name=is_busy,json=isBusy,proto3" json:"is_busy,omitempty"`                                         // 是否正在被调度
	LastScheduleTime uint64     `protobuf:"varint,42,opt,name=last_schedule_time,json=lastScheduleTime,proto3" json:"last_schedule_time,omitempty"`         // 上次被调度时间
	NextScheduleTime uint64     `protobuf:"varint,43,opt,name=next_schedule_time,json=nextScheduleTime,proto3" json:"next_schedule_time,omitempty"`         // 下次可调度时间，调度失败的任务会内置一个调度间隔
	ScheduleCount    uint32     `protobuf:"varint,44,opt,name=schedule_count,json=scheduleCount,proto3" json:"schedule_count,omitempty"`                    // 调度次数，task.schdule_count >= task_delegate.task_max_schdule_count则调度器将任务标记为“已失败”
	FailureCount     uint32     `protobuf:"varint,45,opt,name=failure_count,json=failureCount,proto3" json:"failure_count,omitempty"`                       // 失败次数，task.failure_count >= task_delegate.task_max_failure_count则调度器将任务标记为“已失败”
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *Task) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Task) GetRevision() uint64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *Task) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Task) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *Task) GetDeleteTime() uint64 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}

func (x *Task) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *Task) GetCreatorId() uint64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Task) GetCategory() Task_Category {
	if x != nil {
		return x.Category
	}
	return Task_CATEGORY_UNSPECIFIED
}

func (x *Task) GetType() Task_Type {
	if x != nil {
		return x.Type
	}
	return Task_TYPE_UNSPECIFIED
}

func (x *Task) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Task) GetSingletonId() uint64 {
	if x != nil {
		return x.SingletonId
	}
	return 0
}

func (x *Task) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetUserInfo() *proto1.UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *Task) GetState() Task_State {
	if x != nil {
		return x.State
	}
	return Task_STATE_UNSPECIFIED
}

func (x *Task) GetIsBusy() bool {
	if x != nil {
		return x.IsBusy
	}
	return false
}

func (x *Task) GetLastScheduleTime() uint64 {
	if x != nil {
		return x.LastScheduleTime
	}
	return 0
}

func (x *Task) GetNextScheduleTime() uint64 {
	if x != nil {
		return x.NextScheduleTime
	}
	return 0
}

func (x *Task) GetScheduleCount() uint32 {
	if x != nil {
		return x.ScheduleCount
	}
	return 0
}

func (x *Task) GetFailureCount() uint32 {
	if x != nil {
		return x.FailureCount
	}
	return 0
}

// 任务数据，Delegate自行维护
type TaskData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"PRIMARY_KEY"
	TaskDataId   uint64          `protobuf:"varint,1,opt,name=task_data_id,json=taskDataId,proto3" json:"task_data_id,omitempty" gorm:"PRIMARY_KEY"` // task_id
	Total        uint32          `protobuf:"varint,20,opt,name=total,proto3" json:"total,omitempty"`                                                 // 进度最大值，Prepare填写
	Progress     uint32          `protobuf:"varint,21,opt,name=progress,proto3" json:"progress,omitempty"`                                           // 当前进度，取值[0, total]，Execute增加
	SuccessCount *TaskData_Count `protobuf:"bytes,22,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`                // 成功计数
	FailureCount *TaskData_Count `protobuf:"bytes,23,opt,name=failure_count,json=failureCount,proto3" json:"failure_count,omitempty"`                // 失败计数
	// @inject_tag: gorm:"type:mediumblob"
	Data             []byte            `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty" gorm:"type:mediumblob"`                            // 任务负载数据，一般是proto序列化的二进制 TODO 清理过期数据
	Parameters       *proto2.StringSet `protobuf:"bytes,101,opt,name=parameters,proto3" json:"parameters,omitempty"`                                       // 任务结果参数，查看任务结果时无法解析data，可把关键结果存储到这里
	SuccessAttachUrl string            `protobuf:"bytes,102,opt,name=success_attach_url,json=successAttachUrl,proto3" json:"success_attach_url,omitempty"` // 任务成功下载附件url
	FailureAttachUrl string            `protobuf:"bytes,103,opt,name=failure_attach_url,json=failureAttachUrl,proto3" json:"failure_attach_url,omitempty"` // 任务失败下载附件url；不排除和success_attach_url同时存在的可能性，都定义上
}

func (x *TaskData) Reset() {
	*x = TaskData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskData) ProtoMessage() {}

func (x *TaskData) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskData.ProtoReflect.Descriptor instead.
func (*TaskData) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{3}
}

func (x *TaskData) GetTaskDataId() uint64 {
	if x != nil {
		return x.TaskDataId
	}
	return 0
}

func (x *TaskData) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TaskData) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *TaskData) GetSuccessCount() *TaskData_Count {
	if x != nil {
		return x.SuccessCount
	}
	return nil
}

func (x *TaskData) GetFailureCount() *TaskData_Count {
	if x != nil {
		return x.FailureCount
	}
	return nil
}

func (x *TaskData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TaskData) GetParameters() *proto2.StringSet {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *TaskData) GetSuccessAttachUrl() string {
	if x != nil {
		return x.SuccessAttachUrl
	}
	return ""
}

func (x *TaskData) GetFailureAttachUrl() string {
	if x != nil {
		return x.FailureAttachUrl
	}
	return ""
}

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTaskRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task *Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{8}
}

func (x *GetTaskRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type GetTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task     *Task     `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	TaskData *TaskData `protobuf:"bytes,2,opt,name=task_data,json=taskData,proto3" json:"task_data,omitempty"`
}

func (x *GetTaskResponse) Reset() {
	*x = GetTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskResponse) ProtoMessage() {}

func (x *GetTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskResponse.ProtoReflect.Descriptor instead.
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{9}
}

func (x *GetTaskResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *GetTaskResponse) GetTaskData() *TaskData {
	if x != nil {
		return x.TaskData
	}
	return nil
}

type ListTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category    Task_Category `protobuf:"varint,1,opt,name=category,proto3,enum=ceres.asynctask.Task_Category" json:"category,omitempty"`
	Type        Task_Type     `protobuf:"varint,2,opt,name=type,proto3,enum=ceres.asynctask.Task_Type" json:"type,omitempty"`
	States      []Task_State  `protobuf:"varint,3,rep,packed,name=states,proto3,enum=ceres.asynctask.Task_State" json:"states,omitempty"`
	SingletonId uint64        `protobuf:"varint,4,opt,name=singleton_id,json=singletonId,proto3" json:"singleton_id,omitempty"`
}

func (x *ListTaskRequest) Reset() {
	*x = ListTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskRequest) ProtoMessage() {}

func (x *ListTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskRequest.ProtoReflect.Descriptor instead.
func (*ListTaskRequest) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{10}
}

func (x *ListTaskRequest) GetCategory() Task_Category {
	if x != nil {
		return x.Category
	}
	return Task_CATEGORY_UNSPECIFIED
}

func (x *ListTaskRequest) GetType() Task_Type {
	if x != nil {
		return x.Type
	}
	return Task_TYPE_UNSPECIFIED
}

func (x *ListTaskRequest) GetStates() []Task_State {
	if x != nil {
		return x.States
	}
	return nil
}

func (x *ListTaskRequest) GetSingletonId() uint64 {
	if x != nil {
		return x.SingletonId
	}
	return 0
}

type ListTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks     []*Task              `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	TaskDatas map[uint64]*TaskData `protobuf:"bytes,2,rep,name=task_datas,json=taskDatas,proto3" json:"task_datas,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListTaskResponse) Reset() {
	*x = ListTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskResponse) ProtoMessage() {}

func (x *ListTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskResponse.ProtoReflect.Descriptor instead.
func (*ListTaskResponse) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{11}
}

func (x *ListTaskResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *ListTaskResponse) GetTaskDatas() map[uint64]*TaskData {
	if x != nil {
		return x.TaskDatas
	}
	return nil
}

type RegisterTaskDelegateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskDelegate *TaskDelegate `protobuf:"bytes,1,opt,name=task_delegate,json=taskDelegate,proto3" json:"task_delegate,omitempty"`
}

func (x *RegisterTaskDelegateRequest) Reset() {
	*x = RegisterTaskDelegateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterTaskDelegateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterTaskDelegateRequest) ProtoMessage() {}

func (x *RegisterTaskDelegateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterTaskDelegateRequest.ProtoReflect.Descriptor instead.
func (*RegisterTaskDelegateRequest) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{12}
}

func (x *RegisterTaskDelegateRequest) GetTaskDelegate() *TaskDelegate {
	if x != nil {
		return x.TaskDelegate
	}
	return nil
}

type RegisterTaskDelegateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterTaskDelegateResponse) Reset() {
	*x = RegisterTaskDelegateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterTaskDelegateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterTaskDelegateResponse) ProtoMessage() {}

func (x *RegisterTaskDelegateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterTaskDelegateResponse.ProtoReflect.Descriptor instead.
func (*RegisterTaskDelegateResponse) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{13}
}

type ScheduleTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint64 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *ScheduleTaskRequest) Reset() {
	*x = ScheduleTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleTaskRequest) ProtoMessage() {}

func (x *ScheduleTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleTaskRequest.ProtoReflect.Descriptor instead.
func (*ScheduleTaskRequest) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{14}
}

func (x *ScheduleTaskRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type ScheduleTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ScheduleTaskResponse) Reset() {
	*x = ScheduleTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleTaskResponse) ProtoMessage() {}

func (x *ScheduleTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleTaskResponse.ProtoReflect.Descriptor instead.
func (*ScheduleTaskResponse) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{15}
}

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FooId   uint64 `protobuf:"varint,1,opt,name=foo_id,json=fooId,proto3" json:"foo_id,omitempty"`
	FooName string `protobuf:"bytes,2,opt,name=foo_name,json=fooName,proto3" json:"foo_name,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{16}
}

func (x *Foo) GetFooId() uint64 {
	if x != nil {
		return x.FooId
	}
	return 0
}

func (x *Foo) GetFooName() string {
	if x != nil {
		return x.FooName
	}
	return ""
}

type FooImportData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foos []*Foo `protobuf:"bytes,1,rep,name=foos,proto3" json:"foos,omitempty"`
}

func (x *FooImportData) Reset() {
	*x = FooImportData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooImportData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooImportData) ProtoMessage() {}

func (x *FooImportData) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooImportData.ProtoReflect.Descriptor instead.
func (*FooImportData) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{17}
}

func (x *FooImportData) GetFoos() []*Foo {
	if x != nil {
		return x.Foos
	}
	return nil
}

type FooExportData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foos []*Foo `protobuf:"bytes,1,rep,name=foos,proto3" json:"foos,omitempty"`
}

func (x *FooExportData) Reset() {
	*x = FooExportData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooExportData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooExportData) ProtoMessage() {}

func (x *FooExportData) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooExportData.ProtoReflect.Descriptor instead.
func (*FooExportData) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{18}
}

func (x *FooExportData) GetFoos() []*Foo {
	if x != nil {
		return x.Foos
	}
	return nil
}

// 任务计数
type TaskData_Count struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    uint32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`        // 总数
	Created  uint32 `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`    // 新增操作的
	Updated  uint32 `protobuf:"varint,3,opt,name=updated,proto3" json:"updated,omitempty"`    // 更新操作的
	Deleted  uint32 `protobuf:"varint,4,opt,name=deleted,proto3" json:"deleted,omitempty"`    // 删除操作的
	Reserve1 uint32 `protobuf:"varint,10,opt,name=reserve1,proto3" json:"reserve1,omitempty"` // 任务自定义的计数1
	Reserve2 uint32 `protobuf:"varint,11,opt,name=reserve2,proto3" json:"reserve2,omitempty"` // 任务自定义的计数2
}

func (x *TaskData_Count) Reset() {
	*x = TaskData_Count{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asynctask_proto_asynctask_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskData_Count) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskData_Count) ProtoMessage() {}

func (x *TaskData_Count) ProtoReflect() protoreflect.Message {
	mi := &file_asynctask_proto_asynctask_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskData_Count.ProtoReflect.Descriptor instead.
func (*TaskData_Count) Descriptor() ([]byte, []int) {
	return file_asynctask_proto_asynctask_proto_rawDescGZIP(), []int{3, 0}
}

func (x *TaskData_Count) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TaskData_Count) GetCreated() uint32 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *TaskData_Count) GetUpdated() uint32 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *TaskData_Count) GetDeleted() uint32 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

func (x *TaskData_Count) GetReserve1() uint32 {
	if x != nil {
		return x.Reserve1
	}
	return 0
}

func (x *TaskData_Count) GetReserve2() uint32 {
	if x != nil {
		return x.Reserve2
	}
	return 0
}
