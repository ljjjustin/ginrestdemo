package models

type User struct {
	Id       int    `json:"id" gorm:"AUTO_INCREMENT PRIMARY"`
	Name     string `json:"name" gorm:"UNIQUE NOT NULL VARCHAR(64)"`
	NickName string `json:"nick_name" gorm:"NOT NULL VARCHAR(128)"`
	Age      int    `json:"age" gorm:"NOT NULL DEFAULT 0"`
	Email    string `json:"email" gorm:"NOT NULL DEFAULT ''"`
}

func UserSave(user *User) {
	db.Save(user)
}

func UserUpdates(user User) {
	db.Updates(user)
}

func UserFind(user *User, id int) bool {
	return db.Find(user, id).RecordNotFound()
}

func UserGetAll(users *[]User) []error {
	db.Order("id desc").Find(users)
	return db.GetErrors()
}

func UserDelete(user *User) {
	db.Find(user, user.Id).Delete(user)
}
