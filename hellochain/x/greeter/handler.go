package greeter

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the greeter type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// TODO: Define your msg cases
	case MsgGreet:
		return handleMsgGreet(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName,  msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgGreet does saves the greeting under its recip[ient's address and emits a Greeting event
//TO inett ou keeper function
func handleMsgGreet(ctx sdk.Context, k Keeper)(ctx, msg.ValidatorAddr){
	if err != nil {
		return nil, err
	}
	greeting := NewGreeting(msg.Sender, msg.Body, msg.Recipient)

		keeper.AppendGreeting(msg.Recipient,ctx, greeting)
		kk,.

	// TODO: Define your msg events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddr.String()),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
