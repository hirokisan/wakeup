package usecase

import (
	"fmt"

	"github.com/hirokisan/go-sample-clean-architecture/database/command"
	"github.com/hirokisan/go-sample-clean-architecture/domain"
)

// WakeUpUsecase : 起床
type WakeUpUsecase struct {
	userCommand  command.UserCommand
	alermCommand command.AlermCommand
	bedCommand   command.BedCommand
}

// NewWakeUpUsecase :
func NewWakeUpUsecase() *WakeUpUsecase {
	return &WakeUpUsecase{}
}

// WakeUp : 起床する
func (u *WakeUpUsecase) WakeUp(user domain.User) error {
	if err := u.stopAlert(); err != nil {
		return err
	}
	if err := u.openEye(); err != nil {
		return err
	}
	if err := u.getOutOfBed(); err != nil {
		return err
	}
	return nil
}

func (u *WakeUpUsecase) stopAlert() error {
	if err := u.alermCommand.Stop(); err != nil {
		return fmt.Errorf("open eye: %w", err)
	}
	return nil
}

func (u *WakeUpUsecase) openEye() error {
	if err := u.userCommand.OpenEye(); err != nil {
		return fmt.Errorf("open eye: %w", err)
	}
	return nil
}

func (u *WakeUpUsecase) getOutOfBed() error {
	if err := u.bedCommand.Open(); err != nil {
		return fmt.Errorf("open bed: %w", err)
	}
	return nil
}
