# Find the Index of the First Occurrence in a String

## Key learnings

- Strings in Go are immutable. If you want to mutate them:
``` Go
var nameOfString string
nameOfString += string("something") // string func is to be used if "something is not a string"
```
- To slice out a string in Go:
``` Go
slicedString := stringName[startingIndex : endingIndex]
```

# Search Insert Position

## Key learnings

- One of the algorithms operating on O(log n) time complexity, is Binary Search.
- Implementation of Binary Search:
``` Go
// There are two pointers
left, right := 0, len(slice) - 1 
// loop works until left <= right
for left <= right {
    // define a mid point
    mid := left + (right - left)/2
    if slice[mid] == target {
        return mid // target found at mid
    } else if slice[mid] < target {
        left = mid + 1 // search in rhs
    } else {
        right = mid - 1 // search in lhs 
    }
} 
return left // target not found 
```