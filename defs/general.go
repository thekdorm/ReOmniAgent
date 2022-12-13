package defs

type essences struct {
	AwakenedAir   Commodity
	AwakenedDecay Commodity
	AwakenedEarth Commodity
	AwakenedFire  Commodity
	AwakenedFrost Commodity
	// AwakenedIre   Commodity
	// AwakenedOrder Commodity

	RousingAir   Commodity
	RousingDecay Commodity
	RousingEarth Commodity
	RousingFire  Commodity
	RousingFrost Commodity
	RousingIre   Commodity
	RousingOrder Commodity
}

var Essences = essences{
	AwakenedAir: Commodity{
		id: 190327,
	},
	AwakenedDecay: Commodity{
		id: 190331,
	},
	AwakenedEarth: Commodity{
		id: 190316,
	},
	AwakenedFire: Commodity{
		id: 190321,
	},
	AwakenedFrost: Commodity{
		id: 190329,
	},
	// AwakenedIre: Commodity{
	// 	id: ,
	// },
	// AwakenedOrder: Commodity{
	// 	id: ,
	// },

	RousingAir: Commodity{
		id: 190326,
	},
	RousingDecay: Commodity{
		id: 190330,
	},
	RousingEarth: Commodity{
		id: 190315,
	},
	RousingFire: Commodity{
		id: 190320,
	},
	RousingFrost: Commodity{
		id: 190328,
	},
	RousingIre: Commodity{
		id: 190451,
	},
	RousingOrder: Commodity{
		id: 190322,
	},
}
