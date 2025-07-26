package gorm1

import (
	"fmt"

	"gorm.io/gorm"
)

/***
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。


题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。


题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。


**/

type User struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"size:100;not null"`
	PostNum int    `gorm:default:0`
	Posts   []Post `gorm:"foreignKey:UserID"` // 一对多关系
}

type Post struct {
	ID       int       `gorm:"primaryKey"`
	Title    string    `gorm:"size:200;not null"`
	Content  string    `gorm:"type:text;not null"`
	UserID   int       `gorm:"not null"`
	State    int       `gorm:"default:0"`         //0-无评论  1-有评论
	Comments []Comment `gorm:"foreignKey:PostID"` // 一对多关系
}

type Comment struct {
	ID      int    `gorm:"primaryKey"`
	Content string `gorm:"type:text;not null"`
	PostID  int    `gorm:"not null"`
}

// //为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {

	var user User
	tx.First(&user, p.UserID)
	tx.Model(&User{}).Where("id = ?", user.ID).Update("post_num", user.PostNum+1)
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {

	fmt.Printf("------------------c is %v", c)
	var post Post
	tx.Debug().First(&post, c.PostID)
	if post.State == 0 {
		tx.Debug().Model(&Post{}).Where("id = ?", post.ID).Update("state", 1)
	}

	return nil
}
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {

	fmt.Printf("c is %v", c)
	var post Post
	tx.Debug().Preload("Comments").First(&post, c.PostID)
	if len(post.Comments) == 0 {
		tx.Debug().Model(&Post{}).Where("id = ?", post.ID).Update("state", 0)
	}

	return nil
}

func PostComment() {
	db := initDb()
	db.Debug().Exec("DROP TABLE IF EXISTS users, posts, comments")
	db.Debug().AutoMigrate(&User{})
	db.Debug().AutoMigrate(&Post{})
	db.Debug().AutoMigrate(&Comment{})

	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	db.Debug().CreateBatchSize = 10

	db.Debug().Create([]User{
		{Name: "Alice", Posts: []Post{
			{Title: "Post 1", Content: "Content 1", Comments: []Comment{
				{Content: "Comment 1"},
				{Content: "Comment 2"},
			}},
			{Title: "Post 2", Content: "Content 2", Comments: []Comment{
				{Content: "Comment 1"},
			}},
		}},
		{Name: "Bob", Posts: []Post{
			{Title: "Post 3", Content: "Content 3", Comments: []Comment{
				{Content: "Comment 1"},
				{Content: "Comment 2"},
				{Content: "Comment 3"},
			}},
		}},
	})

	var userId int = 1
	var user User
	var postList []Post
	var commitList []Comment
	// db.Debug().Preload("Posts").Preload("Post.Comments").First(&one, userId)
	db.Debug().First(&user, userId)
	db.Debug().Where("user_id = ?", userId).Find(&postList)

	var postIds []int
	for _, v := range postList {
		postIds = append(postIds, v.ID)

	}

	db.Debug().Where("post_id in (?)", postIds).Find(&commitList)

	fmt.Printf("commitList: %+v\n", commitList)

	for i := 0; i < len(postList); i++ {
		for j := 0; j < len(commitList); j++ {
			if postList[i].ID == commitList[j].PostID {
				postList[i].Comments = append(postList[i].Comments, commitList[j])
			}
		}
	}

	user.Posts = postList

	fmt.Printf("User %d's posts and comments: %+v\n", userId, user)

	// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	var comment Comment
	db.Debug().Where("post_id = ?", 2).First(&comment).Delete(&comment) //一定要先给comment填充值再删，不然删除钩子获取不到值
	// db.Debug().Delete(&comment)

}
