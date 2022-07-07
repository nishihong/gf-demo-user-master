package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf-demo-user/v2/internal/service/internal/do"

	"github.com/gogf/gf-demo-user/v2/internal/model"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"

	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	// sSdkRules is service struct of module SdkRules.
	sSdkRules struct{}
)

var (
	// insSdkRules is the instance of service SdkRules.
	insSdkRules = sSdkRules{}
)

// SdkRules returns the interface of SdkRules service.
func SdkRules() *sSdkRules {
	return &insSdkRules
}

// GetSourceServersNum.
func (s *sSdkRules) GetSourceServersNum(ctx context.Context, id int) (string, error) {

	//当前用户
	userId := Context().Get(ctx).User.Id

	serialInfo, err := dao.SdkUserSerial.Ctx(ctx).Where("id=", id).Where("user_id=", userId).One()
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return "", err
	}
	if serialInfo == nil {
		err = gerror.Newf(`游戏不存在`)
		return "", err
	}

	productInfo, err := dao.SdkUserProduct.Ctx(ctx).Where("id=", serialInfo["sdk_user_product_id"]).Where("user_id=", userId).One()
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return "", err
	}
	if productInfo == nil {
		err = gerror.Newf(`产品不存在`)
		return "", err
	}

	settingInfo, err := dao.SdkUserProductSetting.Ctx(ctx).Where("user_product_id=", productInfo["id"]).One()
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return "", err
	}
	sourceServersNum := "0"
	if settingInfo["source_servers_num"] != nil {
		sourceServersNum = settingInfo["source_servers_num"].String()
	}

	if sourceServersNum == "不限" {
		return sourceServersNum, nil
	}

	info, err := g.Model("yjs_sdk_rules", "a").
		Fields("source_ip").
		LeftJoin("yjs_sdk_user_serial b", "a.user_serial_id=b.id").
		Where("a.user_product_id", productInfo["id"]).
		Group("source_ip").
		All()
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return "", err
	}

	sourceServersNumInt, _ := strconv.Atoi(sourceServersNum)
	//string转int
	infoNum := sourceServersNumInt - len(info)

	//int转string
	infoNumString := strconv.Itoa(infoNum)

	return infoNumString, nil
}

// GetSdkRuleList
func (s *sSdkRules) GetSdkRuleList(ctx context.Context, in model.GetSdkRuleListInput) (list v1.GetSdkRuleListRes, err error) {
	var (
		likeKeyword = `%` + in.Keyword + `%`
		likePort    = `%` + in.Port + `%`
	)

	//当前用户
	userId := Context().Get(ctx).User.Id

	//默认值
	if in.Page == 0 {
		in.Page = 1
	}
	if in.Limit == 0 {
		in.Limit = 10
	}

	list = v1.GetSdkRuleListRes{
		CurrentPage: in.Page,
		PerPage:     in.Limit,
	}

	serialInfo, err := dao.SdkUserSerial.Ctx(ctx).Where("id=", in.UserSerialId).Where("user_id=", userId).One()
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return list, err
	}
	if serialInfo == nil {
		err = gerror.Newf(`游戏不存在`)
		return list, err
	}

	productInfo, err := dao.SdkUserProduct.Ctx(ctx).Where("id=", serialInfo["sdk_user_product_id"]).Where("user_id=", userId).One()
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return list, err
	}
	if productInfo == nil {
		err = gerror.Newf(`产品不存在`)
		return list, err
	}

	//默认方法
	mod := dao.SdkRules.Ctx(ctx)

	//条件
	mod = mod.Where("user_serial_id", in.UserSerialId)
	if in.Type != 0 {
		mod = mod.Where("type=", in.Type)
	}
	if in.Keyword != "" {
		mod = mod.WhereLike("custom_ip", likeKeyword).WhereOrLike("source_ip", likeKeyword)
	}
	if in.Port != "" {
		mod = mod.WhereLike("port", likePort)
	}

	list.Total, err = mod.Count()
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return list, err
	}

	res, err := mod.Order("id desc").
		Page(in.Page, in.Limit).
		All()
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return list, err
	}

	//定义长度 不先定义会有越界的问题
	results := make([]v1.SdkRuleListItem, len(res))

	//数据处理
	for i, v := range res {
		//fmt.Println(i, v)

		results[i].Id = v["id"].Int()
		results[i].UserProductId = v["user_product_id"].Int()
		results[i].Type = v["type"].Int()
		results[i].CustomIp = v["custom_ip"].String()
		results[i].Port = v["name"].Int()
		results[i].SourceIp = v["source_ip"].String()
		results[i].Status = v["status"].Int()
		results[i].CreatedAt = v["created_at"].GTime()
		results[i].UpdatedAt = v["updated_at"].GTime()
		results[i].PortInterval = v["port_interval"].Int()
		results[i].PortType = v["port_type"].Int()
		results[i].PortLimit = v["port_limit"].Int()
		results[i].GetIpWay = v["get_ip_way"].Int()
		results[i].UserSerialId = v["user_serial_id"].Int()

		ruleTypeList := map[int]string{
			1: "PC端",
			2: "移动端",
			3: "PC端+移动端",
		}
		typeName := ruleTypeList[v["type"].Int()]
		results[i].TypeName = typeName
		if v["port_type"].Int() == 2 {
			results[i].PortIntervals = v["port_interval"].String()
		} else {
			results[i].PortIntervals = "--"
		}
	}
	list.Data = results

	return list, nil
}

// EditRule.
func (s *sSdkRules) EditRule(ctx context.Context, id int, name string) error {
	//当前用户
	userId := Context().Get(ctx).User.Id

	mod := dao.SdkUserSerial.Ctx(ctx)

	info, err := mod.Where("id=", id).
		Where("user_id=", userId).
		One()

	// ？？？ code返回值定义问题
	if err != nil {
		return gerror.Newf(`ErrorORM`)
		//return err
	}
	if info == nil {
		return gerror.Newf(`游戏不存在`)
	}

	_, err = mod.Where("id=", id).
		Where("user_id=", userId).
		Data(dao.SdkUserSerial.Columns().Remark, name).
		Update()
	if err != nil {
		return err
	}
	return nil
}

// CreateRule.
func (s *sSdkRules) CreateRule(ctx context.Context, sdk_user_product_id int, name string) error {

	//当前用户
	userId := Context().Get(ctx).User.Id

	name = strings.TrimSpace(name) //去除前后空格

	//判断长度
	length := len(name)
	// ？？？中文的算三个长度，未处理
	if length < 2 || length > 20 {
		// ？？？ 正常错误应该如何返回
		return gerror.Newf(`游戏名长度为2~20个字符`)
	}

	//判断sdk是否属于当前客户
	sdk_user_product_data, err := dao.SdkUserProduct.Ctx(ctx).
		Where("id=", sdk_user_product_id).
		Where("user_id=", userId).
		One()
	//fmt.Println(sdk_user_product_data)
	if err != nil {
		return gerror.Newf(`ErrorORM`)
	}
	if sdk_user_product_data == nil {
		return gerror.Newf(`SDK不存在`)
	}

	//获取配置信息
	setting_data, err := dao.SdkUserProductSetting.Ctx(ctx).
		Where("user_product_id=", sdk_user_product_id).
		One()
	var game_num = 0
	if setting_data["game_num"] != nil {
		game_num = setting_data["game_num"].Int()
	}

	//判断当前是否有超限制
	count, err := dao.SdkUserSerial.Ctx(ctx).Where("sdk_user_product_id=", sdk_user_product_id).
		Count()

	if game_num <= count {
		return gerror.Newf(`游戏授权数已超出限制，请升级套餐。`)
	}

	number := SetProductNum()

	_, err = dao.SdkUserSerial.Ctx(ctx).Data(do.SdkUserSerial{
		UserId:           userId,
		Number:           number,
		SdkUserProductId: sdk_user_product_id,
		ProductId:        sdk_user_product_data["product_id"],
		Remark:           name,
		SdkKey:           GetSdk(number),
		OnOff:            1,
		Status:           1,
		SourceType:       2,
		CreatedAt:        gtime.Now(),
		UpdatedAt:        gtime.Now(),
	}).Insert()

	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return err
	}
	return nil
}

// DeleteRule.
func (s *sSdkRules) DeleteRule(ctx context.Context, id string, userSerialId int) error {
	//当前用户
	userId := Context().Get(ctx).User.Id

	serialInfo, err := dao.SdkUserSerial.Ctx(ctx).Where("id=", userSerialId).Where("user_id=", userId).One()
	if err != nil {
		return gerror.Newf(`ErrorORM`)
	}
	if serialInfo == nil {
		return gerror.Newf(`游戏不存在`)
	}

	productInfo, err := dao.SdkUserProduct.Ctx(ctx).Where("id=", serialInfo["sdk_user_product_id"]).Where("user_id=", userId).One()
	if err != nil {
		return gerror.Newf(`ErrorORM`)
	}
	if productInfo == nil {
		return gerror.Newf(`产品不存在`)
	}

	//explode
	idArr := strings.Split(id, ",")
	fmt.Println(idArr)

	rule, err := dao.SdkRules.Ctx(ctx).
		WhereIn("id", idArr).
		Where("user_serial_id=", userSerialId).
		All()
	if err != nil {
		return err
	}
	if len(rule) != len(idArr) {
		return gerror.Newf(`规则已删除`)
	}

	_, err = dao.SdkRules.Ctx(ctx).WhereIn("id", idArr).
		Delete()
	if err != nil {
		return err
	}

	_, err = dao.SdkUserSerial.Ctx(ctx).
		Where("id=", userSerialId).
		Data(dao.SdkUserSerial.Columns().Guid, time.Now().Unix()).
		Update()
	if err != nil {
		return err
	}

	_, err = dao.SdkUserProductConfigcommon.Ctx(ctx).
		Where("user_product_id=", productInfo["id"]).
		Data(dao.SdkUserProductConfigcommon.Columns().Guid, time.Now().Unix()).
		Update()
	if err != nil {
		return err
	}

	return nil
}
