package ki

type entry struct {
	isDir   bool
	path    string
	entries []*entry
}
