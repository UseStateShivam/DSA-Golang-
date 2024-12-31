package main

func Merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[0:]...)
	msort(nums1, 0, len(nums1)-1)
}

func m1(a []int, l, r, m int) {
	n1 := m - l + 1
	n2 := r - m
	a1 := make([]int, n1)
	a2 := make([]int, n2)
	for i := 0; i < len(a1); i++ {
		a1[i] = a[l+i]
	}
	for i := 0; i < len(a2); i++ {
		a2[i] = a[m+1+i]
	}
	p, q := 0, 0
	w := l
	for p < n1 && q < n2 {
		if a1[p] <= a2[q] {
			a[w] = a1[p]
			p++
		} else {
			a[w] = a2[q]
			q++
		}
        w++
	}
	for p < n1 {
		a[w] = a1[p]
		p++
        w++
	}
	for q < n2 {
		a[w] = a2[q]
		q++
        w++
	}
}

func msort(a []int, l, r int) {
	if l >= r {
		return
	}
	m := l + ((r - l) / 2)
	msort(a, l, m)
	msort(a, m+1, r)
	m1(a, l, r, m)
}

// 0ms
// 4.31mb