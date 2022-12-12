package defs

type herbs struct {
	BubblePoppy QualityCommodity
	Hochenblume QualityCommodity
	Saxifrage   QualityCommodity
	Writhebark  QualityCommodity
}

var Herbs = herbs{
	BubblePoppy: QualityCommodity{
		Tier1: Commodity{
			id: 191467,
		},
		Tier2: Commodity{
			id: 191468,
		},
		Tier3: Commodity{
			id: 191469,
		},
	},
	Hochenblume: QualityCommodity{
		Tier1: Commodity{
			id: 191460,
		},
		Tier2: Commodity{
			id: 191461,
		},
		Tier3: Commodity{
			id: 191462,
		},
	},
	Saxifrage: QualityCommodity{
		Tier1: Commodity{
			id: 191464,
		},
		Tier2: Commodity{
			id: 191465,
		},
		Tier3: Commodity{
			id: 191466,
		},
	},
	Writhebark: QualityCommodity{
		Tier1: Commodity{
			id: 191470,
		},
		Tier2: Commodity{
			id: 191471,
		},
		Tier3: Commodity{
			id: 191472,
		},
	},
}
