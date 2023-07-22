package util

import "sort"

func SortNum(arr []int) (int[]) {
	for i:= 0; i < len(arr); i++ {
		int min = i;

		for j:= i+1; j < len(arr); j++ {
			if (arr[j] < arr[min]) {
				min = j;
			}
		}

		int temp = arr[i];
		arr[i] = arr[min];
		arr[min] = temp;
	}

	return arr;
}

func SortFloat(arr []float64) ([]float64) {
	for i:= 0; i < len(arr); i++ {
		int min = i;

		for j:= i+1; j < len(arr); j++ {
			if (arr[j] < arr[min]) {
				min = j;
			}
		}

		int temp = arr[i];
		arr[i] = arr[min];
		arr[min] = temp;
	}

	return arr;
}

func SortString(arr []String) ([]String) {
	return sort.Strings(arr)
}