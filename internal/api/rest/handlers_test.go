package rest_test

import (
	"testing"

	"github.com/go-openapi/swag"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zergslaw/users/internal/api/rest"
	"github.com/zergslaw/users/internal/api/rest/generated/client/operations"
	"github.com/zergslaw/users/internal/api/rest/generated/models"
	"github.com/zergslaw/users/internal/app"
)

func TestServiceVerificationEmail(t *testing.T) {
	t.Parallel()

	_, shutdown, mockApp, client := testNewServer(t)
	defer shutdown()

	testCases := []struct {
		name   string
		email  string
		appErr error
		want   *models.Error
	}{
		{"success", notExistEmail, nil, nil},
		{"exist", email, app.ErrEmailExist, APIError("email exist")},
		{"any error", email, errAny, APIError("internal error")},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockApp.EXPECT().VerificationEmail(gomock.Any(), tc.email).Return(tc.appErr)

			params := operations.NewVerificationEmailParams().WithEmail(models.Email(tc.email))
			_, err := client.Operations.VerificationEmail(params)
			assert.Equal(t, tc.want, errPayload(err))
		})
	}
}

func TestServiceVerificationUsername(t *testing.T) {
	t.Parallel()

	_, shutdown, mockApp, client := testNewServer(t)
	defer shutdown()

	testCases := []struct {
		name     string
		username string
		appErr   error
		want     *models.Error
	}{
		{"success", notExistUsername, nil, nil},
		{"exist", username, app.ErrUsernameExist, APIError("username exist")},
		{"any error", username, errAny, APIError("internal error")},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockApp.EXPECT().VerificationUsername(gomock.Any(), tc.username).Return(tc.appErr)

			params := operations.NewVerificationUsernameParams().WithUsername(models.Username(tc.username))
			_, err := client.Operations.VerificationUsername(params)
			assert.Equal(t, tc.want, errPayload(err))
		})
	}
}

func TestServiceCreateUser(t *testing.T) {
	t.Parallel()

	_, shutdown, mockApp, client := testNewServer(t)
	defer shutdown()

	testCases := []struct {
		name     string
		email    string
		username string
		password string
		user     *app.User
		token    app.AuthToken
		appErr   error
		want     *models.User
		wantErr  *models.Error
	}{
		{"success", email, username, password,
			&user, authToken, nil, restUser, nil},
		{"email exist", email, username, password,
			nil, "", app.ErrEmailExist, nil, APIError("email exist")},
		{"username exist", email, username, password,
			nil, "", app.ErrUsernameExist, nil, APIError("username exist")},
		{"internal error", email, username, password,
			nil, "", errAny, nil, APIError("internal error")},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockApp.EXPECT().
				CreateUser(gomock.Any(), tc.email, tc.username, tc.password, origin).
				Return(tc.user, tc.token, tc.appErr)

			params := operations.NewCreateUserParams().WithArgs(&models.CreateUserParams{
				Email:    models.Email(tc.email),
				Password: models.Password(tc.password),
				Username: models.Username(tc.username),
			})

			res, err := client.Operations.CreateUser(params)
			if tc.wantErr == nil {
				assert.Nil(t, err)
				assert.Equal(t, rest.User(tc.user), res.Payload)
			} else {
				assert.Nil(t, res)
				assert.Equal(t, tc.wantErr, errPayload(err))
			}
		})
	}
}

func TestServiceLogin(t *testing.T) {
	t.Parallel()

	_, shutdown, mockApp, client := testNewServer(t)
	defer shutdown()

	testCases := []struct {
		name     string
		email    string
		password string
		user     *app.User
		token    app.AuthToken
		appErr   error
		want     *models.User
		wantErr  *models.Error
	}{
		{"success", email, password,
			&user, authToken, nil, restUser, nil},
		{"email not found", email, password,
			nil, "", app.ErrNotFound, nil, APIError("not found")},
		{"not valid password", email, password,
			nil, "", app.ErrNotValidPassword, nil, APIError("not valid password")},
		{"internal error", email, password,
			nil, "", errAny, nil, APIError("internal error")},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockApp.EXPECT().
				Login(gomock.Any(), tc.email, tc.password, origin).
				Return(tc.user, tc.token, tc.appErr)

			params := operations.NewLoginParams().WithArgs(&models.LoginParam{
				Email:    models.Email(tc.email),
				Password: models.Password(tc.password),
			})

			res, err := client.Operations.Login(params)
			if tc.wantErr == nil {
				assert.Nil(t, err)
				assert.Equal(t, rest.User(tc.user), res.Payload)
			} else {
				assert.Nil(t, res)
				assert.Equal(t, tc.wantErr, errPayload(err))
			}
		})
	}
}

func TestServiceLogout(t *testing.T) {
	t.Parallel()

	_, shutdown, mockApp, client := testNewServer(t)
	defer shutdown()

	testCases := []struct {
		name   string
		appErr error
		want   *models.Error
	}{
		{"success", nil, nil},
		{"any error", errAny, APIError("internal error")},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockApp.EXPECT().Logout(gomock.Any(), authUser).Return(tc.appErr)

			params := operations.NewLogoutParams()
			_, err := client.Operations.Logout(params, apiKeyAuth)
			assert.Equal(t, tc.want, errPayload(err))
		})
	}
}

func TestServiceGetUser(t *testing.T) {
	t.Parallel()

	_, shutdown, mockApp, client := testNewServer(t)
	defer shutdown()

	testCases := []struct {
		name    string
		user    *app.User
		appErr  error
		want    *models.User
		wantErr *models.Error
	}{
		{"success", &user, nil, restUser, nil},
		{"not found", nil, app.ErrNotFound, nil, APIError("not found")},
		{"any error", nil, errAny, nil, APIError("internal error")},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockApp.EXPECT().User(gomock.Any(), authUser, authUser.ID).Return(tc.user, tc.appErr)

			params := operations.NewGetUserParams().WithID(int32(user.ID))
			res, err := client.Operations.GetUser(params, apiKeyAuth)
			if tc.wantErr == nil {
				assert.Nil(t, err)
				assert.Equal(t, tc.want, res.Payload)
			} else {
				assert.Nil(t, res)
				assert.Equal(t, tc.wantErr, errPayload(err))
			}
		})
	}
}

func TestServiceDeleteUser(t *testing.T) {
	t.Parallel()

	_, shutdown, mockApp, client := testNewServer(t)
	defer shutdown()

	testCases := []struct {
		name   string
		appErr error
		want   *models.Error
	}{
		{"success", nil, nil},
		{"any error", errAny, APIError("internal error")},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockApp.EXPECT().DeleteUser(gomock.Any(), authUser).Return(tc.appErr)

			params := operations.NewDeleteUserParams()
			_, err := client.Operations.DeleteUser(params, apiKeyAuth)
			assert.Equal(t, tc.want, errPayload(err))
		})
	}
}

func TestServiceUpdatePassword(t *testing.T) {
	t.Parallel()

	_, shutdown, mockApp, client := testNewServer(t)
	defer shutdown()

	testCases := []struct {
		name             string
		oldPass, newPass string
		appErr           error
		want             *models.Error
	}{
		{"success", password, "NewPassword", nil, nil},
		{"not valid password", "notCorrectPass", "NewPassword", app.ErrNotValidPassword, APIError("not valid password")},
		{"any error", password, "NewPassword", errAny, APIError("internal error")},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockApp.EXPECT().UpdatePassword(gomock.Any(), authUser, tc.oldPass, tc.newPass).Return(tc.appErr)

			params := operations.NewUpdatePasswordParams().WithArgs(&models.UpdatePassword{
				New: models.Password(tc.newPass),
				Old: models.Password(tc.oldPass),
			})
			_, err := client.Operations.UpdatePassword(params, apiKeyAuth)
			assert.Equal(t, tc.want, errPayload(err))
		})
	}
}

func TestServiceUpdateUsername(t *testing.T) {
	t.Parallel()

	_, shutdown, mockApp, client := testNewServer(t)
	defer shutdown()

	testCases := []struct {
		name     string
		username string
		appErr   error
		want     *models.Error
	}{
		{"success", username, nil, nil},
		{"username exist", username, app.ErrUsernameExist, APIError("username exist")},
		{"username not different", username, app.ErrUsernameNeedDifferentiate, APIError("username need to differentiate")},
		{"any error", username, errAny, APIError("internal error")},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockApp.EXPECT().UpdateUsername(gomock.Any(), authUser, tc.username).Return(tc.appErr)

			params := operations.NewUpdateUsernameParams().WithUsername(models.Username(tc.username))

			_, err := client.Operations.UpdateUsername(params, apiKeyAuth)
			assert.Equal(t, tc.want, errPayload(err))
		})
	}
}

func TestServiceUpdateEmail(t *testing.T) {
	t.Parallel()

	_, shutdown, mockApp, client := testNewServer(t)
	defer shutdown()

	testCases := []struct {
		name   string
		email  string
		appErr error
		want   *models.Error
	}{
		{"success", email, nil, nil},
		{"email exist", email, app.ErrEmailExist, APIError("email exist")},
		{"email not different", email, app.ErrEmailNeedDifferentiate, APIError("email need to differentiate")},
		{"any error", email, errAny, APIError("internal error")},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockApp.EXPECT().UpdateEmail(gomock.Any(), authUser, tc.email).Return(tc.appErr)

			params := operations.NewUpdateEmailParams().WithEmail(models.Email(tc.email))

			_, err := client.Operations.UpdateEmail(params, apiKeyAuth)
			assert.Equal(t, tc.want, errPayload(err))
		})
	}
}

func TestServiceGetUsers(t *testing.T) {
	t.Parallel()

	_, shutdown, mockApp, client := testNewServer(t)
	defer shutdown()

	testCases := []struct {
		name      string
		username  string
		users     []app.User
		appErr    error
		want      []*models.User
		wantTotal int32
		wantErr   *models.Error
	}{
		{"success", username, []app.User{user}, nil,
			rest.Users([]app.User{user}), 1, nil},
		{"any error", username, nil, errAny,
			nil, 0, APIError("internal error")},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockApp.EXPECT().
				ListUserByUsername(gomock.Any(), authUser, tc.username, app.Page{Limit: 10}).
				Return(tc.users, len(tc.users), tc.appErr)

			params := operations.NewGetUsersParams().WithArgs(&models.ListUsersParams{
				Pagination: &models.Pagination{
					Limit:  swag.Int32(10),
					Offset: swag.Int32(0),
				},
				Username: models.Username(tc.username),
			})

			res, err := client.Operations.GetUsers(params, apiKeyAuth)
			if tc.wantErr == nil {
				assert.Nil(t, err)
				assert.Equal(t, &operations.GetUsersOKBody{
					Total: swag.Int32(int32(len(tc.users))),
					Users: tc.want,
				}, res.Payload)
			} else {
				assert.Nil(t, res)
				assert.Equal(t, tc.wantErr, errPayload(err))
			}
		})
	}
}