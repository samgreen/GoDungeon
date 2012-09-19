package dungeon

type Zone struct {
	id int32;

	Name string

}

func (z *Zone) GetId() int32 {
	return z.id;
}