<template>
  <div class="item-list mt-5">
    <template v-if="list.length > 0">
      <div class="item" v-for="(item, i) in list" :key="i">
        <input
          v-model="item.item"
          class="item__name input is-small is-inline"
          placeholder="Item"
        />

        <input
          v-model="item.description"
          class="item__description input is-small is-inline"
          placeholder="Description"
        />

        <input
          v-model="item.amount"
          class="item__amount input is-small is-inline"
          type="number"
          placeholder="Amount"
        />
        <span
          @click="handleDelete(i)"
          class="item__delete mdi mdi-delete"
        ></span>
      </div>
    </template>
    <section class="is-flex is-justify-content-center" v-else>
      <span class="is-size-4"> Not items to show </span>
    </section>
    <section class="mt-2 is-flex is-justify-content-flex-end">
      <b-button @click="handleAdd" type="item-list__button is-primary">
        <span class="mdi mdi-plus-thick"></span>
        Add</b-button
      >
    </section>
    <section class="mt-3">
      <span class="has-text-weight-bold"> Total of this Invoice: </span>
      {{ totalInvoice }}
    </section>
  </div>
</template>

<script>
export default {
  name: "ItemsList",
  props: {
    itemList: {
      type: Array,
    },
    isSavedClicked: {
      type: Boolean,
    },
  },
  watch: {
    isSavedClicked(newVal) {
      const item = this.list.map((element) => ({
        ...element,
        amount: Number(element.amount),
      }));
      newVal &&
        this.$emit("onSave", {
          item,
          total: Number(this.totalInvoice),
        });
    },
  },
  computed: {
    totalInvoice() {
      return this.list
        .map((item) => Number(item.amount))
        .reduce((total, amount) => total + amount);
    },
  },
  data() {
    return {
      list: this.itemList,
    };
  },
  methods: {
    handleAdd() {
      this.list.push({ name: "", description: "", amount: undefined });
    },
    handleDelete(index) {
      this.list.splice(index, 1);
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/assets/sass/index.scss";
.item-list {
  &__button {
    margin-top: 1em;
  }
}
.item {
  width: 100%;
  display: flex;
  justify-content: space-around;
  align-items: center;
  margin-top: 0.5em;
  &__name {
    flex-basis: 25%;
  }
  &__description {
    flex-basis: 50%;
  }
  &__amount {
    flex-basis: 20%;
  }
  &__delete {
    color: $primary;
    flex-basis: 5%;
    transition: transform 0.5s ease;
    &:hover {
      transform: scale(1.2);
    }
  }
}
</style>
