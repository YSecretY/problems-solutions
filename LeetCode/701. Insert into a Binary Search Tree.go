func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        r := &TreeNode{
            Val: val,
            Left: nil,
            Right: nil,
        }
        root = r
        return root
    }

    if val < root.Val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }

    return root
}
