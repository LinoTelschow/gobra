
package pkg

type node struct {
  value int
  next *node
}

pred infList(headPtr *node) {
  headPtr != nil && acc(headPtr.value) && acc(headPtr.next) && infList(headPtr.next)
}


requires infList(ptr)
requires n >= 0
ensures infList(ptr)
func nth(ptr *node,n int) int {
  unfold infList(ptr)
  if(n == 0) {
    r := ptr.value
    fold infList(ptr)
    return r
  } else {
    r := nth(ptr.next,n-1)
    fold infList(ptr)
    return r
  }
}

