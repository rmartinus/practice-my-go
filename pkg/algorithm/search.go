package algorithm

// LinearSearch searches an item linearly
func LinearSearch(items []int, searchItem int) int {
	for i, item := range items {
		if item == searchItem {
			return i
		}
	}
	return 0
}

// BinarySearch searches an item using binary search algorithm
func BinarySearch(sortedItems []int, searchItem int) int {
	middleIndex := len(sortedItems) / 2
	middleItem := sortedItems[middleIndex]
	if searchItem == middleItem {
		return middleIndex
	} else if searchItem > middleItem {
		return middleIndex + BinarySearch(sortedItems[middleIndex:], searchItem)
	} else if searchItem < middleItem {
		return BinarySearch(sortedItems[0:middleIndex], searchItem) - middleIndex
	}
	return 0
}
