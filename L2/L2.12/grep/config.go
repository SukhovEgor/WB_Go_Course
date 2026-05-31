package grep

type GrepConfig struct {
	AfterContext  int
	BeforeContext int
	Context       int

	Count       bool
	IgnoreCase  bool
	InvertMatch bool
	FixedString bool
	LineNumber  bool
}

func (c *GrepConfig) Contextualize() {
	if c.Context > 0 {
		if c.AfterContext < c.Context {
			c.AfterContext = c.Context
		}

		if c.BeforeContext < c.Context {
			c.BeforeContext = c.Context
		}
	}
}