package cryptarch

import (
	"fmt"

	"github.com/mcollinge/cryptarch/destiny"

	"github.com/mcollinge/cryptarch/destiny/definitions"
)

var endpoints = map[string]string{
	"SearchDestinyPlayer": "/Destiny2/SearchDestinyPlayer/{membershipType}/{displayName}/",
	"GetProfile":          "/Destiny2/{membershipType}/Profile/{destinyMembershipId}/",
}

type Destiny2Api struct {
	bungieService *BungieService
}

func NewDestiny2Api(b *BungieService) *Destiny2Api {
	return &Destiny2Api{bungieService: b}
}

func (d Destiny2Api) SearchPlayer(membershipType definitions.BungieMembershipType, displayName string) ([]destiny.UserInfoCard, error) {
	endPoint := endpoint{
		method: "GET",
		path:   endpoints["SearchDestinyPlayer"],
		pathParameters: map[string]string{
			"membershipType": fmt.Sprintf("%v", membershipType),
			"displayName":    displayName,
		},
		queryParameters: map[string]string{},
	}
	var userInfoCard []destiny.UserInfoCard
	err := d.bungieService.call(endPoint, &userInfoCard)
	return userInfoCard, err
}

func (d Destiny2Api) GetProfile(membershipType definitions.BungieMembershipType, destinyMembershipId int64, componentTypes ...definitions.DestinyComponentType) (*destiny.Profile, error) {
	queryParams := map[string]string{}
	if len(componentTypes) != 0 {
		components := ""
		for i, component := range componentTypes {
			components += fmt.Sprintf("%v", component)
			if i != len(componentTypes)-1 {
				components += ","
			}
		}
		queryParams["components"] = components
	}
	endPoint := endpoint{
		method: "GET",
		path:   endpoints["GetProfile"],
		pathParameters: map[string]string{
			"membershipType":      fmt.Sprintf("%v", membershipType),
			"destinyMembershipId": fmt.Sprintf("%v", destinyMembershipId),
		},
		queryParameters: queryParams,
	}
	var profile destiny.Profile
	err := d.bungieService.call(endPoint, &profile)
	return &profile, err
}
