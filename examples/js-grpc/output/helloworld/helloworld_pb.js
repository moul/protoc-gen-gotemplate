// GENERATED CODE -- DO NOT EDIT!


var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();


goog.exportSymbol('proto.user.HelloRequest', null, global);

goog.exportSymbol('proto.user.HelloReply', null, global);



/**
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.user.HelloRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.user.HelloRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.user.HelloRequest.displayName = 'proto.user.HelloRequest';
}

if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.user.HelloRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.user.HelloRequest.toObject(opt_includeInstance, this);
};

/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.user.HelloRequest} msg The msg instance to transform.
 * @return {!Object}
 */
proto.user.HelloRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    profile: (f = msg.getProfile()) && proto.user.Profile.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}

/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.user.HelloRequest}
 */
proto.user.HelloRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.user.HelloRequest;
  return proto.user.HelloRequest.deserializeBinaryFromReader(msg, reader);
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.user.HelloRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.user.HelloRequest}
 */
proto.user.HelloRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.user.Profile;
      reader.readMessage(value,proto.user.Profile.deserializeBinaryFromReader);
      msg.setProfile(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};

/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.user.HelloRequest} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.user.HelloRequest.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};

/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.user.HelloRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};

/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.user.HelloRequest.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getProfile();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.user.Profile.serializeBinaryToWriter
    );
  }
};

/**
 * optional Profile profile = 1;
 * @return {?proto.user.Profile}
 */
proto.user.HelloRequest.prototype.getProfile = function() {
  return /** @type{?proto.user.Profile} */ (
    jspb.Message.getWrapperField(this, proto.user.Profile, 1));
};

/** @param {?proto.user.Profile|undefined} value */
proto.user.HelloRequest.prototype.setProfile = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};

proto.user.HelloRequest.prototype.clearProfile = function() {
  this.setProfile(undefined);
};

/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.user.HelloRequest.prototype.hasProfile = function() {
  return jspb.Message.getField(this, 1) != null;
};
// --------------------------------









/**
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.user.HelloReply = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.user.HelloReply, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.user.HelloReply.displayName = 'proto.user.HelloReply';
}

if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.user.HelloReply.prototype.toObject = function(opt_includeInstance) {
  return proto.user.HelloReply.toObject(opt_includeInstance, this);
};

/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.user.HelloReply} msg The msg instance to transform.
 * @return {!Object}
 */
proto.user.HelloReply.toObject = function(includeInstance, msg) {
  var f, obj = {
    profile: (f = msg.getProfile()) && proto.user.Profile.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}

/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.user.HelloReply}
 */
proto.user.HelloReply.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.user.HelloReply;
  return proto.user.HelloReply.deserializeBinaryFromReader(msg, reader);
};

/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.user.HelloReply} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.user.HelloReply}
 */
proto.user.HelloReply.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.user.Profile;
      reader.readMessage(value,proto.user.Profile.deserializeBinaryFromReader);
      msg.setProfile(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};

/**
 * Class method variant: serializes the given message to binary data
 * (in protobuf wire format), writing to the given BinaryWriter.
 * @param {!proto.user.HelloReply} message
 * @param {!jspb.BinaryWriter} writer
 */
proto.user.HelloReply.serializeBinaryToWriter = function(message, writer) {
  message.serializeBinaryToWriter(writer);
};

/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.user.HelloReply.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  this.serializeBinaryToWriter(writer);
  return writer.getResultBuffer();
};

/**
 * Serializes the message to binary data (in protobuf wire format),
 * writing to the given BinaryWriter.
 * @param {!jspb.BinaryWriter} writer
 */
proto.user.HelloReply.prototype.serializeBinaryToWriter = function (writer) {
  var f = undefined;
  f = this.getProfile();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.user.Profile.serializeBinaryToWriter
    );
  }
};

/**
 * optional Profile profile = 1;
 * @return {?proto.user.Profile}
 */
proto.user.HelloReply.prototype.getProfile = function() {
  return /** @type{?proto.user.Profile} */ (
    jspb.Message.getWrapperField(this, proto.user.Profile, 1));
};

/** @param {?proto.user.Profile|undefined} value */
proto.user.HelloReply.prototype.setProfile = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};

proto.user.HelloReply.prototype.clearProfile = function() {
  this.setProfile(undefined);
};

/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.user.HelloReply.prototype.hasProfile = function() {
  return jspb.Message.getField(this, 1) != null;
};
// --------------------------------










goog.object.extend(exports, proto.user);
