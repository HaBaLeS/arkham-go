package runtime

import (
	"math/rand"
	"sync"
)

type DRandom struct {
	lk   sync.Mutex
	rand *rand.Rand
}

func NewDeterministicRandom(seed int64) *DRandom {
	src := rand.NewSource(seed)
	return &DRandom{
		sync.Mutex{},
		rand.New(src),
	}
}

func ShuffleDeck[K []any](deck K, rnd *DRandom) {
	rnd.lk.Lock()
	defer rnd.lk.Unlock()
	rnd.rand.Shuffle(len(deck), func(a int, b int) { deck[a], deck[b] = deck[b], deck[a] })
}
