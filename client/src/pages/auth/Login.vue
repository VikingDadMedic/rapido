<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Félix PORTIER <f.portierdev@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <Main>
    <header class="template-header">
      <navbar class="template-header-menu" />
    </header>
    <section class="section template-main">
      <div class="container-fluid h-100">
        <div class="row h-100 align-items-center">
          <div class="col-lg-8 m-auto">
            <card v-bind:title="$t('account__connection')">

              <!-- Register button -->
              <div class="form-group row">
                <span class="col-md-12 d-flex justify-content-end">
                  <router-link tag="button" :to="{ path: '/register', query: $route.query }" class="btn btn-primary">
                    {{ $t('button__create_account') }}
                  </router-link>
                </span>
              </div>
              <hr>

              <form @submit.prevent="login" @keydown="form.onKeydown($event)">

                <div class="form-group row">
                  <label class="col-md-4 col-form-label text-md-right">
                    {{ $t('settings_accounts__email_address') }}
                  </label>
                  <span class="col-md-8">
                    <input v-model="form.username" type="text" name="username" class="form-control">
                  </span>
                </div>

                <!-- Password -->
                <div class="form-group row">
                  <label class="col-md-4 col-form-label text-md-right">
                    {{ $t('settings_accounts__password') }}
                  </label>
                  <span class="col-md-8">
                    <input v-model="form.password" type="password" name="password" class="form-control">
                  </span>
                </div>

                <!-- Remember me -->
                <div class="form-group row">
                  <span class="col-md-8 offset-md-4 d-flex">
                    <checkbox v-model="form.remember" name="remember">
                      {{ $t('account__remember_me') }}
                    </checkbox>
                    <router-link :to="{ name: 'Email' }" class="ml-auto my-auto">
                      {{ $t('account__reset_password') }}
                    </router-link>
                  </span>
                </div>

                <!-- Login button -->
                <div class="form-group row">
                  <span class="col-md-8 offset-md-4 d-flex">
                    <v-button :loading="form.busy || processed" class="btn btn-success">
                      {{ $t('button__connection') }}
                    </v-button>
                  </span>
                </div>

                <!-- Email address -->
              </form>
            </card>
          </div>
        </div>
      </div>
    </section>
  </Main>
</template>

<script>
/* eslint-disable */
import Navbar from "@/components/Navbar.vue";
import Form from "vform";
import axios from "axios";

export default {
  layout: "default",
  components: {
    Navbar,
  },

  metaInfo() {
    return { title: this.$t('account__connection') };
  },

  data: () => ({
    processed: false,
    form: new Form({
      username: "",
      password: "",
      remember: false,
    }),
    regexpPattern: /^.+@.+$/,
    countdown: null,
    title: "",
  }),

  mounted() {
    if (this.$route.query.token) {
      this.send();
    }
  },

  methods: {
    async login() {
      if (!this.form.username || !this.regexpPattern.test(this.form.username.trim())) {
        this.$awn.alert(this.$t('your_email_address_is_invalid'))
      } else if (!this.form.password || this.form.password.trim() === '') {
        this.$awn.alert(this.$t('settings_accounts__invalid_password'))
      } else {
        this.form.username = this.form.username.trim()
        try {
          // Submit the form.
          const { data } = await this.form.post(
            process.env.API_URL + "/api/login"
          );

          // Save the token.
          this.$store.dispatch("auth/saveToken", {
            token: data.token,
            remember: this.form.remember,
          });

          // Fetch the user.
          await this.$store.dispatch("auth/fetchUser");

          // Redirect home.
          if (this.$route.query.redirect) {
            this.$router.push(this.$route.query.redirect);
          } else {
            this.$router.push("/");
          }
        } catch (e) {
          this.$awn.alert(this.$t(e.response.data.message));
        }
      }
    },
    async send() {
      try {
        if (this.$route.query.id && this.$route.query.token) {
          let obj = {
            id: this.$route.query.id,
            activate_token: this.$route.query.token,
          };
          const { data } = await axios.get(process.env.API_URL + "/api/activate_account", {
            params: obj,
          });
          this.$awn.success(this.$t(data.message))
        }
      } catch (e) {
        this.$awn.alert(this.$t(e.response.data.message))
        this.startCountdown();
      }
    },
  },
};

</script>

<style lang="stylus" scoped>
.template-main {
  height: 100vh;
}
</style>
