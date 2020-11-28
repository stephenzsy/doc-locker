
export interface IServerCertificateConfiguration {
    readonly subject: {
        readonly CN: string;
    };
    readonly serial: string;
}

export type YubiKeySlotId = "82" | "83";

export interface IServerCaCertificateConfiguration extends IServerCertificateConfiguration {
    readonly yubikey: {
        readonly slot: YubiKeySlotId;
    }
}

export interface IServerCertificatesSetupConfiguration {
    readonly libPaths: {
        readonly pkcs11: string;
        readonly ykcs11: string;
    }
    readonly ca: {
        readonly root: [IServerCaCertificateConfiguration, IServerCaCertificateConfiguration?]
        readonly deploy: [IServerCaCertificateConfiguration, IServerCaCertificateConfiguration?]
    }
}

export const pkcs11SlotIdMapping: Partial<{ [key in YubiKeySlotId]: string }> = {
    "82": "5"
}

