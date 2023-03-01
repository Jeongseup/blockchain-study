package types

import (
	"errors"
	"fmt"

	"github.com/Jeongseup/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedGame StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
    black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
    return black, sdkerrors.Wrapf(errBlack, errInvalidBlack.Error(), storedGame.Black)
}

func (storedGame StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(storedGame.Red)
	return red, sdkerrors.Wrapf(errRed, errInvalidRed.Error(), storedGame.Red)
}


func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
    board, errBoard := rules.Parse(storedGame.Board)
    if errBoard != nil {
        return nil, sdkerrors.Wrapf(errBoard, errGameNotParseable.Error())
    }
    board.Turn = rules.StringPieces[storedGame.Turn].Player
    if board.Turn.Color == "" {
        return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("Turn: %s", storedGame.Turn)), errGameNotParseable.Error())
    }
    return board, nil
}

func (storedGame StoredGame) Validate() (err error) {
    _, err = storedGame.GetBlackAddress()
    if err != nil {
        return err
    }
    _, err = storedGame.GetRedAddress()
    if err != nil {
        return err
    }
    _, err = storedGame.ParseGame()
    return err
}

var (
    errInvalidBlack     = sdkerrors.Register(ModuleName, 1100, "black address is invalid: %s")
    errInvalidRed       = sdkerrors.Register(ModuleName, 1101, "red address is invalid: %s")
    errGameNotParseable = sdkerrors.Register(ModuleName, 1102, "game cannot be parsed")
)
