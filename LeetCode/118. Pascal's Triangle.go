func generate(numRows int) [][]int {
    res := make([][]int, numRows)
    res[0] = []int{1}

    for i := 0; i < numRows-1; i++ {
        temp := []int{0}
        temp = append(temp, res[i]...)
        temp = append(temp, 0)
        row := make([]int, len(res[i])+1)
        for j := 0; j < len(temp)-1; j++ {
            row[j] = temp[j] + temp[j+1]
        }
        res[i+1] = row
    }

    return res
}
