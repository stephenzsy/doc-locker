import fs from 'fs';
import mustache from 'mustache';
import path from 'path';
import { Configuration, Configurations } from '../lib/common/configuration';

const serverSetupConfig = Configurations.getServerSetupConfiguration();
const serverConfigPathBase = path.join(__dirname, '..', 'templates', 'server-config');

const ldsYamlTemplate: string =
    fs.readFileSync(path.join(serverConfigPathBase, 'lds_yaml.mustache'), 'utf-8');

const ldsYamlContent = mustache.render(ldsYamlTemplate, {
    listenerAddress: serverSetupConfig.proxyListenerConfiguration.address,
    listenerPort: serverSetupConfig.proxyListenerConfiguration.port
});

serverSetupConfig.saveEnvoyConfiguration('lds.yaml', ldsYamlContent);

// certificates
const generateCaTemplate: string =
    fs.readFileSync(path.join(serverConfigPathBase, 'generate-ca_sh.mustache'), 'utf-8');
const caCertificateConfig = serverSetupConfig.certificatesConfiguration.ca;
const configurationDir = Configuration.getConfigurationDir();
const scriptPath = path.join(configurationDir, 'scripts', 'create-ca.sh');
const generateCaContent = mustache.render(generateCaTemplate, {
    slot: caCertificateConfig.root[0].yubikey.slot,
    publicKeyPath: `'${path.join(configurationDir, 'setup', 'ca-public-key.pem')}'`,
    subject: `'${caCertificateConfig.root[0].subject}'`,
    serial: caCertificateConfig.root[0].serial,
    certPath: `'${path.join(configurationDir, 'setup', 'ca-root.pem')}'`
});
fs.writeFileSync(scriptPath, generateCaContent);
fs.chmodSync(scriptPath, '755');
