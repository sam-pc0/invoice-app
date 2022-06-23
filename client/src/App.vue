<template>
  <router-view />
</template>

<script>
import DeviceService from "@/services/device";
import DeviceUtil from "@/plugins/device";

export default {
  name: "App",
  mounted() {
    this.verifyDevice();
  },
  methods: {
    async verifyDevice() {
      if (!DeviceUtil.existsDeviceId()) {
        const deviceId = (await DeviceService.create()).data;
        DeviceUtil.saveDeviceId(deviceId);
      }
    },
  },
};
</script>
