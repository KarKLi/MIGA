// Package param provides a parser for DNN hyperparameter and image process parameter.
package param

// IPParam defines input pre-process parameter.
type IPParam struct {
	// grayScale defines the ratio of grayscaling an image.
	//
	// This operation SHOULD NOT change the number of channel.
	//
	// i.e. If an image have 3 channels (such as RGB).
	//
	// After grayscale operation, this image should still have 3 channels.
	// (same as the origin image)
	grayScale float64

	// vfilping defines the ratio of an image will be vertical filped.
	vfilping float64

	// hfliping defines the ratio of an image will be horizontal filped.
	hfliping float64
}

// IPParamParser is a image pre-process parameter parser interface
type IPParamParser interface {
	ParseIPParam(IPparamFile string) (*IPParam, error)
}

// GetGrayScale returns the grayscale member.
func (ip *IPParam) GetGrayScale() float64 {
	return ip.grayScale
}

// GetVFliping returns the vfliping member.
func (ip *IPParam) GetVFliping() float64 {
	return ip.vfilping
}

// GetHFilping returns the hfliping member.
func (ip *IPParam) GetHFilping() float64 {
	return ip.hfliping
}