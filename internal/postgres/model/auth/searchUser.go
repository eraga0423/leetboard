package auth

import "1337b0rd/internal/types/database"

type userResp struct {
	userName     string
	userImageURL string
}

func (a *Auth) FindUser(f database.FindUserReq) (bool, database.FindUserResp) {
	userSessionID := f.GetSessionID()
	sql := a.db.QueryRow(`
	SELECT 
	name,
	user_avatar
	FROM users
	WHERE session_id =$1`, userSessionID)
	var u userResp
	err := sql.Scan(
		&u.userName,
		&u.userImageURL,
	)
	if err != nil {
		return false, userResp{}
	}
	return true, u
}

func (u userResp) GetUserName() string     { return u.userName }
func (u userResp) GetUserImageURL() string { return u.userImageURL }
