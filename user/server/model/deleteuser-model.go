package usermodel

import (
	"context"

	"github.com/beto-ouverney/blogs-microservices/user/server/entity"
	"github.com/beto-ouverney/blogs-microservices/user/server/errors"
)

func (m *modelSqlx) Delete(id int64) *errors.CustomError {

	var blogPosts []entity.BlogPostResponse
	err := m.sqlx.SelectContext(context.Background(), &blogPosts, `SELECT BlogPosts.id,BlogPosts.userId FROM BlogPosts WHERE BlogPosts.userId = ?`, id)
	if err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.Delete", Err: err}
	}

	tran, err := m.sqlx.BeginTx(context.Background(), nil)
	if err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.Delete", Err: err}
	}
	defer tran.Rollback()

	for _, v := range blogPosts {
		_, err = tran.ExecContext(context.Background(), "DELETE FROM PostCategories WHERE postId = ?", v.ID)
		if err != nil {
			return &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.Delete", Err: err}
		}
	}

	_, err = tran.ExecContext(context.Background(), "DELETE FROM BlogPosts WHERE userId = ?", id)

	if err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.Delete", Err: err}
	}
	_, err = tran.ExecContext(context.Background(), "DELETE FROM Users WHERE id = ?", id)
	if err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "usermodel.Delete", Err: err}
	}

	if err = tran.Commit(); err != nil {
		return &errors.CustomError{Code: errors.EINTERNAL, Op: "blogpostmodel.AddBlogPost", Err: err}
	}

	return nil
}
