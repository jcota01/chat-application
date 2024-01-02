package communicate

type Type uint8

const (
	MsgCont Type = 0
	MsgEnd  Type = 1
	Name    Type = 2
)

type Message struct {
	MsgType Type
	Msg     string
}
