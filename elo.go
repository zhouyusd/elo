package elo

import "github.com/zhouyusd/elo/utils"

type Input struct {
	Rating float64
	Rank   int
}

// CalcNewRating 给定排行榜, 计算每个选手新的Rating
func CalcNewRating(members []Input, k float64) (newRatings []float64) {
	k /= float64(len(members))
	newRatings = make([]float64, len(members))
	for i := range members {
		newRatings[i] = members[i].Rating
	}
	for i := 0; i < len(members); i++ {
		for j := i + 1; j < len(members); j++ {
			e := utils.CalcExpectation(members[i].Rating, members[j].Rating)
			s := 0.0
			if members[i].Rank < members[j].Rank {
				s = 1.0
			} else if members[i].Rank == members[j].Rank {
				s = 0.5
			}
			delta := utils.CalcDelta(e, s, k)
			newRatings[i] += delta
			newRatings[j] -= delta
		}
	}
	return newRatings
}
