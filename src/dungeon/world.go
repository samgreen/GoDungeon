package dungeon

type World struct {
	name		string

	activeZones *[]Zone
	players     *[]Player
}

func NewWorld(name string) *World {
	return &World{name: name, activeZones: make([]Zone, 10, 100)}
}

func (w *World) Name() string {
	return w.name
}

/*
	// NewReaderSize returns a new Reader whose buffer has at least the specified
    43	// size. If the argument io.Reader is already a Reader with large enough
    44	// size, it returns the underlying Reader.
    45	func NewReaderSize(rd io.Reader, size int) *Reader {
    46		// Is it already a Reader?
    47		b, ok := rd.(*Reader)
    48		if ok && len(b.buf) >= size {
    49			return b
    50		}
    51		if size < minReadBufferSize {
    52			size = minReadBufferSize
    53		}
    54		return &Reader{
    55			buf:          make([]byte, size),
    56			rd:           rd,
    57			lastByte:     -1,
    58			lastRuneSize: -1,
    59		}
    60	}
*/
