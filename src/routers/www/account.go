package www

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/syntaqx/echo-middleware/session"
	// "github.com/gin-gonic/gin/binding"

	"models"
	"modules/auth"
)

type LoginForm struct {
	Nickname string `form:"nickname" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginHandler(c *echo.Context) error {
	redirect := c.Query(auth.RedirectParam)

	a := auth.Default(c)
	if a.User.IsAuthenticated() {
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

func LoginPostHandler(c *echo.Context) error {
	redirect := c.Query(auth.RedirectParam)

	a := auth.Default(c)
	if a.User.IsAuthenticated() {
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	loginURL := fmt.Sprintf("/login?%s=%s", auth.RedirectParam, redirect)

	var form LoginForm
	if c.Bind(&form) == nil {
		model := models.Default(c)
		u := model.GetUserByNicknamePwd(form.Nickname, form.Password)
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
		c.Redirect(http.StatusMovedPermanently, loginURL)
		return nil
	}

	return nil
}

func LogoutHandler(c *echo.Context) error {
	session := session.Default(c)
	a := auth.Default(c)
	auth.Logout(session, a.User)

	c.Redirect(http.StatusMovedPermanently, "/")

	return nil
}

func RegisterHandler(c *echo.Context) error {
	redirect := c.Query(auth.RedirectParam)

	a := auth.Default(c)
	if a.User.IsAuthenticated() {
		log.Print("Register IsAuthenticated!")
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

func RegisterPostHandler(c *echo.Context) error {
	redirect := c.Query(auth.RedirectParam)

	a := auth.Default(c)
	if a.User.IsAuthenticated() {
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	registerURL := fmt.Sprintf("/register?%s=%s", auth.RedirectParam, redirect)

	var form LoginForm
	if c.Bind(&form) == nil {
		model := models.Default(c)
		u := model.AddUserWithNicknamePwd(form.Nickname, form.Password)
		if u != nil {
			session := session.Default(c)
			err := auth.AuthenticateSession(session, u)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.Redirect(http.StatusMovedPermanently, redirect)
			return nil
		} else {
			log.Print("Register user add error")
			c.Redirect(http.StatusMovedPermanently, registerURL)
			return nil
		}
	} else {
		log.Print("Register form bind error")
		c.Redirect(http.StatusMovedPermanently, registerURL)
		return nil
	}

	return nil
}
