package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
)

//auto_now 每次 model 保存时都会对时间自动更新
//auto_now_add 第一次保存时才设置时间

type Tangshi300 struct {
	Id      int       `orm:"column(ID);pk"`
	BookID string    `orm:"column(bookID);size(50)"`
	Author   string    `orm:"size(30)"`
	Title   string    `orm:"size(100)"`
	Chaodai string    `orm:"size(100)"`
	//Created time.Time `orm:"auto_now_add;type(date)"`
	//Updated time.Time `orm:"auto_now;type(date)"`
	//Updated time.Time `orm:"auto_now;type(date)"`
	Yuanwen   string    `orm:"column(yuanwen1);size(5500)"`
	Fanyi   string    	`orm:"column(fanyi1);size(5500)"`
	Shangxi   string    	`orm:"column(shangxi1);size(5500)"`
	Beijing   string    	`orm:"column(beijing);size(5500)"`
	Dianping   string    	`orm:"column(dianping);size(5500)"`
	AuthDetail   string    	`orm:"column(authDetail);size(3500)"`
}

func (t *Tangshi300) TableName() string {
	return "BookDetail"
}

func init() {
	orm.RegisterModel(new(Tangshi300))
}

// last inserted Id on success.
func AddTangshi300(m *Tangshi300) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// Id doesn't exist
func GetTangshi300ById(id int) (v *Tangshi300, err error) {
	o := orm.NewOrm()
	v = &Tangshi300{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}


// GetAllBook retrieves all Book matches certain condition. Returns empty list if
// no records exist
func GetAllTangshi300(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Tangshi300))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Tangshi300
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateBook updates Book by Id and returns error if
// the record to be updated doesn't exist
func UpdateTangshi300ById(m *Tangshi300) (err error) {
	o := orm.NewOrm()
	v := Tangshi300{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// the record to be deleted doesn't exist
func DeleteTangshi300(id int) (err error) {
	o := orm.NewOrm()
	v := Tangshi300{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Tangshi300{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
