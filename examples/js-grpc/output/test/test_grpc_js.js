// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var test_pb = require('./test_pb.js');


function serialize_test_TestNoStreamRequest(arg) {
  if (!(arg instanceof test_pb.TestNoStreamRequest)) {
    throw new Error('Expected argument of type test.TestNoStreamRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_test_TestNoStreamRequest(buffer_arg) {
  return test_pb.TestNoStreamRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_test_TestNoStreamReply(arg) {
  if (!(arg instanceof test_pb.TestNoStreamReply)) {
    throw new Error('Expected argument of type test.TestNoStreamReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_test_TestNoStreamReply(buffer_arg) {
  return test_pb.TestNoStreamReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_test_TestStreamRequestRequest(arg) {
  if (!(arg instanceof test_pb.TestStreamRequestRequest)) {
    throw new Error('Expected argument of type test.TestStreamRequestRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_test_TestStreamRequestRequest(buffer_arg) {
  return test_pb.TestStreamRequestRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_test_TestStreamRequestReply(arg) {
  if (!(arg instanceof test_pb.TestStreamRequestReply)) {
    throw new Error('Expected argument of type test.TestStreamRequestReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_test_TestStreamRequestReply(buffer_arg) {
  return test_pb.TestStreamRequestReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_test_TestStreamReplyRequest(arg) {
  if (!(arg instanceof test_pb.TestStreamReplyRequest)) {
    throw new Error('Expected argument of type test.TestStreamReplyRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_test_TestStreamReplyRequest(buffer_arg) {
  return test_pb.TestStreamReplyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_test_TestStreamReplyReply(arg) {
  if (!(arg instanceof test_pb.TestStreamReplyReply)) {
    throw new Error('Expected argument of type test.TestStreamReplyReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_test_TestStreamReplyReply(buffer_arg) {
  return test_pb.TestStreamReplyReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_test_TestStreamBothRequest(arg) {
  if (!(arg instanceof test_pb.TestStreamBothRequest)) {
    throw new Error('Expected argument of type test.TestStreamBothRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_test_TestStreamBothRequest(buffer_arg) {
  return test_pb.TestStreamBothRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_test_TestStreamBothReply(arg) {
  if (!(arg instanceof test_pb.TestStreamBothReply)) {
    throw new Error('Expected argument of type test.TestStreamBothReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_test_TestStreamBothReply(buffer_arg) {
  return test_pb.TestStreamBothReply.deserializeBinary(new Uint8Array(buffer_arg));
}


var TestServiceService = exports.TestServiceService = {
  
  testNoStream: {
    path: '/test.TestService/TestNoStream',
    requestStream: false,
    responseStream: false,
    requestType: test_pb.TestNoStreamRequest,
    responseType: test_pb.TestNoStreamReply,
    requestSerialize: serialize_test_TestNoStreamRequest,
    requestDeserialize: deserialize_test_TestNoStreamRequest,
    responseSerialize: serialize_test_TestNoStreamReply,
    responseDeserialize: deserialize_test_TestNoStreamReply,
  },
  testStreamRequest: {
    path: '/test.TestService/TestStreamRequest',
    requestStream: true,
    responseStream: false,
    requestType: test_pb.TestStreamRequestRequest,
    responseType: test_pb.TestStreamRequestReply,
    requestSerialize: serialize_test_TestStreamRequestRequest,
    requestDeserialize: deserialize_test_TestStreamRequestRequest,
    responseSerialize: serialize_test_TestStreamRequestReply,
    responseDeserialize: deserialize_test_TestStreamRequestReply,
  },
  testStreamReply: {
    path: '/test.TestService/TestStreamReply',
    requestStream: false,
    responseStream: true,
    requestType: test_pb.TestStreamReplyRequest,
    responseType: test_pb.TestStreamReplyReply,
    requestSerialize: serialize_test_TestStreamReplyRequest,
    requestDeserialize: deserialize_test_TestStreamReplyRequest,
    responseSerialize: serialize_test_TestStreamReplyReply,
    responseDeserialize: deserialize_test_TestStreamReplyReply,
  },
  testStreamBoth: {
    path: '/test.TestService/TestStreamBoth',
    requestStream: true,
    responseStream: true,
    requestType: test_pb.TestStreamBothRequest,
    responseType: test_pb.TestStreamBothReply,
    requestSerialize: serialize_test_TestStreamBothRequest,
    requestDeserialize: deserialize_test_TestStreamBothRequest,
    responseSerialize: serialize_test_TestStreamBothReply,
    responseDeserialize: deserialize_test_TestStreamBothReply,
  },
  
}

exports.TestServiceClient = grpc.makeGenericClientConstructor(TestServiceService);
