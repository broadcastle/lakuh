package web

import (
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"broadcastle.co/code/lakuh/pkg/library"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func songView(c echo.Context) error {

	track := library.Audio{}

	if err := track.Echo(c); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := track.View(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, track)

}

func songUpdate(c echo.Context) error {

	track := library.Audio{}

	if err := c.Bind(&track); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := track.Echo(c); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := track.Edit(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, track)

}

func songDelete(c echo.Context) error {

	track := library.Audio{}

	if err := track.Echo(c); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := track.Remove(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "ok")
}

func libraryView(c echo.Context) error {

	tracks, err := library.AllAudio()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, tracks)

}

func libraryAdd(c echo.Context) error {

	file, err := c.FormFile("audio")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer src.Close()

	f := viper.GetString("audio.storage")

	f = path.Join(f, file.Filename)

	dst, err := os.Create(f)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	year, err := strconv.Atoi(c.FormValue("year"))
	if err != nil {
		year = 0
	}

	track := library.Audio{
		Title:  c.FormValue("title"),
		Artist: c.FormValue("artist"),
		Album:  c.FormValue("album"),
		Year:   year,
	}

	go func() {

		defer os.Remove(f)

		if err := track.Import(f); err != nil {
			logrus.Warn(err)
		}

	}()

	return c.JSON(http.StatusOK, "importing")

}
