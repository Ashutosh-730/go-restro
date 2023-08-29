package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type YpRestaurant struct {
	Id             int       `orm:"column(id);auto"`
	MenuId         int       `orm:"column(menuId);null"`
	RestroName     string    `orm:"column(restroName);null"`
	RestroFullName string    `orm:"column(restroFullName);null"`
	RestroGstNo    string    `orm:"column(restroGst);null"`
	RestroAddress  string    `orm:"column(restroAddress);null"`
	RestroCity     string    `orm:"column(restroCity);null"`
	RestroState    string    `orm:"column(restroState);null"`
	RestroPincode  string    `orm:"column(restroPincode);null"`
	RestroPhone    string    `orm:"column(restroPhone);null"`
	RestroEmail    string    `orm:"column(restroEmail);null"`
	CreatedOn      time.Time `orm:"column(createdOn);type(datetime);null"`
	UpdatedOn      time.Time `orm:"column(updatedOn);type(datetime);null"`
}

func (t *YpRestaurant) TableName() string {
	return "yp_restaurant"
}

func init() {
	orm.RegisterModel(new(YpRestaurant))
}
