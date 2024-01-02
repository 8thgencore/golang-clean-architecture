package group

import (
	"app/services/contact/internal/domain/group"
)

func ProtoToGroupResponse(response group.Reader) *GroupResponse {
	return &GroupResponse{
		ID:         response.ID().String(),
		CreatedAt:  response.CreatedAt(),
		ModifiedAt: response.ModifiedAt(),
		Group: Group{
			ShortGroup: ShortGroup{
				Name:        response.Name().Value(),
				Description: response.Description().Value(),
			},
			ContactsAmount: response.ContactCount(),
		},
	}
}

