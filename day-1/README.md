# TwoSum

## Key learnings

- Struct in Go is like the object is js. Strcut is an object that stores the field parameters and corresponding values.
- The type 'nameOfStruct' struct {} is basically the way of making/declaring the strcut.
- Map in Go is like Hashmap in Java.
- If we just need to implement the functionality of checking if a certain value has been seen or not, we can simply assign a key value pair in map. Unlike in other languages in Go we have the ability to store the value as an empty struct. Structs in Go have the ability to take up "Zero" space in the memory.
- The struct{}{} syntax is used to define an empty struct and initialize it with literal zero value.
- In maps, the way to assign a key is like following:
 ``` Go
 mapName[key] = value
 ```
 - In Go, to lookup the map, we use mapName[key] this returns two values, the value of the key itself and a bool stating the key in map was found or not. Thus the syntax looks like follows:
 ``` Go
 _, exists := mapName[key]
 // here _ is used to ignore the first return, ie, the actual value of the key, we are only concerned for the bool of exsistence. 
```
- Make function in Go is used to initialize slices, maps and channels. To initialize a slice:
``` Go
make([]sliceType, minimumSpaceInSlice, totalCapacity)
```
- For loop in Go is as follows:
``` GO
// just like other major languages (not the snake :))
for i := 0; i < totalCapacity; i++{
    // Loop body
}
``` 
- Append function in Go, is a built-in func used to add elements to slices like:
``` Go
append(nameOfSlice, element)
```
- In Go, len(sliceName) is used to find the length of the slice.
- Return a slice in Go? us: 
``` Go
[]sileDataType{values}
```
