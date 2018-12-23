package database

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ServiceComputingGroup/simpleWebServer/entity"
)

func InsertUser(user *entity.User) bool {

	var err error
	if HasUser(user.UserName) {
		return false
	}
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//将参数传递到sql语句中并且执行
	encoded, _ := json.MarshalIndent(&user, "", "\t")
	_, err = stmt.Exec(user.UserName, string(encoded))
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	return true
}
func DeleteUser(username string) error {
	var err error
	if !HasUser(username) {
		err = errors.New("!HasUser")
		return err
	}
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM user WHERE username = ?")
	if err != nil {
		return err
	}
	//设置参数以及执行sql语句
	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}
	//提交事务
	tx.Commit()

	return nil
}

func UpdateUser(user *entity.User) error {
	var err error
	if !HasUser(user.UserName) {
		err = errors.New("!HasUser")
		return err
	}
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE user SET con = ? WHERE username = ?")
	if err != nil {

		return err
	}
	//设置参数以及执行sql语句
	encoded, _ := json.MarshalIndent(&user, "", "\t")
	_, err = stmt.Exec(string(encoded), user.UserName)
	if err != nil {

		return err
	}
	//提交事务
	tx.Commit()

	return nil
}

func HasUser(username string) bool {
	_, b := GetUser(username)
	return b
}
func GetUser(username string) (*entity.User, bool) {
	var data *entity.User
	var id int
	var str string
	err := DB.QueryRow("SELECT * FROM user WHERE username = ?", username).Scan(&id, &str)
	if err != nil {
		fmt.Println("用户不存在数据库中")
		return nil, false
	}
	json.Unmarshal([]byte(str), &data)
	return data, true
}

func GetAllUser() []entity.User {
	//执行查询语句
	rows, err := DB.Query("SELECT * from user")
	if err != nil {
		fmt.Println("查询出错了")
	}
	var datas []entity.User
	//循环读取结果
	for rows.Next() {
		var data entity.User
		var str string
		var id int
		err := rows.Scan(&id, &str)
		json.Unmarshal([]byte(str), &data)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		datas = append(datas, data)
	}
	return datas
}
