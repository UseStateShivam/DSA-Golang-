# Valid Parentheses

## Key learnings

- 'something' here something is byte, while "something" is untyped string.
- map renews the value of the key, if same key provided leter.
``` Go
mapOfSeenParentheses := make(map[byte]int)
s := "(())"

for i := 0; i < len(s); i++ {
    mapOfSeenParentheses[s[i]] = i
}
fmt.Println(mapOfSeenParentheses)
//OUTPUT
//map[40:2 41:3]
```
- Such kinds of problems can be best solved using a stack. However, Go doesn't have a buit in stack structure. Instead we can use slices that are dynamic arrays. 
- Stack operates on LIFO ie, Last in First out. This is used to pop the stack, or, taking out the last inputed value.
- This is how Stack can be implemented in Go, 
``` Go
stack = []byte // declaring a slice with its data type
append(stack, 'any slice type element') // append func is used to push in a desired value
latestElement := stack[len(stack)-1] // latest element in the stack popped out
stack = stack[:(len(stack)-1)] // leaves behind the new stack after popping 
```
