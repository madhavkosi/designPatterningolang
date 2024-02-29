package prototype

import (
	"fmt"

	"testing"
)

func TestProtype(t *testing.T) {
	folder1 := Folder{Name: "folder1"}
	file1 := File{Name: "file1"}
	file2 := File{Name: "file2"}
	file3 := File{Name: "file3"}
	file4 := File{Name: "file4"}
	folder2 := Folder{Name: "folder2"}
	folder3 := Folder{Name: "folder3"}

	folder1.Children = []ClonePrototype{&folder2, &file1, &file2}
	folder2.Children = []ClonePrototype{&folder3, &file3, &file4}
	folder_clone := folder1.Clone()
	fmt.Printf("cloned File %+v\n", folder_clone)
	fmt.Println(folder1)
}
