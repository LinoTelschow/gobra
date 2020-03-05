package trivial

requires 0 <= n
ensures 2*e == (n+1)*n
func sum(n int) (e int) {
  assert e == 0
  invariant 0 <= i && i <= n + 1
  invariant 2*e == i*(i-1)
  for i := 1; i <= n; i++ {
    e = e + i
  }
}

ensures false
func infinite() {
  for {}
}

func shortStatement() {
  if s := sum(4); s == 10 {
  } else { assert s != s }
}

/* switch is not supported yet
func switchStatement() {
  switch s := sum(4); s {
  case 9: assert false;
  case 10: ;
  case 11: assert false;
  };
};
*/