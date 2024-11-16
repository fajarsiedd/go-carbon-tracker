package trip

import (
	"errors"
	"go-carbon-tracker/entities"
	tripMocks "go-carbon-tracker/mocks/repositories/trip"
	vehicleMocks "go-carbon-tracker/mocks/repositories/vehicle"
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
		want    []entities.Trip
		wantErr bool
		repoErr bool
	}{
		// Test Cases
		{
			name: "success_get_all_trips",
			args: args{
				filter: entities.Filter{
					Page:  1,
					Limit: 10,
				},
			},
			want: []entities.Trip{
				{
					Base: entities.Base{
						ID: "trip1",
					},
					DistanceKM:     12,
					CarbonEmission: 333.333,
					Tips:           "Great tips!",
				},
			},
			wantErr: false,
			repoErr: false,
		},
		{
			name:    "failed_get_all_trips",
			args:    args{},
			want:    []entities.Trip{},
			wantErr: true,
			repoErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tripRepo := &tripMocks.TripRepository{}
			vehicleRepo := &vehicleMocks.VehicleRepository{}
			usecase := NewTripUsecase(tripRepo, vehicleRepo)

			if !tt.repoErr {
				tripRepo.On("GetAll", mock.Anything).Return(tt.want, entities.Pagination{}, nil)
			} else {
				tripRepo.On("GetAll", mock.Anything).Return(tt.want, entities.Pagination{}, errors.New("database error"))
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
		want    entities.Trip
		wantErr bool
		repoErr bool
	}{
		// Test Cases
		{
			name: "success_get_by_id",
			args: args{
				id: "trip1",
			},
			want: entities.Trip{
				Base: entities.Base{
					ID: "trip1",
				},
				DistanceKM:     12,
				CarbonEmission: 333.333,
				Tips:           "Great tips!",
			},
			wantErr: false,
			repoErr: false,
		},
		{
			name: "failed_get_by_id",
			args: args{
				id: "trip1",
			},
			want:    entities.Trip{},
			wantErr: true,
			repoErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tripRepo := &tripMocks.TripRepository{}
			vehicleRepo := &vehicleMocks.VehicleRepository{}
			usecase := NewTripUsecase(tripRepo, vehicleRepo)

			if !tt.repoErr {
				tripRepo.On("GetByID", mock.Anything).Return(tt.want, nil)
			} else {
				tripRepo.On("GetByID", mock.Anything).Return(tt.want, errors.New("database error"))
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
		trip entities.Trip
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Trip
		wantErr bool
		repoErr bool
	}{
		// Test Cases
		{
			name: "failed_create",
			args: args{
				trip: entities.Trip{
					VehicleID: "vehicle1",
				},
			},
			want:    entities.Trip{},
			wantErr: true,
			repoErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tripRepo := &tripMocks.TripRepository{}
			vehicleRepo := &vehicleMocks.VehicleRepository{}
			usecase := NewTripUsecase(tripRepo, vehicleRepo)

			if !tt.repoErr {
				vehicleRepo.On("GetByID", tt.args.trip.VehicleID).Return(entities.Vehicle{}, nil)
				tripRepo.On("Create", mock.Anything).Return(tt.want, nil)
			} else {
				vehicleRepo.On("GetByID", tt.args.trip.VehicleID).Return(entities.Vehicle{}, errors.New("database error"))
				tripRepo.On("Create", mock.Anything).Return(tt.want, errors.New("database error"))
			}

			got, err := usecase.Create(tt.args.trip)

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
		trip entities.Trip
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Trip
		wantErr bool
		repoErr bool
	}{
		// Test Cases
		{
			name: "failed_create",
			args: args{
				trip: entities.Trip{
					Base: entities.Base{
						ID: "trip1",
					},
				},
			},
			want:    entities.Trip{},
			wantErr: true,
			repoErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tripRepo := &tripMocks.TripRepository{}
			vehicleRepo := &vehicleMocks.VehicleRepository{}
			usecase := NewTripUsecase(tripRepo, vehicleRepo)

			if !tt.repoErr {
				tripRepo.On("GetByID", tt.args.trip.ID).Return(tt.want, nil)
				tripRepo.On("Update", mock.Anything).Return(tt.want, nil)
			} else {
				tripRepo.On("GetByID", tt.args.trip.ID).Return(tt.want, errors.New("database error"))
				tripRepo.On("Update", mock.Anything).Return(tt.want, errors.New("database error"))
			}

			got, err := usecase.Update(tt.args.trip)

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
		want    entities.Trip
		wantErr bool
		repoErr bool
	}{
		// Test Cases
		{
			name: "success_delete",
			args: args{
				id: "trip1",
			},
			want:    entities.Trip{},
			wantErr: false,
			repoErr: false,
		},
		{
			name: "failed_delete",
			args: args{
				id: "trip1",
			},
			want:    entities.Trip{},
			wantErr: true,
			repoErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tripRepo := &tripMocks.TripRepository{}
			vehicleRepo := &vehicleMocks.VehicleRepository{}
			usecase := NewTripUsecase(tripRepo, vehicleRepo)

			if !tt.repoErr {
				tripRepo.On("Delete", mock.Anything).Return(nil)
			} else {
				tripRepo.On("Delete", mock.Anything).Return(errors.New("database error"))
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
