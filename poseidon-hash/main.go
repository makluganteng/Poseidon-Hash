package main

import (
	"github.com/consensys/gnark/frontend"
)

func SBox(api frontend.API, in frontend.Variable) frontend.Variable {
	inDouble := api.Mul(in, in)
	inQuadruple := api.Mul(inDouble, inDouble)
	return api.Mul(inQuadruple, in)
}

func MatrixMul(api frontend.API, state [3]frontend.Variable) [3]frontend.Variable {
	var out [3]frontend.Variable
	mds_matrix := MdsMatrix()
	for i := 0; i < 3; i++ {
		var tmp frontend.Variable = 0
		for j := 0; j < 3; j++ {
			api.Add(tmp, api.Mul(state[j], mds_matrix[i][j]))
		}
		out[i] = tmp
	}

	return out

}
