package main

//Declaring the struct
type SeenNumbers struct{
    numbers map[int]int
}

//Initializing the struct
func NewSeenNumbers() *SeenNumbers{
    return &SeenNumbers{
        numbers: make(map[int]int),
    }
}

func (s *SeenNumbers) Add(number int, index int) {
    s.numbers[number] = index
}

func (s *SeenNumbers) HasSeenNumbers(number int) (int, bool) {
    value, exists := s.numbers[number]
    return value, exists
}

func TwoSum(nums []int, target int) []int {
    rem := 0
    setOfSeenNumbers := NewSeenNumbers() 
    for i := 0; i < len(nums); i++ {
        rem = target - nums[i]
        remIndex, remExists := setOfSeenNumbers.HasSeenNumbers(rem)
        if (remExists) {
            return []int{i, remIndex}
        }
        setOfSeenNumbers.Add(nums[i], i)
    }
    return []int{}
}

//2ms, O(n)
//5.99mb
// to optimize the runtime performance further, we can directly use the map function inside TwoSum func rather than defining custom functions to add or check the existence.