package subscription_coupon

import (
	"context"
	"database/sql"
	"errors"
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
	GetRole(email string) (string, error)
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
		return nil, fmt.Errorf("GetCouponsGRPC failed: %w", err)
	}

	userLevel, err := p.repository.GetUserSubscriptionLevel(userId)
	if err != nil {
		return nil, fmt.Errorf("GetUserSubscriptionLevel failed: %w", err)
	}

	email, err := p.repository.GetEmailUser(userId)
	if err != nil {
		return nil, fmt.Errorf("GetEmailUser failed: %w", err)
	}

	orgID, err := p.repository.GetOrgId(email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("GetOrgId failed: %w", err)
		}
	}

	orgLevel, err := p.repository.GetOrgSubscriptionLevel(orgID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("GetOrgSubscriptionLevel failed: %w", err)
		}

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

func (p SubscriptionCoupon) GetCouponsStandard() ([]entity.CouponEntity, error) {
	ctx := context.Background()
	coupons, err := p.client.GetCouponsGRPC(ctx, &adminpb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("GetCouponsGRPC failed: %w", err)
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

	return resultCoupons, nil
}

func (p SubscriptionCoupon) GetOrganizationInfo(userId string) (entity.OrganizationEntity, error) {
	ctx := context.Background()
	email, err := p.repository.GetEmailUser(userId)
	if err != nil {
		return entity.OrganizationEntity{}, err
	}
	orgID, err := p.repository.GetOrgId(email)
	if err != nil {
		return entity.OrganizationEntity{}, err
	}
	if orgID == "" {
		return entity.OrganizationEntity{}, fmt.Errorf("user is not a member of organization")
	}
	orgInfo, err := p.client.GetOrganizationInfo(ctx, &adminpb.InfoOrganizationRequest{OrgId: orgID})
	if orgInfo == nil {
		return entity.OrganizationEntity{}, fmt.Errorf("info about company not found")
	}
	var Response = entity.OrganizationEntity{
		Name:              orgInfo.Name,
		ID:                orgInfo.ID,
		EmailAdmin:        orgInfo.EmailAdmin,
		LevelSubscription: int(orgInfo.LevelSubscription),
		ORGN:              orgInfo.Orgn,
		KPP:               orgInfo.Kpp,
		INN:               orgInfo.Inn,
		Address:           orgInfo.Kpp,
		Members:           make([]entity.Member, len(orgInfo.Members)),
	}
	for i, v := range orgInfo.Members {
		Response.Members[i] = entity.Member{
			ID:             v.Id,
			Email:          v.Email,
			FirstName:      v.FirstName,
			SecondName:     v.SecondName,
			OrganizationID: v.OrgID,
			Role:           v.Role,
		}
	}
	return Response, nil
}

func (p SubscriptionCoupon) UpdateOrganizationInfo(organizationEntity entity.OrganizationEntity, role string, userID string) (string, error) {
	ctx := context.Background()
	email, err := p.repository.GetEmailUser(userID)
	if err != nil {
		return "", err
	}
	orgID, err := p.repository.GetOrgId(email)
	if err != nil {
		return "", err
	}
	if orgID == "" {
		return "", fmt.Errorf("user is not a member of organization")
	}
	response, err := p.client.UpdateOrganizationInfo(ctx, &adminpb.UpdateOrganizationRequest{
		ID:                orgID,
		Name:              organizationEntity.Name,
		EmailAdmin:        organizationEntity.EmailAdmin,
		LevelSubscription: int32(organizationEntity.LevelSubscription),
		Orgn:              organizationEntity.ORGN,
		Kpp:               organizationEntity.KPP,
		Inn:               organizationEntity.INN,
		Address:           organizationEntity.Address,
		RoleUser:          role,
	})
	if err != nil {
		return "", err
	}
	var Response = response.GetMessage()

	return Response, nil
}

func (p SubscriptionCoupon) UpdateMembersInfo(members []entity.Member, role string, id string) (string, error) {
	ctx := context.Background()
	email, err := p.repository.GetEmailUser(id)
	if err != nil {
		return "", err
	}
	orgID, err := p.repository.GetOrgId(email)
	if err != nil {
		return "", err
	}
	if orgID == "" {
		return "", fmt.Errorf("user is not a member of organization")
	}
	MembersRequest := make([]*adminpb.MemberInfo, 0)
	for i := range members {
		var MemberProto adminpb.MemberInfo
		fmt.Println("before panic")
		MemberProto.FirstName = members[i].FirstName
		MemberProto.SecondName = members[i].SecondName
		MemberProto.Email = members[i].Email
		MemberProto.Role = members[i].Role
		MembersRequest = append(MembersRequest, &MemberProto)
		fmt.Println("after panic")
	}
	fmt.Println("Members", MembersRequest)
	response, err := p.client.UpdateMembersInfo(ctx, &adminpb.UpdateMembersRequest{
		Members:        MembersRequest,
		OrganizationID: orgID,
		RoleUser:       role,
	})
	if err != nil {
		return "", err
	}
	var Response = response.GetMessage()

	return Response, nil
}

func (p SubscriptionCoupon) GetRole(email string) (string, error) {
	role, err := p.repository.GetRole(email)
	if err != nil {
		return "", err
	}
	return role, nil
}
