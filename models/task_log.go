/*
* @Author: haodaquan
* @Date:   2017-06-21 12:23:22
* @Last Modified by:   haodaquan
* @Last Modified time: 2017-06-22 14:57:13
 */

package models

import (
	"github.com/astaxie/beego/orm"
)

type TaskLog struct {
	Id          int
	TaskId      int
	ServerId    int
	ServerName  string
	Output      string
	Error       string
	Status      int
	ProcessTime int
	CreateTime  int64
}

func (t *TaskLog) TableName() string {
	return TableName("task_log")
}

func TaskLogAdd(t *TaskLog) (int64, error) {
	return orm.NewOrm().Insert(t)
}

func TaskLogGetList(page, pageSize int, filters ...interface{}) ([]*TaskLog, int64) {
	offset := (page - 1) * pageSize

	logs := make([]*TaskLog, 0)

	query := orm.NewOrm().QueryTable(TableName("task_log"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}

	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&logs)

	return logs, total
}

func TaskLogGetById(id int) (*TaskLog, error) {
	obj := &TaskLog{
		Id: id,
	}

	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func TaskLogDelById(id int) error {
	_, err := orm.NewOrm().Delete(&TaskLog{Id: id})
	return err
}

func TaskLogDelByTaskId(taskId int) (int64, error) {
	return orm.NewOrm().QueryTable(TableName("task_log")).Filter("task_id", taskId).Delete()
}

// func GetTodaySuccessNum() (num, error) {
// 	o := orm.NewOrm()
// 	var r RawSeter
// 	r = o.Raw("SELECT COUNT(*) AS num WHERE create_time>=? AND status<0", "")
// }
