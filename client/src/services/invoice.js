import axios from "axios";
import { URL } from "@/services/url";
// import MockService from "@/services/mock";

export default {
  create: async (invoiceData) => {
    console.log(`${URL}/invoices`, invoiceData);
    return axios.post(`${URL}/invoices`, invoiceData);
    // invoiceData.id = 1;
    // return invoiceData;
  },
  getAll: async () => {
    // return axios.get(`${URL}/invoices`);
    return [];
  },
  get: async (invoiceId) => {
    console.log(invoiceId);
    return axios.get(`${URL}/invoices`);
    // console.info(MockService.generateInvoice(invoiceId));
    // return MockService.generateInvoice(invoiceId);
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
