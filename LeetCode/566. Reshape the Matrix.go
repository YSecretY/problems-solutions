func matrixReshape(mat [][]int, r int, c int) [][]int {
    n := len(mat)
    m := len(mat[0])

    if r * c != n * m {
        return mat
    }

    res := make([][]int, r)
    for i := range res {
        res[i] = make([]int, c)
    }

    rInd, cInd := 0, 0
    for _, row := range mat {
        for _, el := range row {
            res[rInd][cInd] = el
            cInd++
            if cInd == c {
                cInd = 0
                rInd++
            }
        }
    }

    return res
}
