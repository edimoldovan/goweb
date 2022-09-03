package utilities

func GetTemplates() []string {
	// currently used templates
	// for now we declare them here upfront
	var templateFiles = []string{
		// layluts
		"./views/layouts/home.html",
		"./views/layouts/bloghome.html",
		"./views/layouts/blogpost.html",
		"./views/layouts/blogposts.html",
		// partials
		"./views/partials/head.html",
		"./views/partials/header.html",
		"./views/partials/products.html",
		"./views/partials/prose.html",
		"./views/partials/tools.html",
		"./views/partials/importmaps.html",
		"./views/partials/wsreload.html",

		"./views/partials/bloghead.html",
		"./views/partials/blogheader.html",
		"./views/partials/blogmain.html",
		"./views/partials/blogpostlist.html",
		"./views/partials/blogfooter.html",
		"./views/partials/blogpostpartial.html",
		"./views/partials/blogpostsheader.html",
	}
	return templateFiles
}
