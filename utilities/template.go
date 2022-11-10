package utilities

func GetTemplates() []string {
	// currently used templates
	// for now we declare them here upfront
	var templateFiles = []string{
		// layluts
		"./templates/layouts/home.html",
		"./templates/layouts/design.html",
		"./templates/layouts/islands.html",
		"./templates/layouts/bloghome.html",
		"./templates/layouts/blogpost.html",
		"./templates/layouts/blogposts.html",
		// partials
		"./templates/partials/designsystem.html",
		"./templates/partials/islandsexamples.html",
		"./templates/partials/head.html",
		"./templates/partials/header.html",
		"./templates/partials/products.html",
		"./templates/partials/prose.html",
		"./templates/partials/tools.html",
		"./templates/partials/importmaps.html",
		"./templates/partials/wsreload.html",

		"./templates/partials/bloghead.html",
		"./templates/partials/blogheader.html",
		"./templates/partials/blogmain.html",
		"./templates/partials/blogpostlist.html",
		"./templates/partials/blogfooter.html",
		"./templates/partials/blogpostpartial.html",
		"./templates/partials/blogpostsheader.html",
	}
	return templateFiles
}
