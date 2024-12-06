# Palindromic Number

## Key learnings

- len() func in Go only works on data structures like arrays, maps and channels.
- In Go, to convert given int to string, we use a package strconv and use Itoa method, as follows:
``` Go
str := strconv.Itoa(x int)
```
- Go does not have native while loops.

# Roman to Int 

## Key learnings

- Syntax for switch in Go:
``` Go
switch {
    case condition:
        return something
    case condition2:
        return something2
}
```
- Using a for loop in Go, string[i] type is uint8 (single byte).
``` Go
	str := "Hello"
	fmt.Printf("Type of str[0]: %T\n", str[0])    // Output: uint8
	fmt.Printf("Value of str[0]: %d\n", str[0])   // Output: 72 (ASCII value of 'H')
	fmt.Printf("Char of str[0]: %c\n", str[0])    // Output: H
```
- Strings in Go are of two types, ASCII, ie normal single byte per char strings, and UTF-8 strings, including special chars or emojis.

# Longest Common Prefix

## Key learnings

- Only key learning is that I need more practice :(