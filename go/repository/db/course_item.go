package repo_db

import (
	"fmt"
	"frascati/comp/queryexec"
	"frascati/exception"
	"frascati/lambda"
	"frascati/obj/converter"
	"frascati/obj/dao"
	"frascati/obj/entity"
	"frascati/typing"
	"frascati/utils/querying"
)

type CourseItemRepository interface {
	FindByCourse(ctx typing.Context, course entity.Course) ([]entity.CourseItem, exception.Exception)
	AddBulk(ctx typing.Context, items []entity.CourseItem) ([]entity.CourseItem, exception.Exception)
}

type courseItemRepositoryImpl struct {
	executor queryexec.QueryExecutor
}

func NewCourseItemRepository(executor queryexec.QueryExecutor) CourseItemRepository {
	return courseItemRepositoryImpl{
		executor: executor,
	}
}

func (r courseItemRepositoryImpl) FindByCourse(ctx typing.Context, course entity.Course) ([]entity.CourseItem, exception.Exception) {
	querystr := `
		SELECT id, name, item_type, item_url, start_time, due_time, description
		FROM course_items
		WHERE
			course_id = $1 AND
			deleted_at IS NULL
	`

	rows, err := r.executor.QueryContext(ctx, querystr, course.ID)
	if err != nil {
		return nil, err
	}

	res, err := querying.ScanForRowsThenTransform(
		rows, dao.NewCourseItemDb,
		func(rows queryexec.Rows, elem dao.CourseItemDb) (dao.CourseItemDb, exception.Exception) {
			scanErr := rows.Scan(
				&elem.ID, &elem.Name, &elem.ItemType, &elem.ItemUrl,
				&elem.StartTime, &elem.DueTime, &elem.Description,
			)

			return elem, scanErr
		},
		converter.CourseItemDbToEntity,
	)

	return res, err
}

func (r courseItemRepositoryImpl) AddBulk(ctx typing.Context, items []entity.CourseItem) ([]entity.CourseItem, exception.Exception) {
	querystr := `
		INSERT INTO course_items(course_id, name, item_type, item_url, start_time, due_time, description, created_at, updated_at)
		VALUES
			%s
		RETURNING id, name, item_type, item_url, start_time, due_time, description
	`

	itemsDb := lambda.MapList(items, converter.CourseItemEntityToDaoDb)

	rowstr, args, _ := querying.BulkValuesFromBeginning(
		itemsDb,
		func(elem dao.CourseItemDb, currentIdx int) querying.DataStrArgs {
			return querying.DataStrArgs{
				RowStr: fmt.Sprintf(
					"(CAST($%d AS BIGINT), $%d, CAST($%d AS SMALLINT), $%d, CAST($%d AS TIMESTAMP), CAST($%d AS TIMESTAMP), $%d, NOW(), NOW())",
					currentIdx, currentIdx+1, currentIdx+2, currentIdx+3, currentIdx+4, currentIdx+5, currentIdx+6,
				),
				Args: []any{elem.Course.ID, elem.Name, elem.ItemType, elem.ItemUrl, elem.StartTime, elem.DueTime, elem.Description},
			}
		},
	)

	querystr = fmt.Sprintf(querystr, rowstr)
	rows, err := r.executor.QueryContext(ctx, querystr, args...)
	if err != nil {
		return nil, err
	}

	res, err := querying.ScanForRowsThenTransform(
		rows, dao.NewCourseItemDb,
		func(rows queryexec.Rows, elem dao.CourseItemDb) (dao.CourseItemDb, exception.Exception) {
			err := rows.Scan(&elem.ID, &elem.Name, &elem.ItemType, &elem.ItemUrl, &elem.StartTime, &elem.DueTime, &elem.Description)
			return elem, err
		},
		converter.CourseItemDbToEntity,
	)

	return res, err
}
