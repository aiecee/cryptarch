package definitions

import (
	"time"
)

// This component contains base properties of the character. You'll probably want to always request this component, but hey you do you.
type CharacterComponent struct {
	// Every Destiny Profile has a membershipId. This is provided on the character as well for convenience.
	MembershipId int64 `json:"membershipId,omitempty,string"`
	// membershipType tells you the platform on which the character plays. Examine the BungieMembershipType enumeration for possible values.
	MembershipType BungieMembershipType `json:"membershipType,omitempty"`
	// The unique identifier for the character.
	CharacterId int64 `json:"characterId,omitempty,string"`
	// The last date that the user played Destiny.
	DateLastPlayed time.Time `json:"dateLastPlayed,omitempty"`
	// If the user is currently playing, this is how long they've been playing.
	MinutesPlayedThisSession int64 `json:"minutesPlayedThisSession,omitempty,string"`
	// If this value is 525,600, then they played Destiny for a year. Or they're a very dedicated Rent fan. Note that this includes idle time, not just time spent actually in activities shooting things.
	MinutesPlayedTotal int64 `json:"minutesPlayedTotal,omitempty,string"`
	// The user's calculated \"Light Level\". Light level is an indicator of your power that mostly matters in the end game, once you've reached the maximum character level: it's a level that's dependent on the average Attack/Defense power of your items.
	Light int32 `json:"light,omitempty"`
	// Your character's stats, such as Agility, Resilience, etc... *not* historical stats.  You'll have to call a different endpoint for those.
	Stats map[int64]int32 `json:"stats,omitempty,string"`
	// Use this hash to look up the character's DestinyRaceDefinition.
	RaceHash uint32 `json:"raceHash,omitempty"`
	// Use this hash to look up the character's DestinyGenderDefinition.
	GenderHash uint32 `json:"genderHash,omitempty"`
	// Use this hash to look up the character's DestinyClassDefinition.
	ClassHash uint32 `json:"classHash,omitempty"`
	// A shortcut path to the user's currently equipped emblem image. If you're just showing summary info for a user, this is more convenient than examining their equipped emblem and looking up the definition.
	EmblemPath string `json:"emblemPath,omitempty"`
	// A shortcut path to the user's currently equipped emblem background image. If you're just showing summary info for a user, this is more convenient than examining their equipped emblem and looking up the definition.
	EmblemBackgroundPath string `json:"emblemBackgroundPath,omitempty"`
	// The hash of the currently equipped emblem for the user. Can be used to look up the DestinyInventoryItemDefinition.
	EmblemHash uint32 `json:"emblemHash,omitempty"`
	// The progression that indicates your character's level. Not their light level, but their character level: you know, the thing you max out a couple hours in and then ignore for the sake of light level.
	LevelProgression Progression `json:"levelProgression,omitempty"`
	// The \"base\" level of your character, not accounting for any light level.
	BaseCharacterLevel int32 `json:"baseCharacterLevel,omitempty"`
	// A number between 0 and 100, indicating the whole and fractional % remaining to get to the next character level.
	PercentToNextLevel float32 `json:"percentToNextLevel,omitempty"`
}
