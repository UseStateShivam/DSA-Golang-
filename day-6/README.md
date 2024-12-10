# Length of Last Word

## Key learnings

- I am improving hehe...

# Add Binary

## Key learnings

- Binary addition is similar to decimal addition but operates in base 2.
- Rules for Binary addition are:
    ``` lua
    0 + 0 = 0
    0 + 1 = 1
    1 + 0 = 1
    1 + 1 = 10 (0 with carry 1)
    1 + 1 + 1 = 11 (1 with carry 1)
    ```
- To align shorter string with longer string:
    ``` Go
    diff := len(a) - len(b)
    b = strings.Repeat("0", diff) + b
    ```
- Converting an int to a string directly yields a string representing the UTF-8 character corresponding to that integer value, not the string of its digits
    ``` Go
    output = string(sum%2+'0') + output // String of UTF-8 character 
    output = strconv.Itoa(sum%2+'0') + output // String of actual int
    ```
- I have to revisit binary addition :(