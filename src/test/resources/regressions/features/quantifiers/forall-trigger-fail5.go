package pkg

type Cell struct {
	val int
}

// invalid trigger: disallowed for now; one just has to write `c` instead of `*c`.
//:: ExpectedOutput(type_error)
requires forall c *Cell :: { *c } acc(c.val)
func foo () { }
