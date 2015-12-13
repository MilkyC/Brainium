package challenges

import (
    "math/rand"
)

type math_challenge struct {
    seed int64
    seeded bool
}

type MathProblem struct {
    Value1 int64
    Value2 int64
    Symbol string
}

func (mc *math_challenge) GetProblem() MathProblem {
    return MathProblem{
        Value1: int64(mc.getRandomInt()),
        Value2: int64(mc.getRandomInt()),
        Symbol: "+",
    }
}

func (mc *math_challenge) getRandomInt() int {
    if !mc.seeded {
        rand.Seed(mc.seed)
        mc.seeded = true
    }
    return rand.Intn(100)
}

func NewMathChallenge(seed int64) math_challenge {
    return math_challenge{seed: seed, seeded: false}
}

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
