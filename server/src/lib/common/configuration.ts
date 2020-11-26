import path from 'path';
import process from 'process';
import fs from 'fs';

export abstract class Configuration<T> {
    public static docLockerConfigDirEnvKey = 'DOC_LOCKER_CONFIG_DIR';
    protected readonly data: T

    public static getConfigurationDir(): string {
        const configDirectory: string = process.env[Configuration.docLockerConfigDirEnvKey] as string;
        if (!configDirectory) {
            throw new Error(`Environment variable not set: ${Configuration.docLockerConfigDirEnvKey}`)
        }
        return configDirectory;
    }

    protected constructor(configurationKeys: string[]) {
        const configDirectory: string = Configuration.getConfigurationDir();
        const filePath = path.join(configDirectory, ...configurationKeys);

        this.data = this.loadFromFile(filePath);
    }

    private loadFromFile(filePath: string): T {
        return require(filePath) as T;
    }
}

export interface IServerSetupConfiguration {
    readonly port: number;
}

export class ServerSetupConfiguration extends Configuration<IServerSetupConfiguration>{

    private static instance: ServerSetupConfiguration | undefined;
    public static getInstance(): ServerSetupConfiguration {
        this.instance = this.instance || new ServerSetupConfiguration();
        return this.instance;
    }

    protected constructor() {
        super(['setup', 'server.json']);
    }

    public get port(): number {
        return this.data.port;
    }

    public saveEnvoyConfiguration(key: 'lds.yaml', content: string) {
        const configDirectory: string = Configuration.getConfigurationDir();
        const filepath = path.join(configDirectory, 'envoy', key);
        fs.writeFileSync(filepath, content);
    }

}

export class Configurations {
    public static getServerSetupConfiguration(): ServerSetupConfiguration {
        return ServerSetupConfiguration.getInstance();
    }
}
