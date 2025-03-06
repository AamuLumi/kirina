package glTools

import (
	"errors"
	"image"

	"github.com/go-gl/gl/v4.1-core/gl"
)

var errTextureNotBound = errors.New("texture not bound")

// Texture contains all elements for an OpenGL texture
type Texture struct {
	handle  uint32
	target  uint32 // same target as gl.BindTexture(<this param>, ...)
	texUnit uint32 // Texture unit that is currently bound to ex: gl.TEXTURE0
}

// NewTexture creates a new texture
func NewTexture(RGBA64 *image.RGBA64, wrapR, wrapS int32) (*Texture, error) {
	var handle uint32
	gl.GenTextures(1, &handle)

	target := uint32(gl.TEXTURE_2D)
	internalFmt := int32(gl.RGBA)
	format := uint32(gl.RGBA)
	width := int32(RGBA64.Rect.Size().X)
	height := int32(RGBA64.Rect.Size().Y)
	pixType := uint32(gl.UNSIGNED_BYTE)
	dataPtr := gl.Ptr(RGBA64.Pix)

	texture := Texture{
		handle: handle,
		target: target,
	}

	texture.Bind(gl.TEXTURE0)

	// set the texture wrapping/filtering options (applies to current bound texture obj)
	// TODO-cs
	gl.TexParameteri(texture.target, gl.TEXTURE_MIN_FILTER, gl.LINEAR) // minification filter
	gl.TexParameteri(texture.target, gl.TEXTURE_MAG_FILTER, gl.LINEAR) // magnification filter
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.TexImage2D(target, 0, internalFmt, width, height, 0, format, pixType, dataPtr)

	gl.GenerateMipmap(texture.handle)

	return &texture, nil
}

// Update updates a texture
func (tex *Texture) Update(RGBA64 *image.RGBA64) {
	tex.Bind(gl.TEXTURE0)

	target := uint32(gl.TEXTURE_2D)
	format := uint32(gl.RGBA)
	width := int32(RGBA64.Rect.Size().X)
	height := int32(RGBA64.Rect.Size().Y)
	pixType := uint32(gl.UNSIGNED_BYTE)

	var data []uint8
	length := RGBA64.Rect.Size().X * RGBA64.Rect.Size().Y * 4
	cur := RGBA64.Pix

	for i := 0; i < length; i++ {
		data = append(data, cur[2*i])
	}

	dataPtr := gl.Ptr(data)

	gl.TexSubImage2D(target, 0, 0, 0, width, height, format, pixType, dataPtr)
}

// Bind enables a texture
func (tex *Texture) Bind(texUnit uint32) {
	gl.ActiveTexture(texUnit)
	gl.BindTexture(tex.target, tex.handle)
	tex.texUnit = texUnit
}

// Unbind disables a texture
func (tex *Texture) Unbind() {
	tex.texUnit = 0
	gl.BindTexture(tex.target, 0)
}

// SetUniform set a uniform texture
func (tex *Texture) SetUniform(uniformLoc int32) error {
	if tex.texUnit == 0 {
		return errTextureNotBound
	}
	gl.Uniform1i(uniformLoc, int32(tex.texUnit-gl.TEXTURE0))
	return nil
}
