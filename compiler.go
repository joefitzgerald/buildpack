package buildpack

type Compiler interface {
	// Compile builds the droplet that will be run by the DEA and will therefore
	// contain all the components necessary to run the application.
	//
	// It is run with two arguments, the build directory for the application and
	// the cache directory, which is a location the buildpack can use to store
	// assets during the build process.
	//
	// During execution all output sent to STDOUT will be relayed back to the end
	// user.
	Compile(buildDir string, cacheDir string) int
}
