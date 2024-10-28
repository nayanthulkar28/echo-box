package repo

import (
	"anon-chat/internal/domain"
	"anon-chat/pkg/postgres"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

const usersTable = "users"

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r *UserRepo) CreateUser(ctx context.Context, u domain.User) error {
	sql, args, err := r.Builder.
		Insert(usersTable).
		Columns("id", "username", "password", "friends", "created_at", "updated_at").
		Values(u.Id, u.Username, u.Password, u.Friends, u.CreatedAt, u.UpdatedAt).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	return err
}

func (r *UserRepo) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	sql, args, err := r.Builder.
		Select("*").
		From(usersTable).
		Where(sq.Eq{"username": username}).
		ToSql()

	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Friends)
	if err != nil {
		if err == pgx.ErrNoRows {
			return &user, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, u *domain.User) error {
	sql, args, err := r.Builder.
		Update(usersTable).
		Set("password", u.Password).
		Set("friends", u.Friends).
		Where(sq.Eq{"id": u.Id}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	return err
}

func (r *UserRepo) GetFriendsByUsername(ctx context.Context, username string) ([]domain.User, error) {
	var users []domain.User
	subQ, subQArgs, err := r.Builder.
		Select("UNNEST(friends)").
		From(usersTable).
		Where(sq.Eq{"username": username}).
		ToSql()

	if err != nil {
		return nil, err
	}

	sql, args, err := r.Builder.
		Select("*").
		From(usersTable).
		Where("username IN ("+subQ+")", subQArgs...).
		ToSql()

	fmt.Println(sql)
	fmt.Println(args...)

	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		fmt.Println(err)
		return []domain.User{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Friends)
		if err != nil {
			return []domain.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
