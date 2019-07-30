const csv = require('csv-parser');
const fs = require('fs');

var rows = [];

function generateTemplate(rows) {
  return `<!DOCTYPE html>
      <html lang="en">
      <head>
          <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
          <title>508 Violations Report</title>
                
          <meta name="viewport" content="width=device-width, initial-scale=1">
          <style>
      
              * {
                  margin: 0;
                  padding: 0;
                  box-sizing: border-box
              }
      
              body, html {
                  height: 100%;
                  font-family: sans-serif
              }
      
              a {
                  margin: 0;
                  transition: all .4s;
                  -webkit-transition: all .4s;
                  -o-transition: all .4s;
                  -moz-transition: all .4s
              }
      
              a:focus {
                  outline: none !important
              }
      
              a:hover {
                  text-decoration: none
              }
      
              h1, h2, h3, h4, h5, h6 {
                  margin: 0
              }
      
              p {
                  margin: 0
              }
      
              ul, li {
                  margin: 0;
                  list-style-type: none
              }
      
              input {
                  display: block;
                  outline: none;
                  border: none !important
              }
      
              textarea {
                  display: block;
                  outline: none
              }
      
              textarea:focus, input:focus {
                  border-color: transparent !important
              }
      
              button {
                  outline: none !important;
                  border: none;
                  background: 0 0
              }
      
              button:hover {
                  cursor: pointer
              }
      
              iframe {
                  border: none !important
              }
      
              .js-pscroll {
                  position: relative;
                  overflow: hidden
              }
      
              .table100 .ps__rail-y {
                  width: 9px;
                  background-color: transparent;
                  opacity: 1 !important;
                  right: 5px
              }
      
              .table100 .ps__rail-y::before {
                  content: "";
                  display: block;
                  position: absolute;
                  background-color: #ebebeb;
                  border-radius: 5px;
                  width: 100%;
                  height: calc(100% - 30px);
                  left: 0;
                  top: 15px
              }
      
              .table100 .ps__rail-y .ps__thumb-y {
                  width: 100%;
                  right: 0;
                  background-color: transparent;
                  opacity: 1 !important
              }
      
              .table100 .ps__rail-y .ps__thumb-y::before {
                  content: "";
                  display: block;
                  position: absolute;
                  background-color: #ccc;
                  border-radius: 5px;
                  width: 100%;
                  height: calc(100% - 30px);
                  left: 0;
                  top: 15px
              }
      
              .limiter {
                  width: 1366px;
                  margin: 0 auto
              }
      
              .container-table100 {
                  width: 100%;
                  min-height: 100vh;
                  background: #fff;
                  display: -webkit-box;
                  display: -webkit-flex;
                  display: -moz-box;
                  display: -ms-flexbox;
                  display: flex;
                  justify-content: center;
                  flex-wrap: wrap;
                  padding: 33px 30px
              }
      
              .wrap-table100 {
                  width: 1170px
              }
      
              .table100 {
                  background-color: #fff
              }
      
              table {
                  width: 100%
              }
      
              th, td {
                  font-weight: unset;
                  padding-right: 10px
              }
      
             .column1 {
                width: 12%;
                padding-left: 40px
              }
          
              .column2 {
                width: 16%
              }
          
              .column3 {
                width: 15%
              }
          
              .column4 {
                width: 15%
              }
          
              .column5 {
                width: 22%
              }
              .column6 {
                width: 20%
              }
      
      
              .table100-head th {
                  padding-top: 18px;
                  padding-bottom: 18px
              }
      
              .table100-body td {
                  padding-top: 16px;
                  padding-bottom: 16px
              }
      
              .table100 {
                  position: relative;
                  padding-top: 60px
              }
      
              .table100-head {
                  position: absolute;
                  width: 100%;
                  top: 0;
                  left: 0
              }
      
              .table100-body {
                  max-height: 585px;
                  overflow: hidden;
              }
      
              .table100.ver1 th {
                  font-family: Lato-Bold;
                  font-size: 18px;
                  color: #fff;
                  line-height: 1.4;
                  background-color: #6c7ae0
              }
      
              .table100.ver1 td {
                  font-family: Lato-Regular;
                  font-size: 15px;
                  line-height: 1.4;
                  color: gray;
              }
      
              .table100.ver1 .table100-body tr:nth-child(even) {
                  background-color: #f8f6ff
              }
      
              .table100.ver1 {
                  border-radius: 10px;
                  overflow: hidden;
                  box-shadow: 0 0 40px 0 rgba(0, 0, 0, .15);
                  -moz-box-shadow: 0 0 40px 0 rgba(0, 0, 0, .15);
                  -webkit-box-shadow: 0 0 40px 0 rgba(0, 0, 0, .15);
                  -o-box-shadow: 0 0 40px 0 rgba(0, 0, 0, .15);
                  -ms-box-shadow: 0 0 40px 0 rgba(0, 0, 0, .15)
              }
      
              .table100.ver1 .ps__rail-y {
                  right: 5px
              }
      
              .table100.ver1 .ps__rail-y::before {
                  background-color: #ebebeb
              }
      
              .table100.ver1 .ps__rail-y .ps__thumb-y::before {
                  background-color: #ccc
              }
              .my-legend .legend-title {
                text-align: left;
                margin-bottom: 5px;
                font-weight: bold;
                font-size: 90%;
                }
              .my-legend .legend-scale ul {
                margin: 0;
                margin-bottom: 5px;
                padding: 0;
                float: left;
                list-style: none;
                }
              .my-legend .legend-scale ul li {
                font-size: 80%;
                list-style: none;
                margin-left: 0;
                line-height: 18px;
                margin-bottom: 2px;
                }
              .my-legend .legend-source {
                font-size: 70%;
                color: #999;
                clear: both;
                }
              .my-legend a {
                color: #777;
                }
      
          </style>
      
          <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.3.1/css/bootstrap.min.css">
          <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.9.0/js/fontawesome.min.js">
          <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/animate.css/3.7.2/animate.min.css">
          <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.7/css/select2.min.css">
          <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/jquery.perfect-scrollbar/1.4.0/css/perfect-scrollbar.min.css">
          <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/jqueryui/1.12.1/themes/base/jquery-ui.min.css">
          <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/tooltipster/3.3.0/css/tooltipster.min.css" />
      
            <h1 align="center" style="padding-top: 20px">508 Violations Report</h1>
            <div class='my-legend' style="padding-left: 45%">
              <div class='legend-title'>Impact Severity</div>
              <div class='legend-scale'>
                <ul class='legend-labels'>
                  <li>Critical</li>
                  <li>Serious</li>
                  <li>Moderate</li>
                  <li>Minor</li>
                </ul>
              </div>
              <div class='legend-source'>Ordered from Most Severe to Least</div>
            </div>
      </head>
      <body>
      <div class="limiter">
          <div class="container-table100">
              <div class="wrap-table100">
                  <div class="table100 ver1 m-b-110">
                      <div class="table100-head">
                          <table>
                              <thead>
                              <tr class="row100 head">
                                  <th class="column1">URL</th>
                                  <th class="column2">Violation</th>
                                  <th class="column3">Impact</th>
                                  <th class="column4">HTML Element</th>
                                  <th class="column5">Messages</th>
                                  <th class="column6s">DOM Element</th>
                              </tr>
                              </thead>
                          </table>
                      </div>
                      <div class="table100-body js-pscroll ps ps--active-y">
                          <table>
                              <tbody>
                              ${rows.map(row => {return row}).join(" ")}
                              </tbody>
                          </table>
                      </div>
                  </div>
      
              </div>
          </div>
      </div>
      
      <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.15.0/esm/popper.min.js" type="text/javascript"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.3.1/js/bootstrap.min.js" type="text/javascript"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/select2/4.0.7/js/select2.min.js" type="text/javascript"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.perfect-scrollbar/1.4.0/perfect-scrollbar.min.js" type="text/javascript"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.4.1/jquery.min.js" type="text/javascript"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/jqueryui/1.12.1/jquery-ui.min.js" type="text/javascript"></script>
      <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/tooltipster/3.3.0/js/jquery.tooltipster.min.js"></script>
      
      <script src="https://cdn.jsdelivr.net/npm/jquery-ellipsis@0.1.6/dist/jquery.ellipsis.min.js" type="text/javascript"></script>
      
      <script type="text/javascript">
          $('.js-pscroll').each(function () {
              var ps = new PerfectScrollbar(this);
      
              $(window).on('resize', function () {
                  ps.update();
              })
          });
      
      
          $('.one-line').ellipsis({ lines: 1 });
          $('.one-line').tooltipster({
              animation: 'fade',
              delay: 200,
              interactive:true
          });
      </script>
      
      </body>
      </html>
      `;
}

fs.createReadStream('test-results.csv')
  .pipe(csv())
  .on('data', (row) => {
    let unparsedEntries = row['HTML Element'].split(",");
    let encodedEntries = [];
    unparsedEntries.forEach(function (value, index, array) {
      unparsedEntries[index] = value.replace(/"/g, "&quot;")
      value = value.replace("<", "&lt");
      value = value.replace(">", "&gt");
      encodedEntries.push(value);
    });

    rows.push(
      `<tr class="row100 body">
    <td class="column1"><a href="${row['URL']}">View page </a></td>
    <td class="column2"><a href="${row['Help']}">${row['Volation Type']} </a></td>
    <td class="column3">${row['Impact']}</td>
    <td class="column4 one-line" title="${unparsedEntries[0]}">${encodedEntries[0]}</td>
    <td class="column5">${unparsedEntries[1]}</td>
    <td class="column6 one-line" title="${unparsedEntries[2]}">${encodedEntries[2]}</td>
   </tr>`);
  })
  .on('end', () => {
    let template = generateTemplate(rows);
    let stream = fs.createWriteStream('./508Report.html');

    stream.once('open', function (fd) {
      stream.end(template);
    });
  });


