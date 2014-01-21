package main

type AuthCommand struct {
	AuthUser    AuthUserCommand    `command:"user" description:"Set an username to use their public data"`
	AuthPrivate AuthPrivateCommand `command:"private" description:"Give access to your private data"`
	AuthDestroy AuthDestroyCommand `command:"destroy" description:"Destroy authorization and delete username"`
}
