package solution

import (
	"sort"
)

// A := []int{10, 2, 5, 1, 8, 12}
// [1, 2, 5, 8, 10, 12]

func Solution(A []int) int {
	N := len(A)
	if N < 3 {
		return 0 // 3つ未満の要素では三角形を作れない
	}

	sort.Ints(A) // ソートして順序を固定

	count := 0

	// PとQを固定し、Rを探索する
	for P := 0; P < N-2; P++ {
		R := P + 2 // Rの初期位置を設定（P, Q, Rの順番を守る）
		for Q := P + 1; Q < N-1; Q++ {
			// R をできるだけ右に動かして条件を満たす範囲を探す
			for R < N && A[P]+A[Q] > A[R] {
				R++
			}
			count += (R - Q - 1) // 有効な R の個数をカウント
		}
	}

	return count
}
