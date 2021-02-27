package command

import "github.com/hirokisan/go-sample-clean-architecture/store"

// BedCommand :
type BedCommand struct {
	bedStore store.BedStore
}

// NewBedCommand :
func NewBedCommand(store store.BedStore) *BedCommand {
	return &BedCommand{store}
}

// Open : 空にする
func (c *BedCommand) Open(id string) error {
	empty := true
	if err := c.bedStore.UpdateID(id, store.UpdateBedParam{
		Empty: &empty,
	}); err != nil {
		return err
	}
	return nil
}
