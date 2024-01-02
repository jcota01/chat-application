package communicate

func sendName(c Connection, name string) {
	msg := Message{MsgType: Name, Msg: name}

	c.Send(msg)
}
