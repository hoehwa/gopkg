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

	/*
		i.callout.titles = []string{
			"Artichoke",
			"Baking Flour",
			"Bananas",
			"Barley",
			"Bean Sprouts",
			"Bitter Melon",
			"Black Cod",
			"Blood Orange",
			"Brown Sugar",
			"Cashew Apple",
			"Cashews",
			"Cat Food",
			"Coconut Milk",
			"Cucumber",
			"Curry Paste",
			"Currywurst",
			"Dill",
			"Dragonfruit",
			"Dried Shrimp",
			"Eggs",
			"Fish Cake",
			"Furikake",
			"Garlic",
			"Gherkin",
			"Ginger",
			"Granulated Sugar",
			"Grapefruit",
			"Green Onion",
			"Hazelnuts",
			"Heavy whipping cream",
			"Honey Dew",
			"Horseradish",
			"Jicama",
			"Kohlrabi",
			"Leeks",
			"Lentils",
			"Licorice Root",
			"Meyer Lemons",
			"Milk",
			"Molasses",
			"Muesli",
			"Nectarine",
			"Niagamo Root",
			"Nopal",
			"Nutella",
			"Oat Milk",
			"Oatmeal",
			"Olives",
			"Papaya",
			"Party Gherkin",
			"Peppers",
			"Persian Lemons",
			"Pickle",
			"Pineapple",
			"Plantains",
			"Pocky",
			"Powdered Sugar",
			"Quince",
			"Radish",
			"Ramps",
			"Star Anise",
			"Sweet Potato",
			"Tamarind",
			"Unsalted Butter",
			"Watermelon",
			"Weißwurst",
			"Yams",
			"Yeast",
			"Yuzu",
			"Snow Peas",
		}

		i.callout.descs = []string{
			"A little weird",
			"Bold flavor",
			"Can’t get enough",
			"Delectable",
			"Expensive",
			"Expired",
			"Exquisite",
			"Fresh",
			"Gimme",
			"In season",
			"Kind of spicy",
			"Looks fresh",
			"Looks good to me",
			"Maybe not",
			"My favorite",
			"Oh my",
			"On sale",
			"Organic",
			"Questionable",
			"Really fresh",
			"Refreshing",
			"Salty",
			"Scrumptious",
			"Delectable",
			"Slightly sweet",
			"Smells great",
			"Tasty",
			"Too ripe",
			"At last",
			"What?",
			"Wow",
			"Yum",
			"Maybe",
			"Sure, why not?",
		}
	*/

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
