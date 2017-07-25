package enums

type ImageVisibility string

const (
	ImageVisibilityPublic    ImageVisibility = "public"
	ImageVisibilityCommunity ImageVisibility = "community"
	ImageVisibilityShared    ImageVisibility = "shared"
	ImageVisibilityPrivate   ImageVisibility = "private"
)

func (iv ImageVisibility) IsValid() bool {
	switch iv {
	case ImageVisibilityPublic, ImageVisibilityCommunity, ImageVisibilityShared, ImageVisibilityPrivate:
		return true
	}

	return false
}
