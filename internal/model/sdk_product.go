package model

type ProductCreateInput struct {
	Name     string
	Password string
	Email    string
}

// ContentGetListOutput 查询列表结果
//type ProductOutput struct {
//Id               uint
//Name             string
//Money            string
//SourceServersNum string
//Bandwidth        string
//DailyActivate    string

//List []ProductOutputItem `json:"list" description:"列表"`
//Page  int                 `json:"page" description:"分页码"`
//Size  int                 `json:"size" description:"分页数量"`
//Total int                 `json:"total" description:"数据总数"`
//}

//type ProductOutput struct {
//	Id               uint
//	Name             string
//	Money            string
//	SourceServersNum string
//	Bandwidth        string
//	DailyActivate    string

//Id               *Id               `json:"product"`
//Name             *Name             `json:"product"`
//Money            *Money            `json:"product"`
//SourceServersNum *SourceServersNum `json:"product"`
//Bandwidth        *Bandwidth        `json:"product"`
//DailyActivate    *DailyActivate    `json:"product"`
//}
