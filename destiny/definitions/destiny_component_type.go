package definitions

type DestinyComponentType int32

type DestinyComponentTypeEnum struct {
	None                  DestinyComponentType
	Profiles              DestinyComponentType
	VendorReceipts        DestinyComponentType
	ProfileInventories    DestinyComponentType
	ProfileCurrencies     DestinyComponentType
	Characters            DestinyComponentType
	CharacterInventories  DestinyComponentType
	CharacterProgressions DestinyComponentType
	CharacterRenderData   DestinyComponentType
	CharacterActivities   DestinyComponentType
	CharacterEquipment    DestinyComponentType
	ItemInstances         DestinyComponentType
	ItemObjectives        DestinyComponentType
	ItemPerks             DestinyComponentType
	ItemRenderData        DestinyComponentType
	ItemStats             DestinyComponentType
	ItemSockets           DestinyComponentType
	ItemTalentGrids       DestinyComponentType
	ItemCommonData        DestinyComponentType
	ItemPlugStates        DestinyComponentType
	Vendors               DestinyComponentType
	VendorCatergories     DestinyComponentType
	VendorSales           DestinyComponentType
	Kiosks                DestinyComponentType
}

var ComponentType = DestinyComponentTypeEnum{
	None:                  DestinyComponentType(0),
	Profiles:              DestinyComponentType(100),
	VendorReceipts:        DestinyComponentType(101),
	ProfileInventories:    DestinyComponentType(102),
	ProfileCurrencies:     DestinyComponentType(103),
	Characters:            DestinyComponentType(200),
	CharacterInventories:  DestinyComponentType(201),
	CharacterProgressions: DestinyComponentType(202),
	CharacterRenderData:   DestinyComponentType(203),
	CharacterActivities:   DestinyComponentType(204),
	CharacterEquipment:    DestinyComponentType(205),
	ItemInstances:         DestinyComponentType(300),
	ItemObjectives:        DestinyComponentType(301),
	ItemPerks:             DestinyComponentType(302),
	ItemRenderData:        DestinyComponentType(303),
	ItemStats:             DestinyComponentType(304),
	ItemSockets:           DestinyComponentType(305),
	ItemTalentGrids:       DestinyComponentType(306),
	ItemCommonData:        DestinyComponentType(307),
	ItemPlugStates:        DestinyComponentType(308),
	Vendors:               DestinyComponentType(400),
	VendorCatergories:     DestinyComponentType(401),
	VendorSales:           DestinyComponentType(402),
	Kiosks:                DestinyComponentType(500),
}
