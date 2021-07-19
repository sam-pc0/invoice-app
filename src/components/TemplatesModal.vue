<template>
  <main class="modal-card">
    <header class="modal-card-head">
      <p class="modal-card-title">
        {{ step === 1 ? "Select a template" : "Document data" }}
      </p>
    </header>
    <section class="modal-card-body">
      <template v-if="step === 1">
        <section class="row columns is-multiline my-1">
          <div
            v-for="(template, i) in templates"
            v-animate-css="getAnimateObject(i)"
            :key="i"
            class="column is-6"
          >
            <div
              @click="handleTemplateClick(template.id)"
              v-animate-css.click="getClickAnimateObject()"
              :class="selectedTemplate === template.id ? 'selected' : ''"
              class="card is-flex is-align-items-center"
            >
              <div class="card-content">
                <p class="title">
                  {{ template.name }}
                </p>
              </div>
            </div>
          </div>
        </section>
      </template>
      <template v-else>
        <section>
          <form class="form__main">
            <b-field v-animate-css="'fadeInRight'" label="Name">
              <b-input
                v-model="invoiceName"
                placeholder="Jhon Due Document"
              ></b-input>
            </b-field>
            <b-field v-animate-css="animatedObject" label="Description">
              <b-input
                v-model="description"
                class="has-fixed-size"
                type="textarea"
                placeholder="A short description about the content of the document"
                maxlength="80"
              >
              </b-input>
            </b-field>
          </form>
        </section>
      </template>
    </section>
    <footer class="modal-card-foot is-justify-content-flex-end">
      <b-button label="Close" @click="$emit('close')" />
      <b-button
        v-if="step === 1"
        @click="step = 2"
        :disabled="selectedTemplate === null"
        label="Next"
        type="is-primary"
      />
      <b-button
        v-else
        @click="handleCreate"
        :disabled="!isFormValid"
        label="Create"
        type="is-primary"
      />
    </footer>
  </main>
</template>

<script>
import { templates } from "@/type";
import InvoiceService from "@/services/invoice";

export default {
  name: "TemplatesModal",
  computed: {
    isFormValid() {
      return this.invoiceName !== "" && this.description !== "";
    },
  },
  data() {
    return {
      step: 1,
      selectedTemplate: null,
      invoiceName: "",
      description: "",
      templates,
      animatedObject: {
        classes: "fadeInRight",
        delay: 100,
        duration: 1000,
      },
    };
  },
  methods: {
    handleCreate() {
      const invoiceData = {
        template: this.selectedTemplate,
        name: this.invoiceName,
        description: this.description,
      };
      InvoiceService.create(invoiceData)
        .then(({ id }) => id && this.$router.push(`/invoices/${id}`))
        .catch(() =>
          this.$toast.error("An error occurred while deleting invoice")
        );
    },
    handleButtonClick() {
      this.step === 1 ? (this.step = 2) : this.handleCreate();
    },
    handleTemplateClick(templateId) {
      this.selectedTemplate = templateId;
    },
    getClickAnimateObject() {
      return {
        classes: "pulse",
        delay: 0,
        duration: 700,
      };
    },
    getAnimateObject(index) {
      return {
        classes: "fadeInRight",
        delay: index * 100,
        duration: 1000,
      };
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/sass/index.scss";

@mixin hoverCard {
  background-color: $primary;
  transition: background-color 0.7s ease-in;
  .title {
    color: $white !important;
    transition: color 0.4s ease-out;
  }
}

.textarea {
  resize: none !important;
}

.modal-card-body {
  overflow-x: hidden;
}

.selected {
  @include hoverCard;
}
.card {
  height: 7em;
  cursor: pointer;
  &:hover {
    @include hoverCard;
  }
}
.modal-card {
  width: 45em;
}
</style>
