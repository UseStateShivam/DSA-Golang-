# H-Index

## Key learnings

- Merge sort has the best time complexity = O(n log(n))
- Merge sort can be implemented as: 
``` Go

func merge(arr []int, left int, mid int, right int) {
    n1 := mid - left + 1
    n2 := right - mid
    arr1 := make([]int, n1)
    arr2 := make([]int, n2)
    for i := 0; i < n1; i++ {
        arr1[i] = arr[left + i]
    }   
    for j := 0; j < n2; j++ {
        arr1[j] = arr[mid + 1 + j]
    } 
    k := 0
    l := 0
    m := left
    for k < n1 && l < n2 {
        if arr1[k] <= arr2[l] {
            arr[m] = arr1[k]
            k++
        } else {
            arr[m] = arr2[l]
            l++
        }
        m++
    }
    for k < n1 {
        arr[m] = arr1[k]
        m++
        k++
    }
    for l < n2 {
        arr[m] = arr2[l]
        m++
        l++
    }
}

func mergeSort(arr []int, left int, right int) {
    if left >= right {
        return
    }
    mid := left + ((right - left) / 2 )
    mergeSort(arr, left, mid)
    mergeSort(arr, mid + 1, right)
    merge(arr, left, mid, right)
}

```

# IsHappy 🙂

## Key learnings

- Me getting pro 😎