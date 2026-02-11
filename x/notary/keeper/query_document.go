package keeper

import (
	"context"

	"obsidian/x/notary/types"

	"cosmossdk.io/store/prefix"
        "github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DocumentAll — Поиск всех документов (с фильтром по владельцу)
func (k Keeper) DocumentAll(goCtx context.Context, req *types.QueryAllDocumentRequest) (*types.QueryAllDocumentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var documents []types.Document
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := prefix.NewStore(runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)), types.KeyPrefix(types.DocumentKeyPrefix))

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var document types.Document
		if err := k.cdc.Unmarshal(value, &document); err != nil {
			return err
		}

		// Если в запросе передан OwnerAddress, фильтруем. Если пустой — выдаем всё.
		if req.OwnerAddress == "" || document.Owner == req.OwnerAddress {
			documents = append(documents, document)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDocumentResponse{Document: documents, Pagination: pageRes}, nil
}

// Document — Поиск одного документа по его индексу (ID)
func (k Keeper) Document(goCtx context.Context, req *types.QueryGetDocumentRequest) (*types.QueryGetDocumentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetDocument(ctx, req.Index)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDocumentResponse{Document: val}, nil
}
