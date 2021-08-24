<template>
  <main class="invoice">
    <div
      class="container is-flex is-justify-content-center is-align-items-center"
    >
      <template v-if="invoiceData.template_code === templatesEnum.BID_PROPOSAL">
        <BID-proposal
          @invoiceDataChange="handleInvoiceDataChange"
          :invoiceData="invoiceData"
        />
      </template>

      <template v-if="invoiceData.template_code === templatesEnum.INVOICE">
        <invoice-template
          @invoiceDataChange="handleInvoiceDataChange"
          :invoiceData="invoiceData"
        />
      </template>

      <template
        v-if="invoiceData.template_code === templatesEnum.INVOICE_CONTRACT"
      >
        <contract-invoice
          @invoiceDataChange="handleInvoiceDataChange"
          :invoiceData="invoiceData"
        />
      </template>

      <template
        v-if="
          invoiceData.template_code === templatesEnum.DAILY_N_MATERIAL_RECORD
        "
      >
        <materials-n-workers-record
          @invoiceDataChange="handleInvoiceDataChange"
          :invoiceData="invoiceData"
        />
      </template>
    </div>
  </main>
</template>

<script>
import InvoiceService from "@/services/invoice";
import { templatesEnum } from "@/type";

import BIDProposal from "@/components/InvoicesTemplates/BIDProposal";
import InvoiceTemplate from "@/components/InvoicesTemplates/InvoiceTemplate";
import ContractInvoice from "@/components/InvoicesTemplates/ContractInvoice";
import MaterialsNWorkersRecord from "@/components/InvoicesTemplates/MaterialsNWorkersRecord";

export default {
  name: "Invoice",
  components: {
    BIDProposal,
    InvoiceTemplate,
    ContractInvoice,
    MaterialsNWorkersRecord,
  },
  mounted() {
    this.templatesEnum = templatesEnum;
    this.setInvoiceInfo();
  },
  data() {
    return {
      invoiceData: {},
      templatesEnum,
    };
  },
  methods: {
    setInvoiceInfo() {
      const invoiceId = this.$route.params.invoiceId;
      InvoiceService.get(invoiceId)
        .then((data) => {
          console.info(data);
          this.invoiceData = data;
        })
        .catch((error) => this.$toast.error(error));
    },
    handleInvoiceDataChange(invoiceData) {
      console.info(invoiceData);
      this.invoiceData = invoiceData;
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/sass/index.scss";
.invoice {
  min-height: 100vh;
  background-image: url("../assets/img/background.jpg");
  background-position: center;
  background-attachment: fixed;
  background-position: cover;
}
</style>
