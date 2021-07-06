<template>
  <main class="container">
    {{ invoiceId }}
  </main>
</template>

<script>
export default {
  name: "Invoice",
  mounted() {
    this.setInvoiceInfo();
  },
  data() {
    return {
      invoiceId: "",
    };
  },
  methods: {
    setInvoiceInfo() {
      this.invoiceId = this.$route.params.invoiceId;
    },
  },
};
</script>

<style scoped></style>
