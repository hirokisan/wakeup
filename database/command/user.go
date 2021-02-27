package command

import (
	"github.com/hirokisan/go-sample-clean-architecture/store"
)

// UserCommand :
type UserCommand struct {
	userStore store.UserStore
}

// NewUserCommand :
func NewUserCommand(store store.UserStore) *UserCommand {
	return &UserCommand{store}
}

// OpenEye : 目を開ける
func (c *UserCommand) OpenEye(id string) error {
	eye := "open"
	if err := c.userStore.UpdateID(id, store.UpdateUserParam{
		Eye: &eye,
	}); err != nil {
		return err
	}
	return nil
}
