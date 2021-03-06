package bucket

import (
	"fmt"
	"Algorithms-Learning-With-Go/sort/utils"
	
	"testing"
)


func TestBucketSort(t *testing.T) {
	list := utils.GetArrayOfSize(100)
			
	Sort(list)

	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println(list)
			t.Error()
		}
	}
}

func benchmarkBucketSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	
	for i := 0; i < b.N; i++ {
		Sort(list)
	}
}

func BenchmarkBucketSort100(b *testing.B) { benchmarkBucketSort(100, b) }
func BenchmarkBucketSort1000(b *testing.B)   { benchmarkBucketSort(1000, b) }
func BenchmarkBucketSort10000(b *testing.B)  { benchmarkBucketSort(10000, b) }
func BenchmarkBucketSort100000(b *testing.B) { benchmarkBucketSort(100000, b) }
