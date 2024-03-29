# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user/v1/user.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12user/v1/user.proto\x12\x07user.v1\"o\n\x0bUserMessage\x12\x0e\n\x02id\x18\x01 \x01(\x03R\x02id\x12\x1d\n\nfirst_name\x18\x02 \x01(\tR\tfirstName\x12\x1b\n\tlast_name\x18\x03 \x01(\tR\x08lastName\x12\x14\n\x05\x65mail\x18\x04 \x01(\tR\x05\x65mail\"S\n\x10\x46\x65tchUserRequest\x12\x13\n\x02id\x18\x01 \x01(\x03H\x00R\x02id\x88\x01\x01\x12\x19\n\x05\x65mail\x18\x02 \x01(\tH\x01R\x05\x65mail\x88\x01\x01\x42\x05\n\x03_idB\x08\n\x06_email\"=\n\x11\x46\x65tchUserResponse\x12(\n\x04user\x18\x01 \x01(\x0b\x32\x14.user.v1.UserMessageR\x04user2S\n\x0bUserService\x12\x44\n\tFetchUser\x12\x19.user.v1.FetchUserRequest\x1a\x1a.user.v1.FetchUserResponse\"\x00\x42\x88\x01\n\x0b\x63om.user.v1B\tUserProtoP\x01Z1github.com/mikedonnici/mono/go/gen/user/v1;userv1\xa2\x02\x03UXX\xaa\x02\x07User.V1\xca\x02\x07User\\V1\xe2\x02\x13User\\V1\\GPBMetadata\xea\x02\x08User::V1b\x06proto3')



_USERMESSAGE = DESCRIPTOR.message_types_by_name['UserMessage']
_FETCHUSERREQUEST = DESCRIPTOR.message_types_by_name['FetchUserRequest']
_FETCHUSERRESPONSE = DESCRIPTOR.message_types_by_name['FetchUserResponse']
UserMessage = _reflection.GeneratedProtocolMessageType('UserMessage', (_message.Message,), {
  'DESCRIPTOR' : _USERMESSAGE,
  '__module__' : 'user.v1.user_pb2'
  # @@protoc_insertion_point(class_scope:user.v1.UserMessage)
  })
_sym_db.RegisterMessage(UserMessage)

FetchUserRequest = _reflection.GeneratedProtocolMessageType('FetchUserRequest', (_message.Message,), {
  'DESCRIPTOR' : _FETCHUSERREQUEST,
  '__module__' : 'user.v1.user_pb2'
  # @@protoc_insertion_point(class_scope:user.v1.FetchUserRequest)
  })
_sym_db.RegisterMessage(FetchUserRequest)

FetchUserResponse = _reflection.GeneratedProtocolMessageType('FetchUserResponse', (_message.Message,), {
  'DESCRIPTOR' : _FETCHUSERRESPONSE,
  '__module__' : 'user.v1.user_pb2'
  # @@protoc_insertion_point(class_scope:user.v1.FetchUserResponse)
  })
_sym_db.RegisterMessage(FetchUserResponse)

_USERSERVICE = DESCRIPTOR.services_by_name['UserService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\013com.user.v1B\tUserProtoP\001Z1github.com/mikedonnici/mono/go/gen/user/v1;userv1\242\002\003UXX\252\002\007User.V1\312\002\007User\\V1\342\002\023User\\V1\\GPBMetadata\352\002\010User::V1'
  _USERMESSAGE._serialized_start=31
  _USERMESSAGE._serialized_end=142
  _FETCHUSERREQUEST._serialized_start=144
  _FETCHUSERREQUEST._serialized_end=227
  _FETCHUSERRESPONSE._serialized_start=229
  _FETCHUSERRESPONSE._serialized_end=290
  _USERSERVICE._serialized_start=292
  _USERSERVICE._serialized_end=375
# @@protoc_insertion_point(module_scope)
