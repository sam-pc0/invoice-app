<template>
  <section v-animate-css="'fadeInRight'" class="page-header p-5 mb-5">
    <div
      class="is-flex is-justify-content-space-between is-align-items-flex-end"
    >
      <div class="page-header__form">
        <input
          class="
            page-header__name
            input
            is-flex
            column
            is-8-mobile is-8-tablet is-size-12-desktop
          "
          v-model="name"
        />
        <textarea
          class="
            page-header__description
            input
            mt-6
            column
            is-8-mobile is-8-tablet is-size-12-desktop  
          "
          v-model="description"
          maxlength="80"
        />
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: "PageHeader",
  props: ["invoiceName", "invoiceDescription", "isSavedClicked"],
  watch: {
    isSavedClicked(newVal) {
      newVal &&
        this.$emit("onSave", {
          name: this.name,
          description: this.description,
        });
    },
  },
  data() {
    return {
      name: this.invoiceName,
      description: this.invoiceDescription,
    };
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
