package repo_db

import (
	"fmt"
	"frascati/comp/queryexec"
	"frascati/exception"
	"frascati/obj/entity"
	repository_exception "frascati/repository/exception"
	"frascati/typing"
	"frascati/utils/querying"
)

type RecordRepository interface {
	FindAll(ctx typing.Context) ([]entity.Record, exception.Exception)
	AddBulk(ctx typing.Context, newData []entity.Record) ([]entity.Record, exception.Exception)
}

type recordRepositoryImpl struct {
	executor queryexec.QueryExecutor
}

func NewRecordRepository(executor queryexec.QueryExecutor) RecordRepository {
	return recordRepositoryImpl{
		executor: executor,
	}
}

func (r recordRepositoryImpl) FindAll(ctx typing.Context) ([]entity.Record, exception.Exception) {
	querystr := `
		SELECT id, name, rand_num, description
		FROM records
		WHERE deleted_at IS NULL
	`

	rows, err := r.executor.QueryContext(ctx, querystr)
	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "record")
	}
	defer r.executor.CloseRows(rows, "record - FindAll")

	fmt.Println(querystr)

	res, err := querying.ScanForRows(
		rows, entity.NewRecord,
		func(rows queryexec.Rows, elem entity.Record) (entity.Record, exception.Exception) {
			err := rows.Scan(&elem.ID, &elem.Name, &elem.RandNum, &elem.Description)
			return elem, err
		},
	)

	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "record")
	}

	return res, nil
}

func (r recordRepositoryImpl) AddBulk(ctx typing.Context, newData []entity.Record) ([]entity.Record, exception.Exception) {
	querystr := `
		INSERT INTO records(name, rand_num, description, created_at, updated_at)
		VALUES
			%s
		RETURNING id, name, rand_num, description
	`
	paramIdxStart := 1
	rowstr, args, _ := querying.BulkValues(
		newData, paramIdxStart,
		func(value entity.Record, currentIdx int) querying.DataStrArgs {
			return querying.DataStrArgs{
				RowStr: fmt.Sprintf("($%d, $%d, $%d, NOW(), NOW())", currentIdx, currentIdx+1, currentIdx+2),
				Args:   []any{&value.Name, &value.RandNum, &value.Description},
			}
		},
	)

	querystr = fmt.Sprintf(querystr, rowstr)
	fmt.Println(querystr)
	rows, err := r.executor.QueryContext(ctx, querystr, args...)
	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "record")
	}
	defer r.executor.CloseRows(rows, "record - AddBulk")

	res, err := querying.ScanForRows(
		rows, entity.NewRecord,
		func(rows queryexec.Rows, elem entity.Record) (entity.Record, exception.Exception) {
			err := rows.Scan(&elem.ID, &elem.Name, &elem.RandNum, &elem.Description)
			return elem, err
		},
	)

	if err != nil {
		return nil, repository_exception.WrapQueryexecException(err, "record")
	}

	return res, nil

}
