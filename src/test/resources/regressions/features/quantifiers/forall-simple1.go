package pkg;

requires forall m int :: true
func test1 () { }

requires forall x int, y int :: x < y
requires forall x int, y int :: { x < y } x < y
func test2 () { }

func test3() {
  assert forall n int :: true
  assert forall n int :: { n } true
}

requires forall x int, y int :: { x, y } x < y
func test4 () { }