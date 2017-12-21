package rcswitch

import "time"

var protocols = []Protocol{
	{
		0,
		350 * time.Microsecond,
		Bit{1, 31},
		Bit{1, 3},
		Bit{3, 1},
		false,
	},
	{
		1,
		650 * time.Microsecond,
		Bit{1, 10},
		Bit{1, 2},
		Bit{2, 1},
		false,
	},
	{
		2,
		100 * time.Microsecond,
		Bit{30, 71},
		Bit{4, 11},
		Bit{9, 6},
		false,
	},
	{
		3,
		380 * time.Microsecond,
		Bit{1, 6},
		Bit{1, 3},
		Bit{3, 1},
		false,
	},
	{
		4,
		500 * time.Microsecond,
		Bit{6, 14},
		Bit{1, 2},
		Bit{2, 1},
		false,
	},
	{ // (HT6P20B)
		5,
		450 * time.Microsecond,
		Bit{23, 1},
		Bit{1, 2},
		Bit{2, 1},
		true,
	},
	{ // (HS2303-PT, i.e. used in AUKEY Remote)
		6,
		150 * time.Microsecond,
		Bit{2, 62},
		Bit{1, 6},
		Bit{6, 1},
		false,
	},
}
