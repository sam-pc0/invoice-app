<template>
  <h2>Invoice works</h2>
</template>

<script>
import { Invoice } from "@/type";

export default {
  name: "Invoice",
  props: {
    invoiceData: {
      required: true,
      type: Invoice,
    },
  },
};
</script>

<style scoped lang="scss"></style>
