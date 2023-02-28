package theasurus

type Theaurus interface {
	Synonyms(term string) ([]string, error)
}
