package listfancy

type Callout struct {
	Titles    []string
	Descs     []string
	Selection *Item
}

func NewCallout(params Callout) *Callout {
	// create a new callout with the given parameters
	// if params.Selection is nil, use default values

	callout := &Callout{
		Titles:    params.Titles,
		Descs:     params.Descs,
		Selection: params.Selection,
	}

	if callout.Selection == nil {
		defaultSelection := Item{
			title:       "Default Title",
			description: "Default Description",
		}
		callout.Selection = &defaultSelection
	}

	return callout
}
