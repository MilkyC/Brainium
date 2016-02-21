package challenges

import (
	"math/rand"
)

// MathChallenge wrapper struct.
type MathChallenge struct {
	seed   int64
	seeded bool
}

// MathProblem wrapper struct.
type MathProblem struct {
	Value1 int64
	Value2 int64
	Symbol string
}

// GetProblem retrieves a MathProblem struct with random values initialized.
func (mc *MathChallenge) GetProblem() MathProblem {
	return MathProblem{
		Value1: int64(mc.getRandomInt()),
		Value2: int64(mc.getRandomInt()),
		Symbol: "+",
	}
}

func (mc *MathChallenge) getRandomInt() int {
	if !mc.seeded {
		rand.Seed(mc.seed)
		mc.seeded = true
	}
	return rand.Intn(100)
}

// MakeMathChallenge wrapper for creating a math challenge.
func MakeMathChallenge(seed int64) MathChallenge {
	return MathChallenge{seed: seed, seeded: false}
}

// Solution retrieves the solution for a MathProblem.
func (mp *MathProblem) Solution() int64 {
	switch mp.Symbol {
	case "+":
		return mp.Value1 + mp.Value2
	case "-":
		return mp.Value1 - mp.Value2
	default:
		panic("unrecognized Symbol")
	}
}
