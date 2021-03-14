// handlers.article.go

package main

// type Data struct {
// 	Title   string
// 	Content string
// 	Picture string
// 	Price   string
// }

// func createArticle(c *gin.Context) {
// 	data := Data{}
// 	data.Title = c.PostForm("title")
// 	data.Content = c.PostForm("content")
// 	data.Picture = c.PostForm("picture")
// 	data.Price = c.PostForm("price")

// 	if _, err := createNewArticle(c, data); err == nil {
// 		c.JSON(200, gin.H{
// 			"message": "Successful creation",
// 		})
// 	} else {
// 		c.JSON(200, gin.H{
// 			"message": "Creation failed",
// 		})
// 	}
// }

// // Creating article
// func createNewArticle(c *gin.Context, data Data) (*article, error) {
// 	a := article{ID: len(getArticleFromDB()) + 1, Title: data.Title, Content: data.Content, Picture: data.Picture, Price: data.Price}

// 	insertArticleToDB(a)

// 	return &a, nil

// }
