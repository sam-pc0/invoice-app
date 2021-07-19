<template>
  <main class="hero">
    <div class="columns login">
      <section class="login__info column">
        <div class="info__header">
          <img class="info__logo" src="@/assets/img/logo.png" alt="logo" />
          <h1 class="is-size-2 has-text-weight-bold">Invoice Generator</h1>
          <h2 class="info__subtitle">Flores's General Contruction, Inc.</h2>
        </div>
        <p class="info__caption">
          <b-icon icon="receipt" class="info__icon"></b-icon>
          Easy document handling.
        </p>
      </section>
      <section
        v-animate-css="'fadeInRight'"
        class="
          login__form
          column
          is-flex is-align-items-center is-justify-content-center
        "
      >
        <div class="login__form-container">
          <h2 class="is-size-3 has-text-weight-bold mb-5">Sign In</h2>
          <form @submit.prevent="handleFormSubmit" class="form__main">
            <b-field>
              <b-input
                v-model="username"
                placeholder="username"
                maxlength="30"
              ></b-input>
            </b-field>
            <b-field>
              <b-input
                v-model="password"
                type="password"
                placeholder="**************"
                password-reveal
              >
              </b-input>
            </b-field>
            <section
              class="
                form__options
                is-flex is-justify-content-space-between is-align-items-center
                mt-6
              "
            >
              <a
                class="is-size-6 has-text-weight-light is-underline"
                href="google.com"
              >
                Trouble login?
              </a>
              <button
                v-animate-css.click="getClickAnimateObject()"
                :disabled="!areFieldsWithData"
                class="button is-primary"
                type="submit"
              >
                Login
              </button>
            </section>
          </form>
        </div>
      </section>
    </div>
  </main>
</template>

<script>
import authService from "@/services/auth";

export default {
  name: "Login",
  data() {
    return {
      username: "",
      password: "",
    };
  },
  computed: {
    areFieldsWithData() {
      return this.username !== "" && this.password !== "";
    },
  },
  methods: {
    handleFormSubmit() {
      console.info(authService);
      authService
        .login(this.username, this.password)
        .then(() => {
          this.$router.push("/invoices");
          this.$toast.success("Login Successful");
        })
        .catch(() => this.$toast.error("Check your login information"));
    },
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

<style lang="scss" scoped>
@import "@/assets/sass/index.scss";

.login {
  height: 100vh;
  margin-top: 0;

  &__form-container {
    width: 30em;
  }

  .info__logo,
  .info__header,
  &__info {
    display: flex;
    align-items: center;
    justify-content: center;
  }
  &__info {
    flex-direction: column;
    //background: linear-gradient(34deg, $primary 0%, $primary-light 100%);
    background-image: url("../assets/img/background.jpg");
    background-position: center;
    color: aliceblue;
    text-align: left;
    .info {
      &__header {
        flex-basis: 100%;
        flex-direction: column;
        margin-top: -1em;
      }
      &__logo {
        $dimension: 9em;
        width: $dimension;
      }
      &__icon {
        color: $primary-light;
      }
      &__subtitle {
        color: $primary-light;
        font-size: 1.4em;
      }
      &__caption {
        color: $primary-light;
        margin-bottom: 3em;
        font-size: 1.1em;
      }
    }
  }
}
</style>
