<template>
  <h2>Contract Invoice works</h2>
</template>

<script>
import { MaterialsNWorkersRecord } from "@/type";
export default {
  name: "MaterialsNWorkersRecord",
  props: {
    invoiceData: {
      required: true,
      type: MaterialsNWorkersRecord,
    },
  },
};
</script>

<style scoped lang="scss"></style>
