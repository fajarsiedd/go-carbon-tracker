package auth

import (
	"errors"
	"go-carbon-tracker/entities"
	"go-carbon-tracker/middlewares"
	mocks "go-carbon-tracker/mocks/repositories/auth"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsecase_Login(t *testing.T) {
	type args struct {
		user entities.User
	}
	tests := []struct {
		name    string
		args    args
		want    entities.User
		wantErr bool
		repoErr bool
	}{
		// Test Cases
		{
			name: "success_login",
			args: args{
				user: entities.User{
					Email:    "sample@mail.com",
					Password: "123456",
				},
			},
			want: entities.User{
				Email:    "sample@mail.com",
				Password: "$argon2id$v=19$m=65536,t=3,p=2$3WE6AyYGQyS/MfbhOgs/mg$9AnBK2OVdaxlCMLiL+RHZX3IJ5AtK8tjzoL1fthsE5M",
			},
			wantErr: false,
			repoErr: false,
		},
		{
			name: "failed_login_invalid_email",
			args: args{
				user: entities.User{
					Email:    "wrong@mail.com",
					Password: "123456",
				},
			},
			want:    entities.User{},
			wantErr: true,
			repoErr: true,
		},
		{
			name: "failed_login_invalid_password",
			args: args{
				user: entities.User{
					Email:    "sample@mail.com",
					Password: "wrongpassword",
				},
			},
			want: entities.User{
				Email:    "sample@mail.com",
				Password: "$argon2id$v=19$m=65536,t=3,p=2$3WE6AyYGQyS/MfbhOgs/mg$9AnBK2OVdaxlCMLiL+RHZX3IJ5AtK8tjzoL1fthsE5M",
			},
			wantErr: true,
			repoErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.AuthRepository{}
			usecase := NewAuthUsecase(repo, &middlewares.JWTConfig{})

			if !tt.repoErr {
				repo.On("Login", mock.Anything).Return(tt.want, nil)
			} else {
				repo.On("Login", mock.Anything).Return(tt.want, errors.New("database error"))
			}

			got, err := usecase.Login(tt.args.user)
			tt.want.Token = got.Token
			if err != nil && !tt.wantErr {
				t.Errorf("usecase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, tt.want, got)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestUsecase_Register(t *testing.T) {
	type args struct {
		user entities.User
	}
	tests := []struct {
		name    string
		args    args
		want    entities.User
		wantErr bool
		repoErr bool
	}{
		// Test Cases
		{
			name: "success_register",
			args: args{
				user: entities.User{
					Name:     "sample",
					Email:    "sample@mail.com",
					Password: "123456",
				},
			},
			want: entities.User{
				Name:     "sample",
				Email:    "sample@mail.com",
				Password: "123456",
			},
			wantErr: false,
			repoErr: false,
		},
		{
			name: "failed_register",
			args: args{
				user: entities.User{
					Name:     "sample",
					Email:    "sample@mail.com",
					Password: "123456",
				},
			},
			want:    entities.User{},
			wantErr: true,
			repoErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.AuthRepository{}
			usecase := NewAuthUsecase(repo, &middlewares.JWTConfig{})

			if !tt.repoErr {
				repo.On("Register", mock.Anything).Return(tt.want, nil)
			} else {
				repo.On("Register", mock.Anything).Return(tt.want, errors.New("database error"))
			}

			got, err := usecase.Register(tt.args.user)
			if err != nil && !tt.wantErr {
				t.Errorf("usecase.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.Equal(t, tt.want, got)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
