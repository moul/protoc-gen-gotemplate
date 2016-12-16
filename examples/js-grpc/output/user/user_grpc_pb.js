// GENERATED CODE -- DO NOT EDIT!


'use strict';
var grpc = require('grpc');
var user_user_pb = require('../user/user_pb.js');



function serialize_user_RegisterRequest(arg) {
  if (!(arg instanceof user_user_pb.RegisterRequest)) {
    throw new Error('Expected argument of type user.RegisterRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_user_RegisterRequest(buffer_arg) {
  return user_user_pb.RegisterRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_RegisterReply(arg) {
  if (!(arg instanceof user_user_pb.RegisterReply)) {
    throw new Error('Expected argument of type user.RegisterReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_user_RegisterReply(buffer_arg) {
  return user_user_pb.RegisterReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_GetUserProfileRequest(arg) {
  if (!(arg instanceof user_user_pb.GetUserProfileRequest)) {
    throw new Error('Expected argument of type user.GetUserProfileRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_user_GetUserProfileRequest(buffer_arg) {
  return user_user_pb.GetUserProfileRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_GetUserProfileReply(arg) {
  if (!(arg instanceof user_user_pb.GetUserProfileReply)) {
    throw new Error('Expected argument of type user.GetUserProfileReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_user_GetUserProfileReply(buffer_arg) {
  return user_user_pb.GetUserProfileReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_UpdateProfileRequest(arg) {
  if (!(arg instanceof user_user_pb.UpdateProfileRequest)) {
    throw new Error('Expected argument of type user.UpdateProfileRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_user_UpdateProfileRequest(buffer_arg) {
  return user_user_pb.UpdateProfileRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_UpdateProfileReply(arg) {
  if (!(arg instanceof user_user_pb.UpdateProfileReply)) {
    throw new Error('Expected argument of type user.UpdateProfileReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_user_UpdateProfileReply(buffer_arg) {
  return user_user_pb.UpdateProfileReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_UploadDocumentRequest(arg) {
  if (!(arg instanceof user_user_pb.UploadDocumentRequest)) {
    throw new Error('Expected argument of type user.UploadDocumentRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_user_UploadDocumentRequest(buffer_arg) {
  return user_user_pb.UploadDocumentRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_user_UploadDocumentReply(arg) {
  if (!(arg instanceof user_user_pb.UploadDocumentReply)) {
    throw new Error('Expected argument of type user.UploadDocumentReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_user_UploadDocumentReply(buffer_arg) {
  return user_user_pb.UploadDocumentReply.deserializeBinary(new Uint8Array(buffer_arg));
}



exports.TestServiceClient = grpc.makeGenericClientConstructor(UserServiceService);


var UserServiceService = exports.UserServiceService = {


  register: {
    path: '/user.UserService/Register',
    requestStream: false,
    responseStream: false,
    requestType: user_user_pb.RegisterRequest,
    responseType: user_user_pb.RegisterReply,
    requestSerialize: serialize_user_RegisterRequest,
    requestDeserialize: deserialize_user_RegisterRequest,
    responseSerialize: serialize_user_RegisterReply,
    responseDeserialize: deserialize_user_RegisterReply,
  },

  getUserProfile: {
    path: '/user.UserService/GetUserProfile',
    requestStream: false,
    responseStream: false,
    requestType: user_user_pb.GetUserProfileRequest,
    responseType: user_user_pb.GetUserProfileReply,
    requestSerialize: serialize_user_GetUserProfileRequest,
    requestDeserialize: deserialize_user_GetUserProfileRequest,
    responseSerialize: serialize_user_GetUserProfileReply,
    responseDeserialize: deserialize_user_GetUserProfileReply,
  },

  updateProfile: {
    path: '/user.UserService/UpdateProfile',
    requestStream: false,
    responseStream: false,
    requestType: user_user_pb.UpdateProfileRequest,
    responseType: user_user_pb.UpdateProfileReply,
    requestSerialize: serialize_user_UpdateProfileRequest,
    requestDeserialize: deserialize_user_UpdateProfileRequest,
    responseSerialize: serialize_user_UpdateProfileReply,
    responseDeserialize: deserialize_user_UpdateProfileReply,
  },

  uploadDocument: {
    path: '/user.UserService/UploadDocument',
    requestStream: false,
    responseStream: false,
    requestType: user_user_pb.UploadDocumentRequest,
    responseType: user_user_pb.UploadDocumentReply,
    requestSerialize: serialize_user_UploadDocumentRequest,
    requestDeserialize: deserialize_user_UploadDocumentRequest,
    responseSerialize: serialize_user_UploadDocumentReply,
    responseDeserialize: deserialize_user_UploadDocumentReply,
  },

}

exports.UserServiceClient = grpc.makeGenericClientConstructor(UserServiceService);

