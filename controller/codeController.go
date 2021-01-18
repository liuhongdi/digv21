package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type CodeController struct{}

func NewCodeController() CodeController {
	return CodeController{}
}
//生成一个二维码
func (g *CodeController) GetOne(c *gin.Context) {
    //保存path
	pngpath:="/data/liuhongdi/digv21/static/images/q.png"

	//生成二维码
	url:="http://www.baidu.com/"
	qrCode, err := qrcode.New(url, qrcode.Highest)
	if err != nil {
		fmt.Println(err)
		return
	}
	qrCode.DisableBorder = true

	//保存成文件
	qrCode.WriteFile(256,pngpath)
    //显示二维码
	pngurl:= "/static/images/q.png"
	html:="<img src='"+pngurl+"' />"
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, html)
	return
}

//生成一个中间带icon的二维码
func (g *CodeController) GetIcon(c *gin.Context) {

	var (
		bgImg  image.Image
		offset  image.Point
		avatarFile *os.File
		avatarImg image.Image
	)
    //png图片的本地保存路径
	pngpath:="/data/liuhongdi/digv21/static/images/q2.png"
	//url,创建二维码
	url:="http://www.baidu.com/"
	qrCode, err := qrcode.New(url, qrcode.Highest)
	if err != nil {
		//return nil, errors.New("创建二维码失败")
		fmt.Println(err)
		return
	}
	qrCode.DisableBorder = true
    bgImg = qrCode.Image(256)
    //icon的路径
	headpath:="/data/liuhongdi/digv21/static/images/head.jpeg"
	avatarFile, err = os.Open(headpath)
	avatarImg, err = jpeg.Decode(avatarFile)
	//修改图片的大小
	avatarImg = resize.Resize(40, 40, avatarImg, resize.Lanczos3)

	//得到背景图的大小
	b := bgImg.Bounds()
	//居中设置icon到二维码图片
	offset = image.Pt((b.Max.X-avatarImg.Bounds().Max.X)/2, (b.Max.Y-avatarImg.Bounds().Max.Y)/2)
	m := image.NewRGBA(b)
	draw.Draw(m, b, bgImg, image.Point{X: 0, Y: 0,}, draw.Src)
	draw.Draw(m, avatarImg.Bounds().Add(offset), avatarImg, image.Point{X: 0, Y: 0}, draw.Over)

    //save image
	errsave:=SaveImage(pngpath,m)
	if (errsave != nil){
		fmt.Println(errsave)
	}
	//显示图片
	pngurl:= "/static/images/q2.png"
	html:="<img src='"+pngurl+"' />"
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, html)
	return
}

//保存image
func SaveImage(p string, src image.Image) error {
	f, err := os.OpenFile(p, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	ext := filepath.Ext(p)
	if strings.EqualFold(ext, ".jpg") || strings.EqualFold(ext, ".jpeg") {
		err = jpeg.Encode(f, src, &jpeg.Options{Quality: 80})
	} else if strings.EqualFold(ext, ".png") {
		err = png.Encode(f, src)
	} else if strings.EqualFold(ext, ".gif") {
		err = gif.Encode(f, src, &gif.Options{NumColors: 256})
	}
	return err
}
