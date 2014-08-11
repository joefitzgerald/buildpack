# buildpack

## Overview

A collection of interfaces representing the contracts required for a buildpack.

### Heroku

From the [Heroku](http://heroku.com) [buildpack API](https://devcenter.heroku.com/articles/buildpacks):

* Detector: Detects whether the buildpack should be run on an app
* Compiler: Compiles an app, building an executable artifact
* Releaser: Releases an app, defining metadata used to execute the app
* Heroku: Encapsulates the [Heroku](http://heroku.com) [buildpack API](https://devcenter.heroku.com/articles/buildpacks)

From the [Cloud Foundry](http://cloudfoundry.org)
[buildpack API]:

* Packager: An extension to the Heroku buildpack API to allow buildpacks to
  create an online / offline artifact representing the buildpack
* CloudFoundry: Encapsulates the [Cloud Foundry](http://cloudfoundry.org)
[buildpack API](http://docs.cloudfoundry.org/buildpacks/)
