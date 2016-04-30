package auth

type Command struct {
	User    UserCommand    `command:"user" alias:"u" description:"Set an username to use their public data"`
	Private PrivateCommand `command:"private" alias:"p" description:"Give access to your private data"`
	Destroy DestroyCommand `command:"destroy" alias:"d" description:"Destroy authorization and delete username"`
}
