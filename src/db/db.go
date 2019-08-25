package db

import (
	"domain"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/satori/go.uuid"
	"kits"
	"log"
)

//数据库文件名称
const dbfile = "data.db"

//数据库对象
type LocalDB struct {
	Db *bolt.DB //数据库对象
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
func (localdb *LocalDB) Add(tableName string, key string, value interface{}) *LocalDB {
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
		err=bucket.Put([]byte(key), buf)
		kits.Panic("插入数据异常！", err)
		return nil
	})
	kits.Panic("新增或更新数据异常！", err)
	return localdb
}

//新增数据
func (localdb *LocalDB) Upadate(tableName string, key string, value interface{}) *LocalDB {
	return localdb.Add(tableName, key, value)
}

//自动生成ID
func (localDB *LocalDB) NextId() string {
	id, err :=uuid.NewV4();
	kits.Panic("生成uuid出现异常：",err)
	fmt.Printf("UUIDv4: %s\n", id)
	return id.String()
}


//删除数据
func (localDB *LocalDB) Delete(tableName string, key string) *LocalDB {

	err := localDB.Db.Update(func(tx *bolt.Tx) error {
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
	kits.Panic("删除异常！", err)
	return localDB
}

//查询数据，返回json数据。
func (localDB *LocalDB) Query(tableName string, key string) []byte {
	var str []byte
	err := localDB.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(tableName))
		 str = bucket.Get([]byte(key))
		return nil
	})
	kits.Panic("查询报错！", err)
	return str
}

//查询用户数据
func (localDB *LocalDB) QueryUser(tableName string,key string) domain.User {
	var user= domain.User{}
	str:=localDB.Query(tableName,key)
	if str!=nil{
		json.Unmarshal(str,&user)
	}
	return user
}
