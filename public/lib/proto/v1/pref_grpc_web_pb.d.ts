import * as grpcWeb from 'grpc-web';

import * as v1_pref_pb from '../v1/pref_pb'; // proto import: "v1/pref.proto"


export class PrefServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  listPrefs(
    request: v1_pref_pb.ListPrefsRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: v1_pref_pb.ListPrefsResponse) => void
  ): grpcWeb.ClientReadableStream<v1_pref_pb.ListPrefsResponse>;

}

export class PrefServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  listPrefs(
    request: v1_pref_pb.ListPrefsRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<v1_pref_pb.ListPrefsResponse>;

}

