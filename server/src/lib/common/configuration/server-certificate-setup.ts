
export interface IServerCertificateConfiguration {
    readonly subject: string;
    readonly serial: string;
    readonly yubikey: {
        readonly slot: "82" | "83";
    }
}

export interface IServerCertificatesSetupConfiguration {
    readonly ca: {
        readonly root: [IServerCertificateConfiguration, IServerCertificateConfiguration?]
    }
}


