'use strict';
/*
 * The requirement for this script is to restore the infrastructure of this project.
 * Note: This script does not support deploying to a region different from the original.
 * 
 * This is to simulate restoring the infrastructure after a disaster. It assumes the following resources still exist.
 *   KMS keys
 *   DNS records and SSL certs
 *   S3 buckets
 *   ECR (docker repositories)
 *   Docker images in the ECR
 * 
 * The script should also support the addition of services.
 * 
 * Known Issues:
 * Must comment out source_code_hash line in each service file because the zip file it refers to doesn't
 * exist locally. The script does this automatically, but this change should not be commited to the repo.
 * 
 * If Elasticseach has to be recreated there may be a race condition setting up trigger mapping requiring second run
 * of the script.
 * 
 * If Cloudfront has to be recreated, during Cloudfront creation there is a race condition with DNS validation that
 *  causes the need to re-run the script.
 * 
 * Jenkins pipeline will no longer be setup and can't be setup via this script (at this time).
 */
const fs           = require('fs');
const { execSync } = require('child_process');
const replace      = require('replace-in-file');
const __basedir    = __dirname;

// Read in configuration file.
const configFile = fs.readFileSync(".pac.json");
let config = JSON.parse(configFile);

/**
 * Description: Executes a shell command on the local host. exexSync is used so that the commands are run
 * syncronously, otherwise there is unexpected behavior from Terraform.
 * 
 * @param {string} c
 */
function cmd(c) {
  execSync(c, {stdio: 'inherit'});
}

/**
 * Description: Navigates to the path passed in and executes terraform init, plan, and apply back-to-back.
 * @param {string} path 
 */
function runTerraform(path) {
  process.chdir(__basedir + '/' + path);

  // Remove tfplan file so Terraform doesn't complain if its state is not up-to-date
  if(fs.existsSync('./tfplan')) {
    fs.unlinkSync('./tfplan');
  }

  cmd('terraform init');
  cmd('terraform plan -out tfplan');
  cmd('terraform apply -auto-approve tfplan');
}

/**
 * Comment out souce_code_hash in service files.
 */
async function replaceTemplateStrings() {
  const options = {
    files: [
      'terraform/**/lambda*.tf'
    ],
    //Replacement to make (string or regex) 
    from: [
      / source_code_hash/g,
    ],
    to: [
      '#source_code_hash'
    ]
  };

  try {
    await replace(options, (error, files) => {
      if (error) {
        return console.error('Error occurred:', error);
      }
      printModifiedFiles(files);
    });
  } catch {
    //carry on
  }
}

function printModifiedFiles(files) {
  let modifiedFiles = "";
  files.forEach((file) => {
    if (file.hasChanged) {
      modifiedFiles += file.file + ", ";
    }
  });
  modifiedFiles = modifiedFiles.substring(0, modifiedFiles.length - 2);
  if (modifiedFiles.length > 0) {
    modifiedFiles = 'Modified files: ' + modifiedFiles;
  } else {
    modifiedFiles = 'No files were modified';
  }
  console.log(modifiedFiles);
}

async function main() {
  await replaceTemplateStrings()

  // Apply Terraform templates for all major Terraform lifecycles
  runTerraform('terraform/management');
  runTerraform('terraform/development');
  runTerraform('terraform/staging');
  runTerraform('terraform/production');

  // Output URL to access site securely and login
  console.log(`Live link: https://development.${config.projectFqdn}\n`);
}

main();
