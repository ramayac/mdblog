package blog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DirNode represents a node in the directory tree for content count reporting.
type DirNode struct {
	Name      string
	Path      string
	FileCount int
	Children  []*DirNode
}

// PrintContentCount scans the post and page directories and prints a treeview.
func (b *Blog) PrintContentCount() {
	fmt.Println("Content Count Treeview:")
	fmt.Println("=======================")

	if b.cfg.PostsDir != "" {
		if node, err := b.buildDirTree(b.cfg.PostsDir); err == nil {
			b.printNode(node, "")
		} else {
			fmt.Printf("%s (error scanning: %v)\n", b.cfg.PostsDir, err)
		}
	}

	fmt.Println()

	if b.cfg.PagesDir != "" {
		if node, err := b.buildDirTree(b.cfg.PagesDir); err == nil {
			b.printNode(node, "")
		} else {
			fmt.Printf("%s (error scanning: %v)\n", b.cfg.PagesDir, err)
		}
	}
}

func (b *Blog) buildDirTree(dirPath string) (*DirNode, error) {
	node := &DirNode{
		Name: filepath.Base(dirPath),
		Path: dirPath,
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			childPath := filepath.Join(dirPath, entry.Name())
			childNode, err := b.buildDirTree(childPath)
			if err == nil {
				node.Children = append(node.Children, childNode)
				node.FileCount += childNode.FileCount
			}
		} else {
			if strings.HasSuffix(strings.ToLower(entry.Name()), ".md") {
				node.FileCount++
			}
		}
	}

	return node, nil
}

func (b *Blog) printNode(node *DirNode, prefix string) {
	fmt.Printf("%s (%d files)\n", node.Name, node.FileCount)
	b.printChildren(node.Children, prefix)
}

func (b *Blog) printChildren(children []*DirNode, prefix string) {
	for i, child := range children {
		isLast := i == len(children)-1
		marker := "├── "
		if isLast {
			marker = "└── "
		}
		fmt.Printf("%s%s%s (%d files)\n", prefix, marker, child.Name, child.FileCount)

		var nextPrefix string
		if isLast {
			nextPrefix = prefix + "    "
		} else {
			nextPrefix = prefix + "│   "
		}
		b.printChildren(child.Children, nextPrefix)
	}
}
