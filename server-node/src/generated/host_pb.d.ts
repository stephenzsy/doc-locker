import { Message } from 'google-protobuf';

export declare class HostStatusRequest extends Message { };
export declare class HostStatusResponse extends Message {
    getStatusjson(): string;
    setStatusjson(statusJson: string): this;
};
