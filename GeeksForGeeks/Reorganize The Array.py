def Rearrange (arr,  n) : 
    hashset = set(arr)
    for i in range(n):
        if i not in hashset:
            arr[i] = -1
        else:
            arr[i] = i
    return arr
