package defs

type essences struct {
	RousingAir   Commodity
	RousingDecay Commodity
	RousingEarth Commodity
	RousingFire  Commodity
	RousingFrost Commodity
	RousingIre   Commodity
	RousingOrder Commodity
}

var Essences = essences{
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
