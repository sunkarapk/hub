package main

type AuthCommand struct {
	AuthUser    AuthUserCommand    `command:"user" alias:"u" description:"Set an username to use their public data"`
	AuthPrivate AuthPrivateCommand `command:"private" alias:"p" description:"Give access to your private data"`
	AuthDestroy AuthDestroyCommand `command:"destroy" alias:"d" description:"Destroy authorization and delete username"`
}
