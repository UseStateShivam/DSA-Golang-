package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	valMap  map[int]int
	valList []int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		valMap:  make(map[int]int),
		valList: []int{},
	}
}

func (randomizedSet *RandomizedSet) Insert(val int) bool {
	if _, exists := randomizedSet.valMap[val]; !exists {
		randomizedSet.valList = append(randomizedSet.valList, val)
		randomizedSet.valMap[val] = len(randomizedSet.valList) - 1
		return true
	} else {
		return false
	}
}

func (randomizedSet *RandomizedSet) Remove(val int) bool {
	if index, exists := randomizedSet.valMap[val]; exists {
		last := randomizedSet.valList[len(randomizedSet.valList)-1]
		randomizedSet.valList[index] = last
		randomizedSet.valMap[last] = index
		randomizedSet.valList = randomizedSet.valList[:len(randomizedSet.valList)-1]
		delete(randomizedSet.valMap, val)
		return true
	} else {
		return false
	}
}

func (randomizedSet *RandomizedSet) GetRandom() int {
	count := len(randomizedSet.valList)
	randIndex := rand.Intn(count)
	return randomizedSet.valList[randIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

 // 50ms
 // 54mb