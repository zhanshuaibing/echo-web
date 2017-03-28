#### ORM支持部分查询缓存
- 查询缓存接口
    - First()
    - Last()
    - Find()
    - Count()

> 使用`Cache(db)`做查询支持缓存
```go
user := User{}
var count int64
db := DB().Where("id = ?", id)

// 连续查询均支持缓存
db := Cache(db).First(&user).Count(&count)  // *CacheDB

// 如果要取消缓存，中间插入.DB
db := Cache(db).First(&user).DB.Count(&count)   // *gorm.DB
```