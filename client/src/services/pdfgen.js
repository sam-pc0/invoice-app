const { http, https } = require("follow-redirects");
var url = require("url");
import { BIDProposal } from "@/type";

function makeRequest(urlEndpoint, method, apiKey, data = null) {
  let d = "";
  if (data != null) d = JSON.stringify(data);
  const uri = url.parse(urlEndpoint);
  const proto = uri.protocol === "https:" ? https : http;
  const opts = {
    method: method,
    headers: {
      "Content-Length": d.length,
      "Content-Type": "application/json",
      "X-API-KEY": apiKey,
    },
  };

  console.log(proto);
  console.log(opts);

  return new Promise((resolve, reject) => {
    const req = proto.request(urlEndpoint, opts, (res) => {
      res.setEncoding("utf8");
      let responseBody = "";

      res.on("data", (chunk) => {
        responseBody += chunk;
      });

      res.on("end", () => {
        resolve(responseBody);
      });
    });

    req.on("error", (err) => {
      reject(err);
    });
    if (data) {
      req.write(d);
    }
    req.end();
  });
}

let apiKey = "64dcNDc2NToxNzc2OlFHYmVpaWwzaFNzTWZjQWU";
let template_id = "04e77b2b1dbde0ca";

export function downloadBID(bid = BIDProposal) {
  return new Promise(async (resolve, reject) => {
    try {
      let resp = await makeRequest(
        "https://api.apitemplate.io/v1/create?template_id=" + template_id,
        "POST",
        apiKey,
        bid
      );
      let ret = JSON.parse(resp);
      alert(ret.download_url);
      //const link = createElement("a");
      //link.href = ret.url;
      //link.click;
      resolve("success");
    } catch (error) {
      reject(error);
    }
  });
}
