package exp

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

type Node struct {
	Name     string
	Path     string // full filesystem path
	IsDir    bool
	Parent   *Node
	Children []*Node
}

func (n *Node) Add(children ...*Node) {
	for _, child := range children {
		child.Parent = n
		n.Children = append(n.Children, child)
	}
}

func (n *Node) Folders() []*Node {
	var result []*Node
	for _, child := range n.Children {
		if child.IsDir {
			result = append(result, child)
		}
	}
	return result
}

func (n *Node) Files() []*Node {
	var result []*Node
	for _, child := range n.Children {
		if !child.IsDir {
			result = append(result, child)
		}
	}
	return result
}

// Entries returns folders first, then files.
func (n *Node) Entries() []*Node {
	return append(n.Folders(), n.Files()...)
}

func BuildTree(rootPath string) (*Node, error) {
	nodes := make(map[string]*Node)

	absRoot, err := filepath.Abs(rootPath)
	if err != nil {
		return nil, err
	}

	root := &Node{
		Name:  filepath.Base(absRoot),
		Path:  absRoot,
		IsDir: true,
	}
	nodes[absRoot] = root

	err = filepath.WalkDir(absRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == absRoot {
			return nil
		}

		node := &Node{
			Name:  d.Name(),
			Path:  path,
			IsDir: d.IsDir(),
		}

		parentPath := filepath.Dir(path)
		parent, ok := nodes[parentPath]
		if !ok {
			return fmt.Errorf("parent not found for %s", path)
		}

		parent.Add(node)
		nodes[path] = node
		return nil
	})

	if err != nil {
		return nil, err
	}
	return root, nil
}