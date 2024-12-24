/**
 * @fileoverview gRPC-Web generated client stub for asaycle.routiq.v1
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v3.19.1
// source: v1/tag.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as v1_tag_pb from '../v1/tag_pb'; // proto import: "v1/tag.proto"


export class TagServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorCreateTag = new grpcWeb.MethodDescriptor(
    '/asaycle.routiq.v1.TagService/CreateTag',
    grpcWeb.MethodType.UNARY,
    v1_tag_pb.CreateTagRequest,
    v1_tag_pb.Tag,
    (request: v1_tag_pb.CreateTagRequest) => {
      return request.serializeBinary();
    },
    v1_tag_pb.Tag.deserializeBinary
  );

  createTag(
    request: v1_tag_pb.CreateTagRequest,
    metadata?: grpcWeb.Metadata | null): Promise<v1_tag_pb.Tag>;

  createTag(
    request: v1_tag_pb.CreateTagRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: v1_tag_pb.Tag) => void): grpcWeb.ClientReadableStream<v1_tag_pb.Tag>;

  createTag(
    request: v1_tag_pb.CreateTagRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: v1_tag_pb.Tag) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/asaycle.routiq.v1.TagService/CreateTag',
        request,
        metadata || {},
        this.methodDescriptorCreateTag,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/asaycle.routiq.v1.TagService/CreateTag',
    request,
    metadata || {},
    this.methodDescriptorCreateTag);
  }

  methodDescriptorListTags = new grpcWeb.MethodDescriptor(
    '/asaycle.routiq.v1.TagService/ListTags',
    grpcWeb.MethodType.UNARY,
    v1_tag_pb.ListTagsRequest,
    v1_tag_pb.ListTagsResponse,
    (request: v1_tag_pb.ListTagsRequest) => {
      return request.serializeBinary();
    },
    v1_tag_pb.ListTagsResponse.deserializeBinary
  );

  listTags(
    request: v1_tag_pb.ListTagsRequest,
    metadata?: grpcWeb.Metadata | null): Promise<v1_tag_pb.ListTagsResponse>;

  listTags(
    request: v1_tag_pb.ListTagsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: v1_tag_pb.ListTagsResponse) => void): grpcWeb.ClientReadableStream<v1_tag_pb.ListTagsResponse>;

  listTags(
    request: v1_tag_pb.ListTagsRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: v1_tag_pb.ListTagsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/asaycle.routiq.v1.TagService/ListTags',
        request,
        metadata || {},
        this.methodDescriptorListTags,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/asaycle.routiq.v1.TagService/ListTags',
    request,
    metadata || {},
    this.methodDescriptorListTags);
  }

}

