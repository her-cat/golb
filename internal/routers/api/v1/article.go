package v1

import "github.com/gin-gonic/gin"

type Article struct {}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {}

// List
// @Summary 获取多个文章
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {}

// Create
// @Summary 新增文章
// @Produce  json
// @Param title body string true "标题" minlength(3) maxlength(100)
// @Param desc body string true "描述" minlength(3) maxlength(100)
// @Param content body string true "内容" minlength(3) maxlength(1000)
// @Param cover_image_url body string true "封面图片" minlength(3) maxlength(255)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {}

// Update
// @Summary 新增文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Param title body string true "标题" minlength(3) maxlength(100)
// @Param desc body string true "描述" minlength(3) maxlength(100)
// @Param content body string true "内容" minlength(3) maxlength(1000)
// @Param cover_image_url body string true "封面图片" minlength(3) maxlength(255)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (a Article) Update(c *gin.Context) {}

// Delete
// @Summary 删除文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/aticles/{id} [delete]
func (a Article) Delete(c *gin.Context) {}
