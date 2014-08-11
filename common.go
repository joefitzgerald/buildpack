package buildpack

// Heroku represents the Heroku buildpack API.
type Heroku interface {
	Detect(buildDir string) int
	Compile(buildDir string, cacheDir string) int
	Release(buildDir string)
}
