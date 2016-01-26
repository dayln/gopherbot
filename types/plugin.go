package types

//Plugin struct defines what a plugin needs to have to be recognized by the bot
type Plugin interface {
	Name() string
	Trigger() string
	Action() ActionType
	CmdExec(cmd string)
}

//ActionType is an enum of plugin action types
type ActionType int

const (
	//MSGRES is a message responcse for action type
	MSGRES ActionType = iota
	//USRADD is used to indicate actions when a new user joins the guild
	USRADD
	//EXTMSG is a plugin action that sends a meesage based of an external event
	EXTMSG
	//RGXRES is a reponse based on the content of the message and not a trigger
	RGXRES
)
