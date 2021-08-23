<template>
  <section class="invoice-template">
    <page-header
      @onSave="handlePageChange"
      :invoiceName="invoice.name"
      :invoiceDescription="invoice.description"
      :isSavedClicked="isSavedClicked"
    />
    <option-buttons
      @saveClicked="handleSave"
      @downladClicked="handleDownload"
    />
    <section v-animate-css="animatedObject" class="invoice">
      <invoice-header invoiceType="Invoice" :number="invoice.number" />
      <owner
        @onSave="handleOwnerChange"
        :invoiceOwner="invoice.owner"
        :isSavedClicked="isSavedClicked"
      />
    </section>
  </section>
</template>

<script>
import PageHeader from "@/components/InvoicesTemplates/shared/PageHeader";
import OptionButtons from "@/components/InvoicesTemplates/shared/OptionButtons";
import InvoiceHeader from "@/components/InvoicesTemplates/shared/InvoiceHeader";
import Owner from "@/components/InvoicesTemplates/shared/Owner";
//import { Invoice } from "@/type";
import { BIDProposal } from "@/type";

export default {
  name: "InvoiceTemplate",
  components: {
    PageHeader,
    OptionButtons,
    InvoiceHeader,
    Owner,
  },
  props: {
    invoiceData: {
      required: true,
      //type: Invoice,
      type: BIDProposal,
    },
  },
  data() {
    return {
      invoice: this.invoiceData,
      isSavedClicked: false,
      animatedObject: {
        classes: "fadeInRight",
        delay: 100,
        duration: 1000,
      },
    };
  },
  methods: {
    handleDownload() {
      if (!this.isSavedClicked) {
        this.handleSave();
      }
    },
    handleSave() {
      if (!this.isSavedClicked) {
        this.isSavedClicked = true;
        setTimeout(() => {
          console.info(this.invoice);
          this.isSavedClicked = false;
        }, 100);
      }
    },
    handlePageChange(data) {
      this.invoice.name = data.name;
      this.invoice.description = data.description;
    },
    handleOwnerChange(owner) {
      this.invoice.owner = owner;
    },
    autoHeight({ target }, minHeight) {
      target.style.height = minHeight;
      target.style.height = 15 + target.scrollHeight + "px";
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/sass/index.scss";
.invoice-template {
  @include documentContainer;
}
</style>
