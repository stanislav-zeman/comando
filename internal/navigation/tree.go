package navigation

import "github.com/stanislav-zeman/comando/internal/config"

type TreeNode struct {
	Name     string
	Command  string
	Children []*TreeNode
	Parent   *TreeNode
	IsFolder bool
}

// ParseTree constructs a tree of TreeNodes from config nodes
func ParseTree(nodes []config.Node) []*TreeNode {
	return buildTreeRecursive(nodes, nil)
}

func buildTreeRecursive(nodes []config.Node, parent *TreeNode) []*TreeNode {
	result := make([]*TreeNode, 0, len(nodes))

	for _, node := range nodes {
		treeNode := &TreeNode{
			Name:     node.Name,
			Command:  node.Command,
			Parent:   parent,
			IsFolder: len(node.Children) > 0,
		}

		if treeNode.IsFolder {
			treeNode.Children = buildTreeRecursive(node.Children, treeNode)
		}

		result = append(result, treeNode)
	}

	return result
}
