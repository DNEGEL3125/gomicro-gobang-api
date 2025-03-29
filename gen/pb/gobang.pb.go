// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: gobang.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 五子棋结果
type GobangResult int32

const (
	GobangResult_Success         GobangResult = 0
	GobangResult_InvalidMove     GobangResult = 1
	GobangResult_GameOver        GobangResult = 2
	GobangResult_RoomNotFound    GobangResult = 3
	GobangResult_PlayerNotInRoom GobangResult = 4
	GobangResult_NotYourTurn     GobangResult = 5
	// 服务器错误
	GobangResult_ServerError GobangResult = 6
)

// Enum value maps for GobangResult.
var (
	GobangResult_name = map[int32]string{
		0: "Success",
		1: "InvalidMove",
		2: "GameOver",
		3: "RoomNotFound",
		4: "PlayerNotInRoom",
		5: "NotYourTurn",
		6: "ServerError",
	}
	GobangResult_value = map[string]int32{
		"Success":         0,
		"InvalidMove":     1,
		"GameOver":        2,
		"RoomNotFound":    3,
		"PlayerNotInRoom": 4,
		"NotYourTurn":     5,
		"ServerError":     6,
	}
)

func (x GobangResult) Enum() *GobangResult {
	p := new(GobangResult)
	*p = x
	return p
}

func (x GobangResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GobangResult) Descriptor() protoreflect.EnumDescriptor {
	return file_gobang_proto_enumTypes[0].Descriptor()
}

func (GobangResult) Type() protoreflect.EnumType {
	return &file_gobang_proto_enumTypes[0]
}

func (x GobangResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GobangResult.Descriptor instead.
func (GobangResult) EnumDescriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{0}
}

// 棋子颜色
type GobangColor int32

const (
	GobangColor_White GobangColor = 0
	GobangColor_Black GobangColor = 1
)

// Enum value maps for GobangColor.
var (
	GobangColor_name = map[int32]string{
		0: "White",
		1: "Black",
	}
	GobangColor_value = map[string]int32{
		"White": 0,
		"Black": 1,
	}
)

func (x GobangColor) Enum() *GobangColor {
	p := new(GobangColor)
	*p = x
	return p
}

func (x GobangColor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GobangColor) Descriptor() protoreflect.EnumDescriptor {
	return file_gobang_proto_enumTypes[1].Descriptor()
}

func (GobangColor) Type() protoreflect.EnumType {
	return &file_gobang_proto_enumTypes[1]
}

func (x GobangColor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GobangColor.Descriptor instead.
func (GobangColor) EnumDescriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{1}
}

// 棋盘方格
type GobangCell int32

const (
	GobangCell_EmptyCell GobangCell = 0
	GobangCell_WhiteCell GobangCell = 1
	GobangCell_BlackCell GobangCell = 2
)

// Enum value maps for GobangCell.
var (
	GobangCell_name = map[int32]string{
		0: "EmptyCell",
		1: "WhiteCell",
		2: "BlackCell",
	}
	GobangCell_value = map[string]int32{
		"EmptyCell": 0,
		"WhiteCell": 1,
		"BlackCell": 2,
	}
)

func (x GobangCell) Enum() *GobangCell {
	p := new(GobangCell)
	*p = x
	return p
}

func (x GobangCell) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GobangCell) Descriptor() protoreflect.EnumDescriptor {
	return file_gobang_proto_enumTypes[2].Descriptor()
}

func (GobangCell) Type() protoreflect.EnumType {
	return &file_gobang_proto_enumTypes[2]
}

func (x GobangCell) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GobangCell.Descriptor instead.
func (GobangCell) EnumDescriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{2}
}

// 游戏状态
type GameStatus int32

const (
	GameStatus_Playing  GameStatus = 0
	GameStatus_WhiteWin GameStatus = 1
	GameStatus_BlackWin GameStatus = 2
	GameStatus_Draw     GameStatus = 3
)

// Enum value maps for GameStatus.
var (
	GameStatus_name = map[int32]string{
		0: "Playing",
		1: "WhiteWin",
		2: "BlackWin",
		3: "Draw",
	}
	GameStatus_value = map[string]int32{
		"Playing":  0,
		"WhiteWin": 1,
		"BlackWin": 2,
		"Draw":     3,
	}
)

func (x GameStatus) Enum() *GameStatus {
	p := new(GameStatus)
	*p = x
	return p
}

func (x GameStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_gobang_proto_enumTypes[3].Descriptor()
}

func (GameStatus) Type() protoreflect.EnumType {
	return &file_gobang_proto_enumTypes[3]
}

func (x GameStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameStatus.Descriptor instead.
func (GameStatus) EnumDescriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{3}
}

// 获取游戏房间请求
type GetGameRoomByEmailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGameRoomByEmailRequest) Reset() {
	*x = GetGameRoomByEmailRequest{}
	mi := &file_gobang_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGameRoomByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameRoomByEmailRequest) ProtoMessage() {}

func (x *GetGameRoomByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gobang_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameRoomByEmailRequest.ProtoReflect.Descriptor instead.
func (*GetGameRoomByEmailRequest) Descriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{0}
}

func (x *GetGameRoomByEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// 获取游戏房间响应
type GetGameRoomByEmailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GameRoom      *GameRoom              `protobuf:"bytes,1,opt,name=game_room,json=gameRoom,proto3" json:"game_room,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetGameRoomByEmailResponse) Reset() {
	*x = GetGameRoomByEmailResponse{}
	mi := &file_gobang_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetGameRoomByEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGameRoomByEmailResponse) ProtoMessage() {}

func (x *GetGameRoomByEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gobang_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGameRoomByEmailResponse.ProtoReflect.Descriptor instead.
func (*GetGameRoomByEmailResponse) Descriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{1}
}

func (x *GetGameRoomByEmailResponse) GetGameRoom() *GameRoom {
	if x != nil {
		return x.GameRoom
	}
	return nil
}

func (x *GetGameRoomByEmailResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// 创建房间请求
type CreateRoomRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Player1Email  string                 `protobuf:"bytes,1,opt,name=player1_email,json=player1Email,proto3" json:"player1_email,omitempty"`
	Player2Email  string                 `protobuf:"bytes,2,opt,name=player2_email,json=player2Email,proto3" json:"player2_email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	mi := &file_gobang_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gobang_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRoomRequest) GetPlayer1Email() string {
	if x != nil {
		return x.Player1Email
	}
	return ""
}

func (x *CreateRoomRequest) GetPlayer2Email() string {
	if x != nil {
		return x.Player2Email
	}
	return ""
}

// 创建房间响应
type CreateRoomResponse struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	RoomId string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Error  string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// 白棋玩家邮箱
	WhitePlayerEmail string `protobuf:"bytes,3,opt,name=white_player_email,json=whitePlayerEmail,proto3" json:"white_player_email,omitempty"`
	// 黑棋玩家邮箱
	BlackPlayerEmail string `protobuf:"bytes,4,opt,name=black_player_email,json=blackPlayerEmail,proto3" json:"black_player_email,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CreateRoomResponse) Reset() {
	*x = CreateRoomResponse{}
	mi := &file_gobang_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomResponse) ProtoMessage() {}

func (x *CreateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gobang_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomResponse.ProtoReflect.Descriptor instead.
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRoomResponse) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *CreateRoomResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateRoomResponse) GetWhitePlayerEmail() string {
	if x != nil {
		return x.WhitePlayerEmail
	}
	return ""
}

func (x *CreateRoomResponse) GetBlackPlayerEmail() string {
	if x != nil {
		return x.BlackPlayerEmail
	}
	return ""
}

// 落子请求
type MoveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	PlayerEmail   string                 `protobuf:"bytes,2,opt,name=player_email,json=playerEmail,proto3" json:"player_email,omitempty"`
	X             int32                  `protobuf:"varint,3,opt,name=x,proto3" json:"x,omitempty"`
	Y             int32                  `protobuf:"varint,4,opt,name=y,proto3" json:"y,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MoveRequest) Reset() {
	*x = MoveRequest{}
	mi := &file_gobang_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveRequest) ProtoMessage() {}

func (x *MoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gobang_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveRequest.ProtoReflect.Descriptor instead.
func (*MoveRequest) Descriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{4}
}

func (x *MoveRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *MoveRequest) GetPlayerEmail() string {
	if x != nil {
		return x.PlayerEmail
	}
	return ""
}

func (x *MoveRequest) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MoveRequest) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

// 落子响应
type MoveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        GobangResult           `protobuf:"varint,1,opt,name=result,proto3,enum=gobang.GobangResult" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MoveResponse) Reset() {
	*x = MoveResponse{}
	mi := &file_gobang_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveResponse) ProtoMessage() {}

func (x *MoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gobang_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveResponse.ProtoReflect.Descriptor instead.
func (*MoveResponse) Descriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{5}
}

func (x *MoveResponse) GetResult() GobangResult {
	if x != nil {
		return x.Result
	}
	return GobangResult_Success
}

// 落子广播消息
type MoveBroadcastMessage struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	RoomId string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	X      int32                  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y      int32                  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	// 进行落子的玩家
	From *Player `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	// 对手
	To            *Player `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MoveBroadcastMessage) Reset() {
	*x = MoveBroadcastMessage{}
	mi := &file_gobang_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MoveBroadcastMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveBroadcastMessage) ProtoMessage() {}

func (x *MoveBroadcastMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gobang_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveBroadcastMessage.ProtoReflect.Descriptor instead.
func (*MoveBroadcastMessage) Descriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{6}
}

func (x *MoveBroadcastMessage) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *MoveBroadcastMessage) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MoveBroadcastMessage) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *MoveBroadcastMessage) GetFrom() *Player {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *MoveBroadcastMessage) GetTo() *Player {
	if x != nil {
		return x.To
	}
	return nil
}

// 玩家
type Player struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Color         GobangColor            `protobuf:"varint,2,opt,name=color,proto3,enum=gobang.GobangColor" json:"color,omitempty"`
	RoomId        string                 `protobuf:"bytes,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Player) Reset() {
	*x = Player{}
	mi := &file_gobang_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_gobang_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{7}
}

func (x *Player) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Player) GetColor() GobangColor {
	if x != nil {
		return x.Color
	}
	return GobangColor_White
}

func (x *Player) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

// 五子棋游戏
type Gobang struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 棋盘大小
	BoardSize int32 `protobuf:"varint,1,opt,name=board_size,json=boardSize,proto3" json:"board_size,omitempty"`
	// 压缩的棋盘
	CompressedBoard []byte `protobuf:"bytes,2,opt,name=compressed_board,json=compressedBoard,proto3" json:"compressed_board,omitempty"`
	// 游戏状态
	Status GameStatus `protobuf:"varint,3,opt,name=status,proto3,enum=gobang.GameStatus" json:"status,omitempty"`
	// 当前落子颜色
	CurrentColor  GobangColor `protobuf:"varint,4,opt,name=current_color,json=currentColor,proto3,enum=gobang.GobangColor" json:"current_color,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Gobang) Reset() {
	*x = Gobang{}
	mi := &file_gobang_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Gobang) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gobang) ProtoMessage() {}

func (x *Gobang) ProtoReflect() protoreflect.Message {
	mi := &file_gobang_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gobang.ProtoReflect.Descriptor instead.
func (*Gobang) Descriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{8}
}

func (x *Gobang) GetBoardSize() int32 {
	if x != nil {
		return x.BoardSize
	}
	return 0
}

func (x *Gobang) GetCompressedBoard() []byte {
	if x != nil {
		return x.CompressedBoard
	}
	return nil
}

func (x *Gobang) GetStatus() GameStatus {
	if x != nil {
		return x.Status
	}
	return GameStatus_Playing
}

func (x *Gobang) GetCurrentColor() GobangColor {
	if x != nil {
		return x.CurrentColor
	}
	return GobangColor_White
}

// 游戏房间
type GameRoom struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	WhitePlayer   *Player                `protobuf:"bytes,2,opt,name=white_player,json=whitePlayer,proto3" json:"white_player,omitempty"`
	BlackPlayer   *Player                `protobuf:"bytes,3,opt,name=black_player,json=blackPlayer,proto3" json:"black_player,omitempty"`
	Gobang        *Gobang                `protobuf:"bytes,4,opt,name=gobang,proto3" json:"gobang,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GameRoom) Reset() {
	*x = GameRoom{}
	mi := &file_gobang_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameRoom) ProtoMessage() {}

func (x *GameRoom) ProtoReflect() protoreflect.Message {
	mi := &file_gobang_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameRoom.ProtoReflect.Descriptor instead.
func (*GameRoom) Descriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{9}
}

func (x *GameRoom) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *GameRoom) GetWhitePlayer() *Player {
	if x != nil {
		return x.WhitePlayer
	}
	return nil
}

func (x *GameRoom) GetBlackPlayer() *Player {
	if x != nil {
		return x.BlackPlayer
	}
	return nil
}

func (x *GameRoom) GetGobang() *Gobang {
	if x != nil {
		return x.Gobang
	}
	return nil
}

// 游戏结束消息
type GameEndMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	GameStatus    GameStatus             `protobuf:"varint,2,opt,name=game_status,json=gameStatus,proto3,enum=gobang.GameStatus" json:"game_status,omitempty"`
	Winner        *Player                `protobuf:"bytes,3,opt,name=winner,proto3" json:"winner,omitempty"`
	Loser         *Player                `protobuf:"bytes,4,opt,name=loser,proto3" json:"loser,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GameEndMessage) Reset() {
	*x = GameEndMessage{}
	mi := &file_gobang_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameEndMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameEndMessage) ProtoMessage() {}

func (x *GameEndMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gobang_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameEndMessage.ProtoReflect.Descriptor instead.
func (*GameEndMessage) Descriptor() ([]byte, []int) {
	return file_gobang_proto_rawDescGZIP(), []int{10}
}

func (x *GameEndMessage) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *GameEndMessage) GetGameStatus() GameStatus {
	if x != nil {
		return x.GameStatus
	}
	return GameStatus_Playing
}

func (x *GameEndMessage) GetWinner() *Player {
	if x != nil {
		return x.Winner
	}
	return nil
}

func (x *GameEndMessage) GetLoser() *Player {
	if x != nil {
		return x.Loser
	}
	return nil
}

var File_gobang_proto protoreflect.FileDescriptor

const file_gobang_proto_rawDesc = "" +
	"\n" +
	"\fgobang.proto\x12\x06gobang\"1\n" +
	"\x19GetGameRoomByEmailRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\"a\n" +
	"\x1aGetGameRoomByEmailResponse\x12-\n" +
	"\tgame_room\x18\x01 \x01(\v2\x10.gobang.GameRoomR\bgameRoom\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error\"]\n" +
	"\x11CreateRoomRequest\x12#\n" +
	"\rplayer1_email\x18\x01 \x01(\tR\fplayer1Email\x12#\n" +
	"\rplayer2_email\x18\x02 \x01(\tR\fplayer2Email\"\x9f\x01\n" +
	"\x12CreateRoomResponse\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error\x12,\n" +
	"\x12white_player_email\x18\x03 \x01(\tR\x10whitePlayerEmail\x12,\n" +
	"\x12black_player_email\x18\x04 \x01(\tR\x10blackPlayerEmail\"e\n" +
	"\vMoveRequest\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\x12!\n" +
	"\fplayer_email\x18\x02 \x01(\tR\vplayerEmail\x12\f\n" +
	"\x01x\x18\x03 \x01(\x05R\x01x\x12\f\n" +
	"\x01y\x18\x04 \x01(\x05R\x01y\"<\n" +
	"\fMoveResponse\x12,\n" +
	"\x06result\x18\x01 \x01(\x0e2\x14.gobang.GobangResultR\x06result\"\x8f\x01\n" +
	"\x14MoveBroadcastMessage\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\x12\f\n" +
	"\x01x\x18\x02 \x01(\x05R\x01x\x12\f\n" +
	"\x01y\x18\x03 \x01(\x05R\x01y\x12\"\n" +
	"\x04from\x18\x04 \x01(\v2\x0e.gobang.PlayerR\x04from\x12\x1e\n" +
	"\x02to\x18\x05 \x01(\v2\x0e.gobang.PlayerR\x02to\"b\n" +
	"\x06Player\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12)\n" +
	"\x05color\x18\x02 \x01(\x0e2\x13.gobang.GobangColorR\x05color\x12\x17\n" +
	"\aroom_id\x18\x03 \x01(\tR\x06roomId\"\xb8\x01\n" +
	"\x06Gobang\x12\x1d\n" +
	"\n" +
	"board_size\x18\x01 \x01(\x05R\tboardSize\x12)\n" +
	"\x10compressed_board\x18\x02 \x01(\fR\x0fcompressedBoard\x12*\n" +
	"\x06status\x18\x03 \x01(\x0e2\x12.gobang.GameStatusR\x06status\x128\n" +
	"\rcurrent_color\x18\x04 \x01(\x0e2\x13.gobang.GobangColorR\fcurrentColor\"\xb1\x01\n" +
	"\bGameRoom\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\x121\n" +
	"\fwhite_player\x18\x02 \x01(\v2\x0e.gobang.PlayerR\vwhitePlayer\x121\n" +
	"\fblack_player\x18\x03 \x01(\v2\x0e.gobang.PlayerR\vblackPlayer\x12&\n" +
	"\x06gobang\x18\x04 \x01(\v2\x0e.gobang.GobangR\x06gobang\"\xac\x01\n" +
	"\x0eGameEndMessage\x12\x17\n" +
	"\aroom_id\x18\x01 \x01(\tR\x06roomId\x123\n" +
	"\vgame_status\x18\x02 \x01(\x0e2\x12.gobang.GameStatusR\n" +
	"gameStatus\x12&\n" +
	"\x06winner\x18\x03 \x01(\v2\x0e.gobang.PlayerR\x06winner\x12$\n" +
	"\x05loser\x18\x04 \x01(\v2\x0e.gobang.PlayerR\x05loser*\x83\x01\n" +
	"\fGobangResult\x12\v\n" +
	"\aSuccess\x10\x00\x12\x0f\n" +
	"\vInvalidMove\x10\x01\x12\f\n" +
	"\bGameOver\x10\x02\x12\x10\n" +
	"\fRoomNotFound\x10\x03\x12\x13\n" +
	"\x0fPlayerNotInRoom\x10\x04\x12\x0f\n" +
	"\vNotYourTurn\x10\x05\x12\x0f\n" +
	"\vServerError\x10\x06*#\n" +
	"\vGobangColor\x12\t\n" +
	"\x05White\x10\x00\x12\t\n" +
	"\x05Black\x10\x01*9\n" +
	"\n" +
	"GobangCell\x12\r\n" +
	"\tEmptyCell\x10\x00\x12\r\n" +
	"\tWhiteCell\x10\x01\x12\r\n" +
	"\tBlackCell\x10\x02*?\n" +
	"\n" +
	"GameStatus\x12\v\n" +
	"\aPlaying\x10\x00\x12\f\n" +
	"\bWhiteWin\x10\x01\x12\f\n" +
	"\bBlackWin\x10\x02\x12\b\n" +
	"\x04Draw\x10\x032\xe6\x01\n" +
	"\vGameService\x12C\n" +
	"\n" +
	"CreateRoom\x12\x19.gobang.CreateRoomRequest\x1a\x1a.gobang.CreateRoomResponse\x125\n" +
	"\bMakeMove\x12\x13.gobang.MoveRequest\x1a\x14.gobang.MoveResponse\x12[\n" +
	"\x12GetGameRoomByEmail\x12!.gobang.GetGameRoomByEmailRequest\x1a\".gobang.GetGameRoomByEmailResponseB\x82\x01\n" +
	"\n" +
	"com.gobangB\vGobangProtoP\x01Z/github.com/DNEGEL3125/gomicro-gobang-api/gen/pb\xa2\x02\x03GXX\xaa\x02\x06Gobang\xca\x02\x06Gobang\xe2\x02\x12Gobang\\GPBMetadata\xea\x02\x06Gobangb\x06proto3"

var (
	file_gobang_proto_rawDescOnce sync.Once
	file_gobang_proto_rawDescData []byte
)

func file_gobang_proto_rawDescGZIP() []byte {
	file_gobang_proto_rawDescOnce.Do(func() {
		file_gobang_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_gobang_proto_rawDesc), len(file_gobang_proto_rawDesc)))
	})
	return file_gobang_proto_rawDescData
}

var file_gobang_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_gobang_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_gobang_proto_goTypes = []any{
	(GobangResult)(0),                  // 0: gobang.GobangResult
	(GobangColor)(0),                   // 1: gobang.GobangColor
	(GobangCell)(0),                    // 2: gobang.GobangCell
	(GameStatus)(0),                    // 3: gobang.GameStatus
	(*GetGameRoomByEmailRequest)(nil),  // 4: gobang.GetGameRoomByEmailRequest
	(*GetGameRoomByEmailResponse)(nil), // 5: gobang.GetGameRoomByEmailResponse
	(*CreateRoomRequest)(nil),          // 6: gobang.CreateRoomRequest
	(*CreateRoomResponse)(nil),         // 7: gobang.CreateRoomResponse
	(*MoveRequest)(nil),                // 8: gobang.MoveRequest
	(*MoveResponse)(nil),               // 9: gobang.MoveResponse
	(*MoveBroadcastMessage)(nil),       // 10: gobang.MoveBroadcastMessage
	(*Player)(nil),                     // 11: gobang.Player
	(*Gobang)(nil),                     // 12: gobang.Gobang
	(*GameRoom)(nil),                   // 13: gobang.GameRoom
	(*GameEndMessage)(nil),             // 14: gobang.GameEndMessage
}
var file_gobang_proto_depIdxs = []int32{
	13, // 0: gobang.GetGameRoomByEmailResponse.game_room:type_name -> gobang.GameRoom
	0,  // 1: gobang.MoveResponse.result:type_name -> gobang.GobangResult
	11, // 2: gobang.MoveBroadcastMessage.from:type_name -> gobang.Player
	11, // 3: gobang.MoveBroadcastMessage.to:type_name -> gobang.Player
	1,  // 4: gobang.Player.color:type_name -> gobang.GobangColor
	3,  // 5: gobang.Gobang.status:type_name -> gobang.GameStatus
	1,  // 6: gobang.Gobang.current_color:type_name -> gobang.GobangColor
	11, // 7: gobang.GameRoom.white_player:type_name -> gobang.Player
	11, // 8: gobang.GameRoom.black_player:type_name -> gobang.Player
	12, // 9: gobang.GameRoom.gobang:type_name -> gobang.Gobang
	3,  // 10: gobang.GameEndMessage.game_status:type_name -> gobang.GameStatus
	11, // 11: gobang.GameEndMessage.winner:type_name -> gobang.Player
	11, // 12: gobang.GameEndMessage.loser:type_name -> gobang.Player
	6,  // 13: gobang.GameService.CreateRoom:input_type -> gobang.CreateRoomRequest
	8,  // 14: gobang.GameService.MakeMove:input_type -> gobang.MoveRequest
	4,  // 15: gobang.GameService.GetGameRoomByEmail:input_type -> gobang.GetGameRoomByEmailRequest
	7,  // 16: gobang.GameService.CreateRoom:output_type -> gobang.CreateRoomResponse
	9,  // 17: gobang.GameService.MakeMove:output_type -> gobang.MoveResponse
	5,  // 18: gobang.GameService.GetGameRoomByEmail:output_type -> gobang.GetGameRoomByEmailResponse
	16, // [16:19] is the sub-list for method output_type
	13, // [13:16] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_gobang_proto_init() }
func file_gobang_proto_init() {
	if File_gobang_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_gobang_proto_rawDesc), len(file_gobang_proto_rawDesc)),
			NumEnums:      4,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gobang_proto_goTypes,
		DependencyIndexes: file_gobang_proto_depIdxs,
		EnumInfos:         file_gobang_proto_enumTypes,
		MessageInfos:      file_gobang_proto_msgTypes,
	}.Build()
	File_gobang_proto = out.File
	file_gobang_proto_goTypes = nil
	file_gobang_proto_depIdxs = nil
}
