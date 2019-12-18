
package list

type node struct {
  value int
  next *node
}

pred list(ptr *node) {
  ptr != nil ==> acc(ptr.value) && acc(ptr.next) && list(ptr.next)
}



requires list(ptr)
requires list(rev)
ensures list(res)
func tail_reverse (ptr, rev *node) (res *node) {
  unfold list(ptr)
  if(ptr == nil) {
    return rev
  } else {
    tmp := ptr
    ptr = ptr.next
    tmp.next = rev
    fold list(tmp)
    return tail_reverse(ptr,tmp)
  }
}

requires list(ptr)
ensures list (res)
func reverse (ptr *node) (res *node) {
  fold list(nil)
  return tail_reverse(ptr,nil)
}




