// source: google/api/launch_stage.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() {
  if (this) { return this; }
  if (typeof window !== 'undefined') { return window; }
  if (typeof global !== 'undefined') { return global; }
  if (typeof self !== 'undefined') { return self; }
  return Function('return this')();
}.call(null));

goog.exportSymbol('proto.google.api.LaunchStage', null, global);
/**
 * @enum {number}
 */
proto.google.api.LaunchStage = {
  LAUNCH_STAGE_UNSPECIFIED: 0,
  EARLY_ACCESS: 1,
  ALPHA: 2,
  BETA: 3,
  GA: 4,
  DEPRECATED: 5
};

goog.object.extend(exports, proto.google.api);
