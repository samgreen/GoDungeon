package dungeon

type Item struct {
	tradeable    bool
	sellable     bool
	name        string
	description string
}

func NewItem(name string) *Item {
	item := &Item{
		name: name,
	}
	return item
}

func (i *Item) String() string {
	return i.name
}

func (i *Item) CanSell() bool {
	return i.sellable
}

func (i *Item) CanTrade() bool {
	return i.tradeable
}

func (i *Item) Name() string {
	return i.name
}

func (i *Item) Description() string {
	return i.description
}
