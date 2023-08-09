package model

type Txt2ImgRequest struct {
	Prompt          string            `json:"prompt,omitempty"`
	NegativePrompt  string            `json:"negative_prompt,omitempty"`
	ModelName       string            `json:"model_name,omitempty"`
	SamplerName     Sampler           `json:"sampler_name,omitempty"`
	BatchSize       int               `json:"batch_size,omitempty"`
	NIter           int               `json:"n_iter,omitempty"`
	Steps           int               `json:"steps,omitempty"`
	CfgScale        float32           `json:"cfg_scale,omitempty"`
	Height          int               `json:"height,omitempty"`
	Width           int               `json:"width,omitempty"`
	Seed            int64             `json:"seed"`
	RestoreFaces    bool              `json:"restore_faces,omitempty"`
	SDVae           string            `json:"sd_vae,omitempty"`
	ClipSkip        int               `json:"clip_skip,omitempty"`
	EnableHr        bool              `json:"enable_hr,omitempty"`
	HrUpscaler      string            `json:"hr_upscaler,omitempty"`
	HrScale         float32           `json:"hr_scale,omitempty"`
	HrResizeX       int               `json:"hr_resize_x,omitempty"`
	HrResizeY       int               `json:"hr_resize_y,omitempty"`
	ControlNetUnits []*ControlNetUnit `json:"controlnet_units,omitempty"`
}

// NewTxt2ImgRequest Get default Txt2ImgRequest. Set value to default value in stable diffusion web ui.
// You can use this method to create Txt2ImgRequest if you don't know each meaning of value,
// then change the value you want to override.
func NewTxt2ImgRequest(prompt, negativePrompt, modelName string) *Txt2ImgRequest {
	return &Txt2ImgRequest{
		Prompt:          prompt,
		NegativePrompt:  negativePrompt,
		ModelName:       modelName,
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

type Sampler string

const (
	EulerA         Sampler = "Euler a"
	Euler          Sampler = "Euler"
	LMS            Sampler = "LMS"
	HEUN           Sampler = "Heun"
	DPM2           Sampler = "DPM2"
	DPM2A          Sampler = "DPM2 a"
	DPM2Karras     Sampler = "DPM2 Karras"
	DPM2AKarras    Sampler = "DPM2 a Karras"
	DPMPPSA        Sampler = "DPM++ 2S a"
	DPMPPM         Sampler = "DPM++ 2M"
	DPMPPSDE       Sampler = "DPM++ SDE"
	DPMPPKarras    Sampler = "DPM++ Karras"
	DPMPPSAKarras  Sampler = "DPM++ 2S a Karras"
	DPMPPMKarras   Sampler = "DPM++ 2M Karras"
	DPMPPSDEKarras Sampler = "DPM++ SDE Karras"
	DDIM           Sampler = "DDIM"
	PLMS           Sampler = "PLMS"
	UNIPC          Sampler = "UniPC"
)

type ControlNetUnit struct {
	Model         string                `json:"model,omitempty"`
	Weight        float32               `json:"weight"`
	Module        ControlNetModule      `json:"module,omitempty"`
	InputImage    string                `json:"input_image,omitempty"`
	ControlMode   ControlNetControlMode `json:"control_mode"`
	ResizeMode    ControlNetResizeMode  `json:"resize_mode"`
	Mask          string                `json:"mask,omitempty"`
	ProcessorRes  int                   `json:"processor_res,omitempty"`
	ThresholdA    int                   `json:"threshold_a,omitempty"`
	ThresholdB    int                   `json:"threshold_b,omitempty"`
	GuidanceStart float32               `json:"guidance_start"`
	GuidanceEnd   float32               `json:"guidance_end"`
	PixelPerfect  bool                  `json:"pixel_perfect,omitempty"`
}

// NewControlNetUnit Get default ControlNetUnit. Set value to default value in stable diffusion web ui.
// You can use this method to create ControlNetUnit if you don't know each meaning of value,
// then change the value you want to override.
func NewControlNetUnit(module ControlNetModule, model string, inputImage string) *ControlNetUnit {
	return &ControlNetUnit{
		Model:         model,
		Weight:        1,
		Module:        module,
		InputImage:    inputImage,
		ControlMode:   Balanced,
		ResizeMode:    CropAndResize,
		GuidanceStart: 0,
		GuidanceEnd:   1,
		Mask:          "",
		PixelPerfect:  false,
	}
}

type ControlNetControlMode int

const (
	Balanced             ControlNetControlMode = 0
	PromptImportance     ControlNetControlMode = 1
	ControlNetImportance ControlNetControlMode = 2
)

type ControlNetResizeMode int

const (
	JustResize    ControlNetResizeMode = 0
	CropAndResize ControlNetResizeMode = 1
	ResizeAndFill ControlNetResizeMode = 2
)

type ControlNetModule string

const (
	None                   ControlNetModule = "none"
	Canny                  ControlNetModule = "canny"
	Depth                  ControlNetModule = "depth"
	DepthLeRes             ControlNetModule = "depth_leres"
	DepthLeResPlusPlus     ControlNetModule = "depth_leres++"
	HED                    ControlNetModule = "hed"
	HEDSafe                ControlNetModule = "hed_safe"
	MediaPipeFace          ControlNetModule = "mediapipe_face"
	MLSD                   ControlNetModule = "mlsd"
	NormalMap              ControlNetModule = "normal_map"
	OpenPose               ControlNetModule = "openpose"
	OpenPoseHand           ControlNetModule = "openpose_hand"
	OpenPoseFace           ControlNetModule = "openpose_face"
	OpenPoseFaceOnly       ControlNetModule = "openpose_faceonly"
	OpenPoseFull           ControlNetModule = "openpose_full"
	ClipVision             ControlNetModule = "clip_vision"
	Color                  ControlNetModule = "color"
	PIDINET                ControlNetModule = "pidinet"
	PIDINESafe             ControlNetModule = "pidinet_safe"
	PIDINESketch           ControlNetModule = "pidinet_sketch"
	PIDINEScribble         ControlNetModule = "pidinet_scribble"
	ScribbleXDOG           ControlNetModule = "scribble_xdog"
	ScribbleHED            ControlNetModule = "scribble_hed"
	Segmentation           ControlNetModule = "segmentation"
	Threshold              ControlNetModule = "threshold"
	DepthZOE               ControlNetModule = "depth_zoe"
	NormalBAE              ControlNetModule = "normal_bae"
	OneFormerCOCO          ControlNetModule = "oneformer_coco"
	OneFormerADE20K        ControlNetModule = "oneformer_ade20k"
	Lineart                ControlNetModule = "lineart"
	LineartCoarse          ControlNetModule = "lineart_coarse"
	LineartAnime           ControlNetModule = "lineart_anime"
	LineartStandard        ControlNetModule = "lineart_standard"
	Shuffle                ControlNetModule = "shuffle"
	TileResample           ControlNetModule = "tile_resample"
	Invert                 ControlNetModule = "invert"
	LineartAnimeDenoise    ControlNetModule = "lineart_anime_denoise"
	ReferenceOnly          ControlNetModule = "reference_only"
	ReferenceADAIN         ControlNetModule = "reference_adain"
	ReferenceADAINPlusAttn ControlNetModule = "reference_adain+attn"
	Inpaint                ControlNetModule = "inpaint"
	InpaintOnly            ControlNetModule = "inpaint_only"
	InpaintOnlyPlusLAMA    ControlNetModule = "inpaint_only+lama"
	TileColorFix           ControlNetModule = "tile_colorfix"
	TileColorFixPlusSharp  ControlNetModule = "tile_colorfix+sharp"
)

type AsyncResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data *AsyncResponseData `json:"data"`
}

func (a AsyncResponse) GetCode() int {
	return a.Code
}

func (a AsyncResponse) GetMsg() string {
	return a.Msg
}

type AsyncResponseData struct {
	TaskID string `json:"task_id"`
	Warn   string `json:"warn"`
}
