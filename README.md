## PAC

###### Pyramid Application Constructor


#### About

The Pyramid App Constructor (PAC) is a toolkit to help jumpstart the application development process, specifically designed for compressed time-duration events like hackathons and tech challenges. PAC will generate scaffolding composed of reusable components, templates, and pipelines to help accelerate development velocity, while ensuring security and quality discipline, to achieve acceptable software hygiene. PAC is an evolving toolkit, and currently supports the MERN stack (MongoDB, Express, React, Node). It leverages Jenkins for pipelines, Auth0 for authentication, AWS as the cloud platform, and is supported by relevant open source libraries


#### Installing (Temporary Instructions)

1. [Install Go](https://golang.org/doc/install)

2. Ensure `$GOPATH` is set (usually ~/go) and $GOPATH/bin is in your `$PATH`

3. `go get -u github.com/PyramidSystemsInc/go/...` installs Pyramid's Golang helper functions

4. `GIT_TERMINAL_PROMPT=1 go get -u github.com/PyramidSystemsInc/pac` clones the project using git credentials

5. `go install github.com/PyramidSystemsInc/pac` installs pac CLI

6. Start by trying `pac create --help`
