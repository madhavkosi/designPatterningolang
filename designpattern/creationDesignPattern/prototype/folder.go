package prototype

type Folder struct {
	Name     string
	Children []ClonePrototype
}

func (f Folder) Clone() ClonePrototype {
	folder := &Folder{Name: f.Name + "_copy"}
	folder.Children = []ClonePrototype{}
	for _, child := range f.Children {
		folder.Children = append(folder.Children, child.Clone())
	}
	return folder
}
