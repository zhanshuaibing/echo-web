package www

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/hobo-go/echo-mw/session"

	"github.com/hobo-go/echo-web/models"
	"github.com/hobo-go/echo-web/modules/auth"
	"github.com/hobo-go/echo-web/modules/log"
)

type LoginForm struct {
	Nickname string `form:"nickname" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginHandler(c echo.Context) error {
	redirect := c.QueryParam(auth.RedirectParam)

	a := auth.Default(c)
	if a.User.IsAuthenticated() {
		if redirect == "" {
			redirect = "/"
		}
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	c.Set("tmpl", "www/login")
	c.Set("data", map[string]interface{}{
		"title":         "Login",
		"redirectParam": auth.RedirectParam,
		"redirect":      redirect,
	})

	return nil
}

func LoginPostHandler(c echo.Context) error {
	redirect := c.QueryParam(auth.RedirectParam)
	if redirect == "" {
		redirect = "/"
	}

	a := auth.Default(c)
	if a.User.IsAuthenticated() {
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	loginURL := c.Request().URI()

	var form LoginForm
	if err := c.Bind(&form); err == nil {
		var User models.User
		u := User.GetUserByNicknamePwd(form.Nickname, form.Password)

		if u != nil {
			session := session.Default(c)
			err := auth.AuthenticateSession(session, u)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.Redirect(http.StatusMovedPermanently, redirect)
			return nil
		} else {
			c.Redirect(http.StatusMovedPermanently, loginURL)
			return nil
		}
	} else {
		log.DebugPrint("Login form params: %v", c.FormParams())
		log.DebugPrint("Login form bind Error: %v", err)
		c.Redirect(http.StatusMovedPermanently, loginURL)
		return nil
	}

	return nil
}

func LogoutHandler(c echo.Context) error {
	session := session.Default(c)
	a := auth.Default(c)
	auth.Logout(session, a.User)

	redirect := c.QueryParam(auth.RedirectParam)
	if redirect == "" {
		redirect = "/"
	}

	c.Redirect(http.StatusMovedPermanently, redirect)

	return nil
}

func RegisterHandler(c echo.Context) error {
	redirect := c.QueryParam(auth.RedirectParam)

	a := auth.Default(c)
	if a.User.IsAuthenticated() {
		if redirect == "" {
			redirect = "/"
		}
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	c.Set("tmpl", "www/register")
	c.Set("data", map[string]interface{}{
		"title":         "Register",
		"redirectParam": auth.RedirectParam,
		"redirect":      redirect,
	})

	return nil
}

func RegisterPostHandler(c echo.Context) error {
	redirect := c.QueryParam(auth.RedirectParam)
	if redirect == "" {
		redirect = "/"
	}

	a := auth.Default(c)
	if a.User.IsAuthenticated() {
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	var form LoginForm
	if err := c.Bind(&form); err == nil {
		var User models.User
		u := User.AddUserWithNicknamePwd(form.Nickname, form.Password)
		if u != nil {
			session := session.Default(c)
			err := auth.AuthenticateSession(session, u)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.Redirect(http.StatusMovedPermanently, redirect)
			return nil
		} else {
			log.DebugPrint("Register user add error")

			s := session.Default(c)
			s.AddFlash("Register user add error", "_error")

			// registerURL := c.Request().URI()
			// c.Redirect(http.StatusMovedPermanently, registerURL)
			c.Set("tmpl", "www/register")
			c.Set("data", map[string]interface{}{
				"title":         "Register",
				"redirectParam": auth.RedirectParam,
				"redirect":      redirect,
			})
			return nil
		}
	} else {
		log.DebugPrint("Register form bind Error: %v", err)

		s := session.Default(c)
		s.AddFlash("Register form bind Error:"+err.Error(), "_error")

		// registerURL := c.Request().URI()
		// c.Redirect(http.StatusMovedPermanently, registerURL)
		c.Set("tmpl", "www/register")
		c.Set("data", map[string]interface{}{
			"title":         "Register",
			"redirectParam": auth.RedirectParam,
			"redirect":      redirect,
		})
		return nil
	}

	return nil
}
