package subscription_coupon

import (
	"context"
	"fmt"
	"os"
	"shop-smart-api/internal/entity"
	"shop-smart-api/pkg"
	adminpb "shop-smart-api/proto"
	"strconv"
)

type SubscriptionCouponInterface interface {
	GetUserSubscriptionLevel(userId string) (int, error)
	GetEmailUser(userID string) (string, error)
	GetOrgId(email string) (string, error)
	GetOrgSubscriptionLevel(orgID string) (int, error)
}

type SubscriptionCoupon struct {
	client     adminpb.AdminServiceClient
	repository SubscriptionCouponInterface
	cfg        *pkg.AppConfig
}

func New(r SubscriptionCouponInterface, cfg *pkg.AppConfig, client adminpb.AdminServiceClient) SubscriptionCoupon {
	return SubscriptionCoupon{
		client:     client,
		repository: r,
		cfg:        cfg,
	}
}

func (p SubscriptionCoupon) GetSubscriptions(userId string) ([]entity.SubscriptionEntity, error) {
	ctx := context.Background()
	subs, err := p.client.GetSubsGRPC(ctx, &adminpb.Empty{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Price : ", subs.Subs[0].Price)
	//userLevel, err := p.repository.GetUserSubscriptionLevel(userId)
	//if err != nil {
	//	return nil, err
	//}
	//email, err := p.repository.GetEmailUser(userId)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println(email)
	//orgID, err := p.repository.GetOrgId(email)
	//if err != nil {
	//	return nil, err
	//}
	//
	//orgLevel, err := p.repository.GetOrgSubscriptionLevel(orgID)
	//if err != nil {
	//	return nil, err
	//}

	resultSubs := make([]entity.SubscriptionEntity, len(subs.Subs))
	for i, v := range subs.Subs {
		resultSubs[i] = entity.SubscriptionEntity{
			Name:        v.Name,
			Description: v.Description,
			Price:       float32(v.Price),
			Level:       v.Level,
		}
	}
	//var subscriptionLevel int32
	//if userLevel >= orgLevel {
	//	subscriptionLevel = int32(userLevel)
	//} else {
	//	subscriptionLevel = int32(orgLevel)
	//}
	//if subscriptionLevel == 0 {
	//	return resultSubs, nil
	//}
	//for i, v := range resultSubs {
	//	if v.Level >= subscriptionLevel {
	//		discount, err := strconv.Atoi(os.Getenv("DISCOUNT"))
	//		if err != nil {
	//			return nil, err
	//		}
	//		if discount < 0 || discount > 100 {
	//			return nil, fmt.Errorf("discount is invalid, it must be >= 0 and <= 100")
	//		}
	//		resultSubs[i].Price *= float32(100-discount) / 100
	//	}
	//}
	return resultSubs, nil
}

func (p SubscriptionCoupon) GetCoupons(userId string) ([]entity.CouponEntity, error) {
	ctx := context.Background()
	coupons, err := p.client.GetCouponsGRPC(ctx, &adminpb.Empty{})
	if err != nil {
		return nil, err
	}

	userLevel, err := p.repository.GetUserSubscriptionLevel(userId)
	if err != nil {
		return nil, err
	}

	email, err := p.repository.GetEmailUser(userId)
	if err != nil {
		return nil, err
	}

	orgID, err := p.repository.GetOrgId(email)
	if err != nil {
		return nil, err
	}

	orgLevel, err := p.repository.GetOrgSubscriptionLevel(orgID)
	if err != nil {
		return nil, err
	}

	resultCoupons := make([]entity.CouponEntity, len(coupons.Coupons))
	for i, v := range coupons.Coupons {
		resultCoupons[i] = entity.CouponEntity{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       float32(v.Price),
			Level:       v.Level,
			Percent:     v.Percent,
			ContentUrl:  v.ContentUrl,
			Media: &entity.Media{
				ID:   v.Media.ID,
				Path: v.Media.Path,
			},
		}
	}
	var subscriptionLevel int32
	if userLevel >= orgLevel {
		subscriptionLevel = int32(userLevel)
	} else {
		subscriptionLevel = int32(orgLevel)
	}
	if subscriptionLevel == 0 {
		return resultCoupons, nil
	}
	for i, v := range resultCoupons {
		if v.Level <= subscriptionLevel {
			resultCoupons[i].Price = 0
		}

		if v.Level > subscriptionLevel {
			discount, err := strconv.Atoi(os.Getenv("DISCOUNT"))
			if err != nil {
				return nil, err
			}
			if discount < 0 || discount > 100 {
				return nil, fmt.Errorf("discount is invalid, it must be >= 0 and <= 100")
			}
			resultCoupons[i].Price *= float32(100-discount) / 100
		}
	}

	return resultCoupons, nil
}