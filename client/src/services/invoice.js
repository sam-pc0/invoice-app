import axios from "axios";
import { URL } from "@/services/url";
import { templatesEnum, BIDProposal, Invoice } from "@/type";
// import MockService from "@/services/mock";
//
const castDataToBill = (data) => {
  const { template_code } = data;
  switch (template_code) {
    case templatesEnum.BID_PROPOSAL:
      return new BIDProposal(data);
    case templatesEnum.INVOICE:
      return new Invoice(data);
  }
};

export default {
  create: async (invoiceData) => {
    return axios.post(`${URL}/bills`, invoiceData);
  },
  getAll: async () => {
    return axios.get(`${URL}/bills`);
  },
  get: async (invoiceId) => {
    return new Promise((resolve, reject) => {
      axios
        .get(`${URL}/bills/bill/${invoiceId}`)
        .then((response) => {
          console.info(response.data);
          const castedData = castDataToBill(response.data);
          if (castedData === null) {
            reject();
          }
          resolve(castedData);
        })
        .catch(reject);
    });
  },
  update: (invoiceData) => {
    return axios.put(`${URL}/bills/bill/${invoiceData.id}`, invoiceData);
  },
  delete: (invoiceId) => {
    return axios.delete(`${URL}/bills/bill/${invoiceId}/`);
  },
};
