package service

import (
	"context"
	"crypto/rc4"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gogf/gf-demo-user/v2/internal/service/internal/do"

	v1 "github.com/gogf/gf-demo-user/v2/api/v1"

	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/os/gtime"
)

type (
	// sSdkProduct is service struct of module Product.
	sSdkUserSerial struct{}
)

var (
	// insSdkProduct is the instance of service Product.
	insSdkUserSerial = sSdkUserSerial{}
)

// Product returns the interface of Product service.
func SdkUserSerial() *sSdkUserSerial {
	return &insSdkUserSerial
}

// GetSdkList
func (s *sSdkUserSerial) GetUserSerialList(ctx context.Context, id int) (list []v1.SdkSerialListItem) {

	//当前用户
	userId := Context().Get(ctx).User.Id

	fmt.Println(userId)

	mod := dao.SdkUserSerial.Ctx(ctx)

	mod = mod.Where("sdk_user_product_id=", id).
		Where("user_id=", userId).
		Order("id asc")

	result, err := mod.All()
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return nil
	}

	//定义长度 不先定义会有越界的问题
	results := make([]v1.SdkSerialListItem, len(result))

	//数据处理
	for i, v := range result {
		//fmt.Println(v)
		host_count, _ := dao.SdkRules.Ctx(ctx).Where("user_serial_id=", v["id"]).Distinct().CountColumn("source_ip")
		results[i].HostCount = host_count

		results[i].Id = v["id"].Int()
		results[i].UserId = v["user_id"].Int()
		results[i].Number = v["number"].String()
		results[i].Name = v["remark"].String()
	}
	list = results

	return list
}

// GetKey.
func (s *sSdkUserSerial) GetKey(ctx context.Context, id int) (string, error) {

	//当前用户
	userId := Context().Get(ctx).User.Id

	mod := dao.SdkUserSerial.Ctx(ctx)

	sdk_key, err := mod.Where("id=", id).Where("user_id=", userId).Value("sdk_key")
	fmt.Println(sdk_key.String())
	if err != nil {
		err = gerror.Newf(`ErrorORM`)
		return "", err
	}

	if sdk_key.String() == "" {
		err = gerror.Newf(`游戏不存在`)
		return "", err

	}

	return sdk_key.String(), nil
}

// EditUserSerial.
func (s *sSdkUserSerial) EditUserSerial(ctx context.Context, id int, name string) error {
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

// EditUserSerial.
func (s *sSdkUserSerial) CreateUserSerial(ctx context.Context, sdk_user_product_id int, name string) error {

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

// DeleteUserSerial.
func (s *sSdkUserSerial) DeleteUserSerial(ctx context.Context, id int) error {
	//当前用户
	userId := Context().Get(ctx).User.Id

	mod := dao.SdkUserSerial.Ctx(ctx)

	info, err := mod.Where("id=", id).
		Where("user_id=", userId).
		One()
	//fmt.Println(len(info))
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
		Delete()
	if err != nil {
		return gerror.Newf(`ErrorORM`)
	}

	return nil
}

//自定义函数
func SetProductNum() string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	b := make([]rune, 2)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(52)]
	}
	return string(b)
}

//自定义函数
func GetSdk(product_num string) string {
	datas := make(map[string]interface{})

	datas["key"] = product_num
	datas["ip"] = []string{
		"127.0.0.1",
		"127.0.0.2",
	}
	datas_string, _ := json.Marshal(datas)

	//？？？？ 目前取法有问题 rc4_string := rc4(config('other.a_key'),$datas)
	// 密钥
	key := []byte("LHlmbUBXQ2LjHOc8")
	rc4_string := make([]byte, len(datas_string))
	cipher1, _ := rc4.NewCipher(key)
	cipher1.XORKeyStream(rc4_string, datas_string)

	//base64
	sdk := base64.StdEncoding.EncodeToString(rc4_string)

	return sdk
}
