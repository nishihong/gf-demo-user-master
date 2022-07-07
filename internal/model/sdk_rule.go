package model

type GetSdkRuleListInput struct {
	UserSerialId int    `json:"user_serial_id"`
	Type         int    `json:"type"`
	Keyword      string `json:"keyword"`
	Port         string `json:"port"`
	Page         int    `json:"page"`
	Limit        int    `json:"limit"`
}

// 组合模型
//type Entity struct {
//	SdkUserProduct        *EntitySdkUserProduct
//	SdkUserProductSetting *EntitySdkUserProductSetting
//	SdkUserSerial         *SdkUserSerial
//}
