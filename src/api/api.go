package api

import (
	"context"
	"encoding/json"

	"github.com/datti-to/signal-cli-grpc-api/client"
	pb "github.com/datti-to/signal-cli-grpc-api/proto"
	"github.com/datti-to/signal-cli-grpc-api/utils"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SignalCliGroupEntry struct {
	Name              string   `json:"name"`
	Id                string   `json:"id"`
	IsMember          bool     `json:"isMember"`
	IsBlocked         bool     `json:"isBlocked"`
	Members           []string `json:"members"`
	PendingMembers    []string `json:"pendingMembers"`
	RequestingMembers []string `json:"requestingMembers"`
	GroupInviteLink   string   `json:"groupInviteLink"`
}

type SignalService struct {
	pb.UnimplementedSignalServiceServer
	signalClient *client.SignalClient
}

func NewSignalService(signalClient *client.SignalClient) *SignalService {
	return &SignalService{
		signalClient: signalClient,
	}
}

func (s *SignalService) About(ctx context.Context, _ *emptypb.Empty) (*pb.AboutResponse, error) {

	a := s.signalClient.About()

	return &pb.AboutResponse{
		Build:                int32(a.BuildNr),
		SupportedApiVersions: a.SupportedApiVersions,
	}, nil
}

func (s *SignalService) Health(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *SignalService) RegisterNumber(ctx context.Context, in *pb.RegisterNumberRequest) (*empty.Empty, error) {
	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	err := s.signalClient.RegisterNumber(in.Number, in.UseVoice, in.Captcha)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &empty.Empty{}, nil
}

func (s *SignalService) VerifyRegisteredNumber(ctx context.Context, in *pb.VerifyRegisteredNumberRequest) (*empty.Empty, error) {
	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	if in.Token == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a verification code")
	}

	err := s.signalClient.VerifyRegisteredNumber(in.Number, in.Token, in.Pin)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	return &empty.Empty{}, nil
}

func (s *SignalService) Send(ctx context.Context, in *pb.SendRequest) (*pb.SendResponse, error) {

	base64Attachments := []string{}
	if in.Base64Attachment != "" {
		base64Attachments = append(base64Attachments, in.Base64Attachment)
	}

	timestamp, err := s.signalClient.SendV1(in.Number, in.Message, in.Recipients, base64Attachments, in.IsGroup)
	if err != nil {
		return nil, err
	}

	return &pb.SendResponse{
		Timestamp: &timestamppb.Timestamp{
			Seconds: timestamp.Timestamp,
		},
	}, nil
}

func (s *SignalService) SendV2(ctx context.Context, in *pb.SendV2Request) (*pb.SendResponse, error) {
	if len(in.Recipients) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Please provide at least one recipient")
	}

	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	timestamp, err := s.signalClient.SendV2(in.Number, in.Message, in.Recipients, in.Base64Attachments)
	if err != nil {
		return nil, err
	}

	return &pb.SendResponse{
		Timestamp: &timestamppb.Timestamp{
			Seconds: (*timestamp)[0].Timestamp,
		},
	}, nil
}

func (s *SignalService) Receive(ctx context.Context, in *pb.ReceiveRequest) (*pb.ReceiveResponse, error) {
	if in.Timeout == 0 {
		in.Timeout = 1
	}

	jsonStr, err := s.signalClient.Receive(in.Number, int64(in.Timeout))
	if err != nil {
		return nil, err
	}

	slice_messages := []string{}

	if err := json.Unmarshal([]byte(jsonStr), &slice_messages); err != nil {
		return nil, err
	}

	return &pb.ReceiveResponse{
		Messages: slice_messages,
	}, nil
}

func (s *SignalService) CreateGroup(ctx context.Context, in *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	if in.Permissions.AddMembers != "" && !utils.StringInSlice(in.Permissions.AddMembers, []string{"every-member", "only-admins"}) {
		return nil, status.Error(codes.InvalidArgument, "Invalid edit group permissions provided - only 'every-member' and 'only-admins' allowed!")
	}

	if in.Permissions.EditGroup != "" && !utils.StringInSlice(in.Permissions.EditGroup, []string{"every-member", "only-admins"}) {
		return nil, status.Error(codes.InvalidArgument, "Invalid add members permission provided - only 'every-member' and 'only-admins' allowed!")
	}

	if in.GroupLink != "" && !utils.StringInSlice(in.GroupLink, []string{"enabled", "enabled-with-approval", "disabled"}) {
		return nil, status.Error(codes.InvalidArgument, "Invalid group link provided - only 'enabled', 'enabled-with-approval' and 'disabled' allowed!")
	}

	editGroupPermission := client.DefaultGroupPermission
	addMembersPermission := client.DefaultGroupPermission
	groupLinkState := client.DefaultGroupLinkState

	groupId, err := s.signalClient.CreateGroup(in.Number, in.Name, in.Members, in.Description, editGroupPermission, addMembersPermission, groupLinkState)
	if err != nil {
		return nil, err
	}

	return &pb.CreateGroupResponse{
		Id: groupId,
	}, nil
}

func (s *SignalService) GetGroups(ctx context.Context, in *pb.GetGroupsRequest) (*pb.GetGroupsResponse, error) {
	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	groups, err := s.signalClient.GetGroups(in.Number)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	grpc_groups := []*pb.GetGroupResponse{}

	for _, group := range groups {
		grpc_groups = append(grpc_groups, &pb.GetGroupResponse{
			Name:            group.Name,
			Id:              group.Id,
			InternalId:      group.InternalId,
			Members:         group.Members,
			Blocked:         group.Blocked,
			PendingInvites:  group.PendingInvites,
			PendingRequests: group.PendingRequests,
			InviteLink:      group.InviteLink,
		})
	}

	return &pb.GetGroupsResponse{
		Groups: grpc_groups,
	}, nil
}

func (s *SignalService) GetGroup(ctx context.Context, in *pb.GroupRequest) (*pb.GetGroupResponse, error) {
	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	if in.Groupid == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a group id")
	}

	group, err := s.signalClient.GetGroup(in.Number, in.Groupid)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if group == nil {
		return nil, status.Error(codes.NotFound, "Group not found")
	}

	return &pb.GetGroupResponse{
		Name:            group.Name,
		Id:              group.Id,
		InternalId:      group.InternalId,
		Members:         group.Members,
		Blocked:         group.Blocked,
		PendingInvites:  group.PendingInvites,
		PendingRequests: group.PendingRequests,
		InviteLink:      group.InviteLink,
	}, nil
}

func (s *SignalService) DeleteGroup(ctx context.Context, in *pb.GroupRequest) (*empty.Empty, error) {
	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	if in.Groupid == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a group id")
	}

	groupId, err := client.ConvertGroupIdToInternalGroupId(in.Groupid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.signalClient.DeleteGroup(in.Number, groupId)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &empty.Empty{}, nil
}

func (s *SignalService) GetQrCodeLink(ctx context.Context, in *pb.GetQrCodeLinkRequest) (*pb.GetQrCodeLinkResponse, error) {
	if in.DeviceName == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a name for the device")
	}

	png, err := s.signalClient.GetQrCodeLink(in.DeviceName)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &pb.GetQrCodeLinkResponse{
		Image: png,
	}, nil
}

func (s *SignalService) GetAttachments(ctx context.Context, _ *empty.Empty) (*pb.GetAttachmentsResponse, error) {
	files, err := s.signalClient.GetAttachments()
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &pb.GetAttachmentsResponse{
		Attachments: files,
	}, nil
}

func (s *SignalService) RemoveAttachment(ctx context.Context, in *pb.RemoveAttachmentRequest) (*empty.Empty, error) {
	err := s.signalClient.RemoveAttachment(in.Attachment)

	if err != nil {
		switch err.(type) {
		case *client.InvalidNameError:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case *client.NotFoundError:
			return nil, status.Error(codes.NotFound, err.Error())
		case *client.InternalError:
			return nil, status.Error(codes.Internal, err.Error())
		default:
			return nil, status.Error(codes.Unknown, err.Error())
		}
	}
	return &empty.Empty{}, nil
}

func (s *SignalService) ServeAttachment(ctx context.Context, in *pb.ServeAttachmentRequest) (*pb.ServeAttachmentResponse, error) {
	attachmentBytes, err := s.signalClient.GetAttachment(in.Attachment)

	if err != nil {
		switch err.(type) {
		case *client.InvalidNameError:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case *client.NotFoundError:
			return nil, status.Error(codes.NotFound, err.Error())
		case *client.InternalError:
			return nil, status.Error(codes.Internal, err.Error())
		default:
			return nil, status.Error(codes.Unknown, err.Error())
		}
	}

	return &pb.ServeAttachmentResponse{
		Attachment: attachmentBytes,
	}, nil
}

func (s *SignalService) UpdateProfile(ctx context.Context, in *pb.UpdateProfileRequest) (*empty.Empty, error) {
	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a profile name")
	}

	err := s.signalClient.UpdateProfile(in.Number, in.Name, in.Base64Avatar)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &empty.Empty{}, nil
}

func (s *SignalService) ListIdentities(ctx context.Context, in *pb.ListIdentitiesRequest) (*pb.ListIdentitiesResponse, error) {
	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	identities, err := s.signalClient.ListIdentities(in.Number)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	grpc_identities := []*pb.ListIdentitiesResponse_ListIdentityResponse{}
	for _, identity := range *identities {
		grpc_identities = append(grpc_identities, &pb.ListIdentitiesResponse_ListIdentityResponse{
			Added:        identity.Added,
			Fingerprint:  identity.Fingerprint,
			Number:       identity.Number,
			SafetyNumber: identity.SafetyNumber,
			Status:       identity.Status,
		})
	}

	return &pb.ListIdentitiesResponse{
		Identities: grpc_identities,
	}, nil
}

func (s *SignalService) TrustIdentity(ctx context.Context, in *pb.TrustIdentityRequest) (*empty.Empty, error) {
	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	if in.NumberToTrust == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number to trust")
	}

	if in.VerifiedSafetyNumber == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a verified safety number")
	}

	err := s.signalClient.TrustIdentity(in.Number, in.NumberToTrust, in.VerifiedSafetyNumber)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &empty.Empty{}, nil
}

func (s *SignalService) SetConfiguration(ctx context.Context, in *pb.SetConfigurationRequest) (*empty.Empty, error) {
	if in.Logging.Level != "" {
		if in.Logging.Level == "debug" {
			log.SetLevel(log.DebugLevel)
		} else if in.Logging.Level == "info" {
			log.SetLevel(log.InfoLevel)
		} else if in.Logging.Level == "warn" {
			log.SetLevel(log.WarnLevel)
		} else {
			return nil, status.Error(codes.InvalidArgument, "Couldn't set log level - invalid log level")
		}
	}

	return &empty.Empty{}, nil
}

func (s *SignalService) GetConfiguration(ctx context.Context, _ *empty.Empty) (*pb.GetConfigurationResponse, error) {
	logLevel := ""
	if log.GetLevel() == log.DebugLevel {
		logLevel = "debug"
	} else if log.GetLevel() == log.InfoLevel {
		logLevel = "info"
	} else if log.GetLevel() == log.WarnLevel {
		logLevel = "warn"
	}

	return &pb.GetConfigurationResponse{
		Logging: &pb.Logging{
			Level: logLevel,
		},
	}, nil
}

func (s *SignalService) BlockGroup(ctx context.Context, in *pb.GroupRequest) (*empty.Empty, error) {
	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	if in.Groupid == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a group id")
	}

	err := s.signalClient.BlockGroup(in.Number, in.Groupid)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &empty.Empty{}, nil
}

func (s *SignalService) JoinGroup(ctx context.Context, in *pb.GroupRequest) (*empty.Empty, error) {
	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	if in.Groupid == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a group id")
	}

	err := s.signalClient.JoinGroup(in.Number, in.Groupid)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &empty.Empty{}, nil
}

func (s *SignalService) QuitGroup(ctx context.Context, in *pb.GroupRequest) (*empty.Empty, error) {
	if in.Number == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a number")
	}

	if in.Groupid == "" {
		return nil, status.Error(codes.InvalidArgument, "Please provide a group id")
	}

	err := s.signalClient.QuitGroup(in.Number, in.Groupid)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &empty.Empty{}, nil
}
