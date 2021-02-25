package command

// AlermCommand :
type AlermCommand struct {
	// store
}

// NewAlermCommand :
func NewAlermCommand() *AlermCommand {
	return &AlermCommand{}
}

// Stop : アラームを停止する
func (c *AlermCommand) Stop() error {
	// TODO : implement
	return nil
}
