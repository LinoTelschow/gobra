package pkg

func foo() {
	//:: ExpectedOutput(parser_error)
	var a [2][2]int = [...][...] { [2]int { 1, 2 }, [2]int { 3, 4 } }
}
