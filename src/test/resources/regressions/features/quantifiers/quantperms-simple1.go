package pkg

type Point struct {
	x int
	y int
}

requires forall p *Point :: acc(p.x) && acc(p.y)
func test1() { }

requires forall q *Point :: q == p ==> acc(q.x)
func test2(p *Point) {
  p.x = 0
}

requires forall q *Point :: acc(q.x)
func test3(p *Point) {
  assert acc(p.x)
}

requires forall q *Point :: acc(q.x) && acc(q.x)
func test4(p *Point) {
  assert acc(p.x) && acc(p.x)
}

requires forall p *Point :: { p.x } acc(p.x)
func test5() { }

requires acc(p.x)
func test6(p *Point) {
  assert forall q *Point :: q == p ==> acc(q.x)
}
