package configurations

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func loadConfigFromFile(filePath string, configData interface{}) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bytes, configData)
	return err
}

type configurations struct {
	configDir        string
	serverSetup      *ServerSetupConfiguration
	serverSetupError error
	serverSetupOnce  sync.Once
}

func (c *configurations) ServerSetup() (*ServerSetupConfiguration, error) {
	c.serverSetupOnce.Do(func() {
		c.serverSetup, c.serverSetupError = newSetupConfiguration(c.configDir)
	})

	return c.serverSetup, c.serverSetupError
}

var (
	config     *configurations
	configOnce sync.Once
)

func newConfigurations() *configurations {
	configDir := os.Getenv("DOCLOCKER_CONFIG_DIR")
	return &configurations{
		configDir: configDir,
	}
}

func Configurations() *configurations {
	configOnce.Do(func() {
		config = newConfigurations()
	})

	return config
}

func (c *configurations) ConfigRootDir() string {
	return c.configDir
}

/*

	import path from 'path';
	import process from 'process';
	import { IServerCertificatesSetupConfiguration } from './server-certificate-setup';

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

	export interface IListenerSetupConfiguration {
		readonly address: string;
		readonly port: number;
	}

	export interface IServerSetupConfiguration {
		readonly proxyListener: IListenerSetupConfiguration;
		readonly serverListener: IListenerSetupConfiguration;
		readonly certificates: IServerCertificatesSetupConfiguration;
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

		public get proxyListenerConfiguration(): IListenerSetupConfiguration {
			return this.data.proxyListener;
		}

		public get serverListenerConfiguration(): IListenerSetupConfiguration {
			return this.data.serverListener;
		}

		public get certificatesConfiguration(): IServerCertificatesSetupConfiguration {
			return this.data.certificates;
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

}

/*
import fs from 'fs';
import path from 'path';
import process from 'process';
import { IServerCertificatesSetupConfiguration } from './server-certificate-setup';

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

export interface IListenerSetupConfiguration {
    readonly address: string;
    readonly port: number;
}

export interface IServerSetupConfiguration {
    readonly proxyListener: IListenerSetupConfiguration;
    readonly serverListener: IListenerSetupConfiguration;
    readonly certificates: IServerCertificatesSetupConfiguration;
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

    public get proxyListenerConfiguration(): IListenerSetupConfiguration {
        return this.data.proxyListener;
    }

    public get serverListenerConfiguration(): IListenerSetupConfiguration {
        return this.data.serverListener;
    }

    public get certificatesConfiguration(): IServerCertificatesSetupConfiguration {
        return this.data.certificates;
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
*/
