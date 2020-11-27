import fs from 'fs';
import mustache from 'mustache';
import path from 'path';
import { Configurations } from '../lib/common/configuration';

const serverSetupConfig = Configurations.getServerSetupConfiguration();

const ldsYamlTemplate: string = fs.readFileSync(path.join(__dirname, '..', 'templates', 'server-config', 'lds_yaml.mustache'), 'utf-8');

const ldsYamlContent = mustache.render(ldsYamlTemplate, {
    listenerAddress: serverSetupConfig.proxyListenerConfiguration.address,
    listenerPort: serverSetupConfig.proxyListenerConfiguration.port
});

serverSetupConfig.saveEnvoyConfiguration('lds.yaml', ldsYamlContent);