/* eslint-disable */
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var udpa_annotations_status_pb = require('../../../udpa/annotations/status_pb.js');
var envoy_config_core_v3_health_check_pb = require('../../../envoy/config/core/v3/health_check_pb.js');
var gogoproto_gogo_pb = require('../../../gogoproto/gogo_pb.js');
goog.exportSymbol('proto.envoy.config.health_checker.http_path.v2.HttpPath', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.envoy.config.health_checker.http_path.v2.HttpPath = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.envoy.config.health_checker.http_path.v2.HttpPath, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.envoy.config.health_checker.http_path.v2.HttpPath.displayName = 'proto.envoy.config.health_checker.http_path.v2.HttpPath';
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
proto.envoy.config.health_checker.http_path.v2.HttpPath.prototype.toObject = function(opt_includeInstance) {
  return proto.envoy.config.health_checker.http_path.v2.HttpPath.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.envoy.config.health_checker.http_path.v2.HttpPath} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.health_checker.http_path.v2.HttpPath.toObject = function(includeInstance, msg) {
  var f, obj = {
    httpHealthCheck: (f = msg.getHttpHealthCheck()) && envoy_config_core_v3_health_check_pb.HealthCheck.HttpHealthCheck.toObject(includeInstance, f)
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
 * @return {!proto.envoy.config.health_checker.http_path.v2.HttpPath}
 */
proto.envoy.config.health_checker.http_path.v2.HttpPath.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.envoy.config.health_checker.http_path.v2.HttpPath;
  return proto.envoy.config.health_checker.http_path.v2.HttpPath.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.envoy.config.health_checker.http_path.v2.HttpPath} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.envoy.config.health_checker.http_path.v2.HttpPath}
 */
proto.envoy.config.health_checker.http_path.v2.HttpPath.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new envoy_config_core_v3_health_check_pb.HealthCheck.HttpHealthCheck;
      reader.readMessage(value,envoy_config_core_v3_health_check_pb.HealthCheck.HttpHealthCheck.deserializeBinaryFromReader);
      msg.setHttpHealthCheck(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.envoy.config.health_checker.http_path.v2.HttpPath.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.envoy.config.health_checker.http_path.v2.HttpPath.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.envoy.config.health_checker.http_path.v2.HttpPath} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.envoy.config.health_checker.http_path.v2.HttpPath.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHttpHealthCheck();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      envoy_config_core_v3_health_check_pb.HealthCheck.HttpHealthCheck.serializeBinaryToWriter
    );
  }
};


/**
 * optional envoy.config.core.v3.HealthCheck.HttpHealthCheck http_health_check = 1;
 * @return {?proto.envoy.config.core.v3.HealthCheck.HttpHealthCheck}
 */
proto.envoy.config.health_checker.http_path.v2.HttpPath.prototype.getHttpHealthCheck = function() {
  return /** @type{?proto.envoy.config.core.v3.HealthCheck.HttpHealthCheck} */ (
    jspb.Message.getWrapperField(this, envoy_config_core_v3_health_check_pb.HealthCheck.HttpHealthCheck, 1));
};


/** @param {?proto.envoy.config.core.v3.HealthCheck.HttpHealthCheck|undefined} value */
proto.envoy.config.health_checker.http_path.v2.HttpPath.prototype.setHttpHealthCheck = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.envoy.config.health_checker.http_path.v2.HttpPath.prototype.clearHttpHealthCheck = function() {
  this.setHttpHealthCheck(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.envoy.config.health_checker.http_path.v2.HttpPath.prototype.hasHttpHealthCheck = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.envoy.config.health_checker.http_path.v2);
