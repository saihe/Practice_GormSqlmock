package service

import (
	"log"
	"practice_gormsqlmock/database"
	"practice_gormsqlmock/domain/entity"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofrs/uuid"
	"github.com/google/go-cmp/cmp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestDepartment_GetDepartment(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal("gormオープン失敗", err)
	}
	database.Db = gdb

	dId := uuid.Must(uuid.NewV4())
	eId1 := uuid.Must(uuid.NewV4())
	eId2 := uuid.Must(uuid.NewV4())

	tests := []struct {
		name    string
		d       DepartmentService
		want    entity.Department
		wantErr bool
	}{
		{
			name:    "失敗",
			d:       Department{},
			want:    entity.Department{},
			wantErr: true,
		},
		{
			name: "成功",
			d:    Department{},
			want: entity.Department{
				ID:   dId,
				Name: "取得したやつ",
				Employee: []entity.Employee{
					{
						ID:           eId1,
						Name:         "一人目",
						DepartmentId: dId,
					},
					{
						ID:           eId2,
						Name:         "二人目",
						DepartmentId: dId,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		if tt.wantErr {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "departments" ORDER BY "departments"."id" LIMIT 1`)).
				WillReturnRows(mock.NewRows([]string{}))
		} else {
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "departments" ORDER BY "departments"."id" LIMIT 1`)).
				WillReturnRows(mock.NewRows([]string{"ID", "Name"}).
					AddRow(dId, tt.want.Name).
					AddRow(dId, tt.want.Name))
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "employees" WHERE "employees"."department_id" = $1`)).
				WillReturnRows(mock.NewRows([]string{"ID", "Name", "DepartmentId"}).
					AddRow(eId1, tt.want.Employee[0].Name, dId).
					AddRow(eId2, tt.want.Employee[1].Name, dId))
		}
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.GetDepartment()
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err.Error())
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("取得結果と期待値が不一致 want=- got=+ [%s]", diff)
			}
		})
	}
}
