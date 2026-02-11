package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "obsidian/testutil/keeper"
	"obsidian/x/notary/keeper"
	"obsidian/x/notary/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDocumentMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.NotaryKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateDocument{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateDocument(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetDocument(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestDocumentMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateDocument
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateDocument{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateDocument{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateDocument{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NotaryKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateDocument{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateDocument(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateDocument(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetDocument(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestDocumentMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteDocument
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteDocument{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteDocument{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteDocument{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.NotaryKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateDocument(ctx, &types.MsgCreateDocument{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteDocument(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetDocument(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
