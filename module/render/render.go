package render

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"

	"github.com/hobo-go/echo-mw/multitemplate"
	"github.com/hobo-go/echo-mw/pongo2echo"
	"github.com/hobo-go/echo-mw/session"

	"echo-web/conf"
	"echo-web/module/auth"
	"echo-web/module/log"
	MT "echo-web/template"
)

func Render() echo.MiddlewareFunc {
	if conf.TMPL_TYPE == conf.PONGO2 {
		return pongo2()
	} else {
		return render()
	}
}

func render() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := next(c); err != nil {
				c.Error(err)
			}

			tmpl, context, err := getContext(c)
			if err == nil {
				c.Render(http.StatusOK, tmpl, context)
			} else {
				log.DebugPrint("Render Error: %v, tmpl %v, content %v", err, tmpl, context)
			}

			return nil
		}
	}
}

func getContext(c echo.Context) (tmpl string, context map[string]interface{}, err error) {
	tmplName := c.Get("tmpl")
	tmplNameValue, isString := tmplName.(string)
	tmplData := c.Get("data")

	// 模板未定义
	if !isString {
		return "", nil, errors.New("No tmpl defined!")
	}

	// 公共模板数据
	commonDatas := getCommonContext(c)

	// 模板数据
	if tmplData != nil {
		contextData, isMap := tmplData.(map[string]interface{})

		if isMap {
			for key, value := range commonDatas {
				contextData[key] = value
			}

			return tmplNameValue, contextData, nil
		}
	}

	return tmplNameValue, commonDatas, nil

}

func getCommonContext(c echo.Context) map[string]interface{} {
	a := auth.Default(c)
	userId := a.User.UniqueId().(uint64)

	// 公共模板数据
	commonDatas := make(map[string]interface{})
	commonDatas["UserId"] = userId
	commonDatas["UserName"] = "用户名"

	commonDatas["DOMAIN_WWW"] = conf.DOMAIN_WWW
	commonDatas["DOMAIN_API"] = conf.DOMAIN_API

	// CSRF
	csrf := c.Get("_csrf")
	commonDatas["_csrf"] = csrf

	// Error
	s := session.Default(c)
	commonDatas["_error"] = s.Flashes("_error")

	path := c.Request().URL.Path
	uri := c.Request().RequestURI

	// 登录、注册、退出页面取已有RedirectParam
	redirect := c.QueryParam(auth.RedirectParam)
	switch path {
	case "/login":
		uri = redirect
	case "/register":
		uri = redirect
	case "/logout":
		uri = redirect
	default:
	}

	log.Print("Path : %v \n", path)
	log.Print("URI : %v \n", uri)

	commonDatas["requestUrl"] = uri

	return commonDatas
}

/**
 * 模板加载
 * 支持文件/Bindata加载模板
 */

func LoadTemplates() echo.Renderer {
	switch conf.TMPL_TYPE {
	case conf.PONGO2:
		return pongo2echo.New(
			pongo2echo.RenderOptions{
				TemplateDir: conf.TMPL_DIR,
				ContentType: "text/html; charset=utf-8",
				Debug:       !conf.RELEASE_MODE,
			})
	case conf.BINDATA:
		return loadTemplatesBindata(conf.TMPL_DIR)
	default:
		return loadTemplatesDefault(conf.TMPL_DIR)
	}
}

func loadTemplatesDefault(templateDir string) *multitemplate.Render {
	r := multitemplate.New()

	layoutDir := templateDir + "/layout/"
	layouts, err := filepath.Glob(layoutDir + "*/*" + conf.TMPL_SUFFIX)
	if err != nil {
		panic(err.Error())
	}

	includeDir := templateDir + "/include/"
	includes, err := filepath.Glob(includeDir + "*" + conf.TMPL_SUFFIX)
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		tmpl := template.Must(template.ParseFiles(files...))
		tmplName := strings.TrimPrefix(layout, layoutDir)
		tmplName = strings.TrimSuffix(tmplName, conf.TMPL_SUFFIX)
		log.DebugPrint("Tmpl add " + tmplName)
		r.Add(tmplName, tmpl)
	}
	return &r
}

func loadTemplatesBindata(templateDir string) *multitemplate.Render {
	r := multitemplate.New()

	layoutDir := templateDir + "/layout"
	layoutDirs, err := MT.AssetDir(layoutDir)
	if err != nil {
		panic(err.Error())
	}

	var layouts []string
	for _, dir := range layoutDirs {
		files, err := MT.AssetDir(layoutDir + "/" + dir)
		if err != nil {
			panic(err.Error())
		}

		// 过滤非.tmpl后缀模板
		layoutFiels, err := tmplsFilter(files, layoutDir+"/"+dir)
		if err != nil {
			panic(err.Error())
		}

		layouts = append(layouts, layoutFiels...)
	}

	includeDir := templateDir + "/include"
	includeFiels, err := MT.AssetDir(includeDir)
	if err != nil {
		panic(err.Error())
	}
	// 过滤非.tmpl后缀模板
	includes, err := tmplsFilter(includeFiels, includeDir)
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		tmpl := template.Must(parseBindataFiles(files...))
		tmplName := strings.TrimPrefix(layout, layoutDir+"/")
		tmplName = strings.TrimSuffix(tmplName, conf.TMPL_SUFFIX)
		log.DebugPrint("Tmpl add " + tmplName)
		r.Add(tmplName, tmpl)
	}
	return &r
}

// 过滤非tmpl后缀模板文件
func tmplsFilter(files []string, dir string) ([]string, error) {
	var tmpls []string
	for _, file := range files {
		if strings.HasSuffix(file, conf.TMPL_SUFFIX) {
			tmpls = append(tmpls, dir+"/"+file)
		}
	}
	return tmpls, nil
}

// parseFiles is the helper for the method and function. If the argument
// template is nil, it is created from the first file.
func parseBindataFiles(filenames ...string) (*template.Template, error) {
	var t *template.Template
	if len(filenames) == 0 {
		// Not really a problem, but be consistent.
		return nil, fmt.Errorf("html/template: no files named in call to ParseFiles")
	}
	for _, filename := range filenames {
		b, err := MT.Asset(filename)
		if err != nil {
			return nil, err
		}
		s := string(b)
		name := filepath.Base(filename)
		// First template becomes return value if not already defined,
		// and we use that one for subsequent New calls to associate
		// all the templates together. Also, if this file has the same name
		// as t, this file becomes the contents of t, so
		//  t, err := New(name).Funcs(xxx).ParseFiles(name)
		// works. Otherwise we create a new template associated with t.
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
