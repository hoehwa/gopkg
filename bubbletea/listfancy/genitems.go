package listfancy

import (
	"math/rand"
	"sync"
)

type itemGenerator struct {
	callout    Callout
	titleIndex int
	descIndex  int
	mtx        *sync.Mutex
	shuffle    *sync.Once
}

func (i *itemGenerator) setCallout(callout Callout) {
	i.callout = callout
}

func (i *itemGenerator) getCalloutLength() int {
	return len(i.callout.Titles)
}

func (i *itemGenerator) reset() {
	i.mtx = &sync.Mutex{}
	i.shuffle = &sync.Once{}

	i.shuffle.Do(func() {
		shuf := func(x []string) {
			rand.Shuffle(len(x), func(i, j int) { x[i], x[j] = x[j], x[i] })
		}
		shuf(i.callout.Titles)
		shuf(i.callout.Descs)
	})
}

func (i *itemGenerator) next() item {
	if i.mtx == nil {
		i.reset()
	}

	i.mtx.Lock()
	defer i.mtx.Unlock()

	idx := item{
		title:       i.callout.Titles[i.titleIndex],
		description: i.callout.Descs[i.descIndex],
	}

	i.titleIndex++
	if i.titleIndex >= len(i.callout.Titles) {
		i.titleIndex = 0
	}

	i.descIndex++
	if i.descIndex >= len(i.callout.Descs) {
		i.descIndex = 0
	}

	return idx
}
