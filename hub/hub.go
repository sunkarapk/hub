package hub

type Command interface {
	Short() string
	Title() string

	Description() string

	Usage() string

	Run(args []string)
}

type Common struct {
	children map[string]Command
}

func (c *Common) Init() {
	c.children = make(map[string]Command)
}

func (c *Common) Add(d Command) {
	c.children[d.Title()] = d
	c.children[d.Short()] = d
}

func ShowHelp(c Command) {
	usage := "hub"
}
