import axios from "axios";
import { URL } from "@/services/url";

export default {
  login: async (username, password) => {
    return axios.post(`${URL}/login`, { username, password });
  },
};
