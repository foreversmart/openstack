package enums

type VolumeActionType string

const (
	Extend VolumeActionType = "extend"
	Attach VolumeActionType = "attach"
	Detach VolumeActionType = "detach"
)

func (v VolumeActionType) IsValid() bool {
	switch v {
	case Extend, Attach, Detach:
		return true
	}

	return false
}
