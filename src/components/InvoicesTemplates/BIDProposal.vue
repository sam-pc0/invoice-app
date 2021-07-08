<template>
  <section class="proposal">
    <page-header
      @headerChange="handlePageChange"
      :invoiceName="invoice.name"
      :invoiceDescription="invoice.description"
    />
    <section v-animate-css="animatedObject" class="invoice">
      <invoice-header invoiceType="BID Proposal" :number="invoice.number" />
      <owner :invoiceOwner="invoice.owner" />
    </section>
  </section>
</template>

<script>
import { BIDProposal } from "@/type";

import PageHeader from "@/components/InvoicesTemplates/shared/PageHeader";
import InvoiceHeader from "@/components/InvoicesTemplates/shared/InvoiceHeader";
import Owner from "@/components/InvoicesTemplates/shared/Owner";

export default {
  name: "BidProposal",
  components: {
    PageHeader,
    InvoiceHeader,
    Owner,
  },
  props: {
    invoiceData: {
      required: true,
      type: BIDProposal,
    },
  },
  data() {
    return {
      invoice: this.invoiceData,
      animatedObject: {
        classes: "fadeInRight",
        delay: 100,
        duration: 1000,
      },
    };
  },
  methods: {
    handlePageChange(data) {
      console.info(data);
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/sass/index.scss";

@mixin documentContainer {
  margin-top: 2em;
  width: 70%;
  .invoice {
    border: 2px solid $primary;
    background-color: rgba($white, 0.85);
    border-radius: 8px;
    min-height: 8em;
    height: 100vh;
    padding: 2em;
  }
}
.proposal {
  @include documentContainer;
}
</style>
