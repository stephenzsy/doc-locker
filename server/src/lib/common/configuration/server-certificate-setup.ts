
export interface IServerCertificateConfiguration {
    readonly subject: {
        readonly CN: string;
    };
    readonly serial: string;
}

export interface IServerCaCertificateConfiguration extends IServerCertificateConfiguration {
    readonly yubikey: {
        readonly slot: "82" | "83";
    }
}

export interface IServerCertificatesSetupConfiguration {
    readonly ca: {
        readonly root: [IServerCaCertificateConfiguration, IServerCaCertificateConfiguration?]
    }
}


