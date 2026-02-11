package keeper

import (
	"context"

	"obsidian/x/notary/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// CreateDocument — Регистрация нового документа с автоматическим сжиганием 25% монет
func (k msgServer) CreateDocument(goCtx context.Context, msg *types.MsgCreateDocument) (*types.MsgCreateDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Проверяем, не занят ли уже этот индекс (ID документа)
	_, isFound := k.GetDocument(ctx, msg.Index)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// 2. Логика сжигания (BURN) 1250 uobs (25% от стандартной комиссии 5000)
	feeAmount := sdk.NewCoins(sdk.NewInt64Coin("uobs", 1250))

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	// Переводим монеты со счета пользователя в модуль для последующего сжигания
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, feeAmount)
	if err != nil {
		// Если на балансе нет 1250 uobs помимо основной комиссии, транзакция отклонится
		return nil, errorsmod.Wrap(err, "insufficient funds for notary burn fee (1250 uobs)")
	}

	// Сжигаем монеты (уничтожаем их навсегда)
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, feeAmount)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to burn coins")
	}

	// 3. Формируем запись документа
	var document = types.Document{
		Creator:   msg.Creator,
		Index:     msg.Index,
		FileHash:  msg.FileHash,
		// Блокчейн сам назначает владельца и время
		Owner:     msg.Creator,
		Timestamp: int32(ctx.BlockTime().Unix()),
	}

	// 4. Сохраняем в базу данных
	k.SetDocument(ctx, document)

	return &types.MsgCreateDocumentResponse{}, nil
}

// UpdateDocument — Заблокировано для безопасности
func (k msgServer) UpdateDocument(goCtx context.Context, msg *types.MsgUpdateDocument) (*types.MsgUpdateDocumentResponse, error) {
	return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Obsidian records are immutable")
}

// DeleteDocument — Заблокировано для безопасности
func (k msgServer) DeleteDocument(goCtx context.Context, msg *types.MsgDeleteDocument) (*types.MsgDeleteDocumentResponse, error) {
	return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Obsidian records are eternal")
}
