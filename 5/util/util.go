package util

type RangeEntry struct  {
	Source int64
	Destination int64
	Width int64
}

func NewRangeEntry(destination int64, source int64, width int64) *RangeEntry {
	return &RangeEntry{Source: source, Destination: destination, Width: width}
}

type RangeMap struct {
	Entries []*RangeEntry
}

func NewRangeMap() *RangeMap {
	return &RangeMap{Entries: []*RangeEntry{}}
}

func (r *RangeMap) AddEntry(entry *RangeEntry) {
	r.Entries = append(r.Entries, entry)
}

func (r *RangeMap) MapSource(source int64) int64 {
	for _, entry := range r.Entries {
		if source >= entry.Source && source < entry.Source + entry.Width {
			return source + (entry.Destination - entry.Source)
		}
	}
	return source
}

func (r *RangeMap) MapDestination(destination int64) int64 {
	for _, entry := range r.Entries {
		if destination >= entry.Destination && destination <= entry.Destination + entry.Width {
			return destination + (entry.Source - entry.Destination)
		}
	}
	return destination
}

type SeedPair struct {
	Start int64
	Length int64
}
