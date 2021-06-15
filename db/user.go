package db

import "ticket/model"

// 用户登录
func UserLogin(mod model.User) (r bool,err error) {
	var newMod = model.User{}
	r, err = mysqlEngine.Where("id=?", mod.Id).Cols("pwd").Get(&newMod)
	if err != nil {
		return r, err
	}
	if mod.Pwd != newMod.Pwd {
		r = false
	}
	return r, nil
}

// 用户注册
func UserRegister(mod model.User) (bool, error) {
	_, err := mysqlEngine.Insert(mod)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 修改密码
func UserModifyPwd(mod model.User, pwd string) (bool, error) {
	return false, nil
}