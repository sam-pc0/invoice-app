<template>
  <div
    v-animate-css="animatedObject"
    class="options is-flex is-flex-direction-column"
  >
    <b-button
      @click="$emit('downloadClicked')"
      v-animate-css.click="clickAnimateObject()"
      color="is-primary-light"
      class="mb-2"
    >
      Download Document
    </b-button>
    <b-button
      @click="$emit('saveClicked')"
      v-animate-css.click="clickAnimateObject()"
      type="is-primary"
    >
      Save Document
    </b-button>
  </div>
</template>

<script>
export default {
  name: "OptionButtons",
  data() {
    return {
      animatedObject: {
        classes: "fadeInRight",
        delay: 300,
        duration: 1000,
      },
      clickAnimateObject() {
        return {
          classes: "pulse",
          delay: 0,
          duration: 400,
        };
      },
    };
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/sass/index.scss";
.options {
  position: fixed;
  top: 3em;
  right: 10em;
  border: 2px solid $primary;
  background-color: rgba($white, 0.85);
  border-radius: 8px;
  padding: 2em;
}
</style>
