import axios from "axios";
import { URL } from "@/services/url";
// import MockService from "@/services/mock";

export default {
  create: async (invoiceData) => {
    return axios.post(`${URL}/bills`, invoiceData);
  },
  getAll: async () => {
    return axios.get(`${URL}/bills`);
  },
  get: async (invoiceId) => {
    console.info(`${URL}/bills/bill/${invoiceId}`);
    return axios.get(`${URL}/bills/bill/${invoiceId}`);
  },
  update: (invoiceData) => {
    // return axios.put(`${URL}/invoices`, invoiceData);
    return invoiceData;
  },
  delete: (invoiceId) => {
    //     return axios.delete(`${URL}/invoices`, {
    // data: { id_Product },
    // });
    return invoiceId;
  },
};
