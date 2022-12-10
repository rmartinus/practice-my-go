package algorithm

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// LinearSearch searches an item linearly
func LinearSearch[T Number](items []T, searchItem T) int {
	for i, item := range items {
		if item == searchItem {
			return i
		}
	}
	return 0
}

// BinarySearch searches an item using binary search algorithm
func BinarySearch[T Number](sortedItems []T, searchItem T) int {
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
