package db

import (
	"github.com/boltdb/bolt"
	"log"
	"encoding/json"
	"kits"
)

//数据库文件名称
const dbfile = "data.db"

//数据库对象
type Table struct {
	TableName string       //表名称
	Db        *bolt.DB     //数据库对象
	Bucket    *bolt.Bucket //表对象
}

//创建数据库
func NewboltDB() *bolt.DB {
	db, err := bolt.Open(dbfile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//根据表名创建数据库对象
func CreateDB(tableName string) *Table {
	//创建数据库对象
	db := Table{TableName: tableName, Db: NewboltDB()}
	return &db
}

//新增数据
func (table *Table) Add(key []byte, value interface{}) *Table {
	err := table.Db.Update(func(tx *bolt.Tx) error {
		tb, err := tx.CreateBucketIfNotExists([]byte(table.TableName))
		if err != nil {
			return err
		}
		//赋予表对象
		table.Bucket = tb
		//转换数据
		buf, err := json.Marshal(value)
		kits.Panic("转换数据异常！", err)
		//存储数据
		tb.Put(key, buf)

		if err := tx.Commit(); err != nil {
			return err
		}
		return nil
	})
	kits.Panic("新增数据异常！", err)
	return table
}

//自动生成ID
func (table *Table) AddAuthForID(value interface{}) *Table  {
	id, _ :=table.Bucket.NextSequence()
	table.Add(kits.UInt64ToBytes(id),value)
	return table
}

//更新数据
func (table *Table) Update(key []byte, value string) *Table {
	return table.Add(key, value)
}

//删除数据
func (table *Table) Delete(key string) *Table {
	table.Db.Update(func(tx *bolt.Tx) error {
		err := table.Bucket.Delete([]byte(key))
		if err != nil {
			log.Panic(err)
			return err
		}
		return nil
	})
	return table
}

//查询数据
func (table *Table) Query(key []byte) []byte {
	var v []byte
	table.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table.TableName))
		v = b.Get(key)
		return nil
	})
	return v
}
