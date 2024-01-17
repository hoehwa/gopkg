package listfancy

import (
	"sync"
)

type itemGenerator struct {
	callout    Callout
	titleIndex int
	descIndex  int
	mtx        *sync.Mutex
}

func (i *itemGenerator) setCallout(callout Callout) {
	i.callout = callout
}

func (i *itemGenerator) getCalloutLength() int {
	return len(i.callout.Titles)
}

func (i *itemGenerator) reset() {
	i.mtx = &sync.Mutex{}
}

func (i *itemGenerator) next() Item {
	if i.mtx == nil {
		i.reset()
	}

	i.mtx.Lock()
	defer i.mtx.Unlock()

	idx := Item{
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
