/* eslint-disable */
// package: consul.options.gloo.solo.io
// file: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/consul/consul.proto

import * as jspb from "google-protobuf";
import * as extproto_ext_pb from "../../../../../../../../../extproto/ext_pb";
import * as github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_pb from "../../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/service_spec_pb";

export class UpstreamSpec extends jspb.Message {
  getServiceName(): string;
  setServiceName(value: string): void;

  clearServiceTagsList(): void;
  getServiceTagsList(): Array<string>;
  setServiceTagsList(value: Array<string>): void;
  addServiceTags(value: string, index?: number): string;

  clearSubsetTagsList(): void;
  getSubsetTagsList(): Array<string>;
  setSubsetTagsList(value: Array<string>): void;
  addSubsetTags(value: string, index?: number): string;

  clearInstanceTagsList(): void;
  getInstanceTagsList(): Array<string>;
  setInstanceTagsList(value: Array<string>): void;
  addInstanceTags(value: string, index?: number): string;

  clearInstanceBlacklistTagsList(): void;
  getInstanceBlacklistTagsList(): Array<string>;
  setInstanceBlacklistTagsList(value: Array<string>): void;
  addInstanceBlacklistTags(value: string, index?: number): string;

  hasServiceSpec(): boolean;
  clearServiceSpec(): void;
  getServiceSpec(): github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_pb.ServiceSpec | undefined;
  setServiceSpec(value?: github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_pb.ServiceSpec): void;

  getConsistencymode(): UpstreamSpec.ConsulConsistencyModesMap[keyof UpstreamSpec.ConsulConsistencyModesMap];
  setConsistencymode(value: UpstreamSpec.ConsulConsistencyModesMap[keyof UpstreamSpec.ConsulConsistencyModesMap]): void;

  getConnectEnabled(): boolean;
  setConnectEnabled(value: boolean): void;

  clearDataCentersList(): void;
  getDataCentersList(): Array<string>;
  setDataCentersList(value: Array<string>): void;
  addDataCenters(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpstreamSpec.AsObject;
  static toObject(includeInstance: boolean, msg: UpstreamSpec): UpstreamSpec.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UpstreamSpec, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpstreamSpec;
  static deserializeBinaryFromReader(message: UpstreamSpec, reader: jspb.BinaryReader): UpstreamSpec;
}

export namespace UpstreamSpec {
  export type AsObject = {
    serviceName: string,
    serviceTagsList: Array<string>,
    subsetTagsList: Array<string>,
    instanceTagsList: Array<string>,
    instanceBlacklistTagsList: Array<string>,
    serviceSpec?: github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_service_spec_pb.ServiceSpec.AsObject,
    consistencymode: UpstreamSpec.ConsulConsistencyModesMap[keyof UpstreamSpec.ConsulConsistencyModesMap],
    connectEnabled: boolean,
    dataCentersList: Array<string>,
  }

  export interface ConsulConsistencyModesMap {
    CONSISTENTMODE: 0;
    DEFAULTMODE: 2;
    STALEMODE: 1;
  }

  export const ConsulConsistencyModes: ConsulConsistencyModesMap;
}
