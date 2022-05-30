package service

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf-demo-user/v2/internal/service/internal/dao"
	"github.com/gogf/gf-demo-user/v2/internal/service/internal/do"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	// sProduct is service struct of module Product.
	sProduct struct{}
)

var (
	// insProduct is the instance of service Product.
	insProduct = sProduct{}
)

// Product returns the interface of Product service.
func Product() *sProduct {
	return &insProduct
}

// Create creates user account.
func (s *sProduct) Create(ctx context.Context, in model.ProductCreateInput) (err error) {
	// If Nickname is not specified, it then uses Passport as its default Nickname.
	if in.Nickname == "" {
		in.Nickname = in.Passport
	}
	var (
		available bool
	)
	// Passport checks.
	available, err = s.IsPassportAvailable(ctx, in.Passport)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Passport "%s" is already token by others`, in.Passport)
	}
	// Nickname checks.
	available, err = s.IsNicknameAvailable(ctx, in.Nickname)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	}
	return dao.Product.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Product.Ctx(ctx).Data(do.Product{
			Passport: in.Passport,
			Password: in.Password,
			Nickname: in.Nickname,
		}).Insert()
		return err
	})
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sProduct) IsSignedIn(ctx context.Context) bool {
	if v := Context().Get(ctx); v != nil && v.Product != nil {
		return true
	}
	return false
}

// SignIn creates session for given user account.
func (s *sProduct) SignIn(ctx context.Context, in model.ProductSignInInput) (err error) {
	var user *entity.Product
	err = dao.Product.Ctx(ctx).Where(do.Product{
		Passport: in.Passport,
		Password: in.Password,
	}).Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New(`Passport or Password not correct`)
	}
	if err = Session().SetProduct(ctx, user); err != nil {
		return err
	}
	Context().SetProduct(ctx, &model.ContextProduct{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
	})
	return nil
}

// SignOut removes the session for current signed-in user.
func (s *sProduct) SignOut(ctx context.Context) error {
	return Session().RemoveProduct(ctx)
}

// IsPassportAvailable checks and returns given passport is available for signing up.
func (s *sProduct) IsPassportAvailable(ctx context.Context, passport string) (bool, error) {
	count, err := dao.Product.Ctx(ctx).Where(do.Product{
		Passport: passport,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
func (s *sProduct) IsNicknameAvailable(ctx context.Context, nickname string) (bool, error) {
	count, err := dao.Product.Ctx(ctx).Where(do.Product{
		Nickname: nickname,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetProfile retrieves and returns current user info in session.
func (s *sProduct) GetProfile(ctx context.Context) *entity.Product {
	return Session().GetProduct(ctx)
}
