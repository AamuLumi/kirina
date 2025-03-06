module kirina

go 1.24.1

require (
	github.com/go-gl/gl v0.0.0-20231021071112-07e5d0ea2e71
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20250301202403-da16c1255728
	kirina/generators v0.0.0-00010101000000-000000000000
	kirina/glTools v0.0.0-00010101000000-000000000000
)

require kirina/tools v0.0.0-00010101000000-000000000000 // indirect

replace kirina/generators => ./generators

replace kirina/glTools => ./glTools

replace kirina/tools => ./tools
