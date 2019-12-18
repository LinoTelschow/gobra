
package pkg

type cell struct{
	val int
}

requires acc(x.val) && acc(y.val)
requires x.val == a && y.val == b
ensures  acc(x.val) && acc(y.val)
ensures  x.val == b && y.val == a
func swap1(x, y *cell, ghost a, b int) {
	x.val, y.val = y.val, x.val
}

requires acc(x.val) && acc(y.val)
ensures  acc(x.val) && acc(y.val)
ensures  x.val == old(y.val) && y.val == old(x.val)
func swap2(x, y *cell) {
	x.val, y.val = y.val, x.val
}

requires acc(self.val) && acc(other.val)
ensures  acc(self.val) && acc(other.val)
ensures  self.val == old(other.val) && other.val == old(self.val)
func (self *cell) swap3(other *cell) {
  self.val, other.val = other.val, self.val
}


func client() {
    x! := cell{42}
    y! := cell{17}
    
    swap1(&x, &y, 42, 17)
    assert x == cell{17} && y.val == 42

    swap2(&x, &y)
    assert x == cell{42} && y == cell{17}

    (&x).swap3(&y)
    assert x == cell{17} && y == cell{42}

    //:: ExpectedOutput(assert_error:assertion_error)
    assert false
}



