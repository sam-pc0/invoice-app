// import axios from "axios";
// import { URL } from "@/services/url";

export default {
  login: async (username, password) => {
    // return axios.post(`${URL}/auth`, { username, password });
    return { username, password };
  },
};
