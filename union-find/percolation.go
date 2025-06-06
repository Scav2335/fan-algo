package union_find

type Site struct {
	parent int
	isOpen bool
}

type Percolator struct {
	Sites []*Site
	Sz    []int
	N     int
}

func NewPercolator(N int) *Percolator {
	c := N*N + 2
	sites := make([]*Site, c)
	sz := make([]int, c)

	for i := 0; i < c; i++ {
		site := &Site{
			parent: i,
			isOpen: false,
		}
		// virtual site should be open
		if i == 0 || i == c-1 {
			site.isOpen = true
		}

		sites[i] = site
		sz[i] = 1
	}

	return &Percolator{
		Sites: sites,
		Sz:    sz,
		N:     N,
	}
}

func (p *Percolator) Root(i int) int {
	for i != p.Sites[i].parent {
		p.Sites[i].parent = p.Sites[p.Sites[i].parent].parent
		i = p.Sites[i].parent
	}
	return i
}

func (p *Percolator) Connected(i, j int) bool {
	return p.Root(i) == p.Root(j)
}

func (p *Percolator) Union(i, j int) {
	ir, jr := p.Root(i), p.Root(j)
	if ir == jr {
		return
	}

	if p.Sz[ir] < p.Sz[jr] {
		p.Sites[ir].parent = jr
		p.Sz[jr] += p.Sz[ir]
	} else {
		p.Sites[jr].parent = ir
		p.Sz[ir] += p.Sz[jr]
	}
}

func (p *Percolator) Open(i int) (ok bool) {
	if p.Sites[i].isOpen {
		return
	}
	p.Sites[i].isOpen = true

	// Is it the top site?
	if i > 0 && i < p.N+1 {
		p.Union(0, i)
	}
	// Is it the bottom site?
	if i > p.N*p.N-p.N && i < p.N*p.N+1 {
		p.Union(i, p.N*p.N+1)
	}

	// up
	if up := i - p.N; up > 0 && up < p.N*p.N+1 && p.Sites[up].isOpen {
		p.Union(up, i)
	}
	// down
	if down := i + p.N; down > 0 && down < p.N*p.N+1 && p.Sites[down].isOpen {
		p.Union(i, down)
	}
	// left
	if left := i - 1; left > 0 && left < p.N*p.N+1 && p.Sites[left].isOpen {
		p.Union(left, i)
	}
	// right
	if right := i + 1; right > 0 && right < p.N*p.N+1 && p.Sites[right].isOpen {
		p.Union(i, right)
	}

	return true
}
