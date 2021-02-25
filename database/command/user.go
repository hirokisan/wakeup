package command

// UserCommand :
type UserCommand struct {
	// store
}

// NewUserCommand :
func NewUserCommand() *UserCommand {
	return &UserCommand{}
}

// OpenEye : 目を開ける
func (c *UserCommand) OpenEye() error {
	// TODO : ユーザの目を開ける(状態変化)
	return nil
}
