<!--
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
-->


<template>
  <Main>
    <section class="section template-main">
      <div class="container-fluid h-100">
        <div class="row h-100 align-items-center">
          <div class="col-lg-8 m-auto">
            <card :title="$t('account__reset_password')">

              <!-- Return to home page button -->
              <div class="form-group row">
                <span class="col-md-12 d-flex justify-content-end">
                  <router-link tag="button" type="button" :to="{ name: 'Login' }" class="btn btn-primary">
                    {{ $t('button__back_to_connection_page') }}
                  </router-link>
                </span>
              </div>
              <hr>


              <form @submit.prevent="reset" @keydown="form.onKeydown($event)">
                <alert-success :form="form" :message="status" />

                <!-- Password -->
                <div class="form-group row">
                  <label class="col-md-4 col-form-label text-md-right">
                    {{ $t('settings_accounts__new_password') }}
                  </label>
                  <span class="col-md-8">
                    <input v-model="form.new_password" ref="new_password" type="password" name="new_password" class="form-control">
                  </span>
                </div>

                <!-- Password Confirmation -->
                <div class="form-group row">
                  <label class="col-md-4 col-form-label text-md-right">
                    {{ $t('settings_accounts__confirm_password') }}
                  </label>
                  <span class="col-md-8">
                    <input v-model="form.password_confirmation" type="password" name="password_confirmation" class="form-control">
                  </span>
                </div>

                <!-- Register button -->
                <div class="form-group row">
                  <span class="col-md-8 offset-md-4 d-flex">
                    <v-button :loading="form.busy" class="btn btn-success">
                      {{ $t('button__save') }}
                    </v-button>
                  </span>
                </div>

              </form>
            </card>
          </div>
        </div>
      </div>
    </section>
  </Main>
</template>

<script>
import Form from 'vform'

export default {
  layout: 'default',

  metaInfo () {
    return { title: this.$t('account__reset_password') }
  },

  data: () => ({
    status: '',
    verified: false,
    form: new Form({
      reminder_token: null,
      id: null,
      new_password: '',
      password_confirmation: ''
    })
  }),

  methods: {
    reset () {
      this.$validator.validateAll()
        .then(async (res) => {
          if (!this.form.new_password || this.form.new_password.trim() === '' || !this.form.password_confirmation || this.form.password_confirmation.trim() === '') {
            this.$awn.alert(this.$t('settings_accounts__invalid_password'))
          } else if (this.form.new_password !== this.form.password_confirmation) {
            this.$awn.alert(this.$t('settings_accounts__passwords_do_not_match'))
          } else {
            if (res) {
              this.form.id = parseInt(this.$route.query.id)
              this.form.reminder_token = this.$route.query.token
              try {
                const { data } = await this.form.post(process.env.API_URL + '/api/new_password')
                this.$awn.success(this.$t(data.message))
                this.$router.push('/login')
              } catch (e) {
                this.$awn.alert(this.$t(e.response.data.message))
                this.form.reset()
              }
            }
          }
        })
    }
  }
}
</script>

<style lang="stylus" scoped>
.template-main
  height: 100vh
</style>
