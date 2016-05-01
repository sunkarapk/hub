package generate

import (
	"github.com/pksunkara/hub/config"
	"github.com/pksunkara/hub/utils"
	"io/ioutil"
	"path/filepath"
)

type ReadmeCommand struct{}

func (r *ReadmeCommand) Execute(args []string) error {
	if len(args) > 1 {
		return &utils.ErrArgument{}
	}

	var user string
	name, err := utils.RepoName()

	if err != nil {
		return err
	}

	if len(args) == 1 {
		user = args[0]
	} else {
		user = config.Get("user")
	}

	name = user + "/" + name

	root, _ := utils.RepoRoot()
	readme := filepath.Join(root, "README.md")

	read, err := ioutil.ReadFile(readme)

	if err != nil {
		return err
	}

	data := string(read[:])

	data = data + "\n## Installation\n\n```bash\n$\n```\n\n"
	data = data + "## Usage\n\n```bash\n$\n```\n\n"
	data = data + "If you like this project, please watch this and follow me.\n\n"
	data = data + "### Testing\n\n```bash\n$\n```\n\n"
	data = data + "## Contributors\nHere is a list of [Contributors](http://github.com/" + name + "/contributors)\n\n"
	data = data + "### TODO\n\n__I accept pull requests and guarantee a reply back within a day__\n\n"
	data = data + "## License\nMIT/X11\n\n"
	data = data + "## Bug Reports\nReport [here](http://github.com/" + name + "/issues). __Guaranteed reply within a day__.\n\n"

	data = data + "## Contact\n" + config.Get("name") + " (" + config.Get("email") + ")\n\n"
	data = data + "Follow me on [github](https://github.com/users/follow?target=" + config.Get("user") + "), [twitter](http://twitter.com/" + config.Get("twitter") + ")\n"

	if err := ioutil.WriteFile(readme, []byte(data), 0644); err != nil {
		return err
	}

	if err := utils.Git([]string{"commit", "-am", "Updated readme"}...); err != nil {
		return err
	}

	utils.HandleInfo("Appended readme and committed successfully")

	return nil
}

func (r *ReadmeCommand) Usage() string {
	return "[<user>]"
}
