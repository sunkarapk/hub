package hub

type Command interface {
	Short() string
	Title() string

	Description() string

	Run(args []string)
}
