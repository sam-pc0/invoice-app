import axios from "axios";
import { URL } from "@/services/url";

export default {
  create: async () => {
    return axios.get(`${URL}/device/create`);
  },
};
