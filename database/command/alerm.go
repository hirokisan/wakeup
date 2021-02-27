package command

import (
	"github.com/hirokisan/go-sample-clean-architecture/store"
)

// AlermCommand :
type AlermCommand struct {
	alermStore store.AlermStore
}

// NewAlermCommand :
func NewAlermCommand(store store.AlermStore) *AlermCommand {
	return &AlermCommand{store}
}

// Stop : アラームを停止する
func (c *AlermCommand) Stop(id string) error {
	stop := true
	if err := c.alermStore.UpdateID(id, store.UpdateAlermParam{
		Stop: &stop,
	}); err != nil {
		return err
	}
	return nil
}
