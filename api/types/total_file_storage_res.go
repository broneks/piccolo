package types

type TotalFileStorage struct {
	UsedMB         float32 `json:"usedMB"`
	UsedPercentage float32 `json:"usedPercentage"`
	TotalMB        float32 `json:"totalMB"`
}
