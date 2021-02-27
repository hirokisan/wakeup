package usecase

import (
	"fmt"

	"github.com/hirokisan/go-sample-clean-architecture/database/command"
	"github.com/hirokisan/go-sample-clean-architecture/database/query"
	"github.com/hirokisan/go-sample-clean-architecture/domain"
)

// WakeUpUsecase : 起床
type WakeUpUsecase struct {
	userCommand  command.UserCommand
	alermCommand command.AlermCommand
	bedCommand   command.BedCommand

	alermQuery query.AlermQuery
	bedQuery   query.BedQuery
}

// NewWakeUpUsecase :
func NewWakeUpUsecase(
	userCommand command.UserCommand,
	bedCommand command.BedCommand,
	alermCommand command.AlermCommand,
	alermQuery query.AlermQuery,
	bedQuery query.BedQuery,
) *WakeUpUsecase {
	return &WakeUpUsecase{
		userCommand:  userCommand,
		bedCommand:   bedCommand,
		alermCommand: alermCommand,
		alermQuery:   alermQuery,
		bedQuery:     bedQuery,
	}
}

// WakeUp : 起床する
func (u *WakeUpUsecase) WakeUp(user domain.User) error {
	alerm, err := u.alermQuery.AlermByUserID(user.ID)
	if err != nil {
		return err
	}
	if err := u.stopAlert(alerm); err != nil {
		return err
	}
	if err := u.openEye(user); err != nil {
		return err
	}
	bed, err := u.bedQuery.BedByUserID(user.ID)
	if err != nil {
		return err
	}
	if err := u.getOutOfBed(bed); err != nil {
		return err
	}
	return nil
}

func (u *WakeUpUsecase) stopAlert(alerm domain.Alerm) error {
	if err := u.alermCommand.Stop(alerm.ID); err != nil {
		return fmt.Errorf("stop alerm: %w", err)
	}
	return nil
}

func (u *WakeUpUsecase) openEye(user domain.User) error {
	if err := u.userCommand.OpenEye(user.ID); err != nil {
		return fmt.Errorf("open eye: %w", err)
	}
	return nil
}

func (u *WakeUpUsecase) getOutOfBed(bed domain.Bed) error {
	if err := u.bedCommand.Open(bed.ID); err != nil {
		return fmt.Errorf("open bed: %w", err)
	}
	return nil
}
