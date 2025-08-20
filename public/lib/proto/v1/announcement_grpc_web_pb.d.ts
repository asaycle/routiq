import * as grpcWeb from 'grpc-web';

import * as v1_announcement_pb from '../v1/announcement_pb'; // proto import: "v1/announcement.proto"


export class AnnouncementServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  listAnnouncements(
    request: v1_announcement_pb.ListAnnouncementsRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_announcement_pb.ListAnnouncementsResponse) => void
  ): grpcWeb.ClientReadableStream<v1_announcement_pb.ListAnnouncementsResponse>;

  getAnnouncement(
    request: v1_announcement_pb.GetAnnouncementRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_announcement_pb.Announcement) => void
  ): grpcWeb.ClientReadableStream<v1_announcement_pb.Announcement>;

}

export class AnnouncementServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  listAnnouncements(
    request: v1_announcement_pb.ListAnnouncementsRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_announcement_pb.ListAnnouncementsResponse>;

  getAnnouncement(
    request: v1_announcement_pb.GetAnnouncementRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_announcement_pb.Announcement>;

}

