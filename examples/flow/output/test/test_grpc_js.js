// @flow
// GENERATED CODE -- DO NOT EDIT!

import base64 from 'base64-js'
import test_pb from './test_pb'


export type TestEnum = {|
    ELEMENT_A?: 0;
    ELEMENT_B?: 1;
|};


export type TestMessage$TestNestedEnum = {|
    ELEMENT_C?: 0;
    ELEMENT_D?: 1;
|};


export type TestMessage$TestNestedMessage = {
    getS?: () => string;
    setS?: (s: string) => void;
    getT?: () => number;
    setT?: (t: number) => void;
};

export type TestMessage = {
    getA?: () => string;
    setA?: (a: string) => void;
    getB?: () => number;
    setB?: (b: number) => void;
    getC?: () => number;
    setC?: (c: number) => void;
    getD?: () => number;
    setD?: (d: number) => void;
    getE?: () => number;
    setE?: (e: number) => void;
    getNList?: () => Array<string>;
    setNList?: (n: Array<string>) => void;
    addN?: (n: string) => void;
    clearNList?: () => void;
    getOList?: () => Array<number>;
    setOList?: (o: Array<number>) => void;
    addO?: (o: number) => void;
    clearOList?: () => void;
    getPList?: () => Array<number>;
    setPList?: (p: Array<number>) => void;
    addP?: (p: number) => void;
    clearPList?: () => void;
    getQList?: () => Array<number>;
    setQList?: (q: Array<number>) => void;
    addQ?: (q: number) => void;
    clearQList?: () => void;
    getRList?: () => Array<number>;
    setRList?: (r: Array<number>) => void;
    addR?: (r: number) => void;
    clearRList?: () => void;
    getU?: () => test$TestEnum;
    setU?: (u: test$TestEnum) => void;
    getV?: () => test$TestMessage$TestNestedEnum;
    setV?: (v: test$TestMessage$TestNestedEnum) => void;
    getWList?: () => Array<test$TestMessage$TestNestedMessage>;
    setWList?: (w: Array<test$TestMessage$TestNestedMessage>) => void;
    addW?: (w: test$TestMessage$TestNestedMessage) => void;
    clearWList?: () => void;
};





export type TestNoStreamRequest = {
    getMessage?: () => test$TestMessage;
    setMessage?: (message: test$TestMessage) => void;
    clearMessage?: () => void;
    hasMessage?: () => boolean;
};





export type TestNoStreamReply = {
    getMessage?: () => test$TestMessage;
    setMessage?: (message: test$TestMessage) => void;
    clearMessage?: () => void;
    hasMessage?: () => boolean;
    getErrMsg?: () => string;
    setErrMsg?: (err_msg: string) => void;
};





export type TestStreamRequestRequest = {
    getMessage?: () => test$TestMessage;
    setMessage?: (message: test$TestMessage) => void;
    clearMessage?: () => void;
    hasMessage?: () => boolean;
};





export type TestStreamRequestReply = {
    getMessage?: () => test$TestMessage;
    setMessage?: (message: test$TestMessage) => void;
    clearMessage?: () => void;
    hasMessage?: () => boolean;
    getErrMsg?: () => string;
    setErrMsg?: (err_msg: string) => void;
};





export type TestStreamReplyRequest = {
    getMessage?: () => test$TestMessage;
    setMessage?: (message: test$TestMessage) => void;
    clearMessage?: () => void;
    hasMessage?: () => boolean;
};





export type TestStreamReplyReply = {
    getMessage?: () => test$TestMessage;
    setMessage?: (message: test$TestMessage) => void;
    clearMessage?: () => void;
    hasMessage?: () => boolean;
    getErrMsg?: () => string;
    setErrMsg?: (err_msg: string) => void;
};





export type TestStreamBothRequest = {
    getMessage?: () => test$TestMessage;
    setMessage?: (message: test$TestMessage) => void;
    clearMessage?: () => void;
    hasMessage?: () => boolean;
};





export type TestStreamBothReply = {
    getMessage?: () => test$TestMessage;
    setMessage?: (message: test$TestMessage) => void;
    clearMessage?: () => void;
    hasMessage?: () => boolean;
    getErrMsg?: () => string;
    setErrMsg?: (err_msg: string) => void;
};

const serializeToBase64 = (byteArray: Uint8Array): string => base64.fromByteArray(byteArray)
const deserializeFromBase64 = (base64Encoded: string): Uint8Array => new Uint8Array(base64.toByteArray(base64Encoded))


function serialize_test_TestNoStreamRequest(arg : TestNoStreamRequest): string {
  if (!(arg instanceof test_pb.TestNoStreamRequest)) {
    throw new Error('Expected argument of type TestNoStreamRequest')
  }
  return serializeToBase64(arg.serializeBinary())
}

function deserialize_test_TestNoStreamRequest(base64Encoded: string): TestNoStreamRequest {
  return test_pb.TestNoStreamRequest.deserializeBinary(deserializeFromBase64(base64Encoded))
}

function serialize_test_TestNoStreamReply(arg : TestNoStreamReply): string {
  if (!(arg instanceof test_pb.TestNoStreamReply)) {
    throw new Error('Expected argument of type TestNoStreamReply')
  }
  return serializeToBase64(arg.serializeBinary())
}

function deserialize_test_TestNoStreamReply(base64Encoded: string): TestNoStreamReply {
  return test_pb.TestNoStreamReply.deserializeBinary(deserializeFromBase64(base64Encoded))
}


function serialize_test_TestStreamRequestRequest(arg : TestStreamRequestRequest): string {
  if (!(arg instanceof test_pb.TestStreamRequestRequest)) {
    throw new Error('Expected argument of type TestStreamRequestRequest')
  }
  return serializeToBase64(arg.serializeBinary())
}

function deserialize_test_TestStreamRequestRequest(base64Encoded: string): TestStreamRequestRequest {
  return test_pb.TestStreamRequestRequest.deserializeBinary(deserializeFromBase64(base64Encoded))
}

function serialize_test_TestStreamRequestReply(arg : TestStreamRequestReply): string {
  if (!(arg instanceof test_pb.TestStreamRequestReply)) {
    throw new Error('Expected argument of type TestStreamRequestReply')
  }
  return serializeToBase64(arg.serializeBinary())
}

function deserialize_test_TestStreamRequestReply(base64Encoded: string): TestStreamRequestReply {
  return test_pb.TestStreamRequestReply.deserializeBinary(deserializeFromBase64(base64Encoded))
}


function serialize_test_TestStreamReplyRequest(arg : TestStreamReplyRequest): string {
  if (!(arg instanceof test_pb.TestStreamReplyRequest)) {
    throw new Error('Expected argument of type TestStreamReplyRequest')
  }
  return serializeToBase64(arg.serializeBinary())
}

function deserialize_test_TestStreamReplyRequest(base64Encoded: string): TestStreamReplyRequest {
  return test_pb.TestStreamReplyRequest.deserializeBinary(deserializeFromBase64(base64Encoded))
}

function serialize_test_TestStreamReplyReply(arg : TestStreamReplyReply): string {
  if (!(arg instanceof test_pb.TestStreamReplyReply)) {
    throw new Error('Expected argument of type TestStreamReplyReply')
  }
  return serializeToBase64(arg.serializeBinary())
}

function deserialize_test_TestStreamReplyReply(base64Encoded: string): TestStreamReplyReply {
  return test_pb.TestStreamReplyReply.deserializeBinary(deserializeFromBase64(base64Encoded))
}


function serialize_test_TestStreamBothRequest(arg : TestStreamBothRequest): string {
  if (!(arg instanceof test_pb.TestStreamBothRequest)) {
    throw new Error('Expected argument of type TestStreamBothRequest')
  }
  return serializeToBase64(arg.serializeBinary())
}

function deserialize_test_TestStreamBothRequest(base64Encoded: string): TestStreamBothRequest {
  return test_pb.TestStreamBothRequest.deserializeBinary(deserializeFromBase64(base64Encoded))
}

function serialize_test_TestStreamBothReply(arg : TestStreamBothReply): string {
  if (!(arg instanceof test_pb.TestStreamBothReply)) {
    throw new Error('Expected argument of type TestStreamBothReply')
  }
  return serializeToBase64(arg.serializeBinary())
}

function deserialize_test_TestStreamBothReply(base64Encoded: string): TestStreamBothReply {
  return test_pb.TestStreamBothReply.deserializeBinary(deserializeFromBase64(base64Encoded))
}


export default {
  
  TestService: {
  
    testnostream: {
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
    teststreamrequest: {
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
    teststreamreply: {
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
    teststreamboth: {
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
  
}
