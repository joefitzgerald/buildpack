package buildpack

// PackageMode defines the modes which can be used during Packaging.
type PackageMode int

const (
	// Offline packaging implies that the buildpack will work without internet
	// connectivity. All buildpack dependencies are included in the package.
	Offline PackageMode = iota
	// Online packaging implies that external dependencies may be downloaded from
	// a network location.
	Online
)

// Packager produces buildpacks archives.
type Packager interface {
	// Package provides buildpack artifacts, which are provided to Cloud Foundry
	// as system buildpacks. Package is intended to provide a way for buildpack
	// developers to package a buildpack with its dependencies. Implementing this
	// is entirely optional. It is a Cloud Foundry extension to the Heroku Buildpack
	// API.
	//
	// The result of running the package script must be a zip file that adheres to
	// the api outlined above, using detect, compile, release.
	//
	// The package script accepts two commands as arguments:
	//
	// 	package online produces a zip of the buildpack, preferably excluding unnecessary content such as tests.
	// 	package offline produces a zip of the buildpack with all dependencies cached.
	// 	package offline promises to produce a buildpack that does not need internet connectivity to run.
	Package(mode PackageMode)
}
