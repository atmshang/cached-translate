package translate

import (
	"errors"
	"github.com/atmshang/plog"
	"github.com/bregydoc/gtranslate"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

func blockTranslate(text string, from string, to string) (string, error) {
	translated, err := gtranslate.TranslateWithParams(text, gtranslate.TranslationParams{
		From: from,
		To:   to,
	})
	if err != nil {
		return "", err
	}
	return translated, nil
}

func createId(text string, from string, to string) string {
	return "|" + text + "|" + from + "|" + to + "|"
}

func translateByDatabase(text string, from string, to string) (string, error) {
	db := getDBInstance()
	var cache Cache
	ret := db.Where("id = ?", createId(text, from, to)).First(&cache)
	if ret.Error != nil {
		return "", ret.Error
	}
	return cache.Translated, nil
}

func insertTranslated(text string, from string, to string, translated string) {
	db := getDBInstance()
	var cache Cache
	cache.ID = createId(text, from, to)
	cache.Text = text
	cache.From = from
	cache.To = to
	cache.Translated = translated
	ret := db.Where("id = ?", createId(text, from, to)).First(&Cache{})
	if errors.Is(ret.Error, gorm.ErrRecordNotFound) {
		ret := db.Create(&cache)
		if ret.Error != nil {
			log.Println("插入翻译", ret.Error)
			return
		}
	} else {
		ret := db.Save(&cache)
		if ret.Error != nil {
			log.Println("更新翻译", ret.Error)
			return
		}
	}
}

func blockTranslateWithCache(text string, from string, to string) string {
	translated, err := translateByDatabase(text, from, to)
	if err == nil {
		return translated
	}
	translated, err = blockTranslate(text, from, to)
	if err == nil {
		insertTranslated(text, from, to, translated)
		return translated
	}
	return text
}

func I18n(text string, from string, to string) string {
	if from == to {
		return text
	}
	return blockTranslateWithCache(text, from, to)
}

func QuickI18nFromRequest(text string, r *http.Request) string {
	form := "zh"
	to := GetPreferredLanguageFromRequest(r)
	return I18n(text, form, to)
}

func Test() {
	text := "Hello World"
	translated := I18n(text, "en", "ja")
	plog.Println("translated:", translated)
}

// GetPreferredLanguageFromRequest 从浏览器的请求头中获取偏好语言，如果没有，返回英语
func GetPreferredLanguageFromRequest(r *http.Request) string {
	acceptLanguage := r.Header.Get("Accept-Language")
	languages := strings.Split(acceptLanguage, ",")
	if len(languages) > 0 {
		preferredLanguage := strings.Split(languages[0], ";")[0]
		return preferredLanguage
	}
	return "en"
}
