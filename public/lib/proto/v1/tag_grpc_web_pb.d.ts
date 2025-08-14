import * as grpcWeb from 'grpc-web';

import * as v1_tag_pb from '../v1/tag_pb'; // proto import: "v1/tag.proto"


export class TagServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  createTag(
    request: v1_tag_pb.CreateTagRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_tag_pb.Tag) => void
  ): grpcWeb.ClientReadableStream<v1_tag_pb.Tag>;

  listTags(
    request: v1_tag_pb.ListTagsRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_tag_pb.ListTagsResponse) => void
  ): grpcWeb.ClientReadableStream<v1_tag_pb.ListTagsResponse>;

}

export class TagServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  createTag(
    request: v1_tag_pb.CreateTagRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_tag_pb.Tag>;

  listTags(
    request: v1_tag_pb.ListTagsRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_tag_pb.ListTagsResponse>;

}

