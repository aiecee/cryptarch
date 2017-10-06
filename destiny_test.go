package cryptarch_test

import (
	"fmt"
	"testing"

	"github.com/mcollinge/cryptarch"
	"github.com/mcollinge/cryptarch/destiny/definitions"
)

const apiKey = "apiKey"

func TestSearchPlayer(t *testing.T) {
	bungieService := cryptarch.NewBungieService(apiKey)
	destinyApi := cryptarch.NewDestiny2Api(bungieService)
	userInfoCards, err := destinyApi.SearchPlayer(definitions.MembershipType.Xbox, "AieCee")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	profile, err := destinyApi.GetProfile(definitions.MembershipType.Xbox, userInfoCards[0].MembershipId, definitions.ComponentType.Characters)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	for key, val := range profile.Characters.Data {
		fmt.Println(key)
		fmt.Println(val.CharacterId)
	}
}
