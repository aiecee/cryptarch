package destiny

import "github.com/mcollinge/cryptarch/destiny/definitions"

type UserInfoCard struct {

	// A platform specific additional display name - ex: psn Real Name, bnet Unique Name, etc.
	SupplementalDisplayName string `json:"supplementalDisplayName,omitempty"`
	// URL the Icon if available.
	IconPath string `json:"iconPath,omitempty"`
	// Type of the membership.
	MembershipType definitions.BungieMembershipType `json:"membershipType,omitempty"`
	// Membership ID as they user is known in the Accounts service
	MembershipId int64 `json:"membershipId,omitempty,string"`
	// Display Name the player has chosen for themselves. The display name is optional when the data type is used as input to a platform API.
	DisplayName string `json:"displayName,omitempty"`
}
