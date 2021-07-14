<template>
  <section class="proposal">
    <page-header
      @headerChange="handlePageChange"
      :invoiceName="invoice.name"
      :invoiceDescription="invoice.description"
    />
    <option-buttons
      @saveClicked="handleSave"
      @downladClicked="handleDownload"
    />
    <section v-animate-css="animatedObject" class="invoice">
      <invoice-header invoiceType="BID Proposal" :number="invoice.number" />
      <owner :invoiceOwner="invoice.owner" />
      <p class="is-size-6 mt-3">
        <b>a. Scope of work: </b> Flores's General Contruction, Inc. hereby
        submits the following specifications and estimates:
      </p>
      <section class="is-flex is-justify-content-center mt-2">
        <textarea
          @keydown="(e) => autoHeight(e, '20em')"
          :model="invoice.specificationNStimates"
          class="lined-input-area input textarea is-small"
        />
      </section>

      <p class="is-size-6 mt-3">
        <b>b. Not Included: </b> This proposal does not include:
      </p>
      <section class="is-flex is-justify-content-center mt-2">
        <textarea
          @keydown="(e) => autoHeight(e, '4em')"
          :model="invoice.notIncluded"
          class="lined-input input textarea is-small"
        />
      </section>
      <section class="columns mt-3">
        <div class="column is-half is-size-7">
          <div class="total">
            <b>c. WE PROPOSE: </b> for furnish materia, equipment and labor in
            accordance with the above specifications for the sum of:
            <span>
              <input
                v-model="invoice.totalSum"
                class="input is-small is-inline"
                type="number"
                style="width: 10em"
                placeholder="total"
              />
            </span>
            dollars.
          </div>

          <div class="total mt-3">
            <b>NOTE: </b> This proposal may be withdrawn if not accepted within
            <span>
              <input
                v-model="invoice.withdrawnDays"
                class="input is-small is-inline"
                type="number"
                style="width: 8em"
                placeholder="days"
              />
            </span>
            days from
            <span class="is-inline-block" style="max-width: 10em">
              <b-datepicker
                v-model="invoice.withdrawnDate"
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
          </div>

          <div class="mt-4">
            Respectfully submitted by:
            <span>
              <input
                v-model="invoice.submittedBy"
                placeholder="Company Representative"
                class="input is-small is-inline"
              />
            </span>
          </div>
        </div>
        <div class="column is-size-7 is-half">
          <b>d. WE ACCEPT: </b> the prices, specifications, and terms as stated
          in this bid proposal are approved. We authorize you to draw up all
          necesary contract documents so work can begin.
          <div
            class="
              approved
              is-flex is-justify-content-center is-flex-direction-column
            "
          >
            <div class="is-flex has-text-centered">
              <span style="width: 60%" class="is-line is-inline-block mt-5 mr-3"
                >approved and accepted (owner or owner's authorized agent)</span
              >
              <span
                class="is-line is-text-align-center is-inline-block mt-5"
                style="width: 40%"
                >date</span
              >
            </div>

            <div class="is-flex has-text-centered">
              <span style="width: 60%" class="is-line is-inline-block mt-5 mr-3"
                >approved and accepted (second owner - if any)</span
              >
              <span
                class="is-line is-text-align-center is-inline-block mt-5"
                style="width: 40%"
                >date</span
              >
            </div>
          </div>
        </div>
      </section>
    </section>
  </section>
</template>

<script>
import { BIDProposal } from "@/type";

import PageHeader from "@/components/InvoicesTemplates/shared/PageHeader";
import OptionButtons from "@/components/InvoicesTemplates/shared/OptionButtons";
import InvoiceHeader from "@/components/InvoicesTemplates/shared/InvoiceHeader";
import Owner from "@/components/InvoicesTemplates/shared/Owner";

export default {
  name: "BidProposal",
  components: {
    PageHeader,
    OptionButtons,
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
    handleDownload() {
      console.info(this.invoice);
    },
    handleSave() {
      console.info(this.invoice);
    },
    handlePageChange(data) {
      console.info(data);
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

@mixin documentContainer {
  margin-top: 2em;
  width: 70%;
  position: relative;
  .invoice {
    border: 2px solid $primary;
    background-color: rgba($white, 0.85);
    border-radius: 8px;
    padding: 2em;
  }
}
.proposal {
  @include documentContainer;
}

.is-line {
  border-top: 1px solid $primary;
}
.lined-input,
.lined-input-area {
  min-width: 0 !important;
  width: 95% !important;
  min-height: 4em;
  line-height: 2.3em;
  background-image: linear-gradient(
    rgba($primary, 0.2) 0.1em,
    transparent 0.1em
  );
  background-size: 100% 2.3em;
}
.lined-input-area {
  min-height: 20em !important;
}
</style>
