# Climbing Stairs

## Key learnings

- The problems associated with the new number depending upon the previous decisions made, can be simplified using the fibonacci algorithm. According to this, every new numebr is the sum of preciding two numbers.
    ``` lua
                |-> 0                [n=0]
        F(n) -> |-> 1                [n=1]
                |-> F(n-1) + F(n-2)  [n>1]
    ```
- The below approach:
    ``` Go
        func F(n int) int {
            if n == 1 {
                return 1
            } else if n == 0 {
                return 0
            } else {
                return F(n-1) + F(n-2)
            }
        }
    ```
Although this is a conceptually correct approach, however, as the n increases, the time complexity of the algorithm increases exponentially.
- To optimize fibonacci recursive computation, we can use a simple for loop:
    ``` Go
        a, b := 1,1 // initializing the first two numbers of the series except 0
        for i := 2; i <= n; i++ {
            a, b = b, a+b
        }
        return b
    ```

# Remove Duplicates from Sorted List

## Key learnings

- I am done with this linked list shi-
- Nah!! even gpt can't stand against linked list
