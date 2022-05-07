package main

import "fmt"

func main(){
	hashMap:=make(map[[26]int]int)
	nums1:=[26]int{1,2,3}
	nums2:=[26]int{1,2,4}
	hashMap[nums1]=10
	hashMap[nums2]=20
	fmt.Printf("%p ; %p\n", &nums1, &nums2)
	fmt.Println(hashMap[nums2])

}


