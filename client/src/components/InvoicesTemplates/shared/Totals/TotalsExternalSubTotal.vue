<template>
  <section class="mt-3 is-flex is-flex-direction-column">
    <span class="has-text-weight-bold">
      Subtotal:
      {{ totals.subTotal ? totals.subTotal : 0 }}
    </span>
    <span class="has-text-weight-bold is-flex" style="max-width: 17.5em">
      <span style="flex-basis: 55%"> Tax Rate: </span>

      <input
        v-model="totals.taxRate"
        default="1"
        class="input is-small"
        type="number"
        placeholder="Tax Rate"
      />
    </span>
    <span class="has-text-weight-bold">
      Taxes:
      {{ getTaxes ? getTaxes : 0 }}
    </span>
    <span class="has-text-weight-bold mb-3">
      Total:
      {{ totals ? getTotal : 0 }}
    </span>
  </section>
</template>

<script>
import { TotalsData } from "@/type";

export default {
  name: "Totals",
  props: {
    invoiceTotals: {
      type: TotalsData,
      required: true,
    },
    isSavedClicked: {
      type: Boolean,
      required: true,
    },
  },
  watch: {
    isSavedClicked(newVal) {
      newVal && this.$emit("onSave", this.getTotals);
    },
    invoiceTotals(newVal) {
      this.totals = newVal;
    },
  },
  computed: {
    getTotals() {
      const { subTotal, taxRate } = this.totals;
      return new TotalsData({
        subTotal,
        taxRate: Number(taxRate),
        tax: this.getTaxes,
        total: this.getTotal,
      });
    },
    getTotal() {
      const { subTotal } = this.totals;
      return this.getTaxes ? subTotal + this.getTaxes : subTotal;
    },
    getTaxes() {
      const { subTotal, taxRate } = this.totals;
      return (subTotal * taxRate) / 100;
    },
  },
  data() {
    return {
      totals: this.invoiceTotals,
    };
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/sass/index.scss";
.input {
  background-color: $input-background;
}
</style>
