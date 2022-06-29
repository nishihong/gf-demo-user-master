package service

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf-demo-user/v2/internal/model/entity"

	"github.com/gogf/gf-demo-user/v2/internal/service/internal/do"
	"golang.org/x/crypto/bcrypt"

	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/frame/g"
)

type (
	// sUser is service struct of module User.
	sUser struct{}
)

var (
	// insUser is the instance of service User.
	insUser = sUser{}
)

// User returns the interface of User service.
func User() *sUser {
	return &insUser
}

// Create creates user account.
//func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {
//	// If Nickname is not specified, it then uses Passport as its default Nickname.
//	if in.Nickname == "" {
//		in.Nickname = in.Passport
//	}
//	var (
//		available bool
//	)
//	// Passport checks.
//	available, err = s.IsPassportAvailable(ctx, in.Passport)
//	if err != nil {
//		return err
//	}
//	if !available {
//		return gerror.Newf(`Passport "%s" is already token by others`, in.Passport)
//	}
//	// Nickname checks.
//	available, err = s.IsNicknameAvailable(ctx, in.Nickname)
//	if err != nil {
//		return err
//	}
//	if !available {
//		return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
//	}
//	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
//		_, err = dao.User.Ctx(ctx).Data(do.User{
//			Passport: in.Passport,
//			Password: in.Password,
//			Nickname: in.Nickname,
//		}).Insert()
//		return err
//	})
//}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sUser) IsSignedIn(ctx context.Context) bool {
	if v := Context().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sUser) IsSignedInRedis(r *ghttp.Request) bool {
	token := r.Request.Header.Get("authorization")

	//fmt.Println(token)
	//fmt.Println(153465)
	//return true

	if token == "" {
		r.Response.WriteStatus(http.StatusBadRequest, "token不能为空！")
		return false
	}

	token = strings.Replace(token, "Bearer ", "", -1)

	var (
		ctxRedis = gctx.New()
	)

	result, err := g.Redis().Do(ctxRedis, "GET", token)
	//fmt.Println(result)

	////验证token是否过期    ？？？判断方式有不对  null 和nil
	if result == nil {
		//fmt.Println("token过期11")
		r.Response.WriteStatus(http.StatusBadRequest, "token过期请重新登录1！")
		return false
	}

	type User struct {
		Id      int
		Name    string
		UserId  int
		SubAuth int
		SubInfo interface{}
		SubType int
	}
	var user_info *User
	if err = result.Struct(&user_info); err != nil {
		r.Response.WriteStatus(http.StatusBadRequest, "解析格式错误！")
		return false
	}
	//fmt.Println(user_info)
	if user_info == nil {
		//fmt.Println("token过期")
		r.Response.WriteStatus(http.StatusBadRequest, "token过期请重新登录！")
		return false
	}

	//user_info := make(map[string]interface{})
	//if err = json.Unmarshal([]byte(result), &user_info); err != nil {
	//	fmt.Println(err)
	//}

	//fmt.Println(user_info.Id)
	//fmt.Println(result)
	//fmt.Println(user_info)

	if err != nil {
		fmt.Println(err)
		r.Response.WriteStatus(http.StatusBadRequest, "redis连接错误！！！")
		return false
	}

	//保持登录状态
	_, err = g.Redis().Do(ctxRedis, "SET", token, user_info)
	if err != nil {
		r.Response.WriteStatus(http.StatusBadRequest, "redis连接错误！")
		return false
	}

	//fmt.Println(result)

	//return true

	// 判断请求权限
	// 获取是否子账号，且为只读（2）权限
	sub_type := 1
	if user_info.SubType != 0 {
		sub_type = user_info.SubType
	}
	sub_auth := 1 //1只读 2 管理
	if user_info.SubAuth != 0 {
		sub_auth = user_info.SubAuth
	}

	if sub_type == 2 && sub_auth == 1 {
		// ？？？？ other写法  in_array写法
		//if in_array(request()->route()->getActionName(), config('other.auth_sub_allow'))) {

		r.Response.WriteStatus(http.StatusBadRequest, "子账号无此操作权限！请联系管理员代为操作！")
		return false
		//}
	}

	return true
}

// SignIn creates session for given user account.
func (s *sUser) SignIn(ctx context.Context, in model.UserSignInInput) (res map[string]interface{}, err error) {

	var (
		username = in.Username
		password = in.Password
	)

	sub_user_id := -1 //数据中心跳转时，需判断是否子账号 ，<0 主账号跳转（默认）
	if in.SubUserId != 0 {
		sub_user_id = in.SubUserId
	}

	if username == "" {
		return nil, gerror.New(`用户名不能为空`)
	}
	if password == "" {
		return nil, gerror.New(`密码不能为空！`)
	}

	//调用接口
	user, err_login := YcLogin(username, password)

	fmt.Println(user)

	if err_login != nil {
		return nil, gerror.New(`接口失败`)
	}
	if user == nil {
		return nil, gerror.New(`账号或密码错误`)
	}

	fmt.Println(3333)

	//if err = Session().SetUser(ctx, user); err != nil {
	//	return err
	//}

	//写入context 随时获取
	//Context().SetUser(ctx, &model.ContextUser{
	//	Id:   user.Id,
	//	Name: user.Name,
	//})

	//结构体转数组 ？？？？
	userArray := make(map[string]interface{})
	userArray["Id"] = user.Id
	userArray["Name"] = user.Name
	userArray["UserId"] = user.UserId

	//设置传递的值，不确定这个是否放这边 session的方式
	if err := Session().SetUser(ctx, userArray); err != nil {
		return nil, gerror.New(`设置数据缓存错误`)
	}

	// token的方式
	token := SetToken(username, userArray, sub_user_id)

	result := make(map[string]interface{})
	result["AccessToken"] = token
	result["TokenType"] = "bearer"
	result["ExpiresIn"] = 3600

	res = result

	return res, nil
}

// SignOut removes the session for current signed-in user.
func (s *sUser) SignOut(ctx context.Context) error {
	return Session().RemoveUser(ctx)
}

// IsPassportAvailable checks and returns given passport is available for signing up.
//func (s *sUser) IsPassportAvailable(ctx context.Context, passport string) (bool, error) {
//	count, err := dao.User.Ctx(ctx).Where(do.User{
//		Passport: passport,
//	}).Count()
//	if err != nil {
//		return false, err
//	}
//	return count == 0, nil
//}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
//func (s *sUser) IsNicknameAvailable(ctx context.Context, nickname string) (bool, error) {
//	count, err := dao.User.Ctx(ctx).Where(do.User{
//		Nickname: nickname,
//	}).Count()
//	if err != nil {
//		return false, err
//	}
//	return count == 0, nil
//}

// GetProfile retrieves and returns current user info in session.
//func (s *sUser) GetProfile(ctx context.Context) *entity.User {
//	return Session().GetUser(ctx)
//}

//自定义函数
func YcLogin(username string, password string) (data *entity.User, err error) {
	// ？？？？ 地址没有用config
	var (
		//url  = config('other.cw_url');
		url_string = "http://121.204.251.6:1111/admin.php"
		//key  = config('other.cw_key');
		key = "BBB1EBABA18DE518AF2CA31E5ADF74F1"
	)
	url_string = url_string + "/yf_api_guo/login?name=" + url.QueryEscape(username) + "&password=" + url.QueryEscape(password) + "&key=" + key

	//fmt.Println(2)

	resp, err := http.Get(url_string)
	if err != nil {
		//fmt.Println(13212)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed, err:%v\n", err)
		//fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return nil, err
	}

	test := make(map[string]interface{})
	//解码
	err = json.Unmarshal([]byte(string(body)), &test)
	if err != nil {
		//fmt.Println(err)
		return nil, err
	}

	//_, ok := test["status"]
	//fmt.Println(ok)

	fmt.Println(test, err)
	//fmt.Println(test["data"].(map[string]interface{})["id"], err)

	// 登录失败返回   ？？？ 值判断有问题
	//if test["status"] != 1 {
	//	fmt.Println(16565)
	//	return nil, err
	//}

	//用户存更新
	info, err := g.DB().Model("yjs_user").
		Where("name=", username).
		One()
	fmt.Println(info)

	if err != nil {
		return nil, gerror.Newf(`ErrorORM`)
		//return err
	}
	//密码处理
	password_string := []byte(password)
	password_string, _ = bcrypt.GenerateFromPassword(password_string, bcrypt.MinCost)
	//fmt.Println(password_string)

	if info == nil {
		_, err = g.DB().Model("yjs_user").Data(do.User{
			Id:       test["data"].(map[string]interface{})["id"],
			UserId:   test["data"].(map[string]interface{})["id"],
			Name:     username,
			Email:    test["data"].(map[string]interface{})["email"],
			Mobile:   test["data"].(map[string]interface{})["mobile"],
			Password: password_string,
		}).Insert()
		if err != nil {
			fmt.Println(222)
			return nil, err
		}
		info, err = g.DB().Model("yjs_user").
			Where("name=", username).
			One()
		if err != nil {
			return nil, gerror.Newf(`ErrorORM`)
		}
	}

	//return nil, nil

	update_info := make(map[string]interface{})
	if info["password"].String() == "" {
		update_info["Password"] = password
	}
	update_info["mobile"] = test["data"].(map[string]interface{})["mobile"]
	update_info["email"] = test["data"].(map[string]interface{})["email"]
	update_info["login_ip"] = "127.0.0.1" //？？？？真实ip还没做

	_, err = g.DB().Model("yjs_user").
		Where("id=", info["id"]).
		Data(update_info).
		Update()
	//fmt.Println(err)
	if err != nil {
		return nil, err
	}

	// ？？？ 更新以后取最新的值 如何简写
	info, err = g.DB().Model("yjs_user").
		Where("name=", username).
		One()
	if err != nil {
		return nil, gerror.Newf(`ErrorORM`)
	}

	//结构体
	_ = info.Struct(&data)

	return data, nil
}

//自定义函数
func SetToken(username string, user map[string]interface{}, sub_user_id int) (token string) {
	//生成token
	str := MD5(MD5(strconv.FormatInt(time.Now().UnixNano()/1e6, 10)))
	//fmt.Println(strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
	token = SHA1(str + username)

	fmt.Println(token)
	fmt.Println(user)

	user["SubType"] = 1 // 1 主账号 2子账号
	user["SubAuth"] = 1 // 1 管理 2只读
	user["SubInfo"] = nil
	// 为子账号，带入子账号数据

	fmt.Println(sub_user_id)

	if sub_user_id > 0 {
		// ？？？ 引用外部接口应该用什么方式
		sub_user_info := SubAccount(user["id"].(int), sub_user_id)

		user["SubType"] = 2 // 1 主账号 2子账号
		if sub_user_info["data"].(map[string]interface{})["is_read_only"] != nil {
			user["SubAuth"] = sub_user_info["data"].(map[string]interface{})["is_read_only"] // 1 管理 2只读
		}
		if sub_user_info["data"].(map[string]interface{})["sub_user"] != nil {
			user["SubInfo"] = sub_user_info["data"].(map[string]interface{})["sub_user"]
		}
	}

	//token写入redis
	var (
		ctx = gctx.New()
	)
	// ？？？ 设置超时时间
	//userStr, _ := json.Marshal(user)
	_, err := g.Redis().Do(ctx, "SET", token, user)
	if err != nil {
		return ""
	}

	//fmt.Println(token)

	return token
}

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}
func SubAccount(user_id int, sub_user_id int) map[string]interface{} {
	url := "http://121.204.251.6:1114" + "/api/sub_account/get_role" + "?user_id=" + string(user_id) + "&sub_user_id=" + string(sub_user_id) + "&type=11"

	headers := make(map[string]string)
	headers["token"] = "yjs_InsYJd9IOjpk"
	headers["unique"] = "21a57ce0b1b0658f80c48d645a779a24"

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, strings.NewReader(string([]byte("test"))))
	if err != nil {
		return nil
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	test := make(map[string]interface{})
	//解码
	err = json.Unmarshal([]byte(string(body)), &test)
	if err != nil {
		//fmt.Println(err)
		return nil
	}

	return test
}
