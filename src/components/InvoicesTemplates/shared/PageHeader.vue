<template>
  <section v-animate-css="'fadeInRight'" class="page-header p-5 mb-5">
    <div
      class="is-flex is-justify-content-space-between is-align-items-flex-end"
    >
      <div class="page-header__form">
        <input class="page-header__name input" v-model="name" />
        <textarea
          class="page-header__description input mt-2"
          v-model="description"
          maxlength="80"
        />
      </div>
      <div class="page-header__buttons is-flex is-flex-direction-column">
        <b-button
          v-animate-css.click="getClickAnimateObject()"
          color="is-primary-light"
          class="mb-2"
        >
          Download Document
        </b-button>
        <b-button
          v-animate-css.click="getClickAnimateObject()"
          type="is-primary"
        >
          Save Document
        </b-button>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: "PageHeader",
  props: ["invoiceName", "invoiceDescription"],
  data() {
    return {
      name: this.invoiceName,
      description: this.invoiceDescription,
    };
  },
  methods: {
    getClickAnimateObject() {
      return {
        classes: "pulse",
        delay: 0,
        duration: 400,
      };
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/sass/index.scss";

.page-header {
  background-color: rgba($white, 0.85);
  border-radius: 8px;
  min-height: 8em;
  &__name,
  &__description {
    word-wrap: break-word;
    word-break: break-all;
    resize: none;
    overflow: auto !important;
  }
  &__name {
    font-size: 1.6em !important;
    min-height: 2em !important;
    font-weight: bold;
    width: 15em;
  }
  &__description {
    min-height: 5em !important;
    width: 35em;
  }
}
</style>
