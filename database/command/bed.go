package command

// BedCommand :
type BedCommand struct {
	// store
}

// NewBedCommand :
func NewBedCommand() *BedCommand {
	return &BedCommand{}
}

// Open : 空にする
func (c *BedCommand) Open(ID string) error {
	// TODO : implement
	return nil
}
