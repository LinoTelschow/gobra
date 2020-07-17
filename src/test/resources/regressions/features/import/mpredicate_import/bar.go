package bar

type Rectangle struct {
    Width, Height int
}

pred (r *Rectangle) RectMem() {
    acc(r.Width) && acc(r.Height)
}

requires r.RectMem()
ensures r.RectMem()
ensures res == unfolding r.RectMem() in (r.Width * r.Height)
ensures unfolding r.RectMem() in (r.Width) == old(unfolding r.RectMem() in r.Width)
ensures unfolding r.RectMem() in (r.Height) == old(unfolding r.RectMem() in r.Height)
func (r *Rectangle) Area() (res int) {
    unfold r.RectMem()
    res := r.Width * r.Height
    fold r.RectMem()
}
