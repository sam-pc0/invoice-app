<template>
  <main class="container">
    <b-modal
      v-model="shouldShowTemplateModal"
      has-modal-card
      trap-focus
      :destroy-on-hide="true"
    >
      <template #default="">
        <TemplatesModal
          @templateSelected="() => console.log('templateSelected')"
          @close="shouldShowTemplateModal = false"
        />
      </template>
    </b-modal>
    <section
      class="
        is-flex is-justify-content-space-between is-align-items-center
        my-5
      "
    >
      <h2 class="is-size-3 has-text-weight-bold">My invoices</h2>
      <b-button @click="handleCreateInvoice" class="is-primary mr-2">
        Create Invoice
      </b-button>
    </section>
    <div class="divider"></div>
    <!-- v-animate-css="'fadeInRight'" -->
    <section class="row columns is-multiline">
      <div
        v-for="(invoice, i) in invoices"
        v-animate-css="getAnimateObject(i)"
        :key="i"
        class="column is-3"
      >
        <div class="card">
          <div class="card-content">
            <h3 class="is-size-3 has-text-weight-bold mb-1">
              {{ invoice.name }}
            </h3>
            <p class="mb-2">
              {{ invoice.description }}
            </p>
            <span class="is-block has-text-weight-medium"> Last edited </span>
            <time datetime="2021-1-1">{{ invoice.timeString }}</time>
          </div>

          <footer class="card-footer">
            <b-button
              @click="handleEditInvoice(invoice.id)"
              class="is-ghost card-footer-item has-text-primary"
              >Edit</b-button
            >
            <b-button
              @click="handleDeleteInvoice"
              class="is-ghost card-footer-item has-text-danger"
              >Delete</b-button
            >
          </footer>
        </div>
      </div>
    </section>
  </main>
</template>

<script>
import InvoiceService from "@/services/invoice";
import TemplatesModal from "@/components/TemplatesModal";

export default {
  name: "Invoices",
  components: { TemplatesModal },
  mounted() {
    this.setInvoices();
  },
  data() {
    return {
      invoices: [],
      shouldShowTemplateModal: false,
    };
  },
  methods: {
    setInvoices() {
      InvoiceService.getAll()
        .then((data) => {
          data = [
            {
              id: "4",
              name: "Testing",
              description:
                "A short description about the information in the invoice",
              timeString: "11:09 PM - 11 Jul 2021",
            },
            {
              id: "4",
              name: "Testing 2",
              description:
                "A short description about the information in the invoice",
              timeString: "11:09 PM - 11 Jul 2021",
            },
            {
              id: "4",
              name: "Testing 3",
              description:
                "A short description about the information in the invoice",
              timeString: "11:09 PM - 11 Jul 2021",
            },
            {
              id: "4",
              name: "Testing 3",
              description:
                "A short description about the information in the invoice",
              timeString: "11:09 PM - 11 Jul 2021",
            },
            {
              id: "4",
              name: "Testing 3",
              description:
                "A short description about the information in the invoice",
              timeString: "11:09 PM - 11 Jul 2021",
            },
            {
              id: "4",
              name: "Testing 3",
              description:
                "A short description about the information in the invoice",
              timeString: "11:09 PM - 11 Jul 2021",
            },
          ];
          this.invoices = data;
        })
        .catch(() =>
          this.$toast.error("An error occurred while uploading invoices")
        );
    },
    getAnimateObject(index) {
      return {
        classes: "fadeInRight",
        delay: index * 100,
        duration: 1000,
      };
    },
    handleCreateInvoice() {
      this.shouldShowTemplateModal = true;
    },
    handleEditInvoice(invoiceId) {
      invoiceId && this.$router.push(`/invoices/${invoiceId}`);
    },
    handleDeleteInvoice(invoiceId) {
      InvoiceService.delete(invoiceId)
        .then(() => {
          this.setInvoices();
          this.$toast.success("Invoice was deleted successfully");
        })
        .catch(() =>
          this.$toast.error("An error occurred while deleting invoice")
        );
    },
  },
};
</script>
<style scoped lang="scss">
@import "@/assets/sass/index.scss";
.divider::before,
.divider::after {
  margin: 0 !important;
}
</style>
