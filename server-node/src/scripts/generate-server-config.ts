import fs from 'fs';
import mustache from 'mustache';
import path from 'path';
import { Configuration, Configurations } from '../lib/common/configuration';
import { pkcs11SlotIdMapping } from '../lib/common/configuration/server-certificate-setup';

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
const caCertificateConfig = serverSetupConfig.certificatesConfiguration.ca;
const configurationDir = Configuration.getConfigurationDir();

// root ca
const generateCaTemplate: string =
    fs.readFileSync(path.join(serverConfigPathBase, 'generate-root-ca_sh.mustache'), 'utf-8');
const generateCaScriptPath = path.join(configurationDir, 'scripts', 'create-ca.sh');
const generateCaContent = mustache.render(generateCaTemplate, {
    slot: caCertificateConfig.root[0].yubikey.slot,
    cnfPath: path.join(configurationDir, 'setup', 'ca-root.cnf'),
    privateKeyPath: path.join(configurationDir, 'setup', 'ca-root-private-key.pem'),
    subjectCn: caCertificateConfig.root[0].subject.CN,
    serial: caCertificateConfig.root[0].serial,
    certPath: `'${path.join(configurationDir, 'setup', 'ca-root.pem')}'`
});
fs.writeFileSync(generateCaScriptPath, generateCaContent);
fs.chmodSync(generateCaScriptPath, '755');

// deployment intermediate ca
const generateDeployCaTemplate: string =
    fs.readFileSync(path.join(serverConfigPathBase, 'generate-int-ca_sh.mustache'), 'utf-8');
const generateIntCaScriptPath = path.join(configurationDir, 'scripts', 'create-int-ca.sh');
const generateIntCaContent = mustache.render(generateDeployCaTemplate, {
    slot: caCertificateConfig.deploy[0].yubikey.slot,
    cnfPath: path.join(configurationDir, 'setup', 'ca-deploy.cnf'),
    csrPath: path.join(configurationDir, 'setup', 'ca-deploy.csr'),
    privateKeyPath: path.join(configurationDir, 'setup', 'ca-deploy-private-key.pem'),
    subjectCn: caCertificateConfig.deploy[0].subject.CN,
    serial: caCertificateConfig.deploy[0].serial,
    caPath: path.join(configurationDir, 'setup', 'ca-root.pem'),
    certPath: `'${path.join(configurationDir, 'setup', 'ca-deploy.pem')}'`,
    pkcs11slotId: pkcs11SlotIdMapping[caCertificateConfig.root[0].yubikey.slot],
    libPaths: serverSetupConfig.certificatesConfiguration.libPaths,
});
fs.writeFileSync(generateIntCaScriptPath, generateIntCaContent);
fs.chmodSync(generateIntCaScriptPath, '755');

// deployment encryption keys with certificate