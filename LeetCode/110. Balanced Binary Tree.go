func dfs(root *TreeNode) (bool, int) {
    if root == nil {
        return true, -1
    }
    ok1, left := dfs(root.Left)
    ok2, right := dfs(root.Right)
    balanced := ok1 && ok2 && math.Abs(float64(left-right)) <= 1

    return balanced, int(math.Max(float64(left), float64(right))) + 1
}

func isBalanced(root *TreeNode) bool {
    res, _ := dfs(root)
    return res
}
