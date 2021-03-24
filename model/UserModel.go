package model

import (
	"context"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"mic_srv_office/db"
	"sync"
	"time"
)

const (
	sqlType = "mysql"
)

type TimeModel struct {
	CreatedAt string `gorm:"column:create_time" json:"create_time"`
	UpdatedAt string `gorm:"column:update_time" json:"update_time"`
}


func (v TimeModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_time", time.Now())
	scope.SetColumn("update_time", time.Now())
	return nil
}

func (v TimeModel) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("update_time", time.Now())
	return nil
}

type UserModel struct {
	TimeModel
	ID int64 `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Pass string `gorm:"column:pass"`
	Age int `gorm:"column:age"`
	Gender int `gorm:column:gender`
	phone string `gorm:"column:phone"`
	Addr string `gorm:"column:addr"`
	IsActive int `gorm:"column:is_active;default:'0'"`
}

func (UserModel)TableName() string {
	return "sys_user"
}


func (u *UserModel)GetUsers()[]UserModel  {
	db,err := gorm.Open(sqlType,db.Getmysql_offConStr())
	if err != nil {
		return nil
	}
	defer db.Close()
	users :=[]UserModel{}
	db.Find(users)
	return users
}

func GetUserByID(id int64)*UserModel  {
	db,err := gorm.Open(sqlType,db.Getmysql_offConStr())
	if err != nil {
		return nil
	}
	defer db.Close()
	user :=UserModel{}
	if nil !=db.Take(&user).Where("id = ?",id).Error{
		return nil
	}
	return &user

}

func AddUser(in UserModel)error  {
	db,err := gorm.Open(sqlType,db.Getmysql_offConStr())
	if err != nil {
		return nil
	}
	defer db.Close()
	if !db.NewRecord(in) {
		return errors.New("insert failed")
	}
	return nil
}

func AddUsers(users []UserModel)error  {
	db,err := gorm.Open(sqlType,db.Getmysql_offConStr())
	if err != nil {
		return nil
	}
	defer db.Close()
	for _,user := range users{
		err =AddUser(user)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteUser(in <-chan int64)error  {
	db,err := gorm.Open(sqlType,db.Getmysql_offConStr())
	if err != nil {
		return nil
	}
	defer db.Close()
	if nil !=db.Where("id = ?",in).Delete(UserModel{}).Error{
		return errors.New("delete failed")
	}
	return nil
}

func DeleteUsersSource(ctx context.Context,users *[]UserModel)(<- chan int64, <-chan error,error)  {
	var chUser chan int64
	var chErr chan error

	chUser=make(chan int64)
	chErr=make(chan error,1)

	go func() {
		defer close(chUser)
		defer close(chErr)
		for _,user := range *users{
			chUser<-user.ID
			select {
			case <-ctx.Done():
				chErr<-errors.New("has been conceled")
				return
			default:
			}
		}
	}()
	return chUser,chErr,nil
}

func DeleteUsersDowork(ctx context.Context,id <-chan int64) (chan error,error){
	errc := make(chan error,1)
	go func() {
		defer close(errc)
		for userid := range id{
			db,err := gorm.Open(sqlType,db.Getmysql_offConStr())
			if err != nil {
				errc<-errors.New("connect db error")
				return
			}
			defer db.Close()
			if nil != db.Delete("id = ?",userid).Error{
				errc<-errors.New(db.Error.Error())
			}
			select {
			case <-ctx.Done():
				errc<-errors.New("has been conceled")
				return
			default:
			}
		}
	}()
	return errc,nil
}

func mergeErrs(err ...<-chan error)chan error  {
	var wg sync.WaitGroup
	errc :=make(chan error,len(err))
	errfun := func(c <-chan error) {
		for in :=range c{
			errc<-in
		}
		wg.Done()
	}
	wg.Add(len(err))
	for _,c := range err{
		go errfun(c)
	}
	return errc
}

func waitingforher(errs ...<-chan error) error {
	err :=mergeErrs(errs...)
	for e :=range err{
		if e !=nil {
			return e
		}
	}
	return nil
}
