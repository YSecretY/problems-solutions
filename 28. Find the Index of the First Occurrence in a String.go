func strStr(haystack string, needle string) int {
    var buf []byte
    for i := 0; i < len(haystack); i++ {
        k := 0
        j := i
        for j < len(haystack) && k < len(needle) && haystack[j] == needle[k] {
            buf = append(buf, haystack[j])
            if string(buf) == needle {
                return i
            }
            k++
            j++
        }
        buf = buf[:0]
    }
    return -1
}
