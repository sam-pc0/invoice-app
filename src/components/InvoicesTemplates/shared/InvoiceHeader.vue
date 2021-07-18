<template>
  <div
    class="
      invoice-header
      is-flex is-justify-content-space-between is-align-items-flex-end
    "
  >
    <img
      class="invoice-header__logo"
      src="@/assets/img/logo-flat.png"
      alt="Logo Flat"
    />
    <div class="has-text-weight-bold is-size-5">
      {{ invoiceType }} NO. {{ number }}
    </div>
  </div>
</template>

<script>
export default {
  name: "InvoiceHeader",
  props: ["invoiceType", "number"],
};
</script>

<style scoped lang="scss">
@import "@/assets/sass/index.scss";

.invoice-header {
  width: 100%;
  &__logo {
    width: 8em;
  }
}

.divider::before,
.divider::after {
  margin: 0 !important;
  background-color: $primary;
}
</style>
