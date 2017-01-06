// @flow
// GENERATED CODE -- DO NOT EDIT!

import grpc from 'grpc'
import pbFile from './pbFile.js'

export type TestEnum = 
  | 'ELEMENT_A'
  | 'ELEMENT_B'
;

export type TestMessage = {|
  a?: string;
  b?: number;
  c?: number;
  d?: number;
  e?: number;
  n?: Array<string>;
  o?: Array<number>;
  p?: Array<number>;
  q?: Array<number>;
  r?: Array<number>;
  s?:
    | 'ELEMENT_C'
    | 'ELEMENT_D'
  ;
|};

export type TestNoStreamRequest = {|
  message?: TestMessage;
|};

export type TestNoStreamReply = {|
  message?: TestMessage;
  err_msg?: string;
|};

export type TestStreamRequestRequest = {|
  message?: TestMessage;
|};

export type TestStreamRequestReply = {|
  message?: TestMessage;
  err_msg?: string;
|};

export type TestStreamReplyRequest = {|
  message?: TestMessage;
|};

export type TestStreamReplyReply = {|
  message?: TestMessage;
  err_msg?: string;
|};

export type TestStreamBothRequest = {|
  message?: TestMessage;
|};

export type TestStreamBothReply = {|
  message?: TestMessage;
  err_msg?: string;
|};


function serialize_test_TestNoStreamRequest(arg: TestNoStreamRequest) {
  if (!(arg instanceof pbFile.TestNoStreamRequest)) {
    throw new Error('Expected argument of type TestNoStreamRequest')
  }
  return new Buffer(arg.serializeBinary())
}

function deserialize_test_TestNoStreamRequest(buffer_arg: Array<number>) {
  return pbFile.TestNoStreamRequest.deserializeBinary(new Uint8Array(buffer_arg))
}

function serialize_test_TestNoStreamReply(arg: TestNoStreamReply) {
  if (!(arg instanceof pbFile.TestNoStreamReply)) {
    throw new Error('Expected argument of type TestNoStreamReply')
  }
  return new Buffer(arg.serializeBinary())
}

function deserialize_test_TestNoStreamReply(buffer_arg: Array<number>) {
  return pbFile.TestNoStreamReply.deserializeBinary(new Uint8Array(buffer_arg))
}

function serialize_test_TestStreamRequestRequest(arg: TestStreamRequestRequest) {
  if (!(arg instanceof pbFile.TestStreamRequestRequest)) {
    throw new Error('Expected argument of type TestStreamRequestRequest')
  }
  return new Buffer(arg.serializeBinary())
}

function deserialize_test_TestStreamRequestRequest(buffer_arg: Array<number>) {
  return pbFile.TestStreamRequestRequest.deserializeBinary(new Uint8Array(buffer_arg))
}

function serialize_test_TestStreamRequestReply(arg: TestStreamRequestReply) {
  if (!(arg instanceof pbFile.TestStreamRequestReply)) {
    throw new Error('Expected argument of type TestStreamRequestReply')
  }
  return new Buffer(arg.serializeBinary())
}

function deserialize_test_TestStreamRequestReply(buffer_arg: Array<number>) {
  return pbFile.TestStreamRequestReply.deserializeBinary(new Uint8Array(buffer_arg))
}

function serialize_test_TestStreamReplyRequest(arg: TestStreamReplyRequest) {
  if (!(arg instanceof pbFile.TestStreamReplyRequest)) {
    throw new Error('Expected argument of type TestStreamReplyRequest')
  }
  return new Buffer(arg.serializeBinary())
}

function deserialize_test_TestStreamReplyRequest(buffer_arg: Array<number>) {
  return pbFile.TestStreamReplyRequest.deserializeBinary(new Uint8Array(buffer_arg))
}

function serialize_test_TestStreamReplyReply(arg: TestStreamReplyReply) {
  if (!(arg instanceof pbFile.TestStreamReplyReply)) {
    throw new Error('Expected argument of type TestStreamReplyReply')
  }
  return new Buffer(arg.serializeBinary())
}

function deserialize_test_TestStreamReplyReply(buffer_arg: Array<number>) {
  return pbFile.TestStreamReplyReply.deserializeBinary(new Uint8Array(buffer_arg))
}

function serialize_test_TestStreamBothRequest(arg: TestStreamBothRequest) {
  if (!(arg instanceof pbFile.TestStreamBothRequest)) {
    throw new Error('Expected argument of type TestStreamBothRequest')
  }
  return new Buffer(arg.serializeBinary())
}

function deserialize_test_TestStreamBothRequest(buffer_arg: Array<number>) {
  return pbFile.TestStreamBothRequest.deserializeBinary(new Uint8Array(buffer_arg))
}

function serialize_test_TestStreamBothReply(arg: TestStreamBothReply) {
  if (!(arg instanceof pbFile.TestStreamBothReply)) {
    throw new Error('Expected argument of type TestStreamBothReply')
  }
  return new Buffer(arg.serializeBinary())
}

function deserialize_test_TestStreamBothReply(buffer_arg: Array<number>) {
  return pbFile.TestStreamBothReply.deserializeBinary(new Uint8Array(buffer_arg))
}


export const TestServiceService = {
  
  testNoStream: {
    path: '/test.TestService/TestNoStream',
    requestStream: false,
    responseStream: false,
    requestType: pbFile.TestNoStreamRequest,
    responseType: pbFile.TestNoStreamReply,
    requestSerialize: serialize_test_TestNoStreamRequest,
    requestDeserialize: deserialize_test_TestNoStreamRequest,
    responseSerialize: serialize_test_TestNoStreamReply,
    responseDeserialize: deserialize_test_TestNoStreamReply,
  },
  testStreamRequest: {
    path: '/test.TestService/TestStreamRequest',
    requestStream: true,
    responseStream: false,
    requestType: pbFile.TestStreamRequestRequest,
    responseType: pbFile.TestStreamRequestReply,
    requestSerialize: serialize_test_TestStreamRequestRequest,
    requestDeserialize: deserialize_test_TestStreamRequestRequest,
    responseSerialize: serialize_test_TestStreamRequestReply,
    responseDeserialize: deserialize_test_TestStreamRequestReply,
  },
  testStreamReply: {
    path: '/test.TestService/TestStreamReply',
    requestStream: false,
    responseStream: true,
    requestType: pbFile.TestStreamReplyRequest,
    responseType: pbFile.TestStreamReplyReply,
    requestSerialize: serialize_test_TestStreamReplyRequest,
    requestDeserialize: deserialize_test_TestStreamReplyRequest,
    responseSerialize: serialize_test_TestStreamReplyReply,
    responseDeserialize: deserialize_test_TestStreamReplyReply,
  },
  testStreamBoth: {
    path: '/test.TestService/TestStreamBoth',
    requestStream: true,
    responseStream: true,
    requestType: pbFile.TestStreamBothRequest,
    responseType: pbFile.TestStreamBothReply,
    requestSerialize: serialize_test_TestStreamBothRequest,
    requestDeserialize: deserialize_test_TestStreamBothRequest,
    responseSerialize: serialize_test_TestStreamBothReply,
    responseDeserialize: deserialize_test_TestStreamBothReply,
  },
  
}

export const TestServiceClient = grpc.makeGenericClientConstructor(TestServiceService)
