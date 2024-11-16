package vehicle

import (
	"errors"
	"go-carbon-tracker/constants/enums"
	"go-carbon-tracker/entities"
	mocks "go-carbon-tracker/mocks/repositories/vehicle"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsecase_GetAll(t *testing.T) {
	type args struct {
		filter entities.Filter
	}
	tests := []struct {
		name    string
		args    args
		want    []entities.Vehicle
		wantErr bool
		repoErr bool
	}{
		// Test Cases
		{
			name: "success_get_all_vehicles",
			args: args{
				filter: entities.Filter{
					Page:  1,
					Limit: 10,
				},
			},
			want: []entities.Vehicle{
				{
					Base: entities.Base{
						ID: "vehicle1",
					},
					Name:           "Honda",
					VehicleType:    enums.MOTOR,
					FuelType:       enums.PREMIUM,
					EmissionFactor: 123.45,
				},
			},
			wantErr: false,
			repoErr: false,
		},
		{
			name:    "failed_get_all_vehicles",
			args:    args{},
			want:    []entities.Vehicle{},
			wantErr: true,
			repoErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.VehicleRepository{}
			usecase := NewVehicleUsecase(repo)

			if !tt.repoErr {
				repo.On("GetAll", mock.Anything).Return(tt.want, entities.Pagination{}, nil)
			} else {
				repo.On("GetAll", mock.Anything).Return(tt.want, entities.Pagination{}, errors.New("database error"))
			}

			got, _, err := usecase.GetAll(tt.args.filter)

			if err != nil && !tt.wantErr {
				t.Errorf("usecase.GetAll() error = %v, wantErr %v", err, tt.wantErr)
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

func TestUsecase_GetByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Vehicle
		wantErr bool
		repoErr bool
	}{
		// Test Cases
		{
			name: "success_get_by_id",
			args: args{
				id: "vehicle1",
			},
			want: entities.Vehicle{
				Base: entities.Base{
					ID: "vehicle1",
				},
				Name:           "Honda",
				VehicleType:    enums.MOTOR,
				FuelType:       enums.PREMIUM,
				EmissionFactor: 123.45,
			},
			wantErr: false,
			repoErr: false,
		},
		{
			name: "failed_get_by_id",
			args: args{
				id: "vehicle1",
			},
			want:    entities.Vehicle{},
			wantErr: true,
			repoErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.VehicleRepository{}
			usecase := NewVehicleUsecase(repo)

			if !tt.repoErr {
				repo.On("GetByID", mock.Anything).Return(tt.want, nil)
			} else {
				repo.On("GetByID", mock.Anything).Return(tt.want, errors.New("database error"))
			}

			got, err := usecase.GetByID(tt.args.id)

			if err != nil && !tt.wantErr {
				t.Errorf("usecase.GetByID() error = %v, wantErr %v", err, tt.wantErr)
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

func TestUsecase_Create(t *testing.T) {
	type args struct {
		vehicle entities.Vehicle
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Vehicle
		wantErr bool
		repoErr bool
	}{
		// Test Cases
		{
			name: "success_create",
			args: args{
				vehicle: entities.Vehicle{
					Name:        "Honda",
					VehicleType: enums.MOTOR,
					FuelType:    enums.PREMIUM,
				},
			},
			want: entities.Vehicle{
				Base: entities.Base{
					ID: "vehicle1",
				},
				Name:           "Honda",
				VehicleType:    enums.MOTOR,
				FuelType:       enums.PREMIUM,
				EmissionFactor: 69.103076,
			},
			wantErr: false,
			repoErr: false,
		},
		{
			name: "failed_create",
			args: args{
				vehicle: entities.Vehicle{
					Name:        "Honda",
					VehicleType: enums.MOTOR,
					FuelType:    enums.PREMIUM,
				},
			},
			want:    entities.Vehicle{},
			wantErr: true,
			repoErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.VehicleRepository{}
			usecase := NewVehicleUsecase(repo)

			if !tt.repoErr {
				repo.On("Create", mock.Anything).Return(tt.want, nil)
			} else {
				repo.On("Create", mock.Anything).Return(tt.want, errors.New("database error"))
			}

			got, err := usecase.Create(tt.args.vehicle)

			if err != nil && !tt.wantErr {
				t.Errorf("usecase.Create() error = %v, wantErr %v", err, tt.wantErr)
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

func TestUsecase_Update(t *testing.T) {
	type args struct {
		vehicle entities.Vehicle
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Vehicle
		wantErr bool
		repoErr bool
	}{
		// Test Cases
		{
			name: "success_update",
			args: args{
				vehicle: entities.Vehicle{
					Name:        "Honda 2",
					VehicleType: enums.MOTOR,
					FuelType:    enums.PREMIUM,
				},
			},
			want: entities.Vehicle{
				Base: entities.Base{
					ID: "vehicle1",
				},
				Name:           "Honda 2",
				VehicleType:    enums.MOTOR,
				FuelType:       enums.PREMIUM,
				EmissionFactor: 69.103076,
			},
			wantErr: false,
			repoErr: false,
		},
		{
			name: "failed_create",
			args: args{
				vehicle: entities.Vehicle{
					Name:        "Honda",
					VehicleType: enums.MOTOR,
					FuelType:    enums.PREMIUM,
				},
			},
			want:    entities.Vehicle{},
			wantErr: true,
			repoErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.VehicleRepository{}
			usecase := NewVehicleUsecase(repo)

			if !tt.repoErr {
				repo.On("Update", mock.Anything).Return(tt.want, nil)
			} else {
				repo.On("Update", mock.Anything).Return(tt.want, errors.New("database error"))
			}

			got, err := usecase.Update(tt.args.vehicle)

			if err != nil && !tt.wantErr {
				t.Errorf("usecase.Update() error = %v, wantErr %v", err, tt.wantErr)
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

func TestUsecase_Delete(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Vehicle
		wantErr bool
		repoErr bool
	}{
		// Test Cases
		{
			name: "success_delete",
			args: args{
				id: "vehicle1",
			},
			want:    entities.Vehicle{},
			wantErr: false,
			repoErr: false,
		},
		{
			name: "failed_delete",
			args: args{
				id: "vehicle1",
			},
			want:    entities.Vehicle{},
			wantErr: true,
			repoErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.VehicleRepository{}
			usecase := NewVehicleUsecase(repo)

			if !tt.repoErr {
				repo.On("Delete", mock.Anything).Return(nil)
			} else {
				repo.On("Delete", mock.Anything).Return(errors.New("database error"))
			}

			err := usecase.Delete(tt.args.id)

			if err != nil && !tt.wantErr {
				t.Errorf("usecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
