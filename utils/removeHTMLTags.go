package utils

import "regexp"

func RemoveHTMLTags(text string) string {
	r, _ := regexp.Compile("<[^>]*>")
	return r.ReplaceAllString(text, "")
}
