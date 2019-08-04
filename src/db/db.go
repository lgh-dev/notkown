package db

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"kits"
	"log"
)

//数据库文件名称
const dbfile = "data.db"

//数据库对象
type LocalDB struct {
	Db        *bolt.DB     //数据库对象
}

//创建数据库
func newboltDB() *bolt.DB {
	db, err := bolt.Open(dbfile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//根据表名创建数据库对象
func CreateLocalDB() *LocalDB {
	//创建数据库对象
	db := LocalDB{Db: newboltDB()}
	return &db
}

//新增数据
func (localdb *LocalDB) Add(tableName string,key string, value interface{}) *LocalDB {
	err := localdb.Db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(tableName))
		if err != nil {
			kits.Panic("创建表异常！", err)
			return err
		}
		//转换json数据
		buf, err := json.Marshal(value)
		kits.Panic("转换数据异常！", err)
		//存储数据
		bucket.Put([]byte(key), buf)
		return nil
	})
	kits.Panic("新增数据异常！", err)
	return localdb
}

//自动生成ID
//func (LocalDB *LocalDB) AddAuthForID(value interface{}) *LocalDB  {
//	id, _ :=LocalDB.Bucket.NextSequence()
//	LocalDB.Add(kits.UInt64ToBytes(id),value)
//	return LocalDB
//}

//更新数据
//func (LocalDB *LocalDB) Update(key []byte, value string) *LocalDB {
//	return LocalDB.Add(key, value)
//}

//删除数据
func (localDB *LocalDB) Delete(tableName string,key string) *LocalDB {

	localDB.Db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(tableName))
		if err != nil {
			kits.Panic("创建表异常！", err)
			return err
		}
		err = bucket.Delete([]byte(key))
		if err != nil {
			log.Panic(err)
			return err
		}
		return nil
	})
	return localDB
}

//查询数据
func (localDB *LocalDB) Query( tableName string,key string) interface{} {
	var v interface{}
	localDB.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(tableName))
		str := bucket.Get([]byte(key))
		err:=json.Unmarshal(str,&v)
		kits.Panic("json转换对象出错！",err)
		return nil
	})
	return v
}
