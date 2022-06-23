const DEVICE_KEY = "invoice-deviceid";

export default {
  existsDeviceId() {
    return !!localStorage.getItem(DEVICE_KEY);
  },

  getDeviceId() {
    return localStorage.getItem(DEVICE_KEY);
  },

  saveDeviceId(deviceId) {
    localStorage.setItem(DEVICE_KEY, deviceId);
  },
};
