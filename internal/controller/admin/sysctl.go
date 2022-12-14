package admin

import (
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nfnt/resize"
	"github.com/zxysilent/utils"
	_ "golang.org/x/image/bmp"
)

// UploadFile doc
// @Tags ctrl-系统相关
// @Summary 上传文件
// @Accept  mpfd
// @Param token query string true "token"
// @Param file formData file true "file"
// @Router /adm/upload/file [post]
func UploadFile(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(utils.Fail("未发现文件", err.Error()))
	}
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(utils.Fail("文件打开失败", err.Error()))

	}
	dir := time.Now().Format("200601/02")
	os.MkdirAll("./static/upload/"+dir[:6], 0755)
	name := "static/upload/" + dir + utils.RandStr(10) + path.Ext(file.Filename)
	dst, err := os.Create(name)
	if err != nil {
		return ctx.JSON(utils.Fail("文件打创建文件失败", err.Error()))

	}
	_, err = io.Copy(dst, src)
	if err != nil {
		return ctx.JSON(utils.Fail("文件保存失败", err.Error()))
	}
	src.Close()
	dst.Close()
	return ctx.JSON(utils.Succ("上传成功", "/"+name))
}

// UploadImage doc
// @Tags ctrl-系统相关
// @Summary 上传图片并裁剪
// @Accept  mpfd
// @Param token query string true "token"
// @Param file formData file true "file"
// @Router /adm/upload/image [post]
func UploadImage(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(utils.ErrIpt(`未发现文件,请重试`, err.Error()))
	}
	if !strings.Contains(file.Header.Get("Content-Type"), "image") {
		return ctx.JSON(utils.ErrIpt("请选择图片文件", file.Header.Get("Content-Type")))
	}
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(utils.ErrIpt(`文件打开失败,请重试`, err.Error()))
	}
	defer src.Close()
	attr := &struct {
		Cut bool `json:"cut" query:"cut" form:"cut"`
		Wd  int  `json:"wd" query:"wd" form:"wd"`
		Hd  int  `json:"hd" query:"hd" form:"hd"`
	}{}
	ctx.Bind(attr)
	dir := time.Now().Format("200601/02")
	os.MkdirAll("./static/upload/"+dir[:6], 0755)
	ext := path.Ext(file.Filename)
	// if conf.App.ImageCut && attr.Cut {
	// 	ext = ".jpg"
	// }
	ext = ".jpg"
	name := "static/upload/" + dir + utils.RandStr(10) + ext
	dst, err := os.Create(name)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("目标文件创建失败,请重试", err.Error()))
	}
	defer dst.Close()
	if attr.Cut { //图片裁剪
		imgSrc, _, err := image.Decode(src)
		// 图片解码
		if err != nil {
			return ctx.JSON(utils.ErrIpt("读取图片失败,请重试", err.Error()))
		}
		if attr.Wd <= 0 {
			attr.Wd = 800
		}
		if attr.Hd <= 0 {
			attr.Hd = 800
		}
		// 宽度>指定宽度 防止负调整
		dx := imgSrc.Bounds().Dx()
		if dx > attr.Wd {
			imgSrc = resize.Resize(uint(attr.Wd), 0, imgSrc, resize.Lanczos3)
		}
		//高度>指定高度 防止负调整
		dy := imgSrc.Bounds().Dy()
		if dy > attr.Hd {
			imgSrc = resize.Resize(0, uint(attr.Hd), imgSrc, resize.Lanczos3)
		}
		err = jpeg.Encode(dst, imgSrc, nil)
		if err != nil {
			return ctx.JSON(utils.ErrIpt("文件写入失败,请重试", err.Error()))
		}
	} else {
		_, err = io.Copy(dst, src)
		if err != nil {
			return ctx.JSON(utils.ErrIpt("文件写入失败,请重试", err.Error()))
		}
	}
	return ctx.JSON(utils.Succ("上传成功", "/"+name))
}
