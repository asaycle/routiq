import * as jspb from 'google-protobuf'

import * as google_api_label_pb from '../../google/api/label_pb'; // proto import: "google/api/label.proto"


export class MetricDescriptor extends jspb.Message {
  getName(): string;
  setName(value: string): MetricDescriptor;

  getType(): string;
  setType(value: string): MetricDescriptor;

  getLabelsList(): Array<google_api_label_pb.LabelDescriptor>;
  setLabelsList(value: Array<google_api_label_pb.LabelDescriptor>): MetricDescriptor;
  clearLabelsList(): MetricDescriptor;
  addLabels(value?: google_api_label_pb.LabelDescriptor, index?: number): google_api_label_pb.LabelDescriptor;

  getMetricKind(): MetricDescriptor.MetricKind;
  setMetricKind(value: MetricDescriptor.MetricKind): MetricDescriptor;

  getValueType(): MetricDescriptor.ValueType;
  setValueType(value: MetricDescriptor.ValueType): MetricDescriptor;

  getUnit(): string;
  setUnit(value: string): MetricDescriptor;

  getDescription(): string;
  setDescription(value: string): MetricDescriptor;

  getDisplayName(): string;
  setDisplayName(value: string): MetricDescriptor;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MetricDescriptor.AsObject;
  static toObject(includeInstance: boolean, msg: MetricDescriptor): MetricDescriptor.AsObject;
  static serializeBinaryToWriter(message: MetricDescriptor, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MetricDescriptor;
  static deserializeBinaryFromReader(message: MetricDescriptor, reader: jspb.BinaryReader): MetricDescriptor;
}

export namespace MetricDescriptor {
  export type AsObject = {
    name: string,
    type: string,
    labelsList: Array<google_api_label_pb.LabelDescriptor.AsObject>,
    metricKind: MetricDescriptor.MetricKind,
    valueType: MetricDescriptor.ValueType,
    unit: string,
    description: string,
    displayName: string,
  }

  export enum MetricKind { 
    METRIC_KIND_UNSPECIFIED = 0,
    GAUGE = 1,
    DELTA = 2,
    CUMULATIVE = 3,
  }

  export enum ValueType { 
    VALUE_TYPE_UNSPECIFIED = 0,
    BOOL = 1,
    INT64 = 2,
    DOUBLE = 3,
    STRING = 4,
    DISTRIBUTION = 5,
    MONEY = 6,
  }
}

export class Metric extends jspb.Message {
  getType(): string;
  setType(value: string): Metric;

  getLabelsMap(): jspb.Map<string, string>;
  clearLabelsMap(): Metric;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Metric.AsObject;
  static toObject(includeInstance: boolean, msg: Metric): Metric.AsObject;
  static serializeBinaryToWriter(message: Metric, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Metric;
  static deserializeBinaryFromReader(message: Metric, reader: jspb.BinaryReader): Metric;
}

export namespace Metric {
  export type AsObject = {
    type: string,
    labelsMap: Array<[string, string]>,
  }
}

