package main

import (
	"fmt"
	"math/rand"
	"time"
)

const Dad = "Dad"
const Mom = "Mom"
const Julie = "Julie"
const Kim = "Kim"
const Tim = "Tim"
const SungAe = "SungAe"
const Andy = "Andy"

type assignment struct {
	who string
	got string
}

func main() {
	var n = 1000000

	var gotSelfTotal = 0
	var gotSpouseButNotSelf = 0

	for i := 0; i < n; i++ {
		givers := pool()
		assignees := randomize(pool())

		assignments := assign(givers, assignees)

		self, spouse := analyze(assignments)

		if self {
			gotSelfTotal++
			continue
		}
		if spouse {
			gotSpouseButNotSelf++
		}
	}

	var roundsNotGettingSelf = n - gotSelfTotal

	fmt.Printf("nobody got themselves but somebody got their spouse: %f\n", float64(gotSpouseButNotSelf)/(float64(roundsNotGettingSelf)))
}

func somebodyGotThemselves(assignments []assignment) bool {
	for _, assn := range assignments {
		if gotSelf(assn.who, assn.got) {
			return true
		}
	}
	return false
}

func somebodyGotTheirSpouse(assignments []assignment) bool {
	for _, assn := range assignments {
		if gotSpouse(assn.who, assn.got) {
			return true
		}
	}
	return false
}

func analyze(assignments []assignment) (bool, bool) {
	var a = somebodyGotThemselves(assignments)
	var b = somebodyGotTheirSpouse(assignments)

	return a, b
}

func assign(pool1, pool2 []string) []assignment {
	var assignments []assignment

	for i, who := range pool1 {
		assignments = append(assignments, assignment{
			who: who,
			got: pool2[i],
		})
	}

	return assignments
}

func randomize(pool []string) []string {
	randomized := make([]string, len(pool))

	var randomness = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i, j := range randomness.Perm(len(pool)) {
		randomized[i] = pool[j]
	}

	return randomized
}

func pool() []string {
	return []string{
		Dad,
		Mom,
		Julie,
		Kim,
		Tim,
		SungAe,
		Andy,
	}
}

func gotSpouse(who string, got string) bool {
	switch who {
	case Kim:
		if got == Tim {
			return true
		}
	case Tim:
		if got == Kim {
			return true
		}
	case Andy:
		if got == SungAe {
			return true
		}
	case SungAe:
		if got == Andy {
			return true
		}
	case Mom:
		if got == Dad {
			return true
		}
	case Dad:
		if got == Mom {
			return true
		}
	}

	return false
}

func gotSelf(who string, got string) bool {
	switch who {
	case Julie:
		if got == Julie {
			return true
		}
	case Kim:
		if got == Kim {
			return true
		}
	case Tim:
		if got == Tim {
			return true
		}
	case Andy:
		if got == Andy {
			return true
		}
	case SungAe:
		if got == SungAe {
			return true
		}
	case Mom:
		if got == Mom {
			return true
		}
	case Dad:
		if got == Dad {
			return true
		}
	}

	return false
}
