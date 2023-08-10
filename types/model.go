package types

import "sort"

type ModelRequest struct {
}

type ModelsResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data *ModelData `json:"data"`
}

type ModelData struct {
	Models []*Model `json:"models"`
}

type Model struct {
	Name                         string          `json:"name"`
	Hash                         string          `json:"hash"`
	SdName                       string          `json:"sd_name"`
	ThirdSource                  string          `json:"third_source"`
	DownloadStatus               int             `json:"download_status"`
	DownloadName                 string          `json:"download_name"`
	DependencyStatus             int             `json:"dependency_status"`
	Type                         ModelType       `json:"type"`
	CivitaiLink                  string          `json:"civitai_link,omitempty"`
	CivitaiModelId               int             `json:"civitai_model_id,omitempty"`
	CivitaiVersionId             int             `json:"civitai_version_id,omitempty"`
	CivitaiNsfw                  bool            `json:"civitai_nsfw"`
	CivitaiImages                []*CivitaiImage `json:"civitai_images,omitempty"`
	CivitaiDownloadUrl           string          `json:"civitai_download_url,omitempty"`
	CivitaiAllowCommercialUse    string          `json:"civitai_allow_commercial_use,omitempty"`
	CivitaiAllowDifferentLicense bool            `json:"civitai_allow_different_license"`
	CivitaiCreateAt              string          `json:"civitai_create_at,omitempty"`
	CivitaiUpdateAt              string          `json:"civitai_update_at,omitempty"`
	CivitaiTrainedWords          string          `json:"civitai_trained_words,omitempty"`
	CivitaiTags                  string          `json:"civitai_tags,omitempty"`
	CivitaiDownloadCount         int64           `json:"civitai_download_count,omitempty"`
	CivitaiFavoriteCount         int64           `json:"civitai_favorite_count,omitempty"`
	CivitaiCommentCount          int64           `json:"civitai_comment_count,omitempty"`
	CivitaiRatingCount           int64           `json:"civitai_rating_count,omitempty"`
	CivitaiRating                float32         `json:"civitai_rating,omitempty"`
	CivitaiDependencyVersionId   int64           `json:"civitai_dependency_version_id,omitempty"`
	CivitaiDependencyModelName   string          `json:"civitai_dependency_model_name,omitempty"`
	OmniUsedCount                int64           `json:"omni_used_count,omitempty"`
	CivitaiImageUrl              string          `json:"civitai_image_url,omitempty"`
	CivitaiImageNsfw             string          `json:"civitai_image_nsfw,omitempty"`
	CivitaiOriginImageUrl        string          `json:"civitai_origin_image_url,omitempty"`
	CivitaiImagePrompt           string          `json:"civitai_image_prompt,omitempty"`
	CivitaiImageNegativePrompt   string          `json:"civitai_image_negative_prompt,omitempty"`
	CivitaiImageSamplerName      string          `json:"civitai_image_sampler_name,omitempty"`
	CivitaiImageHeight           int             `json:"civitai_image_height,omitempty"`
	CivitaiImageWidth            int             `json:"civitai_image_width,omitempty"`
	CivitaiImageSteps            int             `json:"civitai_image_steps,omitempty"`
	CivitaiImageCfgScale         float32         `json:"civitai_image_cfg_scale,omitempty"`
	CivitaiImageSeed             int64           `json:"civitai_image_seed,omitempty"`
}

type CivitaiImage struct {
	URL  string            `json:"url"`
	NSFW string            `json:"nsfw"`
	Meta *CivitaiImageMeta `json:"meta"`
}

type CivitaiImageMeta struct {
	Prompt         string  `json:"prompt"`
	NegativePrompt string  `json:"negative_prompt"`
	SamplerName    string  `json:"sampler_name"`
	Steps          int     `json:"steps"`
	CfgScale       float32 `json:"cfg_scale"`
	Seed           int     `json:"seed"`
	Height         int     `json:"height"`
	Width          int     `json:"width"`
	ModelName      string  `json:"model_name"`
}

func (m ModelsResponse) GetCode() int {
	return m.Code
}

func (m ModelsResponse) GetMsg() string {
	return m.Msg
}

type ModelType string

const (
	Checkpoint    ModelType = "checkpoint"
	Lora          ModelType = "lora"
	VAE           ModelType = "vae"
	ControlNet    ModelType = "controlnet"
	TextInversion ModelType = "textualinversion"
	Upscaler      ModelType = "upscaler"
)

type ModelList []*Model

func (m ModelList) FilterType(typ ModelType) ModelList {
	rs := make([]*Model, 0)
	for _, model := range m {
		if model.Type == typ {
			rs = append(rs, model)
		}
	}
	return rs
}

func (m ModelList) FilterNsfw(nsfw bool) ModelList {
	rs := make([]*Model, 0)
	for _, model := range m {
		if model.CivitaiNsfw == nsfw {
			rs = append(rs, model)
		}
	}
	return rs
}

func (m ModelList) TopN(n int, valueFn func(m *Model) float32) ModelList {
	if n <= 0 {
		return ModelList{}
	}
	if n >= len(m) {
		return m
	}
	sort.Slice(m, func(i, j int) bool {
		return valueFn(m[i]) >= valueFn(m[j])
	})
	return m[:n]
}

func (m ModelList) FilterCivitaiVersionId(civitaiVersionId int) *Model {
	for _, model := range m {
		if model.CivitaiVersionId == civitaiVersionId {
			return model
		}
	}
	return nil
}
