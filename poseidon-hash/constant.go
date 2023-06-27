package main

import (
	"math/big"
)

func MdsMatrix() [][]*big.Int {
	matrix := make([][]*big.Int, 3)

	for i := range matrix {
		matrix[i] = make([]*big.Int, 3)
	}

	var success bool

	matrix[0][0], success = new(big.Int).SetString("92469348809186613947252340883344274339611751744959319352506666082431267346705", 10)
	if !success {
		panic("Failed to parse big int")
	}

	matrix[0][1], success = new(big.Int).SetString("100938028378191533449096235266991198229563815869344032449592738345766724371160", 10)
	if !success {
		panic("Failed to parse big int")
	}

	matrix[0][2], success = new(big.Int).SetString("77486311749148948616988559783475694076613010381924638436641318334458515006661", 10)
	if !success {
		panic("Failed to parse big int")
	}

	matrix[1][0], success = new(big.Int).SetString("110352262556914082363749654180080464794716701228558638957603951672835474954408", 10)
	if !success {
		panic("Failed to parse big int")
	}

	matrix[1][1], success = new(big.Int).SetString("27607004873684391669404739690441550149894883072418944161048725383958774443141", 10)
	if !success {
		panic("Failed to parse big int")
	}

	matrix[1][2], success = new(big.Int).SetString("29671705769502357195586268679831947082918094959101307962374709600277676341325", 10)
	if !success {
		panic("Failed to parse big int")
	}

	matrix[2][0], success = new(big.Int).SetString("77762103796341032609398578911486222569419103128091016773380377798879650228751", 10)
	if !success {
		panic("Failed to parse big int")
	}

	matrix[2][1], success = new(big.Int).SetString("1753012011204964731088925227042671869111026487299375073665493007998674391999", 10)
	if !success {
		panic("Failed to parse big int")
	}

	matrix[2][2], success = new(big.Int).SetString("70274477372358662369456035572054501601454406272695978931839980644925236550307", 10)
	if !success {
		panic("Failed to parse big int")
	}

	return matrix
}
