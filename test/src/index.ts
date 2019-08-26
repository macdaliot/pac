import {
  buildManuallyTriggeredJenkinsPipelines,
  createPacServices,
  installPacBinary,
  modifyTerraformTemplates,
  pushToGitHub,
  stopWebpackDevServer,
  runPacAutomateJenkins,
  runPacCleanCloud,
  runPacDeploy,
  runPacSetup,
  runStressTest,
  runWebpackDevServer,
  validateTestStartConditions,
  verifyDevServer,
  verifyGitHubRepository,
  verifyJenkinsCredentials,
  verifyManuallyTriggeredPipelineSuccesses,
  verifyOnDiskContent,
  verifyPipelineCount,
  verifyWebhookTriggeredPipelineSuccesses
} from './steps';

const main = async () => {
  await validateTestStartConditions();
  installPacBinary();
  const projectName: string = runPacSetup();
  const projectDirectory: string = verifyOnDiskContent(projectName);
  runPacDeploy(projectDirectory);
  await verifyGitHubRepository(projectName);
  await runWebpackDevServer(projectDirectory);
  await verifyDevServer();
  stopWebpackDevServer();
  await verifyJenkinsCredentials(projectName);
  await runPacAutomateJenkins(projectDirectory);
  await verifyPipelineCount(projectName);
  pushToGitHub(projectDirectory);
  await verifyWebhookTriggeredPipelineSuccesses(projectName);
  await buildManuallyTriggeredJenkinsPipelines(projectName);
  verifyManuallyTriggeredPipelineSuccesses(projectName);
  await createPacServices(projectDirectory);
  pushToGitHub(projectDirectory);
  await verifyWebhookTriggeredPipelineSuccesses(projectName);
  await runStressTest(projectDirectory, projectName);
  await modifyTerraformTemplates(projectDirectory);
  runPacCleanCloud(projectDirectory);
};

main().catch(console.error);
