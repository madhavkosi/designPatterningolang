package prototype

type File struct {
	Name string
}

func (f *File) Clone() ClonePrototype {
	return &File{Name: f.Name + "_copy"}
}
