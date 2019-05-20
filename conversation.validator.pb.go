// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: conversation.proto

package proto

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ConvUser) Validate() error {
	return nil
}
func (this *ConversationEntry) Validate() error {
	for _, item := range this.Links {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
			}
		}
	}
	return nil
}
func (this *Conversation) Validate() error {
	for _, item := range this.Users {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
			}
		}
	}
	if this.Link != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Link); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Link", err)
		}
	}
	for _, item := range this.Entries {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Entries", err)
			}
		}
	}
	return nil
}
func (this *GetByUserIDMessage) Validate() error {
	return nil
}
func (this *GetByUserIDResponse) Validate() error {
	for _, item := range this.ConversationList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ConversationList", err)
			}
		}
	}
	return nil
}
func (this *GetByUserMeMessage) Validate() error {
	return nil
}
func (this *GetByUserMeResponse) Validate() error {
	for _, item := range this.ConversationList {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ConversationList", err)
			}
		}
	}
	return nil
}
func (this *GetByIDMessage) Validate() error {
	return nil
}
func (this *GetByIDResponse) Validate() error {
	if this.Conversation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Conversation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Conversation", err)
		}
	}
	return nil
}
func (this *CreateConversationEntryRequest) Validate() error {
	if this.ConversationEntry != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConversationEntry); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConversationEntry", err)
		}
	}
	return nil
}
func (this *CreateConversationEntryResponse) Validate() error {
	if this.ConversationEntry != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConversationEntry); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConversationEntry", err)
		}
	}
	return nil
}
func (this *UpdateConversationEntryRequest) Validate() error {
	return nil
}
func (this *UpdateConversationEntryResponse) Validate() error {
	if this.ConversationEntry != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConversationEntry); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConversationEntry", err)
		}
	}
	return nil
}
func (this *SubmitConversationEntryRequest) Validate() error {
	return nil
}
func (this *SubmitConversationEntryResponse) Validate() error {
	if this.ConversationEntry != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConversationEntry); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConversationEntry", err)
		}
	}
	return nil
}
func (this *GetConversationEntryDraftRequest) Validate() error {
	return nil
}
func (this *GetConversationEntryDraftResponse) Validate() error {
	if this.ConversationEntry != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConversationEntry); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConversationEntry", err)
		}
	}
	return nil
}
func (this *GetConversationEntriesRequest) Validate() error {
	return nil
}
func (this *GetConversationEntriesResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}