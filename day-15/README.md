# Simplify Path

## Key learnings

- In Go, strings are immutable, meaning every concatenation (e.g., dirName += string(path[i]) or r += s[i]) creates a new string. This leads to excessive memory allocations and copying, which makes the program slower.
- Situations where we need to take out certain parts from a string we can use the Split func that helps us to grab a section bw two special chars.
- Usage of split func can also be seen when using JWT tokens for authentication.
- A JWT token is of the form header.payload.formulaForDecryption, in such cases too we use the split func to get the token and payload.