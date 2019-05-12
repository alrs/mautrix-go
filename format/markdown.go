// Copyright 2018 Tulir Asokan
package format

import (
	"strings"

	"gopkg.in/russross/blackfriday.v2"

	"maunium.net/go/mautrix"
)

func RenderMarkdown(text string) mautrix.Content {
	htmlBodyBytes := blackfriday.Run([]byte(text),
		blackfriday.WithExtensions(blackfriday.NoIntraEmphasis|
			blackfriday.Tables|
			blackfriday.FencedCode|
			blackfriday.Strikethrough|
			blackfriday.SpaceHeadings|
			blackfriday.DefinitionLists|
			blackfriday.HardLineBreak),
		blackfriday.WithRenderer(blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
			Flags: blackfriday.UseXHTML,
		})))
	htmlBody := strings.TrimRight(string(htmlBodyBytes), "\n")

	return mautrix.Content{
		FormattedBody: htmlBody,
		Format:        mautrix.FormatHTML,
		MsgType:       mautrix.MsgText,
		Body:          HTMLToText(htmlBody),
	}
}
