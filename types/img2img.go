package types

type Img2ImgRequest struct {
	ModelName              string            `json:"model_name,omitempty"`
	SamplerName            Sampler           `json:"sampler_name,omitempty"`
	InitImages             []string          `json:"init_images,omitempty"`
	Mask                   string            `json:"mask,omitempty"`
	ResizeMode             int               `json:"resize_mode,omitempty"`
	DenoisingStrength      float32           `json:"denoising_strength,omitempty"`
	CfgScale               float32           `json:"cfg_scale,omitempty"`
	MaskBlur               int               `json:"mask_blur,omitempty"`
	InpaintingFill         int               `json:"inpainting_fill,omitempty"`
	InpaintFullRes         int               `json:"inpaint_full_res,omitempty"`
	InpaintFullResPadding  int               `json:"inpaint_full_res_padding,omitempty"`
	InpaintMaskInvert      int               `json:"inpaint_mask_invert,omitempty"`
	InitialNoiseMultiplier float32           `json:"initial_noise_multiplier,omitempty"`
	Prompt                 string            `json:"prompt,omitempty"`
	Seed                   int64             `json:"seed,omitempty"`
	NegativePrompt         string            `json:"negative_prompt,omitempty"`
	BatchSize              int               `json:"batch_size,omitempty"`
	NIter                  int               `json:"n_iter,omitempty"`
	Steps                  int               `json:"steps,omitempty"`
	Width                  int               `json:"width,omitempty"`
	Height                 int               `json:"height,omitempty"`
	RestoreFaces           bool              `json:"restore_face,omitemptys"`
	SDVae                  string            `json:"sd_vae,omitempty"`
	ClipSkip               int               `json:"clip_skip,omitempty"`
	ControlNetUnits        []*ControlNetUnit `json:"controlnet_units,omitempty"`
}

// NewImg2ImgRequest Get default Img2ImgRequest. Set value to default value in stable diffusion web ui.
// You can use this method to create Img2ImgRequest if you don't know each meaning of value,
// then change the value you want to override.
func NewImg2ImgRequest(prompt, negativePrompt, modelName string, initImage string) *Img2ImgRequest {
	return &Img2ImgRequest{
		Prompt:          prompt,
		NegativePrompt:  negativePrompt,
		ModelName:       modelName,
		InitImages:      []string{initImage},
		SamplerName:     EulerA,
		BatchSize:       1,
		NIter:           1,
		Steps:           20,
		CfgScale:        7,
		Height:          512,
		Width:           512,
		Seed:            -1,
		RestoreFaces:    false,
		SDVae:           "",
		ClipSkip:        0,
		ControlNetUnits: nil,
	}
}
