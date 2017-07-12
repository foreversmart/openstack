package enums

type VolumeActionType string

const (
	VolumeActionExtend VolumeActionType = "extend"
	VolumeActionAttach VolumeActionType = "attach"
	VolumeActionDetach VolumeActionType = "detach"
)

func (v VolumeActionType) IsValid() bool {
	switch v {
	case VolumeActionExtend, VolumeActionAttach, VolumeActionDetach:
		return true
	}

	return false
}
