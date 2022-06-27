const SESSION_KEY = "invoice-session";
const VALUES = {
  LOGGED: "logged",
  NOT_LOGGED: "not_logged",
};

export default {
  isLogged() {
    return VALUES.LOGGED == localStorage.getItem(SESSION_KEY);
  },

  setLogged() {
    localStorage.setItem(SESSION_KEY, VALUES.LOGGED);
  },

  setNotLogged() {
    localStorage.setItem(SESSION_KEY, VALUES.NOT_LOGGED);
  },
};
