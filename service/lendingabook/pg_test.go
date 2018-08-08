package lendingabook

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/jinzhu/gorm"

	testutil "github.com/ntp13495/example-go/config/database/pg/util"
	"github.com/ntp13495/example-go/domain"
)

func TestNewPGService(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPGService(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPGService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgService_Create(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	lendingbook := domain.LendingBooks{}
	err = testDB.Create(&lendingbook).Error
	if err != nil {
		t.Fatalf("Faied to create lendingbook by error %v", err)
	}

	user := domain.User{}
	err = testDB.Create(&user).Error
	if err != nil {
		t.Fatalf("Failed to create user by error %v", err)
	}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	type args struct {
		p *domain.LendingBooks
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Create a lendingbook: successful!",
			args: args{
				&domain.LendingBooks{
					Model:  domain.Model{ID: lendingbook.ID},
					BookID: book.ID,
					UserID: user.ID,
					From:   time.Date(2018, 8, 7, 17, 22, 30, 123456789, time.UTC),
					To:     time.Date(2018, 8, 8, 17, 22, 30, 123456789, time.UTC),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			if err := s.Create(context.Background(), tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("pgService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pgService_Update(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	lendingbook := domain.LendingBooks{}
	err = testDB.Create(&lendingbook).Error
	if err != nil {
		t.Fatalf("Faied to create lendingbook by error %v", err)
	}

	user := domain.User{}
	err = testDB.Create(&user).Error
	if err != nil {
		t.Fatalf("Failed to create user by error %v", err)
	}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	fakeLendingBookID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.LendingBooks
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.LendingBooks
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "Update a lendingbook: successful!",
			args: args{
				&domain.LendingBooks{
					Model:  domain.Model{ID: lendingbook.ID},
					BookID: book.ID,
					UserID: user.ID,
					From:   time.Date(2018, 8, 7, 17, 22, 30, 123456789, time.UTC),
					To:     time.Date(2018, 8, 8, 17, 22, 30, 123456789, time.UTC),
				},
			},
		},
		{
			name: "Update a lendingbook: failed!",
			args: args{
				&domain.LendingBooks{
					Model:  domain.Model{ID: fakeLendingBookID},
					BookID: book.ID,
					UserID: user.ID,
					From:   time.Date(2018, 8, 7, 17, 22, 30, 123456789, time.UTC),
					To:     time.Date(2018, 8, 8, 17, 22, 30, 123456789, time.UTC),
				},
			},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			_, err := s.Update(context.Background(), tt.args.p)
			if err != nil && err != tt.wantErr {
				t.Errorf("pgService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.wantErr != nil {
				t.Errorf("pgService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_pgService_Find(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	lendingbook := domain.LendingBooks{}
	err = testDB.Create(&lendingbook).Error
	if err != nil {
		t.Fatalf("Faied to create lendingbook by error %v", err)
	}

	user := domain.User{}
	err = testDB.Create(&user).Error
	if err != nil {
		t.Fatalf("Failed to create user by error %v", err)
	}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	fakeLendingBookID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.LendingBooks
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.LendingBooks
		wantErr error
	}{
		{
			name: "Find a lendingbook: successful!",
			args: args{
				&domain.LendingBooks{
					Model: domain.Model{ID: lendingbook.ID},
				},
			},
			want: &domain.LendingBooks{
				Model: domain.Model{ID: lendingbook.ID},
			},
		},
		{
			name: "Find a lendingbook: failed!",
			args: args{
				&domain.LendingBooks{
					Model: domain.Model{ID: fakeLendingBookID},
				},
			},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			got, err := s.Find(context.Background(), tt.args.p)
			if err != nil && err != tt.wantErr {
				t.Errorf("pgService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.wantErr != nil {
				t.Errorf("pgService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.ID.String() != tt.want.ID.String() {
				t.Errorf("pgService.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgService_FindAll(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		in0 context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.LendingBooks
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: tt.fields.db,
			}
			got, err := s.FindAll(tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("pgService.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pgService.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pgService_Delete(t *testing.T) {
	t.Parallel()
	testDB, _, cleanup := testutil.CreateTestDatabase(t)
	defer cleanup()
	err := testutil.MigrateTables(testDB)
	if err != nil {
		t.Fatalf("Failed to migrate table by error %v", err)
	}

	lendingbook := domain.LendingBooks{}
	err = testDB.Create(&lendingbook).Error
	if err != nil {
		t.Fatalf("Faied to create lendingbook by error %v", err)
	}

	user := domain.User{}
	err = testDB.Create(&user).Error
	if err != nil {
		t.Fatalf("Failed to create user by error %v", err)
	}

	book := domain.Book{}
	err = testDB.Create(&book).Error
	if err != nil {
		t.Fatalf("Failed to create book by error %v", err)
	}

	fakeLendingBookID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	type args struct {
		p *domain.LendingBooks
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
		{
			name: "Update a lendingbook: successful!",
			args: args{
				&domain.LendingBooks{
					Model:  domain.Model{ID: lendingbook.ID},
					BookID: book.ID,
					UserID: user.ID,
					From:   time.Date(2018, 8, 7, 17, 22, 30, 123456789, time.UTC),
					To:     time.Date(2018, 8, 8, 17, 22, 30, 123456789, time.UTC),
				},
			},
		},
		{
			name: "Update a lendingbook: failed!",
			args: args{
				&domain.LendingBooks{
					Model:  domain.Model{ID: fakeLendingBookID},
					BookID: book.ID,
					UserID: user.ID,
					From:   time.Date(2018, 8, 7, 17, 22, 30, 123456789, time.UTC),
					To:     time.Date(2018, 8, 8, 17, 22, 30, 123456789, time.UTC),
				},
			},
			wantErr: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &pgService{
				db: testDB,
			}
			err := s.Delete(context.Background(), tt.args.p)
			if err != nil && err != tt.wantErr {
				t.Errorf("pgService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && tt.wantErr != nil {
				t.Errorf("pgService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
