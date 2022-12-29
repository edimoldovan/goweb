package utilities

// currently used templates
func GetTemplates() []string {
	f := GetExecutable()

	// for now we declare them here upfront
	var templateFiles = []string{
		// layouts
		f + "/templates/layouts/home.html",
		f + "/templates/layouts/design.html",
		f + "/templates/layouts/islands.html",
		f + "/templates/layouts/bloghome.html",
		f + "/templates/layouts/blogpost.html",
		f + "/templates/layouts/blogposts.html",
		// partials
		f + "/templates/partials/designsystem.html",
		f + "/templates/partials/islandsexamples.html",
		f + "/templates/partials/head.html",
		f + "/templates/partials/header.html",
		f + "/templates/partials/products.html",
		f + "/templates/partials/prose.html",
		f + "/templates/partials/tools.html",
		f + "/templates/partials/importmaps.html",
		f + "/templates/partials/wsreload.html",

		f + "/templates/partials/bloghead.html",
		f + "/templates/partials/blogheader.html",
		f + "/templates/partials/blogmain.html",
		f + "/templates/partials/blogpostlist.html",
		f + "/templates/partials/blogfooter.html",
		f + "/templates/partials/blogpostpartial.html",
		f + "/templates/partials/blogpostsheader.html",
		f + "/templates/partials/bloglastpost.html",

		// posts
		f + "/templates/layouts/posts/not-blog-post-layout.html",
		f + "/templates/partials/posts/not-blog-post.html",
	}
	return templateFiles
}
