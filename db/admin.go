package db

import "ticket/model"

// 管理员登录
func AdminLogin(mod model.Admin) (r bool,err error) {
	var newMod = model.Admin{}
	r, err = mysqlEngine.Where("id=?", mod.Id).Cols("pwd").Get(&newMod)
	if err != nil {
		return r, err
	}
	if mod.Pwd != newMod.Pwd {
		r = false
	}
	return r, nil
}