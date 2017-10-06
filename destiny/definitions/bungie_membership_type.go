package definitions

type BungieMembershipType int32

type BungieMembershipTypeEnum struct {
	None        BungieMembershipType
	Xbox        BungieMembershipType
	Playstation BungieMembershipType
	Blizzard    BungieMembershipType
	All         BungieMembershipType
}

var MembershipType = BungieMembershipTypeEnum{
	None:        BungieMembershipType(0),
	Xbox:        BungieMembershipType(1),
	Playstation: BungieMembershipType(2),
	Blizzard:    BungieMembershipType(4),
	All:         BungieMembershipType(-1),
}
