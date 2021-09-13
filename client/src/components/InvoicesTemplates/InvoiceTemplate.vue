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
      @downloadClicked="handleDownload"
    />
    <section v-animate-css="animatedObject" class="invoice">
      <invoice-header invoiceType="Invoice" :number="invoice.id" />
      <owner
        @onSave="handleOwnerChange"
        :invoiceOwner="invoice.owner"
        :isSavedClicked="isSavedClicked"
      />
      <ItemList
        @onSave="handleItemsListChange"
        :itemList="invoice.items"
        :isSavedClicked="isSavedClicked"
      />
      <section class="mt-4">
        <span> Date submitted: </span>
        <span class="is-inline-block" style="max-width: 10em">
          <b-datepicker
            v-model="invoice.dateSubmitted"
            class="is-inline"
            size="is-small"
            locale="en-US"
            placeholder="Select Date"
            icon="calendar-today"
            trap-focus
            position="is-top-right"
            style="width: 10em"
          >
          </b-datepicker>
        </span>
      </section>
      <section class="mt-6">
        <div class="is-flex is-justify-content-space-around has-text-centered">
          <span class="is-size-7 has-text-weight-bold">
            APPROVED BY OWNER
          </span>
          <span style="width: 45%" class="is-line is-inline-block mt-5 mr-3"
            >(Owner's Signature)</span
          >
          <span class="is-size-7 has-text-weight-bold"> SUBMITTED BY </span>
          <span style="width: 45%" class="is-line is-inline-block mt-5 mr-3"
            >(Contractors's Signature)</span
          >
        </div>
      </section>
    </section>
  </section>
</template>

<script>
import InvoiceService from "@/services/invoice";

import PageHeader from "@/components/InvoicesTemplates/shared/PageHeader";
import OptionButtons from "@/components/InvoicesTemplates/shared/OptionButtons";
import InvoiceHeader from "@/components/InvoicesTemplates/shared/InvoiceHeader";
import Owner from "@/components/InvoicesTemplates/shared/Owner";
import ItemList from "@/components/InvoicesTemplates/shared/ItemList";
import { downloadINVO } from "@/services/pdfgeninvo2";
import { Invoice } from "@/type";
const REDIRECT = true;
const NO_REDIRECT = false;
export default {
  name: "InvoiceTemplate",
  components: {
    PageHeader,
    OptionButtons,
    InvoiceHeader,
    Owner,
    ItemList,
  },
  props: {
    invoiceData: {
      required: true,
      type: Invoice,
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
      //!this.isSavedClicked && this.handleSave();

      this.handleSave(NO_REDIRECT)
        .then(() => {
          this.isSavedClicked = true;
          downloadINVO(this.invoice)
            .then(() => this.$toast.success("Document were saved"))
            .catch((error) => this.$toast.error(error));
        })
        .catch((error) => this.$toast.error(error))
        .finally(() => (this.isSavedClicked = false));
    },
    async handleSave() {
      if (!this.isSavedClicked) {
        this.isSavedClicked = true;
        setTimeout(() => {
          const invoice = { ...this.invoice };
          invoice.dateSubmitted = invoice.dateSubmitted.toString();
          InvoiceService.update(this.invoice)
            .then(() => {
              this.isSavedClicked = false;
              this.$toast.success("Document were saved");
              this.$router.replace("/invoices");
            })
            .catch((error) => this.$toast.error(error));
        }, 100);
      }
    },
    handlePageChange(data) {
      this.invoice.name = data.name;
      this.invoice.description = data.description;
      this.invoice.last_edit = new Date();
    },
    handleItemsListChange({ item, total }) {
      this.invoice.items = item;
      this.invoice.total = total;
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
.is-line {
  border-top: 1px solid $primary;
}
.invoice-template {
  @include documentContainer;
}
</style>
