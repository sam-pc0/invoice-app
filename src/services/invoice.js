// import axios from "axios";
// import { URL } from "@/services/url";

export default {
  create: async (invoiceData) => {
    // return axios.post(`${URL}/invoices`, invoiceData);
    invoiceData.id = 5;
    return invoiceData;
  },
  getAll: async () => {
    // return axios.get(`${URL}/invoices`);
    return [];
  },
  get: async () => {
    // return axios.get(`${URL}/invoices`);
    return [];
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
