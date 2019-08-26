// This file is written in JavaScript because the type definitions for
//   AxeReport are provided by the maintainers, but are broken as of
//   2019.07.29

import * as AxeBuilder from 'axe-webdriverjs';
import * as AxeReport from 'axe-reports';
import { sleep } from '../util';
import { config } from '../config';

export const analyzePage = async (url, index, driver) => {
  return new Promise((resolve) => {
    process.stdout.write("\r508 analysis of " + url + " has started ...");
    AxeBuilder(driver)
      .analyze()
      .then((results) => {
        if (results) {
          const reportName = `${config.report.name}-${index}`;
          AxeReport.processResults(results, config.report.type, reportName, true);
          // This is necessary because the command above kicks off an asyncronous
          //   process, but does not return a Promise. This leads to the report
          //   file to be created, but not written by the time it is needed
          sleep(2).then(() => {
            console.log(" DONE");
            resolve();
          });
        }
      }).catch((err) => {
        console.error(err);
      });
  });
}
