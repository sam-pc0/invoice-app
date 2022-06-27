import axios from "axios";
import { URL } from "@/services/url";

export default {
  login: async (username, password, deviceId) => {
    return axios.post(`${URL}/login`, { username, password, deviceId });
  },
};
