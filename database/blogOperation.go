package database

import (
	"opencloud-server/model/dao"
)

func init() {
	db = InitDatabase()
	SyncStruct(db)
}

func AddBlog(blog *dao.Blog) {
	db.Create(blog)
}

func GetBlogs(uid string, blogs *[]dao.Blog) {
	db.Where("uuid=?", uid).Find(blogs)
}
