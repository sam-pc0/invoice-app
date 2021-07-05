// import axios from "axios";
// import { URL } from "@/services/url";

export default {
  getAll: async () => {
    // return axios.get(`${URL}/invoices`);
    return [];
  },
  get: async () => {
    // return axios.get(`${URL}/invoices`);
    return [];
  },
  update: (productObject) => {
    // return axios.put(`${URL}/${URI}`, productObject);
    return productObject;
  },
  delete: (invoiceId) => {
    //     return axios.delete(`${URL}/invoices`, {
    // data: { id_Product },
    // });
    return invoiceId;
  },
};
