package postgres

import (
	"database/sql"
	"fmt"
	pb "users/generated/users"
	"users/model"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) SignUp(in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	var resp pb.SignUpResponse
	err := u.DB.QueryRow(`
			INSERT INTO
			users(
				username, email, password_hash, full_name, user_type)
			VALUES(
				$1, $2, $3, $4, $5)
			RETURNING
				username,
				email,
				password_hash,
				full_name,
				user_type,
				created_at
				`, in.UserName, in.Email, in.Password, in.FullName, in.UserType).Scan(
		&resp.UserName, &resp.Email, &resp.Password, &resp.FullName, &resp.UserType, &resp.CreatedAt)

	return &resp, err
}

func (u *UserRepo) LogIn(in *pb.LogInRequest) (*model.LoginResponse, error) {
	var resp model.LoginResponse

	err := u.DB.QueryRow(`
			SELECT 
				id,
				username,
				email
			FROM users
			WHERE 
				email = $1 and password_hash = $2 and deleted_at=0
	`, in.Email, in.Password).Scan(&resp.Id, &resp.Username, &resp.Email)

	return &resp, err
}

func (u *UserRepo) ViewProfile(id string) (*pb.ViewProfileResponse, error) {

	var resp pb.ViewProfileResponse
	fmt.Println(id)
	err := u.DB.QueryRow(`
	SELECT 
	username,
	email,
	password_hash,
	full_name,
	user_type,
	created_at,
	updated_at
	FROM users
	WHERE id=$1 and deleted_at=0`, id).Scan(
		&resp.UserName, &resp.Email, &resp.Password, &resp.FullName, &resp.UserType, &resp.CreatedAt, &resp.UpdatedAt)

	return &resp, err
}

func (u *UserRepo) EditProfile(in *pb.EditProfileRequeste) (*pb.EditProfileResponse, error) {

	var resp pb.EditProfileResponse

	err := u.DB.QueryRow(`
			UPDATE users
			SET
				full_name=$1,
				bio=$2,
				updated_at=CURRENT_TIMESTAMP
			WHERE 
				id=$3 and deleted_at=0
			RETURNING
				username,
				email,
				password_hash,
				full_name,
				user_type,
				updated_at,
			`, in.FullName, in.FullName, in.Id).Scan(
		resp.UserName, &resp.Email, &resp.Password, &resp.FullName, &resp.UserType, &resp.UpdatedAt)

	return &resp, err

}

func (u *UserRepo) ChangeUserType(in *pb.ChangeUserTypeRequeste) (*pb.ChangeUserTypeResponse, error) {

	var resp pb.ChangeUserTypeResponse

	err := u.DB.QueryRow(`
			UPDATE users
			SET
				user_typ=$1,
				updated_at=CURRENT_TIMESTAMP
			WHERE 
				id=$2 and deleted_at=0
			RETURNING
				id,
				username,
				user_type,
				updated_at,
			`, in.UserType, in.Id).Scan(
		resp.Id, resp.UserName, &resp.UsetType, &resp.UpdatedAt)

	return &resp, err
}

func (u UserRepo) GetAllUsers(in *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	var (
		users []*pb.User
		total int
	)

	err := u.DB.QueryRow("SELECT count(*) FROM users").Scan(&total)
	if err != nil {
		return nil, err
	}

	rows, err := u.DB.Query(`
	SELECT 
	id,
	username,
	user_type,
	full_name
	FROM users offset $1 limit $2
	`, in.Offset, in.Limit)

	for rows.Next() {
		var user pb.User
		err := rows.Scan(&user.Id, &user.UserName, &user.UsetType, &user.FullName)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)

	}

	return &pb.GetAllUsersResponse{
		Users:  users,
		Total:  int64(total),
		Offset: in.Offset,
		Limit:  in.Limit,
	}, err

}

func (u *UserRepo) DeleteUser(id string) (*pb.DeleteUserResponse, error) {

	rows, err := u.DB.Exec(`
			UPDATE users
			SET deleted_at=CURRENT_TIMESTAMP
			Where id=$1 and deleted_at=0`, id)

	if err != nil {
		return &pb.DeleteUserResponse{Success: "hatolik bor"}, err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return &pb.DeleteUserResponse{Success: "hatolik bor"}, err
	}

	return &pb.DeleteUserResponse{Success: "userr ochirildi!"}, nil
}

func (u *UserRepo) PasswordReset(in *pb.PasswordResetRequest) (*pb.PasswordResetResponse, error) {
	_, err := u.DB.Query(`
			UPDATE
			users SET
				password_hash=$1
			WHERE 
				email=$2`, in.PasswordHash, in.Email)
	if err != nil {
		return nil, err
	}

	return &pb.PasswordResetResponse{
		Success: "yangilandi",
	}, nil
}
