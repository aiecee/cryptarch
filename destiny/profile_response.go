package destiny

import "github.com/mcollinge/cryptarch/destiny/definitions"

type Profile struct {

	// Recent, refundable purchases you have made from vendors. When will you use it? Couldn't say...  COMPONENT TYPE: VendorReceipts
	VendorReceipts interface{} `json:"vendorReceipts,omitempty"`
	// The profile-level inventory of the Destiny Profile.  COMPONENT TYPE: ProfileInventories
	ProfileInventory interface{} `json:"profileInventory,omitempty"`
	// The profile-level currencies owned by the Destiny Profile.  COMPONENT TYPE: ProfileCurrencies
	ProfileCurrencies interface{} `json:"profileCurrencies,omitempty"`
	// The basic information about the Destiny Profile (formerly \"Account\").  COMPONENT TYPE: Profiles
	Profile interface{} `json:"profile,omitempty"`
	// Items available from Kiosks that are available Profile-wide (i.e. across all characters)  This component returns information about what Kiosk items are available to you on a *Profile* level. It is theoretically possible for Kiosks to have items gated by specific Character as well. If you ever have those, you will find them on the characterKiosks property.  COMPONENT TYPE: Kiosks
	ProfileKiosks interface{} `json:"profileKiosks,omitempty"`
	// Basic information about each character, keyed by the CharacterId.  COMPONENT TYPE: Characters
	Characters struct {
		Data map[string]definitions.CharacterComponent `json:"data,omitempty"`
	} `json:"characters,omitempty"`
	// The character-level non-equipped inventory items, keyed by the Character's Id.  COMPONENT TYPE: CharacterInventories
	CharacterInventories interface{} `json:"characterInventories,omitempty"`
	// Character-level progression data, keyed by the Character's Id.  COMPONENT TYPE: CharacterProgressions
	CharacterProgressions interface{} `json:"characterProgressions,omitempty"`
	// Character rendering data - a minimal set of info needed to render a character in 3D - keyed by the Character's Id.  COMPONENT TYPE: CharacterRenderData
	CharacterRenderData interface{} `json:"characterRenderData,omitempty"`
	// Character activity data - the activities available to this character and its status, keyed by the Character's Id.  COMPONENT TYPE: CharacterActivities
	CharacterActivities interface{} `json:"characterActivities,omitempty"`
	// The character's equipped items, keyed by the Character's Id.  COMPONENT TYPE: CharacterEquipment
	CharacterEquipment interface{} `json:"characterEquipment,omitempty"`
	// Items available from Kiosks that are available to a specific character as opposed to the account as a whole. It must be combined with data from the profileKiosks property to get a full picture of the character's available items to check out of a kiosk.  This component returns information about what Kiosk items are available to you on a *Character* level. Usually, kiosk items will be earned for the entire Profile (all characters) at once. To find those, look in the profileKiosks property.  COMPONENT TYPE: Kiosks
	CharacterKiosks interface{} `json:"characterKiosks,omitempty"`
	// Information about instanced items across all returned characters, keyed by the item's instance ID.  COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.]
	ItemComponents interface{} `json:"itemComponents,omitempty"`
}
