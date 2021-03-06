package rpc_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zergslaw/boilerplate/internal/api/rpc/pb"
	"github.com/zergslaw/boilerplate/internal/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestService_GetUserByAuthToken(t *testing.T) {
	t.Parallel()

	c, mockApp, shutdown := testNew(t)
	defer shutdown()

	errNotFound := status.Error(codes.NotFound, app.ErrNotFound.Error())
	errDeadline := status.Error(codes.DeadlineExceeded, context.DeadlineExceeded.Error())
	errCanceled := status.Error(codes.Canceled, context.Canceled.Error())
	errInternal := status.Error(codes.Internal, errAny.Error())

	testCases := []struct {
		name    string
		auth    *app.AuthUser
		want    *pb.User
		appErr  error
		wantErr error
	}{
		{"success", &appUser, &rpcUser, nil, nil},
		{"not found", nil, nil, app.ErrNotFound, errNotFound},
		{"deadline", nil, nil, context.DeadlineExceeded, errDeadline},
		{"canceled", nil, nil, context.Canceled, errCanceled},
		{"internal", nil, nil, errAny, errInternal},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockApp.EXPECT().UserByAuthToken(gomock.Any(), app.AuthToken(token)).Return(tc.auth, tc.appErr)

			res, err := c.GetUserByAuthToken(ctx, &pb.AuthInfo{Token: token})
			if tc.wantErr == nil {
				assert.Nil(t, err)
				assert.Equal(t, app.User{
					ID:    app.UserID(res.Id),
					Email: res.Email,
					Name:  res.Username,
				}, app.User{
					ID:    tc.auth.ID,
					Email: tc.auth.Email,
					Name:  tc.auth.Name,
				})
			} else {
				assert.Nil(t, res)
				assert.Equal(t, tc.wantErr, err)
			}
		})
	}
}
