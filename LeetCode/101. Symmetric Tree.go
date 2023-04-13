func Equal(left, right *TreeNode) bool {
    if left == nil || right == nil {
        return left == right
    }
    return left.Val == right.Val && Equal(left.Left, right.Right) && Equal(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
    return Equal(root.Left, root.Right) 
}
