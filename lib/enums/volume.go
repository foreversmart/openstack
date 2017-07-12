package enums

type VolumeActionType string

const (
	VolumeActionExtend VolumeActionType = "extend"
	VolumeActionReset  VolumeActionType = "reset"
	VolumeActionAttach VolumeActionType = "attach"
	VolumeActionDetach VolumeActionType = "detach"
)

func (v VolumeActionType) IsValid() bool {
	switch v {
	case VolumeActionExtend, VolumeActionAttach, VolumeActionDetach, VolumeActionReset:
		return true
	}

	return false
}
