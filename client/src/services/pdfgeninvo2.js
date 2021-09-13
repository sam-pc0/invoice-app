const { http, https } = require("follow-redirects");
var url = require("url");
import { InvoiceTemplate } from "@/type";

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

let apiKey = "54c8NDY5MjoxNzAzOktGNGFBTUtGNGR2U2M4enI";
let template_id = "58377b2b1d71950e";

export function downloadINVO(invo = InvoiceTemplate) {
  return new Promise(async (resolve, reject) => {
    try {
      let resp = await makeRequest(
        "https://api.apitemplate.io/v1/create?template_id=" + template_id,
        "POST",
        apiKey,
        invo
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
