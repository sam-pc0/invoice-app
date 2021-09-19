<template>
  <section class="owner columns mt-3 is-flex">
    <div class="owner__form column is-two-thirds">
      <div class="columns is-variable is-1">
        <div class="column is-two-fifths">
          <input
            v-model="owner.name"
            class="input is-small"
            placeholder="Owner's Name"
          />
        </div>
        <div class="column is-one-third">
          <input
            v-model="owner.address"
            class="input is-small"
            placeholder="Owner's Address"
          />
        </div>
      </div>

      <div class="columns is-variable is-1">
        <div class="column is-half">
          <input
            v-model="owner.location"
            class="input is-small"
            placeholder="Owner's City, State, Zip"
          />
        </div>
        <div class="column is-one-fifth">
          <input
            v-model="owner.phone"
            class="input is-small"
            placeholder="Owner's Phone"
          />
        </div>
        <div class="column is-one-fifth">
          <input
            v-model="owner.altPhone"
            class="input is-small"
            label="Owner's Alt. Phone"
            placeholder="Owner's Alt Phone"
          />
        </div>
      </div>

      <div class="columns is-variable is-1">
        <div class="column is-half">
          <input
            v-model="owner.projectNameNAddress"
            class="input is-small"
            placeholder="Project  Name & Address"
          />
        </div>
        <div class="column is-one-third">
          <input
            v-model="owner.email"
            class="input is-small"
            placeholder="Owner's Email"
          />
        </div>
      </div>
    </div>
    <div class="owner__info">
      <span class="column is-8-mobile is-8-tablet is-8-desktop is-8-widescreen">
        1218 Grand Ave.
      </span>
      <span
        class="
          has-text-right
          is-size-6
          column
          is-8-mobile is-8-tablet is-8-desktop is-8-widescreen
        "
        >Santa Rosa CA 95404
      </span>
      <br />
      <span
        class="
          has-text-right has-text-weight-bold
          column
          is-8-mobile is-8-tablet is-8-desktop is-8-widescreen
        "
      >
        707-236-7078 </span
      ><br />
      <span
        class="
          has-text-right
          is-size-6
          column
          is-8-mobile is-8-tablet is-8-desktop is-8-widescreen
        "
        >CA CSLB License. #1074035
      </span>
      <br />

      <span
        class="
          has-text-justified
          column
          is-4-mobile is-1-tablet is-4-desktop is-4-widescreen
        "
        >hector@floresgeneralconstruction.com </span
      ><br />
    </div>
  </section>
</template>

<script>
export default {
  name: "Owner",
  props: ["invoiceOwner", "isSavedClicked"],
  watch: {
    isSavedClicked(newVal) {
      newVal && this.$emit("onSave", this.invoiceOwner);
    },
  },
  data() {
    return {
      owner: this.invoiceOwner,
    };
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/sass/index.scss";

.input {
  background-color: $input-background;
}

.columns {
  margin-bottom: 0 !important;
}
.owner {
  border-bottom: 1px solid rgba($primary, 0.2);
  &__info {
    text-align: right;
  }
}
</style>
