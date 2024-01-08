package communicate

type Type uint8

const (
	Msg     Type = 0
	AskName Type = 1
	Name    Type = 2
	Ready   Type = 3
)

type Message struct {
	MsgType Type
	Msg     []byte
}
