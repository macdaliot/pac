import * as axios from 'axios';
import { config } from '../config';
import {
  CommandResult,
  deleteFile,
  download,
  fileExists,
  httpGetWithBasicAuth,
  logAndQuit,
  logger,
  runSync,
  sleep
} from '../util';

export interface JenkinsPipelineConfig {
  name: string,
  parameters: HashMap
}

export interface JenkinsPipelineStatus {
  _class: string,
  color: string,
  name: string,
  url: string
}

interface HashMap {
  [key: string]: string
}

export const buildJenkinsPipeline = async (pipelineConfig: JenkinsPipelineConfig, projectName: string): Promise<void> => {
  await ensureJenkinsCli(projectName);
  let buildCmd: string = `${getJenkinsCliStart(projectName)} build ${pipelineConfig.name}`;
  for (const parameter in pipelineConfig.parameters) {
    const value: string = pipelineConfig.parameters[parameter];
    buildCmd += ` -p ${parameter}=${value}`;
  }
  const result: CommandResult = runSync(buildCmd);
  if (result.stderr) {
    logger.error(result.stderr);
  }
}

export const checkForSuccessfulPipelineBuild = async (pipelineName: string, projectName: string): Promise<boolean> => {
  await sleep(20);
  const jenkinsStatusUrl: string = getJenkinsStatusUrl(projectName);
  const jenkinsUsername: string = config.jenkins.credentials.username;
  const jenkinsPassword: string = config.jenkins.credentials.password;
  const result: axios.AxiosResponse<any> = await httpGetWithBasicAuth(jenkinsStatusUrl, jenkinsUsername, jenkinsPassword);
  const pipelineStatuses: Array<JenkinsPipelineStatus> = result.data.jobs;
  let pipelineSuccess: boolean = false;
  pipelineStatuses.forEach((pipelineStatus: JenkinsPipelineStatus) => {
    if (pipelineStatus.color.startsWith("red")) {
      logAndQuit(`ERROR: The ${pipelineStatus.name} pipeline failed`);
    }
    if (pipelineStatus.name === pipelineName && pipelineStatus.color === "blue") {
      pipelineSuccess = true;
    }
  });
  return pipelineSuccess;
}

export const doesJenkinsCredentialExist = async (credentialId: string, projectName: string): Promise<boolean> => {
  await ensureJenkinsCli(projectName);
  const retrieveCredentialsCmd: string = `${getJenkinsCliStart(projectName)} get-credentials-as-xml system::system::jenkins \(global\) ${credentialId}`;
  const result: CommandResult = runSync(retrieveCredentialsCmd);
  if (result.exitCode !== 0) {
    return false;
  }
  return true;
}

export const queryJenkinsBuildStatus = async (projectName: string): Promise<Array<JenkinsPipelineStatus>> => {
  const jenkinsStatusUrl: string = getJenkinsStatusUrl(projectName);
  const jenkinsUsername: string = config.jenkins.credentials.username;
  const jenkinsPassword: string = config.jenkins.credentials.password;
  const result: axios.AxiosResponse<any> = await httpGetWithBasicAuth(jenkinsStatusUrl, jenkinsUsername, jenkinsPassword);
  return result.data.jobs;
}

const ensureJenkinsCli = async (projectName: string): Promise<void> => {
  const corruptJenkinsCliError: string = "Error: Invalid or corrupt jarfile jenkins-cli.jar";
  const jenkinsCliLocation: string = `${__dirname}/../../jenkins-cli.jar`;
  const result: CommandResult = runSync(getJenkinsCliStart(projectName));
  if (result.stderr === corruptJenkinsCliError || !fileExists(jenkinsCliLocation)) {
    if (fileExists(jenkinsCliLocation)) {
      deleteFile(jenkinsCliLocation);
    }
    const jenkinsCliUrl = `${getJenkinsUrl(projectName)}/jnlpJars/jenkins-cli.jar`;
    await download(jenkinsCliUrl);
  }
}

const getJenkinsCliStart = (projectName: string): string => {
  const jenkinsUrl: string = getJenkinsUrl(projectName);
  const jenkinsUsername: string = config.jenkins.credentials.username;
  const jenkinsPassword: string = config.jenkins.credentials.password;
  return `java -jar jenkins-cli.jar -s ${jenkinsUrl} -auth ${jenkinsUsername}:${jenkinsPassword}`;
}

const getJenkinsStatusUrl = (projectName: string): string => {
  return `${getJenkinsUrl(projectName)}/api/json`;
}

const getJenkinsUrl = (projectName: string): string => {
  return `https://jenkins.${projectName}.pac.pyramidchallenges.com`;
}
