package model

type UpscaleRequest struct {
	Image                     string            `json:"image"`
	ResizeMode                UpscaleResizeMode `json:"resize_mode"`
	UpscalingResize           float32           `json:"upscaling_resize,omitempty"`
	UpscalingResizeW          int               `json:"upscaling_resize_w,omitempty"`
	UpscalingResizeH          int               `json:"upscaling_resize_h,omitempty"`
	UpscalingCrop             bool              `json:"upscaling_crop"`
	Upscaler1                 string            `json:"upscaler_1"`
	Upscaler2                 string            `json:"upscaler_2"`
	ExtrasUpscaler2Visibility float32           `json:"extras_upscaler_2_visibility"`
	GfpganVisibility          float32           `json:"gfpgan_visibility"`
	CodeformerVisibility      float32           `json:"codeformer_visibility"`
	CodeformerWeight          float32           `json:"codeformer_weight"`
}

// NewUpscaleRequest Get default UpscaleRequest. Set value to default value in stable diffusion web ui.
// You can use this method to create UpscaleRequest if you don't know each meaning of value,
// then change the value you want to override.
func NewUpscaleRequest(image string, upscalingResize float32) *UpscaleRequest {
	return &UpscaleRequest{
		Image:                     image,
		ResizeMode:                UpscaleBy,
		UpscalingResize:           upscalingResize,
		UpscalingResizeW:          0,
		UpscalingResizeH:          0,
		UpscalingCrop:             false,
		Upscaler1:                 "R-ESRGAN 4x+",
		Upscaler2:                 "",
		ExtrasUpscaler2Visibility: 0,
		GfpganVisibility:          0,
		CodeformerVisibility:      0,
		CodeformerWeight:          0,
	}
}

type UpscaleResizeMode int

const (
	UpscaleBy UpscaleResizeMode = 0
	UpscaleTo UpscaleResizeMode = 1
)
