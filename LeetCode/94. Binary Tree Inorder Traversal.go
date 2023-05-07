func inorderTraversal(root *TreeNode) []int {
    var res []int
    return traversal(root, res)
}

func traversal(root *TreeNode, res []int) []int {
    if root == nil {
        return res
    }

    res = traversal(root.Left, res)
    res = append(res, root.Val)
    res = traversal(root.Right, res)

    return res
}
