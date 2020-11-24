import { Message } from 'google-protobuf';

export declare class HostStatusRequest extends Message { };
export declare class HostStatusResponse extends Message {
    getStatusJson(): string;
    setStatusJson(statusJson: string): this;
};
