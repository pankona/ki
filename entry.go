package ki

type Entry struct {
	isDir   bool
	path    string
	entries []*Entry
}
