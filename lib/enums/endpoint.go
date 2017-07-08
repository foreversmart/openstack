package enums

const (
	EndpointInterfacePublic   EndpointInterface = "public"
	EndpointInterfaceInternal EndpointInterface = "internal"
	EndpointInterfaceAdmin    EndpointInterface = "admin"
)

type EndpointInterface string

func (ei EndpointInterface) IsValid() bool {
	switch ei {
	case EndpointInterfacePublic, EndpointInterfaceInternal, EndpointInterfaceAdmin:
		return true
	}

	return false
}
