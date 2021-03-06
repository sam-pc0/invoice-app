<template>
  <main class="container">
    <b-modal
      v-model="shouldShowTemplateModal"
      has-modal-card
      trap-focus
      :destroy-on-hide="true"
    >
      <template #default="">
        <TemplatesModal @close="shouldShowTemplateModal = false" />
      </template>
    </b-modal>

    <section
      class="
        is-flex is-justify-content-space-between is-align-items-center
        my-5
      "
    >
      <h2 class="is-size-3 has-text-weight-bold">My invoices</h2>
      <div>
        <b-button @click="handleLogOut" class="is-secondary mr-2">
          Log out
        </b-button>
        <b-button @click="handleCreateInvoice" class="is-primary mr-2">
          Create Invoice
        </b-button>
      </div>
    </section>
    <div class="divider"></div>
    <!-- v-animate-css="'fadeInRight'" -->
    <template v-if="invoices.length <= 0"> No invoices to show </template>
    <section v-else class="row columns is-multiline">
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
            <span class="is-block has-text-weight-medium"> Last edit </span>
            <time datetime="2021-1-1">{{
              invoice.last_edit.toLocaleString()
            }}</time>
          </div>

          <footer class="card-footer">
            <b-button
              @click="handleEditInvoice(invoice.id)"
              class="is-ghost card-footer-item has-text-primary"
              >Edit</b-button
            >
            <b-button
              @click="handleDeleteInvoice(invoice.id)"
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
import SessionUtil from "@/plugins/session";

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
        .then((response) => {
          const { data } = response;
          this.invoices = data
            .map((element) => ({
              ...element,
              last_edit: new Date(element.last_edit),
            }))
            .reverse();
        })
        .catch(() => this.$toast.error("An error occurred while get invoices"));
    },
    getAnimateObject(index) {
      return {
        classes: "fadeInRight",
        delay: index * 100,
        duration: 1000,
      };
    },
    handleLogOut() {
      SessionUtil.setNotLogged();
      this.$router.push(`/`);
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
