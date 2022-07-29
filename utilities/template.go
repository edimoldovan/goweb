package utilities

func GetTemplates() []string {
	// currently used templates
	// for now we declare them here upfront
	var templateFiles = []string{
		// pages
		"./views/pages/default.html",
		// layluts
		"./views/layouts/home.html",
		// partials
		"./views/partials/box.html",
	}
	return templateFiles
}
