package ui

import (
	"github.com/creamsensation/cp"
	. "github.com/creamsensation/gox"
)

func Layout(c cp.Control, nodes ...Node) Node {
	return Html(
		Head(
			Title(Text(c.Page().Get().Title())),
			Meta(CharSet("UTF-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0, minimum-scale=1.0")),
			Raw(`<link rel="shortcut icon" href="/.static/icon/favicon.ico">`),
			Raw(
				`
					<link rel="preconnect" href="https://fonts.googleapis.com">
					<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
					<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600;700&display=swap" rel="stylesheet">
			`,
			),
			c.Assets().GetStyles(),
			c.Assets().GetScripts(),
		),
		Body(
			Class("bg-slate-100 font-inter overflow-x-hidden overflow-y-auto h-screen"),
			Fragment(nodes...),
			c.Dev().Tool(),
		),
	)
}
