package pkg

ensures forall i int :: 0 <= i && i < len(a) ==> a[i] == 0
func clear() (a! [64]int) {
  invariant 0 <= i && i <= len(a)
  invariant forall k int :: 0 <= k && k < len(a) ==> acc(a[k])
  invariant forall k int :: 0 <= k && k < i ==> a[k] == 0
	for i := 0; i < len(a); i++ {
    a[i] = 0
	}
}
