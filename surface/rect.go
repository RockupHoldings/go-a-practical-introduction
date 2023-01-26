package surface

// Rect calculates the surface area of a rectangle. This
// function is *exported*.
func Rect(x, y int) int {
	return calculateRectSurface(x, y)
}

// this function is not exported, therefore you cannot import
// it from another package. Go doesn't even generate documentation
// for this function! (one of the Go proverbs is 'documentation is
// for users', meaning internal implementation concerns should not
// be documented).
//
// If you want to see the auto generated docs, you can follow these
// steps:
// - first you need to install the docs tool: go install golang.org/x/tools/cmd/godoc@latest
// - then you can run them: godoc -http=:8080 -index
//
// Open your browser and go to: http://localhost:8080
// You should see your docs under StandardLibrary > [module name], eg. for
// this project, the docs will be under prac_go, since that is my module's
// name.
func calculateRectSurface(x, y int) int {
	return x * y
}
