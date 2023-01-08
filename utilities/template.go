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
	}
	return templateFiles
}
